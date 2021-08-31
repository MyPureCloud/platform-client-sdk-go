package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Reversewhitepageslookupresult
type Reversewhitepageslookupresult struct { 
	// Contacts
	Contacts *[]Externalcontact `json:"contacts,omitempty"`


	// ExternalOrganizations
	ExternalOrganizations *[]Externalorganization `json:"externalOrganizations,omitempty"`

}

func (o *Reversewhitepageslookupresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reversewhitepageslookupresult
	
	return json.Marshal(&struct { 
		Contacts *[]Externalcontact `json:"contacts,omitempty"`
		
		ExternalOrganizations *[]Externalorganization `json:"externalOrganizations,omitempty"`
		*Alias
	}{ 
		Contacts: o.Contacts,
		
		ExternalOrganizations: o.ExternalOrganizations,
		Alias:    (*Alias)(o),
	})
}

func (o *Reversewhitepageslookupresult) UnmarshalJSON(b []byte) error {
	var ReversewhitepageslookupresultMap map[string]interface{}
	err := json.Unmarshal(b, &ReversewhitepageslookupresultMap)
	if err != nil {
		return err
	}
	
	if Contacts, ok := ReversewhitepageslookupresultMap["contacts"].([]interface{}); ok {
		ContactsString, _ := json.Marshal(Contacts)
		json.Unmarshal(ContactsString, &o.Contacts)
	}
	
	if ExternalOrganizations, ok := ReversewhitepageslookupresultMap["externalOrganizations"].([]interface{}); ok {
		ExternalOrganizationsString, _ := json.Marshal(ExternalOrganizations)
		json.Unmarshal(ExternalOrganizationsString, &o.ExternalOrganizations)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Reversewhitepageslookupresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
