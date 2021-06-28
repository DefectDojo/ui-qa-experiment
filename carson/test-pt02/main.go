package main

import (
	"fmt"
	"os"
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

	// Hover over Side Menu Icon
	// TODO: Iffy Selector
	// #side-menu > li:nth-child(2) > a
	p.MustElement("#side-menu > li:nth-child(2) > a").Hover()

	time.Sleep(time.Millisecond * 200)
	fmt.Println("Before the form.")

	// Click on Product Types Listing on the side menu
	// TODO: Bad Selector
	// #side-menu > li:nth-child(2) > ul > li:nth-child(3) > a
	p.MustElement("#side-menu > li:nth-child(2) > ul > li:nth-child(3) > a").MustClick()

	fmt.Println("Made it to the form.")
	time.Sleep(time.Millisecond * 200)

	// Click on wrench icon
	// TODO: Iffy Selector
	// #dropdownMenu1 > span.fa.fa-wrench
	p.MustElement("#dropdownMenu1 > span.fa.fa-wrench").MustClick()

	// Click on "Add Product Type"
	// TODO: Bad Selector
	// #base-content > div > div > div:nth-child(1) > div.panel-heading.tight > h3 > div > ul > li > a
	p.MustElement("#base-content > div > div > div:nth-child(1) > div.panel-heading.tight > h3 > div > ul > li > a").MustClick()

	//Wait for the page to load
	p.WaitLoad()

	// Fill out form
	// Name ID: #id_name
	p.MustElement("#id_name").MustInput("Some Product Type")

	// TODO: Bad Selector
	// Description ID: #base-content > form > div:nth-child(3) > div > div > div.CodeMirror.cm-s-easymde.CodeMirror-wrap > div.CodeMirror-scroll > div.CodeMirror-sizer > div > div > div > div.CodeMirror-code > pre
	p.MustElement(".CodeMirror > div:nth-child(1) > textarea:nth-child(1)").MustInput("Some Product Type Test Description")

	// Click on Submit Button
	p.MustElement("input.btn").MustClick()

	time.Sleep(time.Minute * 5)

	fmt.Printf("PASS - pt01 Add Product Type in %+v\n", time.Since(started))

}