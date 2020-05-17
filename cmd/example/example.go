package main

import (
	"log"
	"time"

	"github.com/hsmtkk/poll_work_review/pkg/cred"
	"github.com/sclevine/agouti"
)

func main() {
	driver := agouti.ChromeDriver()
	defer driver.Stop()

	if err := driver.Start(); err != nil {
		log.Fatal(err)
	}

	page, err := driver.NewPage()
	if err != nil {
		log.Fatal(err)
	}

	userName, password, err := cred.New().LoadCredential()
	if err != nil {
		log.Fatal(err)
	}

	page.Navigate("https://techacademy.jp/mentor/login")

	email := page.FindByID("session_email")
	if err := email.SendKeys(userName); err != nil {
		log.Fatal(err)
	}

	passwordInput := page.FindByID("session_password")
	if err := passwordInput.SendKeys(password); err != nil {
		log.Fatal(err)
	}

	button := page.FindByName("commit")
	if err := button.Click(); err != nil {
		log.Fatal(err)
	}

	page.Navigate("https://techacademy.jp/mentor/all/reports")

	time.Sleep(5 * time.Second)
}
