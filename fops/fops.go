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
	"os"
)

const (
	bufferSize = 32 * 1024
)

const (
	versionString = "v0.0.1"
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
  -h, --help   help for fops`
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

func Run(args []string) {
	if len(args) >= 1 {
		remain := args[1:]
		switch cmd := args[0]; cmd {
		case "-h", "help":
			CmdHelp(remain)
		case "linecount":
			count, err := CmdLineCount(remain)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(count)
			}
		case "checksum":
			fmt.Println(CmdCheckSum(remain))
		case "version":
			fmt.Println("fops " + versionString)
		default:
			fmt.Println("error: undefined command ", cmd)
		}
	} else {
		fmt.Println("error: no find command ")
	}
}

func CmdHelp(args []string) {
	command := "help"
	if len(args) >= 1 {
		command = args[0]
	}
	switch command {
	case "help":
		fmt.Println(helpString)
	case "linecount":
		fmt.Println(linecountString)
	case "checksum":
		fmt.Println(checksumString)
	default:
		fmt.Println(helpString)
	}
}

func CmdLineCount(args []string) (int, error) {
	if len(args) < 2 {
		return 0, errors.New("args not enough")
	}
	switch args[0] {
	case "-f", "--file":
		file, err := os.Open(args[1])
		if err != nil {
			return 0, err
		}
		defer file.Close()
		count, err := linecountBySep(file)
		if err != nil {
			return 0, err
		}
		return count, nil
	default:
		return 0, errors.New("undefined flag " + args[0])
	}
}

func CmdCheckSum(args []string) (string, error) {
	if len(args) < 3 {
		return "", errors.New("args not enough")
	}
	switch args[0] {
	case "-f", "--file":
		file, err := os.Open(args[1])
		if err != nil {
			return "", err
		}
		defer file.Close()
		byteArr, err := checksum(file, args[2])
		if err != nil {
			return "", err
		}
		return hex.EncodeToString(byteArr), nil
	default:
		return "", errors.New("undefined flag " + args[0])
	}
}

func linecountBySep(file io.Reader) (int, error) {
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
			return result, err
		}
	}
}

func checksum(file io.Reader, flag string) ([]byte, error) {
	var hashObj hash.Hash
	switch flag {
	case "--md5":
		hashObj = md5.New()
	case "--sha1":
		hashObj = sha1.New()
	case "--sha256":
		hashObj = sha256.New()
	}
	if _, err := io.Copy(hashObj, file); err != nil {
		return nil, err
	}
	return hashObj.Sum(nil), nil
}

func main() {
	Run(os.Args[1:])
}
