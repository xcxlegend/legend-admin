package common

import "strconv"

// 字符串的数组转int数组 主要用于HTTP请求转
func StringSlice2Int(sl []string) []int {
	isl := []int{}
	for _, s := range sl {
		i, _ := strconv.Atoi(s)
		isl = append(isl, i)
	}
	return isl
}

// 整型数组去重
func UniqueIntSlice(il []int) []int {
	m := map[int]struct{}{}
	newSlice := []int{}
	for _, i := range il{
		if _, ok := m[i]; !ok {
			newSlice = append(newSlice, i)
			m[i] = struct{}{}
		}
	}
	return newSlice
}