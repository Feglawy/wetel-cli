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

	var loginCredentials auth.LoginCredentials

	serviceNum := flag.String("num", "", "service number for login e.g 0238900000")
	password := flag.String("pass", "", "password for login")
	rememberMe := flag.Bool("r", false, "remember me option")
	flag.Parse()

	if *serviceNum == "" || *password == "" { // ask for login credentials
		loginCredentials.AskForLoginData()
	} else { // the user passed login credentials as arguments
		loginCredentials.Number = *serviceNum
		loginCredentials.Pass = *password
	}

	if err := loginCredentials.ConvServiceNum(); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	err := core.Login(loginCredentials.Number, loginCredentials.Pass, app)
	if err != nil {
		fmt.Printf("Login failed: %s\n", err)
		return
	}

	if *rememberMe {
		// if the file that stores the data is not created
		// store the login credintials in a json file

	}

	menu.Menu(app)
}
