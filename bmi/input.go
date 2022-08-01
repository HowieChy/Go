package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/HowieChy/Go/bmi/info"
)

//从标准输入中即命令行中获取
var reader = bufio.NewReader(os.Stdin)

func getUserMetrics() (weight float64, height float64) {
	weight = getUserInput(info.WeightPrompt)
	height = getUserInput(info.HeightPrompt)

	return
}

func getUserInput(promptText string) (value float64) {
	fmt.Print(promptText)

	// \n表示换行的时候去读取
	userInput, _ := reader.ReadString('\n')
	//去掉换行空格
	userInput = strings.Replace(userInput, "\n", "", -1)
	//字符串转数字
	value, _ = strconv.ParseFloat(userInput, 64)

	return
}
