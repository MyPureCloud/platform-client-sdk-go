package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontactlistconfigchangeurireference - A UriReference for a resource
type Dialercontactlistconfigchangeurireference struct { 
	// Id - The ID of the resource
	Id *string `json:"id,omitempty"`


	// Name - The name of the resource
	Name *string `json:"name,omitempty"`

}

func (o *Dialercontactlistconfigchangeurireference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercontactlistconfigchangeurireference
	
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

func (o *Dialercontactlistconfigchangeurireference) UnmarshalJSON(b []byte) error {
	var DialercontactlistconfigchangeurireferenceMap map[string]interface{}
	err := json.Unmarshal(b, &DialercontactlistconfigchangeurireferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DialercontactlistconfigchangeurireferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DialercontactlistconfigchangeurireferenceMap["name"].(string); ok {
		o.Name = &Name
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercontactlistconfigchangeurireference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
