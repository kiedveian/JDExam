package main_test

import (
	"encoding/hex"
	"fmt"
	. "github.com/kiedveian/JDExam/fops"
	"os"
	"runtime"
	"strings"
	"testing"
)

const (
	testCaseString = ` one
two
three
four
five 
`
	testCaseLineCountAnswer = 5
)

func TestLineConut(t *testing.T) {
	testStr := []string{"-f", "testdata/utf8.txt"}
	result, fopsErr := CmdLineCount(testStr)
	if fopsErr != nil {
		t.Errorf(fopsErr.Err.Error())
	} else if result != 4 {
		t.Errorf("input: %s , result: %d , expected: %d ", testStr, result, 4)
	}
	reader := strings.NewReader(testCaseString)
	result, fopsErr = ImpLineCount(reader)
	if fopsErr != nil {
		t.Errorf(fopsErr.Err.Error())
	} else if result != testCaseLineCountAnswer {
		t.Errorf("input: %s , result: %d , expected: %d ", testStr, result, testCaseLineCountAnswer)
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
	md5MyfileAns := "021ca43b2982439db3ab764527abcd28"
	sha1MyFileAns := "fda7c974cd9425e8ec3841125c7b8af9f5577c28"
	sha256MyFileAns := "0e401f46bc43161f9e5a7a2ce9e15a2511d972320dadbef20fa2fec6610f04d0"
	testCmdCheckSum(t, "--md5", md5MyfileAns)
	testCmdCheckSum(t, "--sha1", sha1MyFileAns)
	testCmdCheckSum(t, "--sha256", sha256MyFileAns)

	md5StringAns := "ec0f72e148fa0845cf63bbe75207fc46"
	sha1StringAns := "6fc58d78ef4495ff9caf8c3ef91caf7119f655b2"
	sha256StringAns := "4ebd184271058535645c8e5c962ba63c89b734eab6cceb81dabc9a0127dbda37"
	testStringCheckSum(t, "--md5", testCaseString, md5StringAns)
	testStringCheckSum(t, "--sha1", testCaseString, sha1StringAns)
	testStringCheckSum(t, "--sha256", testCaseString, sha256StringAns)
}

func testLinuxCheckSum(t *testing.T) {
	md5Ans := "a8c5d553ed101646036a811772ffbdd8"
	sha1Ans := "a656582ca3143a5f48718f4a15e7df018d286521"
	sha256Ans := "495a3496cfd90e68a53b5e3ff4f9833b431fe996298f5a28228240ee2a25c09d"
	testCmdCheckSum(t, "--md5", md5Ans)
	testCmdCheckSum(t, "--sha1", sha1Ans)
	testCmdCheckSum(t, "--sha256", sha256Ans)

	md5StringAns := "ec0f72e148fa0845cf63bbe75207fc46"
	sha1StringAns := "6fc58d78ef4495ff9caf8c3ef91caf7119f655b2"
	sha256StringAns := "4ebd184271058535645c8e5c962ba63c89b734eab6cceb81dabc9a0127dbda37"
	testStringCheckSum(t, "--md5", testCaseString, md5StringAns)
	testStringCheckSum(t, "--sha1", testCaseString, sha1StringAns)
	testStringCheckSum(t, "--sha256", testCaseString, sha256StringAns)
}

func testCmdCheckSum(t *testing.T, flag, ans string) {
	testStr := []string{"-f", "testdata/myfile.txt", flag}
	result, fopsErr := CmdCheckSum(testStr)
	if fopsErr != nil {
		t.Errorf(fopsErr.Err.Error())
	} else if result != ans {
		t.Errorf("input: %s \n result: \n %s \n expected: \n %s ", testStr, result, ans)
	}
}

func testStringCheckSum(t *testing.T, flag, input, ans string) {
	reader := strings.NewReader(input)
	byteArr, fopsErr := ImpCheckSum(reader, flag)
	if fopsErr != nil {
		t.Errorf(fopsErr.Err.Error())
	} else if result := hex.EncodeToString(byteArr); result != ans {
		t.Errorf("input: %s \n result: \n %s \n expected: \n %s ", input, result, ans)
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
