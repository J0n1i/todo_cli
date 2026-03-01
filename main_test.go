package main

import (
	"os"
	"testing"
)

func TestWriteAndReadFromFile(t *testing.T) {
	tmpFile := "test_data.txt"
	defer os.Remove(tmpFile)

	// Write tasks
	err := writeToFile(tmpFile, "task1")
	if err != nil {
		t.Fatalf("writeToFile failed: %v", err)
	}
	err = writeToFile(tmpFile, "task2")
	if err != nil {
		t.Fatalf("writeToFile failed: %v", err)
	}

	// Read tasks
	tasks, err := readFromFile(tmpFile)
	if err != nil {
		t.Fatalf("readFromFile failed: %v", err)
	}
	if len(tasks) != 2 || tasks[0] != "task1" || tasks[1] != "task2" {
		t.Errorf("unexpected tasks: %v", tasks)
	}
}

func TestDeleteAllFromFile(t *testing.T) {
	tmpFile := "test_data.txt"
	defer os.Remove(tmpFile)

	_ = writeToFile(tmpFile, "task1")
	_ = writeToFile(tmpFile, "task2")

	err := deleteAllFromFile(tmpFile)
	if err != nil {
		t.Fatalf("deleteAllFromFile failed: %v", err)
	}

	tasks, err := readFromFile(tmpFile)
	if err != nil {
		t.Fatalf("readFromFile failed: %v", err)
	}
	if len(tasks) != 0 {
		t.Errorf("expected 0 tasks, got %d", len(tasks))
	}
}

func TestDeleteFromFile(t *testing.T) {
	tmpFile := "test_data.txt"
	defer os.Remove(tmpFile)

	_ = writeToFile(tmpFile, "task1")
	_ = writeToFile(tmpFile, "task2")
	_ = writeToFile(tmpFile, "task3")

	err := deleteFromFile(tmpFile, 1)
	if err != nil {
		t.Fatalf("deleteFromFile failed: %v", err)
	}

	tasks, err := readFromFile(tmpFile)
	if err != nil {
		t.Fatalf("readFromFile failed: %v", err)
	}
	if len(tasks) != 2 || tasks[0] != "task1" || tasks[1] != "task3" {
		t.Errorf("unexpected tasks after delete: %v", tasks)
	}
}

func TestRemoveAtIndex(t *testing.T) {
	data := []string{"a", "b", "c"}
	newData, err := removeAtIndex(data, 1)
	if err != nil {
		t.Fatalf("removeAtIndex failed: %v", err)
	}
	if len(newData) != 2 || newData[0] != "a" || newData[1] != "c" {
		t.Errorf("unexpected result: %v", newData)
	}

	_, err = removeAtIndex(data, -1)
	if err == nil {
		t.Error("expected error for negative index")
	}
	_, err = removeAtIndex(data, 3)
	if err == nil {
		t.Error("expected error for out-of-range index")
	}
}

func TestReadFromFile_EmptyFile(t *testing.T) {
	tmpFile := "test_empty.txt"
	defer os.Remove(tmpFile)
	os.WriteFile(tmpFile, []byte(""), 0644)

	tasks, err := readFromFile(tmpFile)
	if err != nil {
		t.Fatalf("readFromFile failed: %v", err)
	}
	if len(tasks) != 0 {
		t.Errorf("expected 0 tasks, got %d", len(tasks))
	}
}
