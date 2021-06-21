package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-rod/rod"
)

// Good references:
// https://go-rod.github.io/#/get-started/README
// https://github.com/go-rod/rod/blob/master/examples_test.go
// https://pkg.go.dev/github.com/go-rod/rod#Mouse

func main() {
	// Connect to the demo server
	page := rod.New().MustConnect().MustPage("https://demo.defectdojo.org/")

	// Login via the login page
	page.MustElement("#id_username").MustInput("admin")
	page.MustElement("#id_password").MustInput("defectdojo@demo#appsec")
	fmt.Println("Before login button")
	// CSS selector without the login banner turned on
	// #base-content > form > fieldset > div:nth-child(3) > div.col-sm-offset-1.col-sm-1 > button
	page.MustElement("#base-content > form > fieldset > div:nth-child(4) > div.col-sm-offset-1.col-sm-1 > button").MustClick()
	fmt.Println("After login button")

	// Click on the user side menu - #side-menu > li:nth-child(9) > a > i
	page.MustElement("#side-menu > li:nth-child(9) > a > i").MustClick()

	// Click on the wrench - #dropdownMenu1 > span.fa.fa-wrench
	page.MustElement("#dropdownMenu1 > span.fa.fa-wrench").MustClick()

	// Click on "Add user" - #base-content > div > div > div:nth-child(1) > div.panel-heading.tight > h3 > div > ul > li > a
	page.MustElement("#base-content > div > div > div:nth-child(1) > div.panel-heading.tight > h3 > div > ul > li > a").MustClick()

	// Fill out the User form
	// Username - #id_username
	page.MustElement("#id_username").MustInput("mario")
	page.MustElement("#id_first_name").MustInput("Mario")
	page.MustElement("#id_last_name").MustInput("Plumber")
	page.MustElement("#id_email").MustInput("mario@pipeline.pvt")
	page.MustElement("#id_is_staff").MustClick()
	// Click on the form's button
	page.MustElement("#base-content > form > div > div > input").MustClick()

	// Get new user ID
	usrURL, err := page.Info()
	if err != nil {
		fmt.Printf("Error getting page info was:\n%+v\n", err)
		os.Exit(1)
	}
	uid, err := userFromURL(usrURL.URL)
	if err != nil {
		fmt.Printf("Error getting the user's ID from the URL was:\n%+v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Page info is:\n\t%+v\n", usrURL)
	fmt.Printf("uid is %+v\n", uid)

	// TODO: Set the user's password via Django admin
	//admin := rod.New().MustConnect().MustPage("http://localhost:8888/admin/auth/user/2/change/")
	// #side-menu > li:nth-child(2) > ul > li:nth-child(1) > a
	// #nav-minimize-menu-li
	//page.MustElement("#nav-minimize-menu-li").MustClick()
	// #side-menu > li:nth-child(2) > a > span:nth-child(2)
	//page.MustElement("#side-menu > li:nth-child(2) > a > span:nth-child(2)").MustClick()
	//page.MustElement("#side-menu > li:nth-child(2) > ul > li:nth-child(1) > a").MustClick()
	//page.MustElementR("a", "Add Product").MustClick()

	time.Sleep(time.Hour)
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
