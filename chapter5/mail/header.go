package main

import (
	"fmt"
	"net/mail"
	"strings"
)

// extract header info and print it nicely
func printHeaderInfo(header mail.Header) {

	// this works because we know it's a single address
	// otherwise use ParseAddressList
	toAddress, err := mail.ParseAddress(header.Get("To"))
	if err == nil {
		fmt.Printf("To: %s <%s>\n", toAddress.Name, toAddress.Address)
	}
	fromAddress, err := mail.ParseAddress(header.Get("From"))
	if err == nil {
		fmt.Printf("From: %s <%s>\n", fromAddress.Name, fromAddress.Address)
	}

	fmt.Println("Subject:", header.Get("Subject"))

	// this works for a valid RFC5322 date
	// it does a header.Get("Date"), then a
	// mail.ParseDate(that_result)
	if date, err := header.Date(); err == nil {
		fmt.Println("Date:", date)
	}

	fmt.Println(strings.Repeat("=", 40))
	fmt.Println()
}
