package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textboterrorinputevent - Settings for an input event to the bot flow indicating an error has occurred.
type Textboterrorinputevent struct { 
	// Code - The error code.
	Code *string `json:"code,omitempty"`


	// Message - The error message.
	Message *string `json:"message,omitempty"`

}

func (u *Textboterrorinputevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textboterrorinputevent

	

	return json.Marshal(&struct { 
		Code *string `json:"code,omitempty"`
		
		Message *string `json:"message,omitempty"`
		*Alias
	}{ 
		Code: u.Code,
		
		Message: u.Message,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Textboterrorinputevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
