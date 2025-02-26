package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

const LogFilePath = "logs.txt"
const Rebuiltlog = "logs3.txt"

type Storage interface {
	// a struct that represents a storage interface
	Get(key string) (string, error)
	Set(key string, value string) error
	Delete(key string) error
	Exists(key string) (bool, error)
}

type KeyValue struct {
	// a struct that represents a key-value storage
	mu   sync.RWMutex
	data map[string]string
}

type LogEntry struct {
	// a struct that represents a log entry
	Operation string `json:"operation"`
	Key       string `json:"key"`
	Value     string `json:"value,omitempty"`
}

func CheckFileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !errors.Is(err, os.ErrNotExist)

}
func CreateLog(filepath string) error {
	FileExists := CheckFileExists(filepath)
	if FileExists {
		fmt.Println("The file exist")
		return nil
	}
	_, err := os.Create(filepath)
	if err != nil {
		return err

	}
	fmt.Println("File created successfully")
	return nil
}

func WriteLog(filename string, operation string, key string, value string) error {
	// Write storage operations in log to rebuild the storage in case of a fail
	// Each operation does not erase the original file content
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		return err
	}
	defer file.Close()
	entry := LogEntry{operation, key, value}
	// converting the data structure (such as a struct or a slice) into a JSON string using the json
	data, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("Error marshalling log entry: %v\n", err)
	}
	_, err = file.Write(append(data, '\n'))
	if err != nil {
		return fmt.Errorf("Error writing to file %s: %v\n", filename, err)
	} else {
		fmt.Printf("Successfully wrote to file %s: %s\n", filename, string(data))
		return nil
	}
}
func NewKeyValueStorage() *KeyValue {
	return &KeyValue{data: make(map[string]string)}
}

func (k *KeyValue) Set(key string, value string) error {
	startTime := time.Now()
	WriteLog(LogFilePath, "SET", key, value)
	k.mu.Lock()
	k.data[key] = value
	k.mu.Unlock()
	kvstoreWritesTotal.Inc()
	latency := time.Since(startTime).Seconds()
	kvstoreWritesLatencySeconds.Observe(latency)
	return nil
}

func (k *KeyValue) Get(key string) (string, error) {
	startTime := time.Now()
	k.mu.RLock()
	value, exists := k.data[key]
	defer k.mu.RUnlock()
	if !exists {
		WriteLog(LogFilePath, "GET", key, "")
		return "", fmt.Errorf("key %s not found", key)
	}
	kvstoreReadsTotal.Inc()
	WriteLog(LogFilePath, "GET", key, value)

	latency := time.Since(startTime).Seconds()
	kvstoreReadsLatencySeconds.Observe(latency)
	return value, nil
}

func (k *KeyValue) Delete(key string) error {
	k.mu.Lock()
	defer k.mu.Unlock()
	_, exists := k.data[key]
	if !exists {
		kvstoreErrorsTotal.Inc()
		WriteLog(LogFilePath, "DELETE", key, "")
		return fmt.Errorf("Key %s not found", key)
	}
	delete(k.data, key)
	kvstoreWritesTotal.Inc()
	WriteLog(LogFilePath, "DELETE", key, "")
	return nil
}

func (k *KeyValue) Exists(key string) (bool, error) {
	k.mu.RLock()
	defer k.mu.RUnlock()
	_, exists := k.data[key]
	WriteLog(LogFilePath, "EXISTS", key, "")
	return exists, nil
}

