package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Writabledialercontact
type Writabledialercontact struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// ContactListId - The identifier of the contact list containing this contact.
	ContactListId *string `json:"contactListId,omitempty"`


	// Data - An ordered map of the contact's columns and corresponding values.
	Data *map[string]interface{} `json:"data,omitempty"`


	// Callable - Indicates whether or not the contact can be called.
	Callable *bool `json:"callable,omitempty"`


	// PhoneNumberStatus - A map of phone number columns to PhoneNumberStatuses, which indicate if the phone number is callable or not.
	PhoneNumberStatus *map[string]Phonenumberstatus `json:"phoneNumberStatus,omitempty"`

}

func (u *Writabledialercontact) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Writabledialercontact

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ContactListId *string `json:"contactListId,omitempty"`
		
		Data *map[string]interface{} `json:"data,omitempty"`
		
		Callable *bool `json:"callable,omitempty"`
		
		PhoneNumberStatus *map[string]Phonenumberstatus `json:"phoneNumberStatus,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		ContactListId: u.ContactListId,
		
		Data: u.Data,
		
		Callable: u.Callable,
		
		PhoneNumberStatus: u.PhoneNumberStatus,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Writabledialercontact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
