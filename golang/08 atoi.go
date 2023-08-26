package main

import (
	"math"
	"strconv"
	"unicode"
)

var stateM = map[string][]string{
	"start":{"start","signed","in_number","end"},
	"signed":{"end","end","in_number","end"},
	"in_number":{"end","end","in_number","end"},
	"end" : {"end","end","end","end"},
}

func myAtoi(s string) int  {
    var state = "start"
    var answer = 0
	sign := "+"
	for _, r:= range s{
		if(state == "end"){
			break
		}
		state = stateM[state][getCol(r)]
        if state == "in_number"{
            i,_:=strconv.Atoi(string(r))
            answer = answer*10 + i //溢出 ，直接进一位变成负数 例如 01111111 +1 就变成了11111111
			if answer > math.MaxInt32 && sign == "+"{ 
				return math.MaxInt32
				
			}else if  answer > math.MaxInt32+1 && sign == "-"{ //最小负数的绝对值是math.MaxInt32+1
				return math.MinInt32
			}
        }
		if state == "signed" && string(r)=="-" {
			sign = "-"
		}
	}
	if sign == "-"{
		answer = -answer
	  }
	return answer
}


func getCol(r rune) int {
	if unicode.IsSpace(r){
		return 0
	}else if string(r)=="+" || string(r)== "-" {
		return 1
	}else if unicode.IsDigit(r){
		return 2
	}
	return 3
}

func overflow (i int8) int8{ //溢出 ，直接进一位变成负数 例如 01111111 +1 就变成了11111111
	return i + 1
}

func test(){
	overflow(127)
}