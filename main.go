package main

import "fmt"

var hourlyWage float32 = 24.37
var daysWorkedWeekly float32 = 4
var daysOnVacation float32 = 12
var hoursWorkedDaily float32 = 9.2
var dailyPay = hourlyWage * hoursWorkedDaily

func main() {
	fmt.Print("You will be going on vacation for ", +daysOnVacation, " days. Given your hourly wage of $", +hourlyWage, ", you will be receiving $", dailyPay*daysOnVacation, " less than your usual monthly check.")
}
