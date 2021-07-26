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

	// Shorter name for sess.Page
	p := *sess.Page

	time.Sleep(time.Millisecond * 200)

	// Hover over user icon
	// #side-menu > li:nth-child(9) > a:nth-child(1)
	p.MustElement("#side-menu > li:nth-child(9) > a:nth-child(1)").MustHover()

	// Click on groups
	// #side-menu > li:nth-child(9) > ul:nth-child(2) > li:nth-child(2) > a:nth-child(1)
	p.MustElement("#side-menu > li:nth-child(9) > ul:nth-child(2) > li:nth-child(2) > a:nth-child(1)").MustClick()

	// Edit correct Group
	row := 0
	for j := 2; j <= 15; j++ {
		fmt.Println(j)
		// #groups > tbody:nth-child(1) > tr:nth-child(3) > td:nth-child(2)
		selector := "#groups > tbody:nth-child(1) > tr:nth-child(" + strconv.Itoa(j) + ") > td:nth-child(2)"
		fmt.Println(selector)
		name := p.MustElement(selector).MustText()
		fmt.Println(name)
		if name == "Test Group" {
			// Matched correct username
			fmt.Println("We matched")
			row = j
			j = 15
		}
	}
	fmt.Println("After the loop")
	// Edit correct Group
	// #groups > tbody:nth-child(1) > tr:nth-child(3) > td:nth-child(1) > ul:nth-child(1) > li:nth-child(1) > a:nth-child(1) > b:nth-child(1)
	groupRow := "#groups > tbody:nth-child(1) > tr:nth-child(" + strconv.Itoa(row) + ") > td:nth-child(1) > ul:nth-child(1) > li:nth-child(1) > a:nth-child(1) > b:nth-child(1)"
	p.MustElement(groupRow).MustClick()

	// Click on edit
	// .open > ul:nth-child(2) > li:nth-child(2) > a:nth-child(1)
	p.MustElement(".open > ul:nth-child(2) > li:nth-child(2) > a:nth-child(1)").MustClick()

	// Fill out name
	// #id_name
	p.MustElement("#id_name").MustSelectAllText().MustInput("Test Group 2")

	// Fill out description
	// .CodeMirror > div:nth-child(1) > textarea:nth-child(1)
	p.MustElement(".CodeMirror > div:nth-child(1) > textarea:nth-child(1)").MustSelectAllText().MustInput("Description for Test Group 2")

	// Click on submit button
	// input.btn
	p.MustElement("input.btn").MustClick()

	time.Sleep(time.Millisecond * 200)

	fmt.Printf("PASS - TEST-GR02 Edited selected group in %+v\n", time.Since(started))

}
