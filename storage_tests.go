package main

import (
	"testing"
)

func TestKeyValue_Set(t *testing.T) {
	kv := NewKeyValueStorage()
	err := kv.Set("key1", "value1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if kv.data["key1"] != "value1" {
		t.Fatalf("expected value1, got %v", kv.data["key1"])
	}
}

func TestKeyValue_Get(t *testing.T) {
	kv := NewKeyValueStorage()
	kv.Set("key1", "value1")
	value, err := kv.Get("key1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if value != "value1" {
		t.Fatalf("expected value1, got %v", value)
	}

	_, err = kv.Get("key2")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestKeyValue_Delete(t *testing.T) {
	kv := NewKeyValueStorage()
	kv.Set("key1", "value1")
	err := kv.Delete("key1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if _, exists := kv.data["key1"]; exists {
		t.Fatalf("expected key1 to be deleted")
	}

	err = kv.Delete("key2")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestKeyValue_Exists(t *testing.T) {
	kv := NewKeyValueStorage()
	kv.Set("key1", "value1")
	exists, err := kv.Exists("key1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !exists {
		t.Fatalf("expected key1 to exist")
	}

	exists, err = kv.Exists("key2")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if exists {
		t.Fatalf("expected key2 to not exist")
	}
}
