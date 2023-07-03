package main
// 总是反转一半位数的数字，然后再比大小
// 例如 1223 就反转23 得到32
// 通过取余数得到末尾数字
func isPalindrome(x int) bool {
	if x<0 || (x % 10 == 0 && x != 0) { //结尾是0，不会是回文
		return false
	}
	reversedNum := 0

	for (x>reversedNum) { //相等就停，小于也停，因为翻转超过一半的位数了
		reversedNum =  reversedNum * 10 + x % 10 // 得到从末尾开始的数字
		x = x / 10
	}

	return x == reversedNum || x == reversedNum /10 //奇数就把反转后的数字除10，去除中间数
}