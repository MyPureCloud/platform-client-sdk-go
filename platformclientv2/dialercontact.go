package platformclientv2
import (
	"github.com/leekchan/timeutil"
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


	// ConfigurationOverrides - the priority property within ConfigurationOverides indicates whether or not the contact to be placed in front of the queue or at the end of the queue
	ConfigurationOverrides *Configurationoverrides `json:"configurationOverrides,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Dialercontact) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercontact

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ContactListId *string `json:"contactListId,omitempty"`
		
		Data *map[string]interface{} `json:"data,omitempty"`
		
		CallRecords *map[string]Callrecord `json:"callRecords,omitempty"`
		
		Callable *bool `json:"callable,omitempty"`
		
		PhoneNumberStatus *map[string]Phonenumberstatus `json:"phoneNumberStatus,omitempty"`
		
		ContactColumnTimeZones *map[string]Contactcolumntimezone `json:"contactColumnTimeZones,omitempty"`
		
		ConfigurationOverrides *Configurationoverrides `json:"configurationOverrides,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		ContactListId: u.ContactListId,
		
		Data: u.Data,
		
		CallRecords: u.CallRecords,
		
		Callable: u.Callable,
		
		PhoneNumberStatus: u.PhoneNumberStatus,
		
		ContactColumnTimeZones: u.ContactColumnTimeZones,
		
		ConfigurationOverrides: u.ConfigurationOverrides,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialercontact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
