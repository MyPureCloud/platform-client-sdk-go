package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Widgetclientconfigv1
type Widgetclientconfigv1 struct { 
	// WebChatSkin
	WebChatSkin *string `json:"webChatSkin,omitempty"`


	// AuthenticationUrl
	AuthenticationUrl *string `json:"authenticationUrl,omitempty"`

}

func (o *Widgetclientconfigv1) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Widgetclientconfigv1
	
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

func (o *Widgetclientconfigv1) UnmarshalJSON(b []byte) error {
	var Widgetclientconfigv1Map map[string]interface{}
	err := json.Unmarshal(b, &Widgetclientconfigv1Map)
	if err != nil {
		return err
	}
	
	if WebChatSkin, ok := Widgetclientconfigv1Map["webChatSkin"].(string); ok {
		o.WebChatSkin = &WebChatSkin
	}
    
	if AuthenticationUrl, ok := Widgetclientconfigv1Map["authenticationUrl"].(string); ok {
		o.AuthenticationUrl = &AuthenticationUrl
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Widgetclientconfigv1) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
