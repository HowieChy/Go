package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	birthDate   string
	createdDate time.Time
}

func (user *User) outputDetails() {
	//user.firstName go会在后台转换成(*user).firstName
	fmt.Printf("My name is %v %v (born on %v)", user.firstName, user.lastName, user.birthDate)
}

//加*User 代表返回指针
func NewUser(fName string, lName string, bDate string) *User {
	created := time.Now()

	user := User{
		firstName:   fName,
		lastName:    lName,
		birthDate:   bDate,
		createdDate: created,
	}

	//返回数据的内存地址
	return &user
}

var reader = bufio.NewReader(os.Stdin)

func main() {
	var newUser *User

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (YYYY/MM/DD): ")

	newUser = NewUser(firstName, lastName, birthdate)

	// ... do something awesome with that gathered data!

	newUser.outputDetails()
	//可以使用time内置的方法 返回创建的哪周几
	fmt.Println(newUser.createdDate.Weekday())
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')

	cleanedInput := strings.Replace(userInput, "\n", "", -1)

	return cleanedInput
}
