package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dncpatchemailsrequest
type Dncpatchemailsrequest struct { 
	// Action - The action to perform
	Action *string `json:"action,omitempty"`


	// EmailAddresses - The list of email addresses to Add to / Remove from the DNC list 
	EmailAddresses *[]string `json:"emailAddresses,omitempty"`


	// ExpirationDateTime - Expiration date for DNC email addresses in yyyy-MM-ddTHH:mmZ format
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`

}

func (o *Dncpatchemailsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dncpatchemailsrequest
	
	return json.Marshal(&struct { 
		Action *string `json:"action,omitempty"`
		
		EmailAddresses *[]string `json:"emailAddresses,omitempty"`
		
		ExpirationDateTime *string `json:"expirationDateTime,omitempty"`
		*Alias
	}{ 
		Action: o.Action,
		
		EmailAddresses: o.EmailAddresses,
		
		ExpirationDateTime: o.ExpirationDateTime,
		Alias:    (*Alias)(o),
	})
}

func (o *Dncpatchemailsrequest) UnmarshalJSON(b []byte) error {
	var DncpatchemailsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &DncpatchemailsrequestMap)
	if err != nil {
		return err
	}
	
	if Action, ok := DncpatchemailsrequestMap["action"].(string); ok {
		o.Action = &Action
	}
    
	if EmailAddresses, ok := DncpatchemailsrequestMap["emailAddresses"].([]interface{}); ok {
		EmailAddressesString, _ := json.Marshal(EmailAddresses)
		json.Unmarshal(EmailAddressesString, &o.EmailAddresses)
	}
	
	if ExpirationDateTime, ok := DncpatchemailsrequestMap["expirationDateTime"].(string); ok {
		o.ExpirationDateTime = &ExpirationDateTime
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dncpatchemailsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
