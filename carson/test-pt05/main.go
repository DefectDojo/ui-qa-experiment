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

	// Login and start a session with DefectDojo
	var sess ddl.DDLogin
	err := sess.SetAndLogin("https://demo.defectdojo.org/", "admin", "defectdojo@demo#appsec", true, false)
	if err != nil {
		fmt.Printf("Error logging into DefectDojo. Error was:\n\t%+v\n", err)
		fmt.Printf("FAILED Log-In")
		os.Exit(1)

	}

	// Shorter name for sess.Page
	p := *sess.Page

	// Hover over Products on the side menu
	// TODO: bad selector
	// #side-menu > li:nth-child(2) > a
	// New ID: #product-side-menu
	p.MustElement("#side-menu > li:nth-child(2) > a").Hover()

	time.Sleep(time.Millisecond * 200)

	// Click on Product Types Listing on the side menu
	// TODO: Bad Selector
	// #side-menu > li:nth-child(2) > ul > li:nth-child(3) > a
	// New ID: #all-product-types-list
	p.MustElement("#side-menu > li:nth-child(2) > ul > li:nth-child(3) > a").MustClick()

	// Edit correct Product Type
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
	// #dropdown
	// TODO: Bad Selector
	// #product_types > tbody:nth-child(2) > tr:nth-child(" + strconv.Itoa(row) + ") > td:nth-child(1) > ul:nth-child(1) > li:nth-child(1) > a:nth-child(1)
	// New ID: #product-type-options
	productRow := "#product_types > tbody:nth-child(2) > tr:nth-child(" + strconv.Itoa(row) + ") > td:nth-child(1) > ul:nth-child(1) > li:nth-child(1) > a:nth-child(1)"
	p.MustElement(productRow).MustClick()
	// TODO: Bad Selector
	// .open > ul:nth-child(2) > li:nth-child(2) > a:nth-child(1)
	// New ID: #edit-button
	p.MustElement(".open > ul:nth-child(2) > li:nth-child(2) > a:nth-child(1)").MustClick()

	time.Sleep(time.Millisecond * 200)

	// Fill out form
	// Name ID: #id_name
	p.MustElement("#id_name").MustSelectAllText().MustInput("A Product Type")

	// TODO: Bad Selector
	// Description ID:.CodeMirror > div:nth-child(1) > textarea:nth-child(1)
	// New ID: #description
	p.MustElement(".CodeMirror > div:nth-child(1) > textarea:nth-child(1)").MustSelectAllText().MustInput("Description of product type")

	// Critical Product Type ID: #id_critical_product
	p.MustElement("#id_critical_product").MustClick()
	// Key Product Type ID: #id_key_product
	// p.MustElement("#id_key_product").MustClick()

	//Click Submit Button
	// #base-content > form > div:nth-child(6) > div > input
	// New ID: #submit-button
	p.MustElement("#base-content > form > div:nth-child(6) > div > input").MustClick()
	// Current ID (row above)

	fmt.Printf("PASS - TEST-PT05 Edited all elements for a Product Type in %+v\n", time.Since(started))

}
