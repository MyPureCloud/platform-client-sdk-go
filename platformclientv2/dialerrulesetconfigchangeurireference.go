package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerrulesetconfigchangeurireference - A UriReference for a resource
type Dialerrulesetconfigchangeurireference struct { 
	// Id - The ID of the resource
	Id *string `json:"id,omitempty"`


	// Name - The name of the resource
	Name *string `json:"name,omitempty"`

}

func (o *Dialerrulesetconfigchangeurireference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerrulesetconfigchangeurireference
	
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

func (o *Dialerrulesetconfigchangeurireference) UnmarshalJSON(b []byte) error {
	var DialerrulesetconfigchangeurireferenceMap map[string]interface{}
	err := json.Unmarshal(b, &DialerrulesetconfigchangeurireferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DialerrulesetconfigchangeurireferenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DialerrulesetconfigchangeurireferenceMap["name"].(string); ok {
		o.Name = &Name
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerrulesetconfigchangeurireference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
