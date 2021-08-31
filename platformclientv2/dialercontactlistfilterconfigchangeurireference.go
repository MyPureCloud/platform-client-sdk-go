package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontactlistfilterconfigchangeurireference
type Dialercontactlistfilterconfigchangeurireference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`

}

func (o *Dialercontactlistfilterconfigchangeurireference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercontactlistfilterconfigchangeurireference
	
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

func (o *Dialercontactlistfilterconfigchangeurireference) UnmarshalJSON(b []byte) error {
	var DialercontactlistfilterconfigchangeurireferenceMap map[string]interface{}
	err := json.Unmarshal(b, &DialercontactlistfilterconfigchangeurireferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DialercontactlistfilterconfigchangeurireferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DialercontactlistfilterconfigchangeurireferenceMap["name"].(string); ok {
		o.Name = &Name
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercontactlistfilterconfigchangeurireference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
