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
	p.MustElement("#side-menu > li:nth-child(9) > a > i").MustClick()

	// Wait for page to load
	p.WaitLoad()

	// If username column equals correct username, then click on three dots to edit that user
	// Click on the three dots next to the user wanting to be changed
	// Admin ID: #users > tbody > tr:nth-child(2) > td:nth-child(4) > a

	row := 0
	for j := 2; j <= 15; j++ {
		fmt.Println(j)
		selector := "#users > tbody > tr:nth-child(" + strconv.Itoa(j) + ") > td:nth-child(4) > a"
		fmt.Println(selector)
		name := p.MustElement(selector).MustText()
		fmt.Println(name)
		if name == "Username" {
			// Matched correct username
			fmt.Println("We matched")
			row = j
			j = 15
		}
	}
	fmt.Println("After the loop")
	// #users > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(1) > ul:nth-child(1) > li:nth-child(1) > a:nth-child(1)
	// New ID: #edit-icon
	userRow := "#users > tbody:nth-child(1) > tr:nth-child(" + strconv.Itoa(row) + ") > td:nth-child(1) > ul:nth-child(1) > li:nth-child(1) > a:nth-child(1)"
	p.MustElement(userRow).MustClick()
	// .open > ul:nth-child(2) > li:nth-child(6) > a:nth-child(1)
	// #delete-button
	// New id for user edit? (row above)
	p.MustElement(".open > ul:nth-child(2) > li:nth-child(6) > a:nth-child(1)").MustClick()
	// #base-content > div > form > div > button
	// New ID: #confirm-deletion-button
	p.MustElement("#base-content > div > form > div > button").MustClick()

	fmt.Printf("PASS - TEST-U06 Removed user in %+v\n", time.Since(started))

}
