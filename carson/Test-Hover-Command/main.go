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

	// Wait for the page to load
	p.WaitLoad()

	// Hover over Engangement Menu
	// TODO: Iffy Selector
	// #side-menu > li:nth-child(3) > a
	p.MustElement("#side-menu > li:nth-child(3) > a").Hover()

	// Click on All Engangements
	// TODO: Bad Selector
	// #side-menu > li:nth-child(3) > ul > li:nth-child(2) > a
	p.MustElement("#side-menu > li:nth-child(3) > ul > li:nth-child(2) > a").MustClick()

	fmt.Printf("PASS - .Hover() Test in %+v\n", time.Since(started))

}
