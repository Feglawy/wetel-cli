package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/Feglawy/wetel-cli/config"
	"github.com/Feglawy/wetel-cli/internal/api"
	"github.com/Feglawy/wetel-cli/internal/auth"
	"github.com/Feglawy/wetel-cli/internal/models"
	"github.com/Feglawy/wetel-cli/internal/ui"
	"github.com/Feglawy/wetel-cli/pkg/app"
	"github.com/Feglawy/wetel-cli/pkg/core"
	"github.com/Feglawy/wetel-cli/utils"
)

func main() {
	// Parse CLI flags
	serviceNum := flag.String("num", "", "service number for login e.g 0238900000")
	password := flag.String("pass", "", "password for login")
	remember := flag.Bool("r", false, "remember the login creds")
	flag.Parse()

	fmt.Println(config.LOGO)
	run(serviceNum, password, remember)
}

func run(serviceNum, password *string, remember *bool) {
	// Initialize core components
	client := app.NewClient()
	apiHandler := api.NewAPI(client)
	coreHandler := core.NewCore(apiHandler)

	// Get login data
	loginCredentials := auth.GetLoginData(*serviceNum, *password)
	if loginCredentials == nil {
		return
	}

	// Attempt login
	if err := coreHandler.Login(loginCredentials.Number, loginCredentials.Pass); err != nil {
		fmt.Printf("Login failed: %s\n", err)
		return
	}
	fmt.Println("Login successful!")

	// Optionally store credentials
	if *remember {
		if err := auth.StoreLoginData(*loginCredentials); err != nil {
			fmt.Printf("Error saving credentials: %s\n", err)
			return
		}
	}

	wg := sync.WaitGroup{}
	balanceChan := make(chan utils.Result[float64], 1)
	planChan := make(chan utils.Result[*models.Plan], 1)

	utils.RunTask(coreHandler.GetBalance, client.GetUserInfo().AccountId, &wg, balanceChan)
	utils.RunTask(coreHandler.GetPlans, client.GetUserInfo().SubscriberId, &wg, planChan)

	wg.Wait()

	balanceRes := <-balanceChan
	planRes := <-planChan

	if balanceRes.Err != nil {
		fmt.Printf("Error fetching balance: %s\n", balanceRes.Err)
		return
	}
	if planRes.Err != nil {
		fmt.Printf("Error fetching plans: %s\n", planRes.Err)
		return
	}

	userInfo := client.GetUserInfo()
	ui.Overview(&userInfo, balanceRes.Val, planRes.Val)

	// Start main menu loop
	for {
		switch showMenuAndGetChoice() {
		case DETAILED_PLANS:
			ui.OfferUsageOverview(planRes.Val)
		case RENEW_MAIN_PLAN:
			msg, err := coreHandler.RenewMainOffer(userInfo.ServNumber, userInfo.SubscriberId)
			if err != nil {
				fmt.Printf("an error happened during renewing the main plan err: %v \n", err)
			}
			fmt.Println(msg)
		case SUBSCRIBE_TO_ADDON:
			addonOffers, err := coreHandler.GetAddonOffers(userInfo.ServNumber)
			if err != nil {
				fmt.Printf("an error happed while getting the addons offers err: %v", err)
			} else {
				addonOffer := ui.ChooseAnADDON(addonOffers)
				if addonOffer != nil {
					msg, err := coreHandler.SubscribeToAPlan(userInfo.SubscriberId, *addonOffer)
					fmt.Println(msg)
					if err != nil {
						fmt.Printf("an error happed while subscribing to the addon offer err: %v", err)
					}
				}
			}
		case TOGGLE_REMEMBER_ME:
			if _, err := auth.RetriveLoginData(); err != nil {
				auth.StoreLoginData(*loginCredentials)
				fmt.Println("Your credentials has been stored")
			} else {
				auth.ClearLoginData()
				fmt.Println("Your credentials has been cleared")
			}
		case REFRESH:
			fmt.Println("Refreshing...")
			run(serviceNum, password, remember)
			return
		case EXIT:
			fmt.Println("Good Bye!")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}

type options int

const (
	DETAILED_PLANS options = iota + 1
	RENEW_MAIN_PLAN
	SUBSCRIBE_TO_ADDON
	TOGGLE_REMEMBER_ME
	REFRESH
	EXIT
)

func showMenuAndGetChoice() options {
	fmt.Println()
	fmt.Println("Menu:")
	fmt.Println("1. Show detailed plan overview")
	fmt.Println("2. Renew main plan")
	fmt.Println("3. Subscribe to an AddOn")
	fmt.Println("4. Toggle remember me")
	fmt.Println("5. Refresh")
	fmt.Println("6. Exit")

	var opt int
	fmt.Print("Enter choice: ")
	_, err := fmt.Scan(&opt)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return 0
	}
	return options(opt)
}
