package main_test

import (
	"fmt"
	. "github.com/kiedveian/JDExam/fops"
	"os"
	"runtime"
	"testing"
)

func TestLineConut(t *testing.T) {
	testStr := []string{"-f", "testdata/utf8.txt"}
	result, fopsErr := CmdLineCount(testStr)
	if fopsErr != nil {
		t.Errorf(fopsErr.Err.Error())
	} else if result != 4 {
		t.Errorf("input: %s , result: %d , expected: %d ", testStr, result, 4)
	}
}

func TestCheckSum(t *testing.T) {
	switch runtime.GOOS {
	case "windows":
		testWindowsCheckSum(t)
	default:
		testLinuxCheckSum(t)
	}
}

func testWindowsCheckSum(t *testing.T) {
	md5Ans := "021ca43b2982439db3ab764527abcd28"
	sha1Ans := "fda7c974cd9425e8ec3841125c7b8af9f5577c28"
	sha256Ans := "0e401f46bc43161f9e5a7a2ce9e15a2511d972320dadbef20fa2fec6610f04d0"
	testCheckSumImp(t, "--md5", md5Ans)
	testCheckSumImp(t, "--sha1", sha1Ans)
	testCheckSumImp(t, "--sha256", sha256Ans)
}

func testLinuxCheckSum(t *testing.T) {
	md5Ans := "a8c5d553ed101646036a811772ffbdd8"
	sha1Ans := "a656582ca3143a5f48718f4a15e7df018d286521"
	sha256Ans := "495a3496cfd90e68a53b5e3ff4f9833b431fe996298f5a28228240ee2a25c09d"
	testCheckSumImp(t, "--md5", md5Ans)
	testCheckSumImp(t, "--sha1", sha1Ans)
	testCheckSumImp(t, "--sha256", sha256Ans)
}

func testCheckSumImp(t *testing.T, flag, ans string) {
	testStr := []string{"-f", "testdata/myfile.txt", flag}
	result, fopsErr := CmdCheckSum(testStr)
	if fopsErr != nil {
		t.Errorf(fopsErr.Err.Error())
	} else if result != ans {
		t.Errorf("input: %s \n result: \n %s \n expected: \n %s ", testStr, result, ans)
	}
}

func TestFileError(t *testing.T) {
	nonExistFilename := "non-exist-file.ttt"
	directoryFilename := "testdata/"

	_, fopsErr := CmdLineCount([]string{"-f", nonExistFilename})
	checkNotExistErr(t, fopsErr, fmt.Sprintf("linecout -f %s", nonExistFilename))

	_, fopsErr = CmdCheckSum([]string{"-f", nonExistFilename, "--md5"})
	checkNotExistErr(t, fopsErr, fmt.Sprintf("checksum -f %s --md5", nonExistFilename))

	commandString := fmt.Sprintf("linecout -f %s", directoryFilename)
	_, fopsErr = CmdLineCount([]string{"-f", directoryFilename})
	if fopsErr == nil {
		t.Errorf("cmd: %s, result: <nil>, expected a error ", commandString)
	} else if fopsErr.TypeId != ErrIsDir {
		t.Errorf("cmd: %s, result: %s, expected error: Expected file got directory", commandString, fopsErr.Err)
	}
}

func checkNotExistErr(t *testing.T, fopsErr *FopsError, command string) {
	if fopsErr == nil {
		t.Errorf("cmd: %s, result: <nil>, expected a error ", command)
	} else if !os.IsNotExist(fopsErr.Err) {
		t.Errorf("cmd: %s, result: %s, expected error: the file is not exist ", command, fopsErr.Err)
	}
}
