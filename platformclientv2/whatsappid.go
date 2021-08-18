package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Whatsappid - User information for a WhatsApp account
type Whatsappid struct { 
	// PhoneNumber - The phone number associated with this WhatsApp account
	PhoneNumber *Phonenumber `json:"phoneNumber,omitempty"`


	// DisplayName - The displayName of this person's account in WhatsApp
	DisplayName *string `json:"displayName,omitempty"`

}

func (u *Whatsappid) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Whatsappid

	

	return json.Marshal(&struct { 
		PhoneNumber *Phonenumber `json:"phoneNumber,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		*Alias
	}{ 
		PhoneNumber: u.PhoneNumber,
		
		DisplayName: u.DisplayName,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Whatsappid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
