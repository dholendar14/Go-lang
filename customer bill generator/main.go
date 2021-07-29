package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func userInput(prompt string, r *bufio.Reader)(string,error){
	fmt.Println(prompt)
	input,err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill{
	reader := bufio.NewReader(os.Stdin)
	name,_ := userInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Created the bill -", b.name)
	return b
}

func promptOption(b bill) {
	reader := bufio.NewReader(os.Stdin)
	option,_ := userInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	switch option {
	case "a":
		name, _ := userInput("Item name:", reader)
		price, _ := userInput("Item Price", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number....")
			promptOption(b)
		}
		b.addItems(name, p)

		fmt.Println("item added -", name, price)
		promptOption(b)

	case "t":
		tip, _ := userInput("Enter tip amount ($)", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number...")
			promptOption(b)
		}
		b.updateTip(t)
		fmt.Println("tip has been updated to", tip)
		promptOption(b)
	case "s":
		b.save()
		fmt.Println("bill has been saved as", b.name)
	default:
		fmt.Println("That was not a valid option...")
		promptOption(b)
	}
}


func main() {
	mybill := createBill()
	promptOption(mybill)

}

