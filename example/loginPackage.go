package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	ddl "github.com/DefectDojo/ui-qa-experiment/login"
)

// Good references:
// https://go-rod.github.io/#/get-started/README
// https://github.com/go-rod/rod/blob/master/examples_test.go
// https://pkg.go.dev/github.com/go-rod/rod#Mouse

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

	// We should now be on the main DefectDojo page aka /dashboard
	// Click on the user side menu - #side-menu > li:nth-child(9) > a > i
	p.MustElement("#side-menu > li:nth-child(9) > a > i").MustClick()

	// Click on the wrench - #dropdownMenu1 > span.fa.fa-wrench
	p.MustElement("#dropdownMenu1 > span.fa.fa-wrench").MustClick()

	// Click on "Add user" - #base-content > div > div > div:nth-child(1) > div.panel-heading.tight > h3 > div > ul > li > a
	p.MustElement("#base-content > div > div > div:nth-child(1) > div.panel-heading.tight > h3 > div > ul > li > a").MustClick()

	// Wait for page to load
	p.WaitLoad()

	// Fill out the User form
	p.MustElement("#id_username").MustInput("b.ross-da-boss")
	p.MustElement("#id_first_name").MustInput("Bob")
	p.MustElement("#id_last_name").MustInput("Ross")
	p.MustElement("#id_email").MustInput("bob.ross@happytrees.com")
	p.MustElement("#id_is_staff").MustClick()
	// Click on the form's button
	p.MustElement("#base-content > form > div > div > input").MustClick()

	// Wait for page to load
	p.WaitLoad()

	// Wait a bit more for page to update
	time.Sleep(time.Millisecond * 600)

	// Get new user ID
	pageInfo, err := p.Info()
	if err != nil {
		fmt.Printf("Error getting page info was:\n%+v\n", err)
		fmt.Println("FAILED - Add User")
		os.Exit(1)
	}

	// If the URL looks like this, something bad happened - https://demo.defectdojo.org/user/add
	if strings.Contains(pageInfo.URL, "/user/add") {
		fmt.Println("Creating the user unsuccessful - likely the user already exits")
		fmt.Println("FAILED - Add User")
		os.Exit(1)
	}

	// Extract the new user's ID from the page
	uid, err := userFromURL(pageInfo.URL)
	if err != nil {
		fmt.Printf("Error getting the user's ID from the URL was:\n%+v\n", err)
		fmt.Println("FAILED - Add User")
		os.Exit(1)
	}

	//fmt.Printf("Page info is:\n\t%+v\n", pageInfo)
	fmt.Printf("uid is %+v\n", uid)

	// User created, now set the user's password via Django admin

	// Click on the link to the Django Admin page:
	p.MustElement("div.alert:nth-child(2) > a:nth-child(1)").MustClick()

	// Wait for page to load
	time.Sleep(time.Millisecond * 300)
	//p.WaitLoad()

	// Click on the form to set a password
	p.MustElement(".field-password > div:nth-child(1) > div:nth-child(3) > a:nth-child(1)").MustClick()

	// Wait for page to load
	time.Sleep(time.Millisecond * 300)
	//p.WaitLoad()

	// Fill out the two password fields
	p.MustElement("#id_password1").MustInput("p41nt3r$")
	p.MustElement("#id_password2").MustInput("p41nt3r$")

	// Click on change password button
	p.MustElement(".default").MustClick()

	fmt.Printf("PASS - Add User in %+v\n", time.Since(started))
	os.Exit(0)
}

func userFromURL(rawURL string) (uint64, error) {
	// Take the URL for a user detail page and get the user's ID
	// e.g. http://localhost:8888/user/3/edit
	u, err := url.Parse(rawURL)
	if err != nil {
		fmt.Printf("Error parsing URL - %+v was:\n\t%+v\n", rawURL, err)
		return 0, err
	}
	uidStr := strings.Replace(strings.Replace(u.Path, "/user/", "", 1), "/edit", "", 1)

	id, err := strconv.ParseUint(uidStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}
