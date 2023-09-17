package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	numLetterMap := map[byte]string{
		'2' : "abc",
		'3' :  "def",
		'4' : "ghi",
		'5' : "jkl",
		'6' : "mno",
		'7' : "pqrs",
		'8' : "tuv",
		'9' : "wxyz",
	}
	result := []string{}
	var dfs func(int) 
	combo := []byte{}
	// 不用记录visit，因为no+1，就是往下一层找，不会出现ad，da的情况
	// 而且当出现22的情况，用了visit会导致没有aa的情况
	dfs = func(no int) {
		if len(combo) == len(digits) {
			result = append(result, string(combo))
			return
		}
 
		num := digits[no]
		for i := 0 ; i< len(numLetterMap[num]) ;i++{
			letter := numLetterMap[num][i]
			combo = append(combo,letter)
			dfs(no+1)
			combo = combo[:len(combo) - 1]
		}
	}
	dfs(0)
	return result
}
