package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Configurationoverrides
type Configurationoverrides struct { 
	// Priority - Indicates whether or not the contact will be placed in front of the queue or at the end of the queue.
	Priority *bool `json:"priority,omitempty"`

}

func (o *Configurationoverrides) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Configurationoverrides
	
	return json.Marshal(&struct { 
		Priority *bool `json:"priority,omitempty"`
		*Alias
	}{ 
		Priority: o.Priority,
		Alias:    (*Alias)(o),
	})
}

func (o *Configurationoverrides) UnmarshalJSON(b []byte) error {
	var ConfigurationoverridesMap map[string]interface{}
	err := json.Unmarshal(b, &ConfigurationoverridesMap)
	if err != nil {
		return err
	}
	
	if Priority, ok := ConfigurationoverridesMap["priority"].(bool); ok {
		o.Priority = &Priority
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Configurationoverrides) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
