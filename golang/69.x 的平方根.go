package main
func mySqrt(x int) int {
	l,r := 0,x
	ans := -1 
	for l <= r{
		mid := l + (r-l) / 2
		if mid * mid <= x {
			ans = mid
			l = mid +1 
		}else{
			r = mid - 1
		}
	}
	return ans
}

// for i:= 0 ;i<x ;i ++{
// 	if i * i <= x && (i+1) * (i+1) > x {
// 		return i
// 	}
// }
// return 0