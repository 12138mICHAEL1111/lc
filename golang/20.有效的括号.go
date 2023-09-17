package main

func isLeft(s byte)bool{
	if s == '(' || s == '{' || s == '['{
		return true
	}

	return false
}

func isValid(s string) bool {
	stack := []byte{}
	m := map[byte]byte {
		')':'(',
		']':'[',
		'}':'{',
	}
	
	for i:=0; i<len(s); i++{
		if isLeft(s[i]){
			stack = append(stack, s[i])
		}else{
			if len(stack) != 0 && m[s[i]] == stack[len(stack)-1]{
				stack = stack[:len(stack)-1]
			}else{
				return false
			}
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}