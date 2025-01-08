package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Storage interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Delete(key string) error
	Exists(key string) (bool, error)
}

type KeyValue struct {
	data map[string]string
}

type LogEntry struct {
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
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	entry := LogEntry{operation, key, value}
	// converting the data structure (such as a struct or a slice) into a JSON string using the json
	data, err := json.Marshal(entry)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file.Write(append(data, '\n'))
	if err != nil {
		fmt.Println(err)
	}

}
func NewKeyValueStorage() *KeyValue {
	return &KeyValue{data: make(map[string]string)}
}

func (k *KeyValue) Set(key string, value string) error {
	WriteLog("logs.txt", "SET", key, value)
	k.data[key] = value
	return nil
}

func (k *KeyValue) Get(key string) (string, error) {
	value, exists := k.data[key]
	if !exists {
		return "", fmt.Errorf("key %s not found", key)
	}
	return value, nil
}

func (k *KeyValue) Delete(key string) error {
	_, exists := k.data[key]
	if !exists {
		return fmt.Errorf("Key %s not found", key)
	}
	delete(k.data, key)
	return nil
}

func (k *KeyValue) Exists(key string) (bool, error) {
	_, exists := k.data[key]
	return exists, nil
}
