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

func (o *Whatsappid) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Whatsappid
	
	return json.Marshal(&struct { 
		PhoneNumber *Phonenumber `json:"phoneNumber,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		*Alias
	}{ 
		PhoneNumber: o.PhoneNumber,
		
		DisplayName: o.DisplayName,
		Alias:    (*Alias)(o),
	})
}

func (o *Whatsappid) UnmarshalJSON(b []byte) error {
	var WhatsappidMap map[string]interface{}
	err := json.Unmarshal(b, &WhatsappidMap)
	if err != nil {
		return err
	}
	
	if PhoneNumber, ok := WhatsappidMap["phoneNumber"].(map[string]interface{}); ok {
		PhoneNumberString, _ := json.Marshal(PhoneNumber)
		json.Unmarshal(PhoneNumberString, &o.PhoneNumber)
	}
	
	if DisplayName, ok := WhatsappidMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Whatsappid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
