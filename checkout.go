package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	products       []string
	productQuantity []int
	prices         []float64
	priceTotal     []float64
	continuePrompt string
	date           = time.Now()
	scanner        = bufio.NewScanner(os.Stdin)
	vat            = 7.5
	pricePaid      float64
	discount       float64
)

func cashierPrompts() {
	fmt.Println("What did the user buy?")
	scanner.Scan()
	productName := scanner.Text()
	products = append(products, productName)

	fmt.Println("How many pieces?")
	scanner.Scan()
	quantity, _ := strconv.Atoi(scanner.Text())
	productQuantity = append(productQuantity, quantity)

	fmt.Println("How much per unit?")
	scanner.Scan()
	productPrice, _ := strconv.ParseFloat(scanner.Text(), 64)
	prices = append(prices, productPrice)
}

func getPriceTotal() []float64 {
	for i := 0; i < len(prices); i++ {
		priceTotal = append(priceTotal, float64(productQuantity[i])*prices[i])
	}
	return priceTotal
}

func calculateTotal() float64 {
	var total float64
	for i := 0; i < len(prices); i++ {
		total += priceTotal[i]
	}
	return total
}

func discountTotal(discount float64) float64 {
	return calculateTotal() * (discount / 100)
}

func vatDiscount() float64 {
	return calculateTotal() * (vat / 100)
}

func paymentAmountCheck() bool {
	return pricePaid > ((calculateTotal() + vatDiscount()) - discountTotal(discount))
}

func itemListDisplay() {
	for i := 0; i < len(products); i++ {
		fmt.Printf("%20s%10d%10.2f%15.2f\n", products[i], productQuantity[i], prices[i], getPriceTotal()[i])
	}
}

func receipt() {
	fmt.Printf("%40s: %.2f\n", "Amount Paid", pricePaid)
	fmt.Printf("%40s: %.2f\n", "Balance", pricePaid-billTotal())
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("\t\t\tTHANK YOU FOR YOUR PATRONAGE")
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("--------------------------------------------------------------")
}

func billTotal() float64 {
	return (calculateTotal() + vatDiscount()) - discountTotal(discount)
}

func paymentLine() {
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("--------------------------------------------------------------")
	fmt.Printf("THIS IS NOT AN RECEIPT KINDLY PAY %.2f\n", billTotal())
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("--------------------------------------------------------------")
}

func displayPaymentSlip(cashierName, customerName string, discount float64) {
	fmt.Println("SEMICOLON STORES\nMAIN BRANCH\nLOCATION: 312, HERBERT MACAULAY WAY, SABO YABA, LAGOS.")
	fmt.Printf("Date: %s\n", date)
	fmt.Printf("Cashier: %s\n", cashierName)
	fmt.Printf("Customer Name: %s\n", customerName)
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("--------------------------------------------------------------")
	fmt.Printf("%20s%10s%10s%15s\n", "ITEM", "QTY", "PRICE", "TOTAL(NGN)")
	fmt.Println("--------------------------------------------------------------")
	itemListDisplay()
	fmt.Println("--------------------------------------------------------------")
	fmt.Printf("%40s: %.2f\n", "Sub Total", calculateTotal())
	fmt.Printf("%40s: %.2f\n", "Discount", discountTotal(discount))
	fmt.Printf("%40s: %.2f\n", "VAT @ 7.50%", vatDiscount())
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("--------------------------------------------------------------")
	fmt.Printf("%40s: %.2f\n", "Bill Total", billTotal())
}

func main() {
	fmt.Println("What is the customer's name?")
	scanner.Scan()
	customerName := scanner.Text()

	for {
		cashierPrompts()
		fmt.Println("Add more Items?")
		scanner.Scan()
		continuePrompt = scanner.Text()
		if strings.ToLower(continuePrompt) != "yes" {
			break
		}
	}

	fmt.Println("What is your name?")
	scanner.Scan()
	cashierName := scanner.Text()

	fmt.Println("How much discount will the customer get?")
	scanner.Scan()
	discount, _ = strconv.ParseFloat(scanner.Text(), 64)
	fmt.Println("")

	displayPaymentSlip(cashierName, customerName, discount)
	paymentLine()
	fmt.Println("\n\n\n\n")

	fmt.Println("How much did the customer give to you?")
	scanner.Scan()
	pricePaid, _ = strconv.ParseFloat(scanner.Text(), 64)
	for !paymentAmountCheck() {
		fmt.Println("Pay the correct amount")
		scanner.Scan()
		pricePaid, _ = strconv.ParseFloat(scanner.Text(), 64)
		fmt.Println("")
		if paymentAmountCheck() {
			displayPaymentSlip(cashierName, customerName, discount)
			receipt()
		}
	}
}
