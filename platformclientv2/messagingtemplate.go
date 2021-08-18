package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagingtemplate - The messaging template identifies a structured message templates supported by a messaging channel.
type Messagingtemplate struct { 
	// WhatsApp - Defines a messaging template for a WhatsApp messaging channel
	WhatsApp *Whatsappdefinition `json:"whatsApp,omitempty"`

}

func (u *Messagingtemplate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagingtemplate

	

	return json.Marshal(&struct { 
		WhatsApp *Whatsappdefinition `json:"whatsApp,omitempty"`
		*Alias
	}{ 
		WhatsApp: u.WhatsApp,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Messagingtemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
