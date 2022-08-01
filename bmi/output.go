package main

import (
	"strconv"

	"github.com/common-nighthawk/go-figure"
)

func printBMI(bmi float64) {
	// fmt.Printf("Your BMI: %.2f", bmi)
	myFigure := figure.NewFigure(strconv.FormatFloat(bmi, 'f', 2, 32), "", true)
	myFigure.Print()
}
