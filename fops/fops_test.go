package fops_test

import (
	. "github.com/kiedveian/JDExam/fops"
	"testing"
)

func TestLineConut(t *testing.T) {
	testStr := []string{"-f", "testdata/myfile.txt"}
	result := CmdLineCount(testStr)
	if result != "4" {
		t.Errorf("input: %s , result: %s , expected: %s ", testStr, result, "4")
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
	result := CmdCheckSum(testStr)
	if result != ans {
		t.Errorf("input: %s \n result: \n %s \n expected: \n %s ", testStr, result, ans)
	}
}
