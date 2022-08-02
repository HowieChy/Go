package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func (prod *Product) store() {
	file, _ := os.Create(prod.id + ".txt")

	content := fmt.Sprintf(
		"ID: %v\nTitle: %v\nDescription: %v\nPrice: USD %.2f\n",
		prod.id,
		prod.title,
		prod.description,
		prod.price,
	)

	file.WriteString(content)

	file.Close()
}

//输入
func (prod *Product) printData() {
	fmt.Printf("ID: %v\n", prod.id)
	fmt.Printf("Title: %v\n", prod.title)
	fmt.Printf("Description: %v\n", prod.description)
	fmt.Printf("Price: USD %.2f\n", prod.price)
}

//返回一个指针
func NewProduct(id string, t string, d string, p float64) *Product {
	//指针保存数据的内存地址
	return &Product{id, t, d, p}
}

func main() {
	createdProduct := getProduct()

	createdProduct.printData()
	createdProduct.store()
}

func getProduct() *Product {
	fmt.Println("Please enter the product data.")
	fmt.Println("-----------------------------")

	//读取标准命令行
	reader := bufio.NewReader(os.Stdin)

	idInput := readUserInput(reader, "Product ID: ")
	titleInput := readUserInput(reader, "Product Title: ")
	descriptionInput := readUserInput(reader, "Product Description: ")
	priceInput := readUserInput(reader, "Product Price: ")

	//字符串转数值
	priceValue, _ := strconv.ParseFloat(priceInput, 64)

	product := NewProduct(idInput, titleInput, descriptionInput, priceValue)

	return product
}

func readUserInput(reader *bufio.Reader, promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	//去回车空格
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}
