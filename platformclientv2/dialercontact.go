package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontact
type Dialercontact struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ContactListId - The identifier of the contact list containing this contact.
	ContactListId *string `json:"contactListId,omitempty"`


	// Data - An ordered map of the contact's columns and corresponding values.
	Data *map[string]interface{} `json:"data,omitempty"`


	// CallRecords - A map of call records for the contact phone columns.
	CallRecords *map[string]Callrecord `json:"callRecords,omitempty"`


	// Callable - Indicates whether or not the contact can be called.
	Callable *bool `json:"callable,omitempty"`


	// PhoneNumberStatus - A map of phone number columns to PhoneNumberStatuses, which indicate if the phone number is callable or not.
	PhoneNumberStatus *map[string]Phonenumberstatus `json:"phoneNumberStatus,omitempty"`


	// ContactColumnTimeZones - Map containing data about the timezone the contact is mapped to. This will only be populated if the contact list has automatic timezone mapping turned on. The key is the column name. The value is the timezone it mapped to and the type of column: Phone or Zip
	ContactColumnTimeZones *map[string]Contactcolumntimezone `json:"contactColumnTimeZones,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercontact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
