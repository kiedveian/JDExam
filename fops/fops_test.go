package fops_test


import (
    . "github.com/kiedveian/JDExam/fops"
    "testing"
)


func TestLineConut(t *testing.T) {
    testStr  := []string{"-f", "testdata/myfile.txt"}
    result := CmdLineCount(testStr)
    if result != "4"{
        t.Errorf("input: %s , result: %s , expected: %s ", testStr, result, "4")
    }
}