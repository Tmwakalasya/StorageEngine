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

func WriteLog(filename string, operation string, key string, value string) {
	// Write storage operations in log to rebuild the storage in case of fail
	// Each operation does not erase the original file content
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		return
	}
	defer file.Close()
	entry := LogEntry{operation, key, value}
	// converting the data structure (such as a struct or a slice) into a JSON string using the json
	data, err := json.Marshal(entry)
	if err != nil {
		fmt.Printf("Error marshalling log entry: %v\n", err)
		return
	}
	_, err = file.Write(append(data, '\n'))
	if err != nil {
		fmt.Printf("Error writing to file %s: %v\n", filename, err)
	} else {
		fmt.Printf("Successfully wrote to file %s: %s\n", filename, string(data))
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
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(file)))
	// Process each json object as a whole
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		var Entry LogEntry
		err := json.Unmarshal([]byte(line), &Entry)

		if err != nil {
			fmt.Println("Error parsing line", err)
			continue
		}
		fmt.Println("Entry: ", Entry)
		if Entry.Operation != "GET" {
			switch Entry.Operation {
			case "SET":
				k.data[Entry.Key] = Entry.Value
			case "DELETE":
				delete(k.data, Entry.Key)
			}
			WriteLog(Rebuiltlog, Entry.Operation, Entry.Key, Entry.Value)
		}
	}
}

func (k *KeyValue) Replication() {

}
