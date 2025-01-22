package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
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
	// Write storage operations in log to rebuild the storage in case of fail
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
	WriteLog(LogFilePath, "SET", key, value)
	k.mu.Lock()
	k.data[key] = value
	k.mu.Unlock()
	return nil
}

func (k *KeyValue) Get(key string) (string, error) {
	k.mu.RLock()
	value, exists := k.data[key]
	defer k.mu.RUnlock()
	if !exists {
		WriteLog(LogFilePath, "GET", key, "")
		return "", fmt.Errorf("key %s not found", key)
	}
	WriteLog(LogFilePath, "GET", key, value)

	return value, nil
}

func (k *KeyValue) Delete(key string) error {
	k.mu.Lock()
	defer k.mu.Unlock()
	_, exists := k.data[key]
	if !exists {
		WriteLog(LogFilePath, "DELETE", key, "")
		return fmt.Errorf("Key %s not found", key)
	}
	delete(k.data, key)

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

}
