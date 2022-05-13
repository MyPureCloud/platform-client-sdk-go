package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Servicecontext
type Servicecontext struct { 
	// Name - Unused field for the purpose of ensuring a Swagger definition is created for a class with only @JsonIgnore members.
	Name *string `json:"name,omitempty"`

}

func (o *Servicecontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Servicecontext
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		Alias:    (*Alias)(o),
	})
}

func (o *Servicecontext) UnmarshalJSON(b []byte) error {
	var ServicecontextMap map[string]interface{}
	err := json.Unmarshal(b, &ServicecontextMap)
	if err != nil {
		return err
	}
	
	if Name, ok := ServicecontextMap["name"].(string); ok {
		o.Name = &Name
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Servicecontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
