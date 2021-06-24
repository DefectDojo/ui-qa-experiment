package login

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/go-rod/rod"
)

type DDLogin struct {
	BaseURL  *url.URL     // The base URL for the DefectDojo install
	Username string       // The user to login with
	Password string       // The password for the user logging in
	BannerOn bool         // True for login banner being on, false when not
	Browser  *rod.Browser // Rod browser for this login session
	Page     *rod.Page    // Rod browser page
	Debug    bool         // True causes a 5 second delay before clicking the login button
}

// SetURL takes a string and creates a url.URL struct
// returns an error if the URL provided cannot be parsed successfully
// String should be a URL like https://demo.defectdojo.org
func (l *DDLogin) SetURL(u string) error {
	validURL, err := url.Parse(u)
	if err != nil {
		return err
	}

	l.BaseURL = validURL

	return nil
}

// SetUsername taks a string and sets the username to use for logging into DefectDojo
func (l *DDLogin) SetUsername(usr string) {
	l.Username = usr
}

// SetPassword takes a string and sets the password to use for logging into DefectDojo
func (l *DDLogin) SetPassword(p string) {
	l.Password = p
}

// SetBanner takes a boolean (true/false) to set if the login banner is on (true) or
// off (false)
func (l *DDLogin) SetBanner(b bool) {
	l.BannerOn = b
}

// SetDebug takes a boolean.  If set to true, it will pause for 5 seconds before it submits
// the provided credentials allowing you to see the login screen and optionally change the
// dev tools layout
func (l *DDLogin) SetDebug(d bool) {
	l.Debug = d
}

// NewSession logs into DefectDojo and creates a new Rod session to use to 'browse'
// DefectDojo
func (l *DDLogin) NewSession() error {
	// Make sure we have the necessary struct values
	if len(l.BaseURL.Host) < 3 {
		return errors.New("DefectDojo URL is empty or too short")
	}
	if len(l.Username) <= 0 {
		return errors.New("DefectDojo username to login with is empty")
	}
	if len(l.Password) <= 0 {
		return errors.New("DefectDojo password to login with is empty")
	}

	// Create a new Rod 'browser' session
	fmt.Print(l.BaseURL.String())
	l.Page = rod.New().MustConnect().MustPage(l.BaseURL.String())

	// TODO: Incomplete experiement below
	// Attempt to avoid the panics of MustX calls
	//l.Browser = rod.New()
	//err := l.Browser.Connect()
	//if err != nil {
	//	return err
	//}

	//// Set the initial URL to open in the browser
	//var tar proto.TargetCreateTarget
	//tar.URL = l.stringURL()
	//p, err := l.Browser.Page(tar)
	//if err != nil {
	//	return err
	//}
	//l.Browser.Page = &p

	//// Enter credentials into login page
	l.Page.MustElement("#id_username").MustInput("admin")
	l.Page.MustElement("#id_password").MustInput("defectdojo@demo#appsec")

	// Optionally pause for debugging
	if l.Debug {
		time.Sleep(time.Second * 5)
	}

	// Click the login button
	// TODO: Add a proper CSS selector for the login botton so this isn't necessary
	if l.BannerOn {
		// CSS selector with the login banner turned on
		l.Page.MustElement("#base-content > form > fieldset > div:nth-child(4) > div.col-sm-offset-1.col-sm-1 > button").MustClick()
	} else {
		// CSS selector without the login banner turned off
		l.Page.MustElement("#base-content > form > fieldset > div:nth-child(3) > div.col-sm-offset-1.col-sm-1 > button").MustClick()
	}

	return nil
}

// SetAndLogin is used to set the necessary values to login to DefectDojo using Rod.
// It takes three strings: a URL (u), a username (usr) and a password (p) and two
// booleans: b should be set to true if the login banner is on and d should be set
// to true if you want a 5 second delay before clicking the login button (for debug)
func (l *DDLogin) SetAndLogin(u, usr, p string, b, d bool) error {
	l.SetURL(u)
	l.SetUsername(usr)
	l.SetPassword(p)
	l.SetBanner(b)
	l.SetDebug(d)
	err := l.NewSession()
	if err != nil {
		return err
	}
	return nil
}
