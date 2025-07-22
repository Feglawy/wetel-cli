package app

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Feglawy/wetel-cli/internal/models"
	"github.com/Feglawy/wetel-cli/utils"
)

func (a *Client) HandleResponse(resp *http.Response, out models.Scannable) (string, error) {
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if utils.IsRespSuccessful(string(body)) && resp.StatusCode < 400 {
		if out != nil {
			out.ScanJson(string(body))
		}
	} else {
		return string(body), fmt.Errorf("error response:\n%s", utils.GetIndentedJson(body))
	}
	return string(body), nil
}
