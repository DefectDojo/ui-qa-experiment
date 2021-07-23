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

	// Hover over user menu
	// #side-menu > li:nth-child(9) > a:nth-child(1)
	p.MustElement("#side-menu > li:nth-child(9) > a:nth-child(1)").MustHover()

	// Click on groups
	// #side-menu > li:nth-child(9) > ul:nth-child(2) > li:nth-child(2) > a:nth-child(1)
	// New ID: #group-listings
	p.MustElement("#side-menu > li:nth-child(9) > ul:nth-child(2) > li:nth-child(2) > a:nth-child(1)").MustClick()

	// Click wrench icon
	// #dropdownMenu1
	p.MustElement("#dropdownMenu1").MustClick()

	// Click on add new group
	// .dropdown-menu-right > li:nth-child(1) > a:nth-child(1)
	// New Id: #add-new-group
	p.MustElement(".dropdown-menu-right > li:nth-child(1) > a:nth-child(1)").MustClick()

	// Fill out form
	// Name: #id_name
	p.MustElement("#id_name").MustInput("New Group")

	// Description: .CodeMirror > div:nth-child(1) > textarea:nth-child(1)
	p.MustElement(".CodeMirror > div:nth-child(1) > textarea:nth-child(1)").MustInput("description for New Group")

	// Submit Button: input.btn
	p.MustElement("input.btn").MustClick()

	time.Sleep(time.Millisecond * 200)

	fmt.Printf("PASS - TEST-GR01 Added a new group in %+v\n", time.Since(started))

}
