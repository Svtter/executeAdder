package executeadder

import (
	"os"
	"testing"
)

func TestAddToFile(t *testing.T) {
	AddToFile("hello\n", "test.txt", "new-test.txt")

	f, err := os.Open("new-test.txt")
	if err != nil {
		t.Error("Error opening file")
	}
	data := make([]byte, 10000)
	count, err := f.Read(data)
	if err != nil {
		t.Error("Error reading file")
	}

	if string(data[:count]) != "hello\nworld" {
		t.Errorf("Error writing file, content is not correct: %s", string(data[:count]))
	}
}
