package main

import (
	"fmt"
	"io"
	"log"
	"net/mail"
	"os"
	"strings"
)

// an example email message
const msg string = `Date: Thu, 24 Jul 2019 08:00:00 -0700
From: Aaron <fake_sender@example.com>
To: Reader <fake_receiver@example.com>
Subject: Gophercon 2019 is going to be awesome!

Feel free to share my book with others if you're attending.
This recipe can be used to process and parse email information.
`

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

func main() {
	r := strings.NewReader(msg)
	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatal(err)
	}

	printHeaderInfo(m.Header)

	// after printing the header, dump the body to stdout
	if _, err := io.Copy(os.Stdout, m.Body); err != nil {
		log.Fatal(err)
	}
}
