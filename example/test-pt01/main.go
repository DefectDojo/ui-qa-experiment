package main

import (
	"fmt"
	"os"
	"time"

	ddl "github.com/DefectDojo/ui-qa-experiment/login"
)

func main() {
	// Start the timer
	started := time.Now()

	// Login and start a session with DefectDojo
	var sess ddl.DDLogin
	err := sess.SetAndLogin("https://demo.defectdojo.org/", "admin", "defectdojo@demo#appsec", true, false)
	if err != nil {
		fmt.Printf("Error logging into DefectDojo. Error was:\n\t%+v\n", err)
		fmt.Println("FAILED - Add User")
		os.Exit(1)
	}

	// Make a shorter name for sess.Page
	p := *sess.Page

	// Hover over Products on the side menu
	// TODO: bad selector
	// EXAMPLE - PR4797
	// #side-menu > li:nth-child(2) > a
	p.MustElement("#side-menu > li:nth-child(2) > a").Hover()

	// Click on Add Product Type
	// TODO: bad selector
	// Example - PR4799
	// #side-menu > li:nth-child(2) > ul:nth-child(2) > li:nth-child(4) > a:nth-child(1)
	p.MustElement("#side-menu > li:nth-child(2) > ul:nth-child(2) > li:nth-child(4) > a:nth-child(1)").MustClick()

	// Wait for page to load
	p.WaitLoad()

	// Fill in the form
	// #id_name (Product Type name)
	p.MustElement("#id_name").MustInput("QA Product Type")
	// TODO: bad selector
	// .CodeMirror > div:nth-child(1) > textarea:nth-child(1) (Description)
	p.MustElement(".CodeMirror > div:nth-child(1) > textarea:nth-child(1)").MustInput("QA Test Description")
	// #id_critical_product
	p.MustElement("#id_critical_product").MustClick()
	// #id_key_product
	p.MustElement("#id_key_product").MustClick()

	// Click submit
	// TODO: Iffy selector
	// input.btn
	p.MustElement("input.btn").MustClick()

	// Uncomment for debugging
	//time.Sleep(time.Minute * 30)

	fmt.Printf("PASS - pt01 Add Product Type in %+v\n", time.Since(started))
}
