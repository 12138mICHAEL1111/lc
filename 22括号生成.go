package main


func generateParenthesis(n int) []string {
	var result []string
	var dfs func(string,int, int)
	dfs = func(s string,l int, r int){
		if len(s) == 2 * n{
			result = append(result, s)
			return
		}

		if l < n{
			dfs(s+"(",l+1,r)
		}

		if r < l {
			dfs(s+")",l,r+1)
		}
	}
	dfs("",0,0)
	return result
}