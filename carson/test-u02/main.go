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

	// Go to the User's Page
	// Iffy Selector
	// #side-menu > li:nth-child(9) > a > i
	// New ID: #user-icon
	p.MustElement("#side-menu > li:nth-child(9) > a > i").MustClick()

	// Wait for page to load
	p.WaitLoad()

	// Click on wrench (top right)
	// TODO: Iffy selector
	// #dropdownMenu1 > span.fa.fa-wrench
	// New ID: #wrench-icon
	p.MustElement("#dropdownMenu1 > span.fa.fa-wrench").MustClick()

	// Click on New User
	// TODO: bad selector
	// New ID: #new-user-icon
	// #base-content > div > div > div:nth-child(1) > div.panel-heading.tight > h3 > div > ul > li > a
	p.MustElement("#base-content > div > div > div:nth-child(1) > div.panel-heading.tight > h3 > div > ul > li > a").MustClick()

	// Wait for page to load
	p.WaitLoad()

	// Fill out form
	// Username ID: #id_username
	p.MustElement("#id_username").MustInput("Username1")

	// First Name ID: #id_first_name
	p.MustElement("#id_first_name").MustInput("First Name1")

	// Last Name ID: #id_last_name
	p.MustElement("#id_last_name").MustInput("Last Name1")

	// Email Address ID: #id_email
	p.MustElement("#id_email").MustInput("emailaddress1@emailaddress.com")

	// Select Staff status
	// #id_is_staff
	p.MustElement("#id_is_staff").MustClick()

	// Click submit Button
	// Iffy Selector
	// #base-content > form > div > div > input
	// New ID: #submit-button
	p.MustElement("#base-content > form > div > div > input").MustClick()

	fmt.Printf("PASS - TEST-U02 Added an staff user in %+v\n", time.Since(started))

}
