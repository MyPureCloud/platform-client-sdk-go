package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Copyworkplanrotationrequest
type Copyworkplanrotationrequest struct { 
	// Name - Name to apply to the new copy of the work plan rotation
	Name *string `json:"name,omitempty"`

}

func (o *Copyworkplanrotationrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Copyworkplanrotationrequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		Alias:    (*Alias)(o),
	})
}

func (o *Copyworkplanrotationrequest) UnmarshalJSON(b []byte) error {
	var CopyworkplanrotationrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CopyworkplanrotationrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := CopyworkplanrotationrequestMap["name"].(string); ok {
		o.Name = &Name
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Copyworkplanrotationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
