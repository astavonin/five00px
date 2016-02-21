package five00px

import (
	"fmt"
	"testing"
)

var (
	oauthResponce         string = "/?oauth_token=0U3iDlwSGMRAx0XiejCNgCtMPVZ7yLVRr4XPHdYF&oauth_verifier=vNjuWrmN3mewcAnPIUAD"
	oauthTokenExpected    string = "0U3iDlwSGMRAx0XiejCNgCtMPVZ7yLVRr4XPHdYF"
	oauthVerifierExpected string = "vNjuWrmN3mewcAnPIUAD"
)

func TestOAuthResponseParsing(t *testing.T) {
	oauthToken, oauthVerifier, err := parseAccessToken(oauthResponce)
	if err != nil {
		t.Fatal(err)
	}
	if oauthTokenExpected != oauthToken {
		t.Errorf(fmt.Sprintf("Invalid toket, expectd value %s, real value %s",
			oauthTokenExpected, oauthToken))
	}
	if oauthVerifierExpected != oauthVerifier {
		t.Errorf(fmt.Sprintf("Invalid verifier, expected value %s, real value %s",
			oauthVerifierExpected, oauthVerifier))
	}
}
