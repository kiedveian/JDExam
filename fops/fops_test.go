package main_test

import (
	. "github.com/kiedveian/JDExam/fops"
	"runtime"
	"testing"
)

type args struct {
	args []string
}

type checkSumTestCase struct {
	name    string
	args    args
	want    string
	wantErr *FopsError
}

func compareError(lhs, rhs *FopsError) bool {
	if lhs == nil && rhs == nil {
		return true
	} else if lhs == nil {
		return false
	} else if rhs == nil {
		return false
	}
	return lhs.TypeId == rhs.TypeId
}

func TestRunLineConut(t *testing.T) {
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr *FopsError
	}{
		{
			"linecout utf8 file",
			args{[]string{"-f", "testdata/utf8.txt"}},
			4,
			nil,
		},
		{
			"no find file error",
			args{[]string{"-f", "non-exist-file.ttt"}},
			0,
			&FopsError{ErrStd, nil},
		},
		{
			"file is directory error",
			args{[]string{"-f", "testdata/"}},
			0,
			&FopsError{ErrIsDir, nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RunLineCount(tt.args.args)
			if got != tt.want {
				t.Errorf("RunLineCount() got = %v, want %v", got, tt.want)
			}
			if !compareError(err, tt.wantErr) {
				t.Errorf("RunLineCount() err = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRunCheckSum(t *testing.T) {
	var tests []checkSumTestCase
	switch runtime.GOOS {
	case "windows":
		tests = getWindowsCheckSumTests()
	default:
		tests = getLinuxCheckSumTests()
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RunCheckSum(tt.args.args)
			if got != tt.want {
				t.Errorf("RunCheckSum() got = %v, want %v", got, tt.want)
			}
			if !compareError(err, tt.wantErr) {
				t.Errorf("RunCheckSum() err = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func getWindowsCheckSumTests() []checkSumTestCase {
	tests := []checkSumTestCase{
		{
			"Windows md5",
			args{[]string{"-f", "testdata/myfile.txt", "--md5"}},
			"021ca43b2982439db3ab764527abcd28",
			nil,
		},
		{
			"Windows sha1",
			args{[]string{"-f", "testdata/myfile.txt", "--sha1"}},
			"fda7c974cd9425e8ec3841125c7b8af9f5577c28",
			nil,
		},
		{
			"Windows sha256",
			args{[]string{"-f", "testdata/myfile.txt", "--sha256"}},
			"0e401f46bc43161f9e5a7a2ce9e15a2511d972320dadbef20fa2fec6610f04d0",
			nil,
		},
		{
			"no find file error",
			args{[]string{"-f", "non-exist-file.ttt", "--md5"}},
			"",
			&FopsError{ErrStd, nil},
		},
		{
			"file is directory error",
			args{[]string{"-f", "testdata/", "--md5"}},
			"",
			&FopsError{ErrIsDir, nil},
		},
	}
	return tests
}

func getLinuxCheckSumTests() []checkSumTestCase {
	tests := []checkSumTestCase{
		{
			"Linux md5",
			args{[]string{"-f", "testdata/myfile.txt", "--md5"}},
			"a8c5d553ed101646036a811772ffbdd8",
			nil,
		},
		{
			"Linux sha1",
			args{[]string{"-f", "testdata/myfile.txt", "--sha1"}},
			"a656582ca3143a5f48718f4a15e7df018d286521",
			nil,
		},
		{
			"Linux sha256",
			args{[]string{"-f", "testdata/myfile.txt", "--sha256"}},
			"495a3496cfd90e68a53b5e3ff4f9833b431fe996298f5a28228240ee2a25c09d",
			nil,
		},
		{
			"no find file error",
			args{[]string{"-f", "non-exist-file.ttt" , "--md5"}},
			"",
			&FopsError{ErrStd, nil},
		},
		{
			"file is directory error",
			args{[]string{"-f", "testdata/", "--md5"}},
			"",
			&FopsError{ErrIsDir, nil},
		},
	}
	return tests
}
