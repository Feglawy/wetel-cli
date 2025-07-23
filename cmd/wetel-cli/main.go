package main

import (
	"flag"
	"fmt"

	"github.com/Feglawy/wetel-cli/internal/auth"
	"github.com/Feglawy/wetel-cli/internal/menu"
	"github.com/Feglawy/wetel-cli/pkg/app"
	"github.com/Feglawy/wetel-cli/pkg/core"
)

func main() {
	app := app.NewClient()
	// env := config.GetEnv()

	var loginCredentials *auth.LoginCredentials

	serviceNum := flag.String("num", "", "service number for login e.g 0238900000")
	password := flag.String("pass", "", "password for login")
	remember := flag.Bool("r", false, "remember the login creds")
	flag.Parse()

	loginCredentials = auth.GetLoginData(*serviceNum, *password)
	if loginCredentials == nil {
		return
	}
	err := core.Login(loginCredentials.Number, loginCredentials.Pass, app)
	if err != nil {
		fmt.Printf("Login failed: %s\n", err)
		return
	}
	fmt.Println("Login successful!")
	
	if *remember {
		err := auth.StoreLoginData(*loginCredentials)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}

	menu.Menu(app)
}