func (k *KeyValue) ReBuildStore(filename string) {
	fmt.Println("The rebuild method has been called")
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File read successfully")
	scanner := bufio.NewScanner(strings.NewReader(string(file)))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Read line:", line)
		var Entry LogEntry
		err := json.Unmarshal([]byte(line), &Entry)
		if err != nil {
			fmt.Println("Error unmarshalling line:", err, line)
			continue
		}
		fmt.Printf("Parsed entry: %v\n", Entry)

		if Entry.Operation != "GET" {
			fmt.Println("Non-GET operation detected")
		} else {
			fmt.Println("Skipping the GET operation")
			continue
		}

		switch Entry.Operation {
		case "SET":
			k.data[Entry.Key] = Entry.Value
			fmt.Println("Set operation: key =", Entry.Key, "value =", Entry.Value)
		case "DELETE":
			delete(k.data, Entry.Key)
			fmt.Println("Delete operation: key =", Entry.Key)
		}
		fmt.Println("About to write to the log")
		err = WriteLog(Rebuiltlog, Entry.Operation, Entry.Key, Entry.Value)
		if err != nil {
			fmt.Println("Error writing log:", err)
		} else {
			fmt.Println("Log written successfully")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
	fmt.Println("ReBuildStore method completed")
}

func (k *KeyValue) Replication() {
	// create a new instance of the  replica kv store
	replica := NewKeyValueStorage()
	fmt.Println(replica)
	if replica == nil {
		fmt.Println("Failed to create a replica")
	} else {
		fmt.Println("Replica created successfully.")
	}
	replica.ReBuildStore(LogFilePath)
	fmt.Println("State of the replica: ", replica)

	if k.CompareReplica(replica) == false {
		fmt.Println("Replica is not a copy of the original store")
	}
	fmt.Println("Replica matches the original store")

}

func (k *KeyValue) CompareReplica(Replica *KeyValue) bool {
	k.mu.RLock()
	defer k.mu.RUnlock()
	Replica.mu.RLock()
	defer Replica.mu.RUnlock()

	b := fmt.Sprintf("%v", k.data) == fmt.Sprintf("%v", Replica.data)
	return b
}

func DelayAdd(duration time.Duration) {
	start := time.Now()
	for time.Since(start) < duration {

	}
}

func TrackFileChanges(filename string) {
	// Create a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Failed to create watcher: %v", err)
	}
	defer watcher.Close()

	// Add file to watcher
	err = watcher.Add(filename)
	if err != nil {
		log.Fatalf("Failed to add file to watcher: %v", err)
	}
	fmt.Println("Tracking file changes on:", filename)

	// Initialize key-value store
	kv := NewKeyValueStorage()
	var wg sync.WaitGroup
	var lastOffset int64

	for {
		select {
		// Handle file events
		case ev := <-watcher.Events:
			log.Println("Event:", ev)
			if ev.Op&fsnotify.Write == fsnotify.Write {
				log.Println("File modified")
			}
			file, err := os.Open(filename)
			if err != nil {
				log.Printf("Error opening file %s: %v", filename, err)
				continue
			}
			defer file.Close()

			_, err = file.Seek(lastOffset, io.SeekStart)
			if err != nil {
				log.Printf("Error seeking to end of file %s: %v", filename, err)
				continue
			}
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				fmt.Println("New log entry:", line)
			}
			if err := scanner.Err(); err != nil {
				log.Println("Error reading new content from file %s: %v ", filename, err)
			}

			lastOffset, err := file.Seek(0, io.SeekEnd)
			if err != nil {
				log.Printf("Failed to update the last offset in file %s: %v", filename, err)
			}
			fmt.Println("Updated last offset:", lastOffset)

			if ev.Op&fsnotify.Rename == fsnotify.Rename || ev.Op&fsnotify.Remove == fsnotify.Remove {
				wg.Add(1)
				go func() {
					defer wg.Done()
					for {
						time.Sleep(1 * time.Second)
						err = watcher.Add(filename)
						if err != nil {
							log.Printf("Successfully re-added file %s to watcher", filename)
							break
						}
						log.Printf("Retrying to add file %s to watcher: %v", filename, err)
					}

				}()

			}
			log.Println("File modified, replication triggered")
			kv.ReBuildStore(filename)

		// Handle watcher errors
		case err := <-watcher.Errors:
			log.Println("Error:", err)
		}
	}
	wg.Wait()
}
