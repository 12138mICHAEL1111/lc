package main

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack){
		return -1
	}
	index := -1
	for i:=0;i<=len(haystack)-len(needle) ;i++{	
		if haystack[i] == needle[0]{
			index = i
			for j:=1;j<len(needle);j++{
				if haystack[i+j] != needle[j]{
					index = -1
					break
				}
			}
			if index != -1{
				break
			}
		}
	}
	return index
}