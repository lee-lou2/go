package oauth2

import (
	"fmt"
	"net/url"
	"os"
)

// GetRedirectUri 로그인 페이지로 리다이렉트할 URI를 반환
func GetRedirectUri() (string, error) {
	codeVerifier, err := getCodeVerifier()
	if err != nil {
		return "", err
	}
	codeChallenge := getCodeChallenge(codeVerifier)

	redirectURI := fmt.Sprintf(
		"%s/oauth2/authorize/?"+
			"response_type=code&"+
			"code_challenge=%s&"+
			"code_challenge_method=S256&"+
			"client_id=%s&"+
			"redirect_uri=%s&"+
			"state=%s",
		os.Getenv("OAUTH2_SSO_SERVER_HOST"),
		url.QueryEscape(codeChallenge),
		url.QueryEscape(os.Getenv("OAUTH2_CLIENT_ID")),
		url.QueryEscape(os.Getenv("OAUTH2_CALLBACK_URL")),
		url.QueryEscape(codeVerifier),
	)
	return redirectURI, nil
}
