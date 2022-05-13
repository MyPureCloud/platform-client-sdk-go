package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contextintent
type Contextintent struct { 
	// Name - The name of the intent.
	Name *string `json:"name,omitempty"`

}

func (o *Contextintent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contextintent
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		Alias:    (*Alias)(o),
	})
}

func (o *Contextintent) UnmarshalJSON(b []byte) error {
	var ContextintentMap map[string]interface{}
	err := json.Unmarshal(b, &ContextintentMap)
	if err != nil {
		return err
	}
	
	if Name, ok := ContextintentMap["name"].(string); ok {
		o.Name = &Name
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contextintent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
