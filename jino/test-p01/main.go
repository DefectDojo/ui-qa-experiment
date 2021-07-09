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

	// hover over product icon
	// TODO: Iffy Selector
	// #side-menu > li:nth-child(2) > a:nth-child(1)
	p.MustElement("#side-menu > li:nth-child(2) > a:nth-child(1)").Hover()

	//click on "Add Product"
	// TODO: Bad Selector
	// #side-menu > li:nth-child(2) > ul:nth-child(2) > li:nth-child(2) > a:nth-child(1)
	p.MustElement("#side-menu > li:nth-child(2) > ul:nth-child(2) > li:nth-child(2) > a:nth-child(1)").MustClick()

	// input product name
	//TODO: Bad Selector
	//#id_name
	p.MustElement("#id_name").MustInput("test-product")

	// input product description
	//TODO: Bad Selector
	//.CodeMirror > div:nth-child(1) > textarea:nth-child(1)
	p.MustElement(".CodeMirror > div:nth-child(1) > textarea:nth-child(1)").MustInput("made for test")

	// click on product type bar
	// TODO: Bad Selector
	// #id_prod_type
	p.MustElement("#id_prod_type").MustClick()
	fmt.Println("after must click")

	//click on "Research and Development"
	// TODO: Bad Selector
	// #id_prod_type > option:nth-child(4)
	//p.MustElement("#id_prod_type > option:nth-child(4)").MustClick()
	//#id_prod_type
	nl := '\t'
	p.MustElement("#id_prod_type").MustInput("Research and Development").MustPress(nl)
	fmt.Println("after r and d click")

	time.Sleep(time.Minute * 4)
	//click on user icon
	// TODO: Iffy selector
	//#side-menu > li:nth-child(9) > a:nth-child(1) > i:nth-child(1)
	//	p.MustElement("#side-menu > li:nth-child(9) > a:nth-child(1) > i:nth-child(1)").MustClick()

	fmt.Printf("PASS - pt06 ++++++Add Product Type++++ in %+v\n", time.Since(started))

	///home/jinot/Documents/Go-Programs/ui-qa-experiment/jino/test-p01/main.go:59 +0x437
	//move must press to new line
}
