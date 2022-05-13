package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Cursors
type Cursors struct { 
	// Before
	Before *string `json:"before,omitempty"`


	// After
	After *string `json:"after,omitempty"`

}

func (o *Cursors) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Cursors
	
	return json.Marshal(&struct { 
		Before *string `json:"before,omitempty"`
		
		After *string `json:"after,omitempty"`
		*Alias
	}{ 
		Before: o.Before,
		
		After: o.After,
		Alias:    (*Alias)(o),
	})
}

func (o *Cursors) UnmarshalJSON(b []byte) error {
	var CursorsMap map[string]interface{}
	err := json.Unmarshal(b, &CursorsMap)
	if err != nil {
		return err
	}
	
	if Before, ok := CursorsMap["before"].(string); ok {
		o.Before = &Before
	}
    
	if After, ok := CursorsMap["after"].(string); ok {
		o.After = &After
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Cursors) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
