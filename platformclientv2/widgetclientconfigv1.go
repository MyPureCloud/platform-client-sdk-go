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

func (u *Widgetclientconfigv1) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Widgetclientconfigv1

	

	return json.Marshal(&struct { 
		WebChatSkin *string `json:"webChatSkin,omitempty"`
		
		AuthenticationUrl *string `json:"authenticationUrl,omitempty"`
		*Alias
	}{ 
		WebChatSkin: u.WebChatSkin,
		
		AuthenticationUrl: u.AuthenticationUrl,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Widgetclientconfigv1) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
