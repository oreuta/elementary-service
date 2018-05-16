package luckyTickets

import (
	"errors"
)

func validate(obj *TaskContext) (err error){
	if (obj.Min - 99999) < 0 || (obj.Max - 99999) < 0 || obj.Min > obj.Max || obj.Min > 999999 || obj.Max > 999999{
		return errors.New("incorrect data")
	}
	return
}

type TaskContext struct {
	Min, Max int
}

// Winner is a struct for return winner
type Winner struct {
	Method                              int
	CountFirstMethod, CountSecondMethod int
}

func simpleMethod(obj *TaskContext) int {
	var firstThreeLettersSum, lastThreeLettersSum, count int
	for current := obj.Min; current <= obj.Max; current++ {
		firstThreeLettersSum += (current / 100000)%10
		firstThreeLettersSum += (current / 10000)%10
		firstThreeLettersSum += (current / 1000)%10

		lastThreeLettersSum += (current / 100)%10
		lastThreeLettersSum += (current / 10)%10
		lastThreeLettersSum += (current / 1)%10
		if firstThreeLettersSum == lastThreeLettersSum {
			count++
		}
		firstThreeLettersSum, lastThreeLettersSum = 0, 0
	}
	return count
}

func hardMethod(obj *TaskContext) int {
	var even, odd, count int
	for current := obj.Min; current <= obj.Max; current++ {
		for i := 100000; i >= 1; i /= 10{
			if ((current / i)%10)%2 == 0{
				odd += (current / i)%10
			}else{
					even += (current / i)%10
				}
			}
			if odd == even{
				count++
			}
			odd, even = 0, 0
		}
	return count
}

// GetLuckyTickets is a function for get the lucky tickets by simple and hard methods
func GetLuckyTickets(obj *TaskContext) (win Winner , err error){
	err = validate(obj)
	if err != nil{
		win = Winner {0, 0 ,0}
		return win, err
	}
	win = Winner {0, simpleMethod(obj), hardMethod(obj)}
	if win.CountFirstMethod > win.CountSecondMethod{
		win.Method = 1
	}else if win.CountFirstMethod < win.CountSecondMethod{
		win.Method = 2
	}else{
		win.Method = 3
	}
	return
}