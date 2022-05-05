package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dynamiccontactqueueingsettings
type Dynamiccontactqueueingsettings struct { 
	// Sort - Whether to sort contacts dynamically
	Sort *bool `json:"sort,omitempty"`

}

func (o *Dynamiccontactqueueingsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dynamiccontactqueueingsettings
	
	return json.Marshal(&struct { 
		Sort *bool `json:"sort,omitempty"`
		*Alias
	}{ 
		Sort: o.Sort,
		Alias:    (*Alias)(o),
	})
}

func (o *Dynamiccontactqueueingsettings) UnmarshalJSON(b []byte) error {
	var DynamiccontactqueueingsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &DynamiccontactqueueingsettingsMap)
	if err != nil {
		return err
	}
	
	if Sort, ok := DynamiccontactqueueingsettingsMap["sort"].(bool); ok {
		o.Sort = &Sort
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dynamiccontactqueueingsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
