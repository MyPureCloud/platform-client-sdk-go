package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Fieldconfigs
type Fieldconfigs struct { 
	// Org
	Org *Fieldconfig `json:"org,omitempty"`


	// Person
	Person *Fieldconfig `json:"person,omitempty"`


	// Group
	Group *Fieldconfig `json:"group,omitempty"`


	// ExternalContact
	ExternalContact *Fieldconfig `json:"externalContact,omitempty"`

}

func (o *Fieldconfigs) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Fieldconfigs
	
	return json.Marshal(&struct { 
		Org *Fieldconfig `json:"org,omitempty"`
		
		Person *Fieldconfig `json:"person,omitempty"`
		
		Group *Fieldconfig `json:"group,omitempty"`
		
		ExternalContact *Fieldconfig `json:"externalContact,omitempty"`
		*Alias
	}{ 
		Org: o.Org,
		
		Person: o.Person,
		
		Group: o.Group,
		
		ExternalContact: o.ExternalContact,
		Alias:    (*Alias)(o),
	})
}

func (o *Fieldconfigs) UnmarshalJSON(b []byte) error {
	var FieldconfigsMap map[string]interface{}
	err := json.Unmarshal(b, &FieldconfigsMap)
	if err != nil {
		return err
	}
	
	if Org, ok := FieldconfigsMap["org"].(map[string]interface{}); ok {
		OrgString, _ := json.Marshal(Org)
		json.Unmarshal(OrgString, &o.Org)
	}
	
	if Person, ok := FieldconfigsMap["person"].(map[string]interface{}); ok {
		PersonString, _ := json.Marshal(Person)
		json.Unmarshal(PersonString, &o.Person)
	}
	
	if Group, ok := FieldconfigsMap["group"].(map[string]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	
	if ExternalContact, ok := FieldconfigsMap["externalContact"].(map[string]interface{}); ok {
		ExternalContactString, _ := json.Marshal(ExternalContact)
		json.Unmarshal(ExternalContactString, &o.ExternalContact)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Fieldconfigs) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
