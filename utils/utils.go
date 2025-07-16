package utils

import (
	"encoding/json"
	"net/http"

	"github.com/tidwall/gjson"
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
