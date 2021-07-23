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

	time.Sleep(time.Millisecond * 400)

	// Hover over Product Menu
	// TODO: Iffy Selector
	// #side-menu > li:nth-child(2) > a
	// New ID: #product-side-menu
	p.MustElement("#side-menu > li:nth-child(2) > a").Hover()

	time.Sleep(time.Millisecond * 400)

	// Hover over Add Product Type
	// TODO: Bad Selector
	// #side-menu > li:nth-child(2) > ul > li:nth-child(4) > a
	// New ID: #add-product-type
	p.MustElement("#side-menu > li:nth-child(2) > ul > li:nth-child(4) > a").Hover()

	// Click Add Product Type
	// TODO: Bad Selector
	// #side-menu > li:nth-child(2) > ul > li:nth-child(4) > a
	// New ID: #add-product-type
	p.MustElement("#side-menu > li:nth-child(2) > ul > li:nth-child(4) > a").MustClick()

	//Wait for the page to load
	p.WaitLoad()

	// Fill out Form
	// Name ID: #id_name
	p.MustElement("#id_name").MustInput("Critical Product")

	// TODO: Bad Selector
	// Description ID: ##base-content > form > div:nth-child(3) > div > div > div.CodeMirror.cm-s-easymde.CodeMirror-wrap > div.CodeMirror-scroll > div.CodeMirror-sizer > div > div > div > div.CodeMirror-code > pre
	// New ID: #description
	p.MustElement("#base-content > form > div:nth-child(3) > div > div > div.CodeMirror.cm-s-easymde.CodeMirror-wrap > div.CodeMirror-scroll > div.CodeMirror-sizer > div > div > div > div.CodeMirror-code > pre")

	// Select Critical Product Type
	// #id_critical_product
	p.MustElement("#id_critical_product").MustClick()

	// Submit Form
	// TODO: Bad Selector
	// #base-content > form > div:nth-child(6) > div > input
	// New ID: // New ID: #submit-button
	p.MustElement("#base-content > form > div:nth-child(6) > div > input").MustClick()

	fmt.Printf("PASS - TEST-PT03 Added a Critical Product Type in %+v\n", time.Since(started))

}
