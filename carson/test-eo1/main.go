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

	time.Sleep(time.Millisecond * 200)

	// Hover over product menu
	// TODO: Iffy Selector
	// #side-menu > li:nth-child(2) > a:nth-child(1)
	// New ID: #product-side-menu
	p.MustElement("#side-menu > li:nth-child(2) > a:nth-child(1)").MustHover()

	time.Sleep(time.Millisecond * 200)

	// Click on all products
	// TODO: Bad Selector
	// #side-menu > li:nth-child(2) > ul:nth-child(2) > li:nth-child(1) > a:nth-child(1)
	// New ID: #all-product-listings
	p.MustElement("#side-menu > li:nth-child(2) > ul:nth-child(2) > li:nth-child(1) > a:nth-child(1)").MustClick()

	// Wait for the page to load
	p.WaitLoad()

	// Insert loop code
	row := 0
	for j := 2; j <= 15; j++ {
		fmt.Println(j)
		// tr.odd:nth-child(" + strconv.Itoa(j) + ") > td:nth-child(2) > a:nth-child(1) > b:nth-child(1)
		selector := "tr.odd:nth-child(" + strconv.Itoa(j) + ") > td:nth-child(2) > a:nth-child(1) > b:nth-child(1)"
		fmt.Println(selector)
		name := p.MustElement(selector).MustText()
		fmt.Println(name)
		if name == "Test Product" {
			// Matched correct username
			fmt.Println("We matched")
			row = j
			j = 15
		}
	}
	fmt.Println("After the loop")
	// Click on edit button for wanted product
	// TODO: Bad Selector
	// tr.odd:nth-child(3) > td:nth-child(" + strconv.Itoa(row) + ") > div:nth-child(1) > div:nth-child(1) > a:nth-child(1) > b:nth-child(1)
	productRow := "tr.odd:nth-child(3) > td:nth-child(" + strconv.Itoa(row) + ") > div:nth-child(1) > div:nth-child(1) > a:nth-child(1) > b:nth-child(1)"
	p.MustElement(productRow).MustClick()

	// Click on add new engagement
	// TODO: Bad Selector
	// .open > ul:nth-child(2) > li:nth-child(5) > a:nth-child(1)
	// New ID: #add-new-engagement
	p.MustElement(".open > ul:nth-child(2) > li:nth-child(5) > a:nth-child(1)").MustClick()

	// Fill out name
	// #id_name
	p.MustElement("#id_name").MustInput("Test Engagement")

	// Fill out description
	// TODO: Bad Selector
	// .CodeMirror > div:nth-child(1) > textarea:nth-child(1)
	p.MustElement(".CodeMirror > div:nth-child(1) > textarea:nth-child(1)").MustInput("test engagement description")

	// Submit form
	// TODO: Iffy Selector
	// input.btn:nth-child(3)
	// New ID: #done-button
	p.MustElement("input.btn:nth-child(3)").MustClick()

	time.Sleep(time.Millisecond * 200)

	fmt.Printf("PASS - TEST-EO1 Added an interactive engagement in %+v\n", time.Since(started))

}
