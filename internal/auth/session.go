package auth

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/Feglawy/wetel-cli/config"
	"golang.org/x/term"
)

type LoginCredentials struct {
	Number string
	Pass   string
}

func (l *LoginCredentials) AskForLoginData() {

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
	f, err := os.Create(config.LOGIN_INFO_FILE)
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(info)
	if err != nil {
		return err
	}
	err = f.Sync()
	if err != nil {
		return err
	}
	return nil
}

func ClearLoginData() error {
	return os.Remove(config.LOGIN_INFO_FILE)
}

func RetriveLoginData() (*LoginCredentials, error) {
	file, err := os.Open(config.LOGIN_INFO_FILE)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var login LoginCredentials
	err = json.NewDecoder(file).Decode(&login)
	if err != nil {
		return nil, err
	}
	return &login, nil
}

func GetLoginData(serviceNumflag, passflag string) *LoginCredentials {
	var creds LoginCredentials

	// Check if we should use remembered credentials
	if serviceNumflag == "" && passflag == "" {
		remembered, _ := RetriveLoginData()
		if remembered != nil {
			fmt.Println("Using remembered credentials.")
			return remembered
		}
	}
	reader := bufio.NewReader(os.Stdin)
	if serviceNumflag == "" {
		fmt.Print("Service number: ")
		number, _ := reader.ReadString('\n')
		creds.Number = strings.TrimSpace(number)
	}
	if passflag == "" {
		fmt.Print("Password: ")
		bytePass, _ := term.ReadPassword(int(syscall.Stdin))
		creds.Pass = strings.TrimSpace(string(bytePass))
		fmt.Println() // For a clean newline after password
	}

	if err := creds.ConvServiceNum(); err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil
	}

	return &creds
}
