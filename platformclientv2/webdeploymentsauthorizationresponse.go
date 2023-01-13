package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webdeploymentsauthorizationresponse
type Webdeploymentsauthorizationresponse struct { 
	// RefreshToken - Refresh token used to issue a new JWT.
	RefreshToken *string `json:"refreshToken,omitempty"`


	// Jwt
	Jwt *string `json:"jwt,omitempty"`

}

func (o *Webdeploymentsauthorizationresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webdeploymentsauthorizationresponse
	
	return json.Marshal(&struct { 
		RefreshToken *string `json:"refreshToken,omitempty"`
		
		Jwt *string `json:"jwt,omitempty"`
		*Alias
	}{ 
		RefreshToken: o.RefreshToken,
		
		Jwt: o.Jwt,
		Alias:    (*Alias)(o),
	})
}

func (o *Webdeploymentsauthorizationresponse) UnmarshalJSON(b []byte) error {
	var WebdeploymentsauthorizationresponseMap map[string]interface{}
	err := json.Unmarshal(b, &WebdeploymentsauthorizationresponseMap)
	if err != nil {
		return err
	}
	
	if RefreshToken, ok := WebdeploymentsauthorizationresponseMap["refreshToken"].(string); ok {
		o.RefreshToken = &RefreshToken
	}
    
	if Jwt, ok := WebdeploymentsauthorizationresponseMap["jwt"].(string); ok {
		o.Jwt = &Jwt
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Webdeploymentsauthorizationresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
