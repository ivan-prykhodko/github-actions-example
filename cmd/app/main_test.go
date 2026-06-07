package main

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestLyrics(t *testing.T) {
	expected := []string{
		"Hello",
		"Is it me you're looking for?",
	}
	actual := lyrics()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestMainOutput(t *testing.T) {
	// Keep track of the original stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run main
	main()

	// Restore stdout and close the pipe
	err := w.Close()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = old

	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		t.Fatal(err)
	}
	output := buf.String()

	expected := "Hello\nIs it me you're looking for?\n"
	if output != expected {
		t.Errorf("expected %q, got %q", expected, output)
	}
}
