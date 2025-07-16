package payload

type AuthPayload struct {
	ServiceNum      string `json:"acctId"`
	Password        string `json:"password"`
	AppLocale       string `json:"appLocale"`
	IsSelfcare      string `json:"isSelfcare"`
	IsMobile        string `json:"isMobile"`
	RecaptchaTocken string `json:"recaptchaToken"`
}

func NewAuthPayload(ServiceNum, password string) *AuthPayload {
	return &AuthPayload{
		ServiceNum:      ServiceNum,
		Password:        password,
		AppLocale:       "en-US",
		IsSelfcare:      "Y",
		IsMobile:        "N",
		RecaptchaTocken: "",
	}
}
