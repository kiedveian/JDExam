package main_test

import (
	. "github.com/kiedveian/JDExam/fops"
	"testing"
)

func TestLineConut(t *testing.T) {
	testStr := []string{"-f", "testdata/myfile.txt"}
	result, err := CmdLineCount(testStr)
	if err != nil {
		t.Errorf(err.Error())
	} else if result != 4 {
		t.Errorf("input: %s , result: %d , expected: %d ", testStr, result, 4)
	}
}

func TestCheckSum(t *testing.T) {
	md5Ans := "021ca43b2982439db3ab764527abcd28"
	sha1Ans := "fda7c974cd9425e8ec3841125c7b8af9f5577c28"
	sha256Ans := "0e401f46bc43161f9e5a7a2ce9e15a2511d972320dadbef20fa2fec6610f04d0"
	testCheckSumImp(t, "--md5", md5Ans)
	testCheckSumImp(t, "--sha1", sha1Ans)
	testCheckSumImp(t, "--sha256", sha256Ans)
}

func testCheckSumImp(t *testing.T, flag, ans string) {
	testStr := []string{"-f", "testdata/myfile.txt", flag}
	result, err := CmdCheckSum(testStr)
	if err != nil {
		t.Errorf(err.Error())
	} else if result != ans {
		t.Errorf("input: %s \n result: \n %s \n expected: \n %s ", testStr, result, ans)
	}
}

func TestError(t *testing.T) {
	filename := "non-exist-file.ttt"
	if _, err := CmdLineCount([]string{"-f", filename}); err == nil {
		t.Errorf("cmd: linecout -f %s, result: <nil>, expected a error ", filename)
	}
	if _, err := CmdCheckSum([]string{"-f", filename, "--md5"}); err == nil {
		t.Errorf("cmd: checksum -f %s --md5, result: <nil>, expected a error ", filename)
	}
}
