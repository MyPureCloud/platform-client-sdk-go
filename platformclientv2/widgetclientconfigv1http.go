package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Widgetclientconfigv1http
type Widgetclientconfigv1http struct { 
	// WebChatSkin
	WebChatSkin *string `json:"webChatSkin,omitempty"`


	// AuthenticationUrl
	AuthenticationUrl *string `json:"authenticationUrl,omitempty"`

}

func (o *Widgetclientconfigv1http) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Widgetclientconfigv1http
	
	return json.Marshal(&struct { 
		WebChatSkin *string `json:"webChatSkin,omitempty"`
		
		AuthenticationUrl *string `json:"authenticationUrl,omitempty"`
		*Alias
	}{ 
		WebChatSkin: o.WebChatSkin,
		
		AuthenticationUrl: o.AuthenticationUrl,
		Alias:    (*Alias)(o),
	})
}

func (o *Widgetclientconfigv1http) UnmarshalJSON(b []byte) error {
	var Widgetclientconfigv1httpMap map[string]interface{}
	err := json.Unmarshal(b, &Widgetclientconfigv1httpMap)
	if err != nil {
		return err
	}
	
	if WebChatSkin, ok := Widgetclientconfigv1httpMap["webChatSkin"].(string); ok {
		o.WebChatSkin = &WebChatSkin
	}
    
	if AuthenticationUrl, ok := Widgetclientconfigv1httpMap["authenticationUrl"].(string); ok {
		o.AuthenticationUrl = &AuthenticationUrl
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Widgetclientconfigv1http) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
