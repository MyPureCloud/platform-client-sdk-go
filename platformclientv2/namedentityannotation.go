package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Namedentityannotation
type Namedentityannotation struct { 
	// Name - The name of the annotated named entity.
	Name *string `json:"name,omitempty"`

}

func (o *Namedentityannotation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Namedentityannotation
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		Alias:    (*Alias)(o),
	})
}

func (o *Namedentityannotation) UnmarshalJSON(b []byte) error {
	var NamedentityannotationMap map[string]interface{}
	err := json.Unmarshal(b, &NamedentityannotationMap)
	if err != nil {
		return err
	}
	
	if Name, ok := NamedentityannotationMap["name"].(string); ok {
		o.Name = &Name
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Namedentityannotation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
