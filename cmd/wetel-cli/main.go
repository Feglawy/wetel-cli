package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/Feglawy/wetel-cli/internal/app"
	"github.com/Feglawy/wetel-cli/internal/core"
	"github.com/Feglawy/wetel-cli/internal/models"
	"golang.org/x/term"
)

type LoginCredentials struct {
	number string
	pass   string
}

func askForLoginData() LoginCredentials {
	data := LoginCredentials{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Service number: ")
	number, _ := reader.ReadString('\n')
	data.number = strings.TrimSpace(number)

	fmt.Print("Password: ")
	bytePass, _ := term.ReadPassword(int(syscall.Stdin))
	data.pass = strings.TrimSpace(string(bytePass))
	fmt.Println() // For a clean newline after password

	return data
}

func (l *LoginCredentials) ConvServiceNum() error {
	num, err := strconv.Atoi(l.number)
	if err != nil {
		return fmt.Errorf("wrong format for the service number")
	}
	l.number = "FBB" + strconv.Itoa(num)
	return nil
}

func overview(user *models.User, balance float64, plan *models.Plan) {
	fmt.Printf("\nHello, %s\n", user.CustomerName)
	fmt.Printf("Current balance: %.2f\n", balance)
	fmt.Println("______________________________")
	fmt.Println("Usage overview")
	fmt.Printf("Used: %.2f GB\n", plan.Used())
	fmt.Printf("Remaining: %.2f GB\n", plan.Remain())
}

func offerUsageOverview(plan *models.Plan) {
	fmt.Println("______________________________")
	fmt.Println("Offers usage overview")
	for _, offer := range plan.Offers {
		fmt.Printf("Offer: %s\n", offer.Name)
		fmt.Printf("Total: %.2f GB\n", offer.InitialAmount)
		fmt.Printf("Remaining: %.2f GB\n", offer.CurrentAmount)
		fmt.Printf("Used: %.2f GB\n", offer.InitialAmount-offer.CurrentAmount)
		fmt.Println("---------------------------")
	}
}

func menu(app *app.App) {
	balance, err := core.GetBalance(app)
	if err != nil {
		fmt.Printf("Error fetching balance: %s\n", err)
		return
	}

	plan, err := core.GetPlans(app)
	if err != nil {
		fmt.Printf("Error fetching plans: %s\n", err)
		return
	}

	overview(&app.UserInfo, balance, plan)

	for {
		fmt.Print(`
______________________________
Menu:
1. Show detailed plan overview
2. Renew main plan
3. Refresh
4. Exit
>> `)
		var option int
		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch option {
		case 1:
			offerUsageOverview(plan)
		case 2:
			body, err := core.RenewMainOffer(app)
			if err != nil {
				fmt.Printf("Failed to renew: %s\n", err)
			} else {
				fmt.Println(body)
			}
		case 3:
			menu(app)
			return
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Unknown option. Please try again.")
		}
	}
}

func main() {
	app := app.NewApp()
	// env := config.GetEnv()

	var loginCredentials LoginCredentials

	serviceNum := *flag.String("num", "", "service number for login")
	password := *flag.String("pass", "", "password for login")
	flag.Parse()

	// serviceNum := env.ServiceNumber
	// password := env.Password

	if serviceNum == "" || password == "" {
		data := askForLoginData()
		loginCredentials = data
	} else {
		loginCredentials.number = serviceNum
		loginCredentials.pass = password
	}

	if err := loginCredentials.ConvServiceNum(); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	err := core.Login(loginCredentials.number, loginCredentials.pass, app)
	if err != nil {
		fmt.Printf("Login failed: %s\n", err)
		return
	}

	menu(app)
}
