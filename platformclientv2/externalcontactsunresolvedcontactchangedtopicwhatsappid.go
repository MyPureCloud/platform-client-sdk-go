package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactsunresolvedcontactchangedtopicwhatsappid
type Externalcontactsunresolvedcontactchangedtopicwhatsappid struct { 
	// PhoneNumber
	PhoneNumber *Externalcontactsunresolvedcontactchangedtopicphonenumber `json:"phoneNumber,omitempty"`


	// DisplayName
	DisplayName *string `json:"displayName,omitempty"`

}

func (o *Externalcontactsunresolvedcontactchangedtopicwhatsappid) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalcontactsunresolvedcontactchangedtopicwhatsappid
	
	return json.Marshal(&struct { 
		PhoneNumber *Externalcontactsunresolvedcontactchangedtopicphonenumber `json:"phoneNumber,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		*Alias
	}{ 
		PhoneNumber: o.PhoneNumber,
		
		DisplayName: o.DisplayName,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalcontactsunresolvedcontactchangedtopicwhatsappid) UnmarshalJSON(b []byte) error {
	var ExternalcontactsunresolvedcontactchangedtopicwhatsappidMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactsunresolvedcontactchangedtopicwhatsappidMap)
	if err != nil {
		return err
	}
	
	if PhoneNumber, ok := ExternalcontactsunresolvedcontactchangedtopicwhatsappidMap["phoneNumber"].(map[string]interface{}); ok {
		PhoneNumberString, _ := json.Marshal(PhoneNumber)
		json.Unmarshal(PhoneNumberString, &o.PhoneNumber)
	}
	
	if DisplayName, ok := ExternalcontactsunresolvedcontactchangedtopicwhatsappidMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactsunresolvedcontactchangedtopicwhatsappid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
