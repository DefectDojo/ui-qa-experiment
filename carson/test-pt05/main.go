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

	// Brackets need to be fixed
	// Shorter name for sess.Page
	p := *sess.Page

	// Hover over Products on the side menu
	// TODO: bad selector
	// #side-menu > li:nth-child(2) > a
	p.MustElement("#side-menu > li:nth-child(2) > a").Hover()

	time.Sleep(time.Millisecond * 200)

	// Click on Product Types Listing on the side menu
	// TODO: Bad Selector
	// #side-menu > li:nth-child(2) > ul > li:nth-child(3) > a
	p.MustElement("#side-menu > li:nth-child(2) > ul > li:nth-child(3) > a").MustClick()

	// Edit correct Product Type
	row := 0
	for j := 2; j <= 15; j++ {
		fmt.Println(j)
		// #product_types > tbody > tr:nth-child(3) > td:nth-child(2)
		selector := "#product_types > tbody > tr:nth-child(" + strconv.Itoa(j) + ") > td:nth-child(2)
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
	// #users > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(1) > ul:nth-child(1) > li:nth-child(1) > a:nth-child(1)
	// Need to get correct selector for "User Row line" (Below)
	productRow := "#users > tbody:nth-child(1) > tr:nth-child(" + strconv.Itoa(row) + ") > td:nth-child(1) > ul:nth-child(1) > li:nth-child(1) > a:nth-child(1)"
	p.MustElement(productRow).MustClick()
	// .open > ul:nth-child(2) > li:nth-child(6) > a:nth-child(1)
	// #editUser
	p.MustElement(".open > ul:nth-child(2) > li:nth-child(6) > a:nth-child(1)").MustClick()

	time.Sleep(time.Millisecond * 200)

	// Fill out form
	// Name ID: #id_name
	p.MustElement("#id_name").MustInput("A Product Type")

	// TODO: Bad Selector
	// Description ID:#base-content > form > div:nth-child(3) > div > div > div.CodeMirror.cm-s-easymde.CodeMirror-wrap > div.CodeMirror-scroll > div.CodeMirror-sizer > div > div > div > div.CodeMirror-code > pre
	p.MustElement("#base-content > form > div:nth-child(3) > div > div > div.CodeMirror.cm-s-easymde.CodeMirror-wrap > div.CodeMirror-scroll > div.CodeMirror-sizer > div > div > div > div.CodeMirror-code > pre").MustInput("Description of product type")

	//Click Submit Button
	// ID: #base-content > form > div:nth-child(6) > div > input
	p.MustElement("#base-content > form > div:nth-child(6) > div > input").MustClick()

	fmt.Printf("PASS - TEST-PT05 Edited all elements for a Product Type in %+v\n", time.Since(started))

}
