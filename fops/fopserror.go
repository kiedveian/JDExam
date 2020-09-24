package fops

import (
	"errors"
)

type ErrorType uint

type FopsError struct {
	TypeId ErrorType
	Err    error
}

const (
	ErrUndefined ErrorType = iota
	ErrStd
	ErrArgsNotEnough
	ErrUndefinedFlag
	ErrIsDir
	ErrNotText
)

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

func CompareErrorType(lhs, rhs *FopsError) bool {
	if lhs == nil && rhs == nil {
		return true
	} else if lhs == nil {
		return false
	} else if rhs == nil {
		return false
	}
	return lhs.TypeId == rhs.TypeId
}
