package main


const (
    TARGET = 24
    EPSILON = 1e-6 
    ADD, MULTIPLY, SUBTRACT, DIVIDE = 0, 1, 2, 3
)


func judgePoint24(cards []int) bool {
	list := []float64{}
	for i:=0;i<4;i++{
		list = append(list, float64(cards[i])) 
	}
	return solve(list)
}

func solve(list []float64) bool {
	if len(list) == 1 {
		if abs(TARGET - list[0]) < EPSILON{  //因为计算是浮点数，会出现无穷小数的情况
			return true
		}
	}

	for i := 0;i <len(list);i++{ // 选第一个数
		for j := 0 ; j <len(list) ; j ++ { // 选第二个数
			if i != j{ // 避免选相同的数
				residualList := []float64{}
				for x := 0 ; x < len(list) ; x ++{
					if x != i && x != j{
						residualList = append(residualList, list[x]) // 添加剩余的数
					}
				}

				for y:= 0; y < 4 ; y++{ //循环操作符
					if y < 2 && i > j { // + 和 *, 1 +2 和2+1是一样的，没必要重复计算
						continue
					}
					var num float64
					switch y{
					case ADD:
						num = list[i] + list[j]
					case MULTIPLY:
						num = list[i] * list[j]
					case SUBTRACT:
						num = list[i] - list[j]
					case DIVIDE:
						if abs(list[j]) < EPSILON{ // 除数不能为0
							continue
						}
						num = list[i] / list[j]
					}
					residualList = append(residualList, num)
					result := solve(residualList)
					if result == true{
						return true
					}
					residualList = residualList[:len(residualList)-1]
				}
			}
		}
	}
	return false
}

func abs(num float64)float64{
	if num < 0 {
		return - num
	}
	return num
}