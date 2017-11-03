package fileUtil

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"runtime"
	"strings"

	. "github.com/Ivandolchevic/goapis/pkg/models/util"
)

var pathseparator = "/"

// MultipartFileToDisk write a multipart file into the specified folder on the disk
func MultipartFileToDisk(content multipart.File, header *multipart.FileHeader, folderpath string) error {
	var buffer bytes.Buffer

	defer buffer.Reset()

	// Get the file name
	name := strings.Split(header.Filename, ".")

	// Display the file name
	fmt.Printf("File name %s\n", name[0])

	// Copy the file data to my buffer
	io.Copy(&buffer, content)

	// create the full file path
	fullpath := PathBuilder([]string{folderpath, name[0]})

	// Create a file on the disk
	file, err := os.Create(fullpath)

	defer file.Close()

	// Writes bytes into the file
	err = ioutil.WriteFile(fullpath, buffer.Bytes(), 0644)

	return err
}

func PathBuilder(paths []string) string {
	if runtime.GOOS == "windows" {
		pathseparator = "\\"
	}

	result := ""

	for i := 0; i < len(paths); i++ {
		for strings.HasSuffix(paths[i], pathseparator) {
			paths[i] = strings.TrimSuffix(paths[i], pathseparator)
		}

		if i > 0 {
			for strings.HasPrefix(paths[i], pathseparator) {
				paths[i] = strings.TrimPrefix(paths[i], pathseparator)
			}

			result = result + pathseparator + paths[i]
			// firt occurence
		} else {
			result = paths[i]
		}
	}

	return result
}

// WriteStringToDisk writes the given string into a file
func WriteStringToDisk(filepath string, data string) *APIError {
	// Open the file
	file, err := CreateOrOpen(filepath)

	if err != nil {
		return &APIError{Error: err, Message: "create a file error", Code: http.StatusInternalServerError}
	}

	defer file.Close()

	// Writes bytes into the file
	if _, err = file.WriteString("\n" + data); err != nil {
		return &APIError{Error: err, Message: "write in file error", Code: http.StatusInternalServerError}
	}

	return nil
}

// CreateOrOpen creates a file if not exists or just open it otherwise
func CreateOrOpen(filepath string) (*os.File, error) {
	var file *os.File
	var err error

	// Check if the file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		// Create a file on the disk
		file, err = os.Create(filepath)
	} else {
		file, err = os.OpenFile(filepath, os.O_APPEND, 0666)
	}

	return file, err
}

// WriteByteToDisk writes the given array of byte into a file
func WriteByteToDisk(filepath string, data []byte) *APIError {
	// Open the file
	file, err := CreateOrOpen(filepath)

	if err != nil {
		return &APIError{Error: err, Message: "create a file error", Code: http.StatusInternalServerError}
	}

	// Create the buffer writer
	defer file.Close()

	// Writes bytes into the file
	if _, err = file.Write(data); err != nil {
		return &APIError{Error: err, Message: "write in file error", Code: http.StatusInternalServerError}
	}

	return nil
}
