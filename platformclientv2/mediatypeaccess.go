package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Mediatypeaccess - Media type access definitions
type Mediatypeaccess struct { 
	// Inbound - List of media types allowed for inbound messages from customers. If inbound messages from a customer contain media that is not in this list, the media will be dropped from the outbound message.
	Inbound *[]Mediatype `json:"inbound,omitempty"`


	// Outbound - List of media types allowed for outbound messages to customers. If an outbound message is sent that contains media that is not in this list, the message will not be sent.
	Outbound *[]Mediatype `json:"outbound,omitempty"`

}

func (u *Mediatypeaccess) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mediatypeaccess

	

	return json.Marshal(&struct { 
		Inbound *[]Mediatype `json:"inbound,omitempty"`
		
		Outbound *[]Mediatype `json:"outbound,omitempty"`
		*Alias
	}{ 
		Inbound: u.Inbound,
		
		Outbound: u.Outbound,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Mediatypeaccess) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
