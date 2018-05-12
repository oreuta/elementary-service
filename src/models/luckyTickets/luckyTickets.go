package luckyTickets

import (
	"fmt"
	"errors"
)

func validate(obj *context) (err error){
	if (obj.min - 99999) < 0 || (obj.max - 99999) < 0 || obj.min > obj.max || obj.min > 999999 || obj.max > 999999{
		return errors.New("incorrect data!")
	}
	return
}
type errorInstruction struct {
	Status string
	Reason string
}

type context struct {
	min, max int
}

type winner struct {
	method                              string
	countFirstMethod, countSecondMethod int
}

func simpleMethod(obj *context) int {
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

func hardMethod(obj *context) int {
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

func GetLuckyTickets(obj *context) (win winner, err error){
	check := validate(obj)
	if check != nil{
		win = winner{"incorrect data!", 0 ,0}
		return win, check
	}
	win = winner{"", simpleMethod(obj), hardMethod(obj)}
	if win.countFirstMethod > win.countSecondMethod{
		win.method = "Simple method is win!"
	}else if win.countFirstMethod < win.countSecondMethod{
		win.method = "Hard method is win!"
	}else{
		win.method = "Methods are equally!"
	}
	return
}

func enterData() (obj *context){
	obj = new(context)
	fmt.Println("Enter the min-value (length > 6):")
	fmt.Scan(&obj.min)
	fmt.Println("Enter the max-value (length > 6):")
	fmt.Scan(&obj.max)
	return obj
}
func main() {
	obj := enterData()
	winner, err := GetLuckyTickets(obj)
	fmt.Println(winner, err)
}