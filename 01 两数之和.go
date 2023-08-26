package main


func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i,v := range(nums){
		if index,ok := hashMap[target-v]; ok{
			return []int{i,index}
		}
		hashMap[v] = i
	}
	return []int{}
}