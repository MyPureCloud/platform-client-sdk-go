package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contextentity
type Contextentity struct { 
	// Name - The name of the entity.
	Name *string `json:"name,omitempty"`

}

func (o *Contextentity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contextentity
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		Alias:    (*Alias)(o),
	})
}

func (o *Contextentity) UnmarshalJSON(b []byte) error {
	var ContextentityMap map[string]interface{}
	err := json.Unmarshal(b, &ContextentityMap)
	if err != nil {
		return err
	}
	
	if Name, ok := ContextentityMap["name"].(string); ok {
		o.Name = &Name
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contextentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
