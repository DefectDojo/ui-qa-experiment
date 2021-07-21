package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	ddl "github.com/DefectDojo/ui-qa-experiment/login"
)

func main() {
	// Start the time
	started := time.Now()

	// Login into Defect Dojo
	var sess ddl.DDLogin
	err := sess.SetAndLogin("https://demo.defectdojo.org/", "admin", "defectdojo@demo#appsec", true, false)
	if err != nil {
		fmt.Printf("Error logging into DefectDojo. Error was:\n\t%+v\n", err)
		fmt.Printf("FAILED Log-In")
		os.Exit(1)

	}

	// Shorter name for sess.Page
	p := *sess.Page

	// Hover over product menu
	// TODO: Bad Selector
	// #side-menu > li:nth-child(2) > a:nth-child(1)
	// New ID: #product-side-menu
	p.MustElement("#side-menu > li:nth-child(2) > a:nth-child(1)").MustHover()
	time.Sleep(time.Millisecond * 200)

	// click on product type listings
	// TODO: Bad Selector
	// #side-menu > li:nth-child(2) > ul:nth-child(2) > li:nth-child(3) > a:nth-child(1)
	// New ID: #product-type-listings
	p.MustElement("#side-menu > li:nth-child(2) > ul:nth-child(2) > li:nth-child(3) > a:nth-child(1)").MustClick()

	// Wait for the page to load
	p.WaitLoad()

	// Delete correct product type
	row := 0
	for j := 2; j <= 15; j++ {
		fmt.Println(j)
		// #product_types > tbody > tr:nth-child(3) > td:nth-child(2)
		selector := "#product_types > tbody > tr:nth-child(" + strconv.Itoa(j) + ") > td:nth-child(2)"
		fmt.Println(selector)
		name := p.MustElement(selector).MustText()
		fmt.Println(name)
		if name == "Some Product Type" {
			// Matched correct username
			fmt.Println("We matched")
			row = j
			j = 15
		}
	}
	fmt.Println("After the loop")

	// click on edit product type icon
	// TODO: Bad Selector
	// #product_types > tbody:nth-child(2) > tr:nth-child(" + strconv.Itoa(row) + ") > td:nth-child(1) > ul:nth-child(1) > li:nth-child(1) > a:nth-child(1)
	// New ID: #edit-product-type
	productRow := "#product_types > tbody:nth-child(2) > tr:nth-child(" + strconv.Itoa(row) + ") > td:nth-child(1) > ul:nth-child(1) > li:nth-child(1) > a:nth-child(1)"
	p.MustElement(productRow).MustClick()

	// click on delete
	// TODO: Bad Selector
	// .open > ul:nth-child(2) > li:nth-child(8) > a:nth-child(1)
	// New ID: #delete-button
	p.MustElement(".open > ul:nth-child(2) > li:nth-child(8) > a:nth-child(1)").MustClick()

	// confirm deletion
	// TODO: Weird name selector
	// .btn-danger
	// New ID: #confirm-deletion
	p.MustElement(".btn-danger").MustClick()

	fmt.Printf("PASS - TEST-PT06 Deleted a Product Type in %+v\n", time.Since(started))

}
