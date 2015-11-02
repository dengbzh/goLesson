/*
实现 WordCount。它应当返回一个含有 s 中每个 “词” 个数的 map。函数 wc.Test 针对这个函数执行一个测试用例，并输出成功还是失败。

你会发现 strings.Fields 很有帮助。
*/

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

var f = func(c rune) bool {
	return c == ' '
}

func WordCount(s string) map[string]int {

	split := strings.FieldsFunc(s, f)
	wc := make(map[string]int)

	for _, v := range split {
		_, ok := wc[v]
		if !ok {
			wc[v] = 1
		} else {
			wc[v]++
		}
	}
	return wc
}

func main() {
	wc.Test(WordCount)
}
