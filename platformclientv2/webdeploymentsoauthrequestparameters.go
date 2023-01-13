package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webdeploymentsoauthrequestparameters
type Webdeploymentsoauthrequestparameters struct { 
	// Code - The authorization code to be sent to the authentication server during the token request.  Refer to https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
	Code *string `json:"code,omitempty"`


	// RedirectUri - Redirect URI sent in the \"Authentication Request\"Refer to https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
	RedirectUri *string `json:"redirectUri,omitempty"`


	// Nonce - Required if provided in the \"Authentication Request\". Otherwise should be empty.String value used to associate a Client session with an ID Token, and to mitigate replay attacks. The value is passed through unmodified from the Authentication Request to the ID Token. Refer to https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
	Nonce *string `json:"nonce,omitempty"`


	// MaxAge - Required if provided in the  \"Authentication Request\". Otherwise should be empty.Specifies the allowable elapsed time in seconds since the last time the End-User was actively authenticated.Refer to https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
	MaxAge *int `json:"maxAge,omitempty"`


	// CodeVerifier - Required if authorizing using Proof Key for Code Exchange (PKCE). Otherwise should be empty.Random URL-safe string with a minimum length of 43 characters generated at start of authorization flow to mitigate the threat of having the authorization code intercepted. Refer to https://datatracker.ietf.org/doc/html/rfc7636
	CodeVerifier *string `json:"codeVerifier,omitempty"`


	// Iss - Optional parameter. Set it if authorization server discovery metadata authorization_response_iss_parameter_supported is enabled. Refer to https://datatracker.ietf.org/doc/html/rfc9207
	Iss *string `json:"iss,omitempty"`

}

func (o *Webdeploymentsoauthrequestparameters) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webdeploymentsoauthrequestparameters
	
	return json.Marshal(&struct { 
		Code *string `json:"code,omitempty"`
		
		RedirectUri *string `json:"redirectUri,omitempty"`
		
		Nonce *string `json:"nonce,omitempty"`
		
		MaxAge *int `json:"maxAge,omitempty"`
		
		CodeVerifier *string `json:"codeVerifier,omitempty"`
		
		Iss *string `json:"iss,omitempty"`
		*Alias
	}{ 
		Code: o.Code,
		
		RedirectUri: o.RedirectUri,
		
		Nonce: o.Nonce,
		
		MaxAge: o.MaxAge,
		
		CodeVerifier: o.CodeVerifier,
		
		Iss: o.Iss,
		Alias:    (*Alias)(o),
	})
}

func (o *Webdeploymentsoauthrequestparameters) UnmarshalJSON(b []byte) error {
	var WebdeploymentsoauthrequestparametersMap map[string]interface{}
	err := json.Unmarshal(b, &WebdeploymentsoauthrequestparametersMap)
	if err != nil {
		return err
	}
	
	if Code, ok := WebdeploymentsoauthrequestparametersMap["code"].(string); ok {
		o.Code = &Code
	}
    
	if RedirectUri, ok := WebdeploymentsoauthrequestparametersMap["redirectUri"].(string); ok {
		o.RedirectUri = &RedirectUri
	}
    
	if Nonce, ok := WebdeploymentsoauthrequestparametersMap["nonce"].(string); ok {
		o.Nonce = &Nonce
	}
    
	if MaxAge, ok := WebdeploymentsoauthrequestparametersMap["maxAge"].(float64); ok {
		MaxAgeInt := int(MaxAge)
		o.MaxAge = &MaxAgeInt
	}
	
	if CodeVerifier, ok := WebdeploymentsoauthrequestparametersMap["codeVerifier"].(string); ok {
		o.CodeVerifier = &CodeVerifier
	}
    
	if Iss, ok := WebdeploymentsoauthrequestparametersMap["iss"].(string); ok {
		o.Iss = &Iss
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Webdeploymentsoauthrequestparameters) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
