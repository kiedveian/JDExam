package fops

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
)

const (
	bufferSize = 32 * 1024
)

func Run(args []string) {
	if len(args) >= 1 {
		remain := args[1:]
		switch cmd := args[0]; cmd {
		case "help":
			CmdHelp(remain)
		case "linecount":
			fmt.Println(CmdLineCount(remain))
		case "checksum":
			fmt.Println(CmdCheckSum(remain))
		default:
			fmt.Println("undefined command ", cmd)
		}
	}
}

func CmdHelp(args []string) {
	fmt.Println("test help command")
	fmt.Println("args: ", args)
}

func CmdLineCount(args []string) string {
	switch args[0] {
	case "-f", "--file":
		file, err := os.Open(args[1])
		if err != nil {
			return fmt.Sprint(err)
		}
		count, err := linecount(file)
		if err != nil {
			return fmt.Sprint(err)
		}
		return fmt.Sprint(count)
	default:
		return "undefined error"
	}
}

func CmdCheckSum(args []string) string {
	switch args[0] {
	case "-f", "--file":
		file, err := os.Open(args[1])
		if err != nil {
			return fmt.Sprint(err)
		}
		defer file.Close()
		byteArr, err := checksum(file, args[2])
		if err != nil {
			return fmt.Sprint(err)
		}
		return hex.EncodeToString(byteArr)
	default:
		return "undefined error"
	}
}

func linecount(flie io.Reader) (int, error) {
	buf := make([]byte, bufferSize)
	result := 0
	lineSep := []byte{'\n'}

	for {
		count, err := flie.Read(buf)
		result += bytes.Count(buf[:count], lineSep)
		switch {
		case err == io.EOF:
			return result, nil
		case err != nil:
			return result, err
		}
	}
}

func checksum(flie io.Reader, flag string) ([]byte, error) {
	var hashObj hash.Hash
	switch flag {
	case "--md5":
		hashObj = md5.New()
	case "--sha1":
		hashObj = sha1.New()
	case "--sha256":
		hashObj = sha256.New()
	}
	if _, err := io.Copy(hashObj, flie); err != nil {
		return nil, err
	}
	return hashObj.Sum(nil), nil
}
