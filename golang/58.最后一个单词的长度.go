package main
func lengthOfLastWord(s string) int {
	l := 0
	
	for i:=len(s)-1;i >= 0;i--{
		if string(s[i]) != " "{
			l++
			if string(s[i-1]) == " "  {
				break
			}
		} 
	}
	return l
}