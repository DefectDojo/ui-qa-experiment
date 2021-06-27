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

	// Go to the User's Page
	// TODO: Iffy Selector
	// #side-menu > li:nth-child(9) > a > i
	p.MustElement("#side-menu > li:nth-child(9) > a > i").MustClick()

	// Wait for page to load
	p.WaitLoad()

	// Click on the three dots next to the user wanting to be changed
	// Editing user "First Name 3"
	// #dropdownMenuUser > b
	p.MustElement("#dropdownMenuUser > b").MustClick()

	// Click on edit
	// #editUser
	p.MustElement("#editUser").MustClick()
	p.WaitLoad()

	// Fill out form
	// Modify Username Information
	// Username ID: #id_username
	p.MustElement("#id_username").MustInput("Username4")

	// Modify First Name Information
	// First Name ID: #id_first_name
	p.MustElement("#id_first_name").MustInput("First Name4")

	// Modify last Name Information
	// Last Name ID: #id_last_name
	p.MustElement("#id_last_name").MustInput("Last Name4")

	// Modify Email Address Information
	// Email Address ID: #id_email
	p.MustElement("#id_email").MustInput("emailaddress4@emailaddress.com")

	// Select desired status (Active Status selected as default)
	// Keep Staff status
	// #id_is_staff
	// p.MustElement("#id_is_staff").MustClick()

	// Remove Super status
	// #id_is_superuser
	p.MustElement("#id_is_superuser").MustClick()

	// Click submit button
	// TODO: Iffy Selector
	// #base-content > form > div > div > input
	p.MustElement("#base-content > form > div > div > input").MustClick()

	fmt.Printf("PASS - TEST-U05 Updated User's Information in %+v\n", time.Since(started))

}
