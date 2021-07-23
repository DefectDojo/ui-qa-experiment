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

	// Go to the User's Page
	// TODO: Iffy Selector
	// #side-menu > li:nth-child(9) > a > i
	// New ID: #user-icon
	p.MustElement("#side-menu > li:nth-child(9) > a > i").MustClick()

	// Wait for page to load
	p.WaitLoad()

	// If username column equals correct username, then click on three dots to edit that user
	// Click on the three dots next to the user wanting to be changed
	// Admin ID: #users > tbody > tr:nth-child(2) > td:nth-child(4) > a

	row := 0
	for j := 2; j <= 15; j++ {
		// fmt.Println(j)
		selector := "#users > tbody > tr:nth-child(" + strconv.Itoa(j) + ") > td:nth-child(4) > a"
		// fmt.Println(selector)
		name := p.MustElement(selector).MustText()
		// fmt.Println(name)
		if name == "Test_Name" {
			// Matched correct username
			row = j
			j = 15
		}
	}
	// #users > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(1) > ul:nth-child(1) > li:nth-child(1) > a:nth-child(1)
	// New ID: #user-option-button
	userRow := "#users > tbody:nth-child(1) > tr:nth-child(" + strconv.Itoa(row) + ") > td:nth-child(1) > ul:nth-child(1) > li:nth-child(1) > a:nth-child(1)"
	p.MustElement(userRow).MustClick()
	// .open > ul:nth-child(2) > li:nth-child(2) > a:nth-child(1)
	// #edit-user-button
	// New ID for Edit User? (row above)
	p.MustElement(".open > ul:nth-child(2) > li:nth-child(2) > a:nth-child(1)").MustClick()

	// Fill out form
	// Modify Username Information
	// Username ID: #id_username
	p.MustElement("#id_username").MustSelectAllText().MustInput("Username4")

	// p.MustElement("#id_username").MustInput("Username4")

	// Modify First Name Information
	// First Name ID: #id_first_name
	p.MustElement("#id_first_name").MustSelectAllText().MustInput("First Name4")

	// Modify last Name Information
	// Last Name ID: #id_last_name
	p.MustElement("#id_last_name").MustSelectAllText().MustInput("Last Name4")

	// Modify Email Address Information
	// Email Address ID: #id_email
	p.MustElement("#id_email").MustSelectAllText().MustInput("emailaddress4@emailaddress.com")

	// Select desired status (Active Status selected as default)
	// Select, Keep, or Remove Staff Status
	// #id_is_staff
	// p.MustElement("#id_is_staff").MustClick()

	// Select, Keep, or Remove Super status (Selecting)
	// #id_is_superuser
	p.MustElement("#id_is_superuser").MustClick()

	// Submit changes
	// TODO: Iffy Selector
	// #base-content > form > div > div > input
	// New ID: #submit-button
	p.MustElement("#base-content > form > div > div > input").MustClick()

	fmt.Printf("PASS - TEST-U05 Updated User's Information in %+v\n", time.Since(started))

}
