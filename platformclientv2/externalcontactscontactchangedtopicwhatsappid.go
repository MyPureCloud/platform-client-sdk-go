package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactscontactchangedtopicwhatsappid
type Externalcontactscontactchangedtopicwhatsappid struct { 
	// PhoneNumber
	PhoneNumber *Externalcontactscontactchangedtopicphonenumber `json:"phoneNumber,omitempty"`


	// DisplayName
	DisplayName *string `json:"displayName,omitempty"`

}

func (o *Externalcontactscontactchangedtopicwhatsappid) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalcontactscontactchangedtopicwhatsappid
	
	return json.Marshal(&struct { 
		PhoneNumber *Externalcontactscontactchangedtopicphonenumber `json:"phoneNumber,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		*Alias
	}{ 
		PhoneNumber: o.PhoneNumber,
		
		DisplayName: o.DisplayName,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalcontactscontactchangedtopicwhatsappid) UnmarshalJSON(b []byte) error {
	var ExternalcontactscontactchangedtopicwhatsappidMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactscontactchangedtopicwhatsappidMap)
	if err != nil {
		return err
	}
	
	if PhoneNumber, ok := ExternalcontactscontactchangedtopicwhatsappidMap["phoneNumber"].(map[string]interface{}); ok {
		PhoneNumberString, _ := json.Marshal(PhoneNumber)
		json.Unmarshal(PhoneNumberString, &o.PhoneNumber)
	}
	
	if DisplayName, ok := ExternalcontactscontactchangedtopicwhatsappidMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactscontactchangedtopicwhatsappid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
