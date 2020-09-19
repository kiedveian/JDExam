package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"io"
	"net/http"
	"os"
	"strings"
)

type ErrorType uint

type FopsError struct {
	TypeId ErrorType
	Err    error
}

const (
	bufferSize = 32 * 1024
)

const (
	ErrUndefined ErrorType = iota
	ErrStd
	ErrArgsNotEnough
	ErrUndefinedFlag
	ErrIsDir
	ErrNotText
)

const (
	FlagHelpShort = "-h"
	FlagHelpLong  = "--help"
	FlagFileShort = "-f"
	FlagFileLong  = "--file"
	FlagVersion   = "-v"
	FlagMd5       = "--md5"
	FlagSha1      = "--sha1"
	FlagSha256    = "--sha256"
)

const (
	CmdHelp      = "help"
	CmdLineCount = "linecount"
	CmdVersion   = "version"
	CmdCheckSum  = "checksum"
)

const (
	helpString = `File Ops
Usage:
  fops [flags]
  fops [command]
Available Commands:
  linecount    Print line count of file
  checksum     Print checksum of file
  version      Show the version info
  help         Help about commands
Flags:
  -h, --help   help for fops
  -v           Show the version tag`
)

const (
	linecountString = `Print line count of file
Usage:
  fops linecount [flags]
Flags:
  -f, --file   the input file`
)

const (
	checksumString = `Print checksum of file
Usage:
  fops checksum [flags]
Flags:
  -f, --file   the input file 
  --md5
  --sha1
  --sha256`
)

const versionStringTemplate = "fops %s"

const (
	// error template strings
	undefinedCmdErrTemplate  = "error: undefined command '%s'"
	noFindCmdErrTemplate     = "error: no find command "
	argNotEnoughTemplate     = "args not enough"
	undefinedFlagErrTamplate = "undefined flag '%s'"
	fileIsDirErrTamplate     = "Expected file got directory '%s'"
	fileTypeErrTamplate      = "Cannot do linecount (detect content type: %s)"
)

var Version = "No Version Provided"

func CreateStdErr(err error) *FopsError {
	result := new(FopsError)
	result.TypeId = ErrStd
	result.Err = err
	return result
}

func CreateFopsErr(typeId ErrorType, message string) *FopsError {
	result := new(FopsError)
	result.TypeId = typeId
	result.Err = errors.New(message)
	return result
}

func RunFops(args []string) {
	if len(args) >= 1 {
		remain := args[1:]
		switch cmd := args[0]; cmd {
		case CmdHelp, FlagHelpShort, FlagHelpLong:
			RunHelp(remain)
		case CmdLineCount:
			count, err := RunLineCount(remain)
			if err != nil {
				fmt.Println(err.Err.Error())
			} else {
				fmt.Println(count)
			}
		case CmdCheckSum:
			str, err := RunCheckSum(remain)
			if err != nil {
				fmt.Println(err.Err.Error())
			} else {
				fmt.Println(str)
			}
		case CmdVersion:
			fmt.Println(versionStringTemplate, Version)
		case FlagVersion:
			fmt.Println(Version)
		default:
			fmt.Println(undefinedCmdErrTemplate, cmd)
		}
	} else {
		fmt.Println(noFindCmdErrTemplate)
	}
}

func RunHelp(args []string) {
	command := CmdHelp
	if len(args) >= 1 {
		command = args[0]
	}
	switch command {
	case CmdHelp:
		fmt.Println(helpString)
	case CmdLineCount:
		fmt.Println(linecountString)
	case CmdCheckSum:
		fmt.Println(checksumString)
	default:
		fmt.Println(helpString)
	}
}

func RunLineCount(args []string) (int, *FopsError) {
	if len(args) < 2 {
		return 0, CreateFopsErr(ErrArgsNotEnough, argNotEnoughTemplate)
	}
	switch args[0] {
	case FlagFileShort, FlagFileLong:
		file, fopsError := CheckOpenFile(args[1], nil)
		if fopsError != nil {
			return 0, fopsError
		}
		defer file.Close()
		count, fopsError := ImpLineCount(file)
		if fopsError != nil {
			return 0, fopsError
		}
		return count, nil
	default:
		errStr := fmt.Sprint(undefinedFlagErrTamplate, args[0])
		return 0, CreateFopsErr(ErrUndefinedFlag, errStr)
	}
}

func RunCheckSum(args []string) (string, *FopsError) {
	if len(args) < 3 {
		return "", CreateFopsErr(ErrArgsNotEnough, argNotEnoughTemplate)
	}
	switch args[0] {
	case FlagFileShort, FlagFileLong:
		file, fopsError := CheckOpenFile(args[1], map[ErrorType]bool{ErrNotText: true})
		if fopsError != nil {
			return "", fopsError
		}
		defer file.Close()
		byteArr, fopsError := ImpCheckSum(file, args[2])
		if fopsError != nil {
			return "", fopsError
		}
		return hex.EncodeToString(byteArr), nil
	default:
		errStr := fmt.Sprint(undefinedFlagErrTamplate, args[0])
		return "", CreateFopsErr(ErrUndefinedFlag, errStr)
	}
}

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
		errStr := fmt.Sprint(fileIsDirErrTamplate, filename)
		return nil, CreateFopsErr(ErrIsDir, errStr)
	}
	fileType, fopsError := getFileContentType(file)
	if fopsError != nil {
		defer file.Close()
		return nil, fopsError
	}
	if !skipError[ErrNotText] && !strings.Contains(fileType, "text") {
		errStr := fmt.Sprint(fileTypeErrTamplate, fileType)
		return nil, CreateFopsErr(ErrNotText, errStr)
	}
	return file, nil
}

func ImpLineCount(file io.Reader) (int, *FopsError) {
	buf := make([]byte, bufferSize)
	result := 0
	lineSep := []byte{'\n'}

	for {
		count, err := file.Read(buf)
		result += bytes.Count(buf[:count], lineSep)
		switch {
		case err == io.EOF:
			return result, nil
		case err != nil:
			return result, CreateStdErr(err)
		}
	}
}

func ImpCheckSum(file io.Reader, flag string) ([]byte, *FopsError) {
	var hashObj hash.Hash
	switch flag {
	case FlagMd5:
		hashObj = md5.New()
	case FlagSha1:
		hashObj = sha1.New()
	case FlagSha256:
		hashObj = sha256.New()
	}
	if _, err := io.Copy(hashObj, file); err != nil {
		return nil, CreateStdErr(err)
	}
	return hashObj.Sum(nil), nil
}

func main() {
	RunFops(os.Args[1:])
}
