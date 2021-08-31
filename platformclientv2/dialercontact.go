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


	// LatestSmsEvaluations - A map of SMS records for the contact phone columns.
	LatestSmsEvaluations *map[string]Messageevaluation `json:"latestSmsEvaluations,omitempty"`


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

func (o *Dialercontact) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercontact
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ContactListId *string `json:"contactListId,omitempty"`
		
		Data *map[string]interface{} `json:"data,omitempty"`
		
		CallRecords *map[string]Callrecord `json:"callRecords,omitempty"`
		
		LatestSmsEvaluations *map[string]Messageevaluation `json:"latestSmsEvaluations,omitempty"`
		
		Callable *bool `json:"callable,omitempty"`
		
		PhoneNumberStatus *map[string]Phonenumberstatus `json:"phoneNumberStatus,omitempty"`
		
		ContactColumnTimeZones *map[string]Contactcolumntimezone `json:"contactColumnTimeZones,omitempty"`
		
		ConfigurationOverrides *Configurationoverrides `json:"configurationOverrides,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ContactListId: o.ContactListId,
		
		Data: o.Data,
		
		CallRecords: o.CallRecords,
		
		LatestSmsEvaluations: o.LatestSmsEvaluations,
		
		Callable: o.Callable,
		
		PhoneNumberStatus: o.PhoneNumberStatus,
		
		ContactColumnTimeZones: o.ContactColumnTimeZones,
		
		ConfigurationOverrides: o.ConfigurationOverrides,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercontact) UnmarshalJSON(b []byte) error {
	var DialercontactMap map[string]interface{}
	err := json.Unmarshal(b, &DialercontactMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DialercontactMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DialercontactMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if ContactListId, ok := DialercontactMap["contactListId"].(string); ok {
		o.ContactListId = &ContactListId
	}
	
	if Data, ok := DialercontactMap["data"].(map[string]interface{}); ok {
		DataString, _ := json.Marshal(Data)
		json.Unmarshal(DataString, &o.Data)
	}
	
	if CallRecords, ok := DialercontactMap["callRecords"].(map[string]interface{}); ok {
		CallRecordsString, _ := json.Marshal(CallRecords)
		json.Unmarshal(CallRecordsString, &o.CallRecords)
	}
	
	if LatestSmsEvaluations, ok := DialercontactMap["latestSmsEvaluations"].(map[string]interface{}); ok {
		LatestSmsEvaluationsString, _ := json.Marshal(LatestSmsEvaluations)
		json.Unmarshal(LatestSmsEvaluationsString, &o.LatestSmsEvaluations)
	}
	
	if Callable, ok := DialercontactMap["callable"].(bool); ok {
		o.Callable = &Callable
	}
	
	if PhoneNumberStatus, ok := DialercontactMap["phoneNumberStatus"].(map[string]interface{}); ok {
		PhoneNumberStatusString, _ := json.Marshal(PhoneNumberStatus)
		json.Unmarshal(PhoneNumberStatusString, &o.PhoneNumberStatus)
	}
	
	if ContactColumnTimeZones, ok := DialercontactMap["contactColumnTimeZones"].(map[string]interface{}); ok {
		ContactColumnTimeZonesString, _ := json.Marshal(ContactColumnTimeZones)
		json.Unmarshal(ContactColumnTimeZonesString, &o.ContactColumnTimeZones)
	}
	
	if ConfigurationOverrides, ok := DialercontactMap["configurationOverrides"].(map[string]interface{}); ok {
		ConfigurationOverridesString, _ := json.Marshal(ConfigurationOverrides)
		json.Unmarshal(ConfigurationOverridesString, &o.ConfigurationOverrides)
	}
	
	if SelfUri, ok := DialercontactMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercontact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
