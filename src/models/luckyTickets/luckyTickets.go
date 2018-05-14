package luckyTickets

import (
	"fmt"
	"errors"
)

func validate(obj *taskContext) (err error){
	if (obj.min - 99999) < 0 || (obj.max - 99999) < 0 || obj.min > obj.max || obj.min > 999999 || obj.max > 999999{
		return errors.New("incorrect data!")
	}
	return
}
type errorInstruction struct {
	Status string
	Reason string
}

type taskContext struct {
	min, max int
}

type winner struct {
	method                              int
	countFirstMethod, countSecondMethod int
}

func simpleMethod(obj *taskContext) int {
	var firstThreeLettersSum, lastThreeLettersSum, count int
	for current := obj.min; current <= obj.max; current++ {
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

func hardMethod(obj *taskContext) int {
	var even, odd, count int
	for current := obj.min; current <= obj.max; current++ {
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

func GetLuckyTickets(obj *taskContext) (win winner, err error){
	err = validate(obj)
	if err != nil{
		win = winner{0, 0 ,0}
		return win, err
	}
	win = winner{0, simpleMethod(obj), hardMethod(obj)}
	if win.countFirstMethod > win.countSecondMethod{
		win.method = 1
	}else if win.countFirstMethod < win.countSecondMethod{
		win.method = 2
	}else{
		win.method = 3
	}
	return
}

func enterData() (obj *taskContext){
	obj = new(taskContext)
	fmt.Println("Enter the min-value (length > 6):")
	fmt.Scan(&obj.min)
	fmt.Println("Enter the max-value (length > 6):")
	fmt.Scan(&obj.max)
	return obj
}