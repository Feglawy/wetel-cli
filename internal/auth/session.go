package auth

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"

	"golang.org/x/term"
)

type LoginCredentials struct {
	Number string
	Pass   string
}

func (l *LoginCredentials) AskForLoginData() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Service number: ")
	number, _ := reader.ReadString('\n')
	l.Number = strings.TrimSpace(number)

	fmt.Print("Password: ")
	bytePass, _ := term.ReadPassword(int(syscall.Stdin))
	l.Pass = strings.TrimSpace(string(bytePass))
	fmt.Println() // For a clean newline after password
}

func (l *LoginCredentials) ConvServiceNum() error {
	num, err := strconv.Atoi(l.Number)
	if err != nil {
		return fmt.Errorf("wrong format for the service number")
	}
	l.Number = "FBB" + strconv.Itoa(num)
	return nil
}

func StoreLoginData(info LoginCredentials) error {
	// if it doesn't exist create it
	// write the data in the file rewrite it
	return nil
}

func RetriveLoginData() (*LoginCredentials, error) {
	return nil, nil
}
