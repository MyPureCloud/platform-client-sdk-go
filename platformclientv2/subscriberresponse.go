package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Subscriberresponse
type Subscriberresponse struct { 
	// MessageReturned - Suggested valid addresses
	MessageReturned *[]string `json:"messageReturned,omitempty"`


	// Status - http status
	Status *string `json:"status,omitempty"`

}

func (u *Subscriberresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Subscriberresponse

	

	return json.Marshal(&struct { 
		MessageReturned *[]string `json:"messageReturned,omitempty"`
		
		Status *string `json:"status,omitempty"`
		*Alias
	}{ 
		MessageReturned: u.MessageReturned,
		
		Status: u.Status,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Subscriberresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
