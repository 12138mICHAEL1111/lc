package main

var romanMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X' :10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	l := len(s)
	result := romanMap[s[l-1]]
	for i:=l-2; i>=0;i-- {
		if(romanMap[s[i]]<romanMap[s[i+1]]){
			result = result - romanMap[s[i]]
		}else{
			result = result + romanMap[s[i]]
		}
	}
	return result
}
