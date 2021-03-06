package parser

import (
	"errors"
	"strings"
)

// CommandType is CommandType
type CommandType int

const (
	CArithmetic CommandType = iota
	CPush
	CPop
	CLabel
	CGoto
	CIf
	CFunction
	CReturn
	CCall
)

func (command CommandType) String() string {
	switch command {
	case CArithmetic:
		return "C_ARITHMETIC"
	case CPush:
		return "C_PUSH"
	case CPop:
		return "C_POP"
	case CLabel:
		return "C_LABEL"
	case CGoto:
		return "C_GOTO"
	case CIf:
		return "C_IF"
	case CFunction:
		return "C_FUNCTION"
	case CReturn:
		return "C_RETURN"
	case CCall:
		return "C_CALL"
	default:
		return "Unknown"
	}
}

// GetCommandType GetCommandType
func GetCommandType(commandStr string) (c CommandType, err error) {
	s := strings.TrimSpace(commandStr)
	arithmeticCommand := map[string]int{"add": 1, "sub": 1, "neg": 1, "eq": 1, "gt": 1, "lt": 1, "and": 1, "or": 1, "not": 1}
	if _, ok := arithmeticCommand[s]; ok {
		return CArithmetic, nil
	}
	if strings.HasPrefix(s, "push") {
		return CPush, nil
	}
	if strings.HasPrefix(s, "pop") {
		return CPop, nil
	}
	if strings.HasPrefix(s, "label") {
		return CLabel, nil
	}
	if strings.HasPrefix(s, "goto") {
		return CGoto, nil
	}
	if strings.HasPrefix(s, "if-goto") {
		return CGoto, nil
	}
	if strings.HasPrefix(s, "function") {
		return CFunction, nil
	}
	if strings.HasPrefix(s, "call") {
		return CCall, nil
	}
	if s == "return" {
		return CReturn, nil
	}
	return 100, errors.New("Invalid CommandType")
}

// GetArg1 GetArg1
func GetArg1(commandStr string) (arg1 string, err error) {
	s := strings.TrimSpace(commandStr)
	commandType, _ := GetCommandType(s)
	if commandType == CArithmetic {
		return s, nil
	}
	if commandType == CPush || commandType == CPush || commandType == CPop || commandType == CLabel || commandType == CGoto || commandType == CIf || commandType == CFunction || commandType == CReturn || commandType == CCall {
		return strings.Split(s, " ")[1], nil
	}
	return "", errors.New("Command has no symbol")
}

// GetArg2 a
func GetArg2(commandStr string) (arg2 string, err error) {
	s := strings.TrimSpace(commandStr)
	commandType, _ := GetCommandType(s)
	if commandType == CPush || commandType == CPush || commandType == CPop || commandType == CLabel || commandType == CGoto || commandType == CIf || commandType == CFunction || commandType == CReturn || commandType == CCall {
		return strings.Split(s, " ")[2], nil
	}
	return "", errors.New("Command has no symbol")
}
