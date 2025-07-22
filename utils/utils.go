package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/tidwall/gjson"
	"golang.org/x/term"
)

const (
	SUCCESSFUL_RETCODE string = "0"
)

func SetHeaders(req *http.Request, csrfToken string) {
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("csrftoken", csrfToken)
	req.Header.Set("languageCode", "en-US")
	req.Header.Set("isMobile", "false")
	req.Header.Set("isCoporate", "false")
	req.Header.Set("isSelfcare", "true")
	req.Header.Set("channelId", "702")
}

func IsRespSuccessful(Json string) bool {
	res := gjson.Get(Json, "header.retCode").String()
	return res == SUCCESSFUL_RETCODE
}

func GetIndentedJson(raw []byte) string {
	var obj map[string]interface{}
	json.Unmarshal(raw, &obj)
	jsonRaw, _ := json.MarshalIndent(obj, "", "	")
	return string(jsonRaw)

}

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
