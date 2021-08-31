package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainrole
type Domainrole struct { 
	// Id - The ID of the role
	Id *string `json:"id,omitempty"`


	// Name - The name of the role
	Name *string `json:"name,omitempty"`

}

func (o *Domainrole) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domainrole
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		Alias:    (*Alias)(o),
	})
}

func (o *Domainrole) UnmarshalJSON(b []byte) error {
	var DomainroleMap map[string]interface{}
	err := json.Unmarshal(b, &DomainroleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DomainroleMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DomainroleMap["name"].(string); ok {
		o.Name = &Name
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Domainrole) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
