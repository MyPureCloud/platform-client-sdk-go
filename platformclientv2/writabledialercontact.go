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


	// LatestSmsEvaluations - A map of SMS records for the contact phone columns.
	LatestSmsEvaluations *map[string]Messageevaluation `json:"latestSmsEvaluations,omitempty"`


	// Callable - Indicates whether or not the contact can be called.
	Callable *bool `json:"callable,omitempty"`


	// PhoneNumberStatus - A map of phone number columns to PhoneNumberStatuses, which indicate if the phone number is callable or not.
	PhoneNumberStatus *map[string]Phonenumberstatus `json:"phoneNumberStatus,omitempty"`


	// ContactableStatus - A map of media types (Voice, SMS and Email) to ContactableStatus, which indicates if the contact can be contacted using the specified media type.
	ContactableStatus *map[string]Contactablestatus `json:"contactableStatus,omitempty"`

}

func (o *Writabledialercontact) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Writabledialercontact
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ContactListId *string `json:"contactListId,omitempty"`
		
		Data *map[string]interface{} `json:"data,omitempty"`
		
		LatestSmsEvaluations *map[string]Messageevaluation `json:"latestSmsEvaluations,omitempty"`
		
		Callable *bool `json:"callable,omitempty"`
		
		PhoneNumberStatus *map[string]Phonenumberstatus `json:"phoneNumberStatus,omitempty"`
		
		ContactableStatus *map[string]Contactablestatus `json:"contactableStatus,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		ContactListId: o.ContactListId,
		
		Data: o.Data,
		
		LatestSmsEvaluations: o.LatestSmsEvaluations,
		
		Callable: o.Callable,
		
		PhoneNumberStatus: o.PhoneNumberStatus,
		
		ContactableStatus: o.ContactableStatus,
		Alias:    (*Alias)(o),
	})
}

func (o *Writabledialercontact) UnmarshalJSON(b []byte) error {
	var WritabledialercontactMap map[string]interface{}
	err := json.Unmarshal(b, &WritabledialercontactMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WritabledialercontactMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ContactListId, ok := WritabledialercontactMap["contactListId"].(string); ok {
		o.ContactListId = &ContactListId
	}
    
	if Data, ok := WritabledialercontactMap["data"].(map[string]interface{}); ok {
		DataString, _ := json.Marshal(Data)
		json.Unmarshal(DataString, &o.Data)
	}
	
	if LatestSmsEvaluations, ok := WritabledialercontactMap["latestSmsEvaluations"].(map[string]interface{}); ok {
		LatestSmsEvaluationsString, _ := json.Marshal(LatestSmsEvaluations)
		json.Unmarshal(LatestSmsEvaluationsString, &o.LatestSmsEvaluations)
	}
	
	if Callable, ok := WritabledialercontactMap["callable"].(bool); ok {
		o.Callable = &Callable
	}
    
	if PhoneNumberStatus, ok := WritabledialercontactMap["phoneNumberStatus"].(map[string]interface{}); ok {
		PhoneNumberStatusString, _ := json.Marshal(PhoneNumberStatus)
		json.Unmarshal(PhoneNumberStatusString, &o.PhoneNumberStatus)
	}
	
	if ContactableStatus, ok := WritabledialercontactMap["contactableStatus"].(map[string]interface{}); ok {
		ContactableStatusString, _ := json.Marshal(ContactableStatus)
		json.Unmarshal(ContactableStatusString, &o.ContactableStatus)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Writabledialercontact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
