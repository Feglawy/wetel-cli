package menu

import (
	"fmt"
	"sync"

	"github.com/Feglawy/wetel-cli/internal/models"
	"github.com/Feglawy/wetel-cli/pkg/app"
	"github.com/Feglawy/wetel-cli/pkg/core"
)

type result[T any] struct {
	Val T
	err error
}

func runTask[T any](
	app *app.Client,
	fn func(a *app.Client) (val T, err error),
	wg *sync.WaitGroup,
	ch chan<- result[T],
) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		val, err := fn(app)
		ch <- result[T]{val, err}
	}()
}

func Menu(app *app.Client) {
	wg := sync.WaitGroup{}

	balanceChan := make(chan result[float64], 1)
	planChan := make(chan result[*models.Plan], 1)

	runTask(app, core.GetBalance, &wg, balanceChan)
	runTask(app, core.GetPlans, &wg, planChan)

	wg.Wait()
	balanceRes := <-balanceChan
	planRes := <-planChan

	if balanceRes.err != nil {
		fmt.Printf("Error fetching balance: %s\n", balanceRes.err)
		return
	}
	if planRes.err != nil {
		fmt.Printf("Error fetching plans: %s\n", planRes.err)
		return
	}
	overview(&app.UserInfo, balanceRes.Val, planRes.Val)

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
			offerUsageOverview(planRes.Val)
		case 2:
			body, err := core.RenewMainOffer(app)
			if err != nil {
				fmt.Printf("Failed to renew: %s\n", err)
			} else {
				fmt.Println(body)
			}
		case 3:
			Menu(app)
			return
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Unknown option. Please try again.")
		}
	}
}

func overview(user *models.User, balance float64, plan *models.Plan) {
	fmt.Printf("\nHello, %s\n", user.CustomerName)
	fmt.Printf("Current balance: %.2f\n", balance)

	fmt.Println("______________________________")

	fmt.Println("Usage overview")
	fmt.Printf("Used: %.2f GB\n", plan.Used())
	fmt.Printf("Remaining: %.2f GB\n", plan.Remain())
	fmt.Printf("Remaining days: %v", plan.GetRemainingDays())
}

func offerUsageOverview(plan *models.Plan) {
	fmt.Println("______________________________")
	fmt.Println("Offers usage overview")
	for _, offer := range plan.Offers {
		fmt.Printf("Offer: %s\n", offer.Name)
		fmt.Printf("Remaining days: %v\n", offer.RemainingDays)

		fmt.Printf("Total: %.2f GB\n", offer.InitialAmount)
		fmt.Printf("Used: %.2f GB\n", offer.InitialAmount-offer.CurrentAmount)
		fmt.Printf("Remaining: %.2f GB\n", offer.CurrentAmount)
		fmt.Println("---------------------------")
	}
}
