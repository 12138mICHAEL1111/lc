package main

func combinationSum(candidates []int, target int) [][]int {
	combo := []int{}
	result := [][]int{}
	var dfs func(int, int)
	dfs = func(begin, sum int) {
		if sum == target {
			result = append(result, append([]int(nil), combo...))
			return
		}

		if sum > target {
			return
		}

		for i := begin; i < len(candidates); i++ { // begin相当于去重，例如2，3，4，223和322是一样的，但在遍历3的时候i已经为1了就不会遍历2了
			combo = append(combo, candidates[i])

			dfs(i, sum+candidates[i])

			combo = combo[:len(combo)-1]
		}
	}
	dfs(0, 0)
	return result
}
