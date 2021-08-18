package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queuemessagingaddresses
type Queuemessagingaddresses struct { 
	// SmsAddress
	SmsAddress *Domainentityref `json:"smsAddress,omitempty"`

}

func (u *Queuemessagingaddresses) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queuemessagingaddresses

	

	return json.Marshal(&struct { 
		SmsAddress *Domainentityref `json:"smsAddress,omitempty"`
		*Alias
	}{ 
		SmsAddress: u.SmsAddress,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queuemessagingaddresses) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
