package fops

import (
	"os"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getFileContentType(file *os.File) (string, *FopsError) {
	buffer := make([]byte, 512)
	pos, err := file.Seek(0, 1)
	if err != nil {
		return "", CreateStdErr(err)
	}
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		return "", CreateStdErr(err)
	}
	_, err = file.Seek(pos, 0)
	if err != nil {
		return "", CreateStdErr(err)
	}
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}

func CheckOpenFile(filename string, skipError map[ErrorType]bool) (*os.File, *FopsError) {
	if skipError == nil {
		skipError = map[ErrorType]bool{}
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, CreateStdErr(err)
	}
	info, err := file.Stat()
	if err != nil {
		return nil, CreateStdErr(err)
	}
	if !skipError[ErrIsDir] && info.IsDir() {
		defer file.Close()
		errStr := fmt.Sprintf(fileIsDirErrTamplate, filename)
		return nil, CreateFopsErr(ErrIsDir, errStr)
	}
	fileType, fopsError := getFileContentType(file)
	if fopsError != nil {
		defer file.Close()
		return nil, fopsError
	}
	if !skipError[ErrNotText] && !strings.Contains(fileType, "text") {
		errStr := fmt.Sprintf(fileTypeErrTamplate, fileType)
		return nil, CreateFopsErr(ErrNotText, errStr)
	}
	return file, nil
}
