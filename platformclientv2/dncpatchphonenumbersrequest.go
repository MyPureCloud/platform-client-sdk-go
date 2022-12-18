package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dncpatchphonenumbersrequest
type Dncpatchphonenumbersrequest struct { 
	// Action - The action to perform
	Action *string `json:"action,omitempty"`


	// PhoneNumbers - The list of phone numbers to Add to / Remove from the DNC list 
	PhoneNumbers *[]string `json:"phoneNumbers,omitempty"`


	// ExpirationDateTime - Expiration date for DNC phone numbers in yyyy-MM-ddTHH:mmZ format
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`

}

func (o *Dncpatchphonenumbersrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dncpatchphonenumbersrequest
	
	return json.Marshal(&struct { 
		Action *string `json:"action,omitempty"`
		
		PhoneNumbers *[]string `json:"phoneNumbers,omitempty"`
		
		ExpirationDateTime *string `json:"expirationDateTime,omitempty"`
		*Alias
	}{ 
		Action: o.Action,
		
		PhoneNumbers: o.PhoneNumbers,
		
		ExpirationDateTime: o.ExpirationDateTime,
		Alias:    (*Alias)(o),
	})
}

func (o *Dncpatchphonenumbersrequest) UnmarshalJSON(b []byte) error {
	var DncpatchphonenumbersrequestMap map[string]interface{}
	err := json.Unmarshal(b, &DncpatchphonenumbersrequestMap)
	if err != nil {
		return err
	}
	
	if Action, ok := DncpatchphonenumbersrequestMap["action"].(string); ok {
		o.Action = &Action
	}
    
	if PhoneNumbers, ok := DncpatchphonenumbersrequestMap["phoneNumbers"].([]interface{}); ok {
		PhoneNumbersString, _ := json.Marshal(PhoneNumbers)
		json.Unmarshal(PhoneNumbersString, &o.PhoneNumbers)
	}
	
	if ExpirationDateTime, ok := DncpatchphonenumbersrequestMap["expirationDateTime"].(string); ok {
		o.ExpirationDateTime = &ExpirationDateTime
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dncpatchphonenumbersrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
