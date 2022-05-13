package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Columnstatus
type Columnstatus struct { 
	// Contactable - Indicates whether or not an individual contact method column is contactable.
	Contactable *bool `json:"contactable,omitempty"`

}

func (o *Columnstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Columnstatus
	
	return json.Marshal(&struct { 
		Contactable *bool `json:"contactable,omitempty"`
		*Alias
	}{ 
		Contactable: o.Contactable,
		Alias:    (*Alias)(o),
	})
}

func (o *Columnstatus) UnmarshalJSON(b []byte) error {
	var ColumnstatusMap map[string]interface{}
	err := json.Unmarshal(b, &ColumnstatusMap)
	if err != nil {
		return err
	}
	
	if Contactable, ok := ColumnstatusMap["contactable"].(bool); ok {
		o.Contactable = &Contactable
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Columnstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
