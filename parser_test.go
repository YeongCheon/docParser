package main

import (
	"os"
	"testing"
)

func TestGetDocumentContent(t *testing.T) {
	docContent := GetDocumentContent("sample_test.docx")
	if docContent == "" {
		t.Error("content is null")
	}
}

func TestSaveTextFile(t *testing.T) {
	SaveTextFile("result_test.txt", "result_test content")
	if _, err := os.Stat("result_test.txt"); os.IsNotExist(err) {
		t.Error("file is not found")
	}
}
