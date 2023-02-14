package util

import "regexp"

// CheckNum 检查字符串0-9，固定长度，如果长度-1，仅检查字符串规则aA0
func CheckNum(str string, strLen int) bool {
	if len(str) != strLen && strLen != -1 {
		return false
	}
	match, _ := regexp.MatchString("^[0-9]+$", str)
	return match
}

// CheckStr 检查字符串Aa0，固定长度，如果长度-1，仅检查字符串规则aA0
func CheckStr(str string, strLen int) bool {
	if len(str) != strLen && strLen != -1 {
		return false
	}
	match, _ := regexp.MatchString("^[A-Za-z0-9]+$", str)
	return match
}

// CheckStrLen 检查字符串Aa0，判断字符串长度为1-strLen之间
func CheckStrLen(str string, strLen int) bool {
	if len(str) > strLen || len(str) == 0 {
		return false
	}
	match, _ := regexp.MatchString("^[A-Za-z0-9]+$", str)
	return match
}

// CheckStrMaxLen 检查字符串Aa0，长度最大strLen
func CheckStrMaxLen(str string, strLen int) bool {
	if len(str) > strLen {
		return false
	}
	match, _ := regexp.MatchString("^[A-Za-z0-9]+$", str)
	return match
}

// CheckStrMinLen 检查字符串Aa0，长度最小strLen
func CheckStrMinLen(str string, strLen int) bool {
	if len(str) < strLen {
		return false
	}
	match, _ := regexp.MatchString("^[A-Za-z0-9]+$", str)
	return match
}
