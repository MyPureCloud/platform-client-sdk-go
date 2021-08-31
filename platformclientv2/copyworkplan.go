package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Copyworkplan - Information associated with a work plan thats created as a copy
type Copyworkplan struct { 
	// Name - Name of the copied work plan
	Name *string `json:"name,omitempty"`

}

func (o *Copyworkplan) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Copyworkplan
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		Alias:    (*Alias)(o),
	})
}

func (o *Copyworkplan) UnmarshalJSON(b []byte) error {
	var CopyworkplanMap map[string]interface{}
	err := json.Unmarshal(b, &CopyworkplanMap)
	if err != nil {
		return err
	}
	
	if Name, ok := CopyworkplanMap["name"].(string); ok {
		o.Name = &Name
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Copyworkplan) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
