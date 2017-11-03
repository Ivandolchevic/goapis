package fileHandler

import (
	"net/http"

	. "github.com/Ivandolchevic/goapis/pkg/models/business"
	. "github.com/Ivandolchevic/goapis/pkg/models/util"
	fileUtil "github.com/Ivandolchevic/goapis/pkg/utils/fileUtil"
)

// Put upload a file an store it on the disk
func Put(w http.ResponseWriter, r *http.Request) *APIError {

	// Get file object from the request form
	file, header, err := r.FormFile("file")
	if err != nil {
		return &APIError{Code: http.StatusNoContent, Error: err, Message: "No file found"}
	}

	// Write the file to the disk
	err = fileUtil.MultipartFileToDisk(file, header, "E:\\WORKSPACE\\Temp")

	return nil
}
