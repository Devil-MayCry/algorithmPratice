package main

/*
给定一个字符串str,str全部由数字宇符组成，
如果str中某一个或某相邻两个字符组成的子串值在l~26之间，则这个子串可以转换为一个字母。
规定”1”转换为”A”，”2”转换为“B”,..."26"转换为“Z”。
写一个函数，求str有多少种不同的转换结果，并返回种数。
*/

func main() {

}

func num1(str string) int {
	if str == "" {
		return 0
	}
	chs := []rune(str)
	return process(chs, 0)
}

func process(chs []rune, i int) int {
	if i == len(chs) {
		return 1
	}
	if chs[i] == '0' {
		return 0
	}
	res := process(chs, i+1)
	if i+1 < len(chs) && (chs[i]-'0'*10+chs[i+1]-'0' < 27) {
		res += process(chs, i+2)
	}
	return res
}
