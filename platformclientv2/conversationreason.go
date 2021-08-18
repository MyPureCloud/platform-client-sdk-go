package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationreason - Reasons for a failed message receipt.
type Conversationreason struct { 
	// Code - The reason code for the failed message receipt.
	Code *string `json:"code,omitempty"`


	// Message - Description of the reason for the failed message receipt.
	Message *string `json:"message,omitempty"`

}

func (u *Conversationreason) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationreason

	

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
func (o *Conversationreason) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
