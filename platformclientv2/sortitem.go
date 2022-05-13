package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Sortitem
type Sortitem struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// Ascending
	Ascending *bool `json:"ascending,omitempty"`

}

func (o *Sortitem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Sortitem
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Ascending *bool `json:"ascending,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Ascending: o.Ascending,
		Alias:    (*Alias)(o),
	})
}

func (o *Sortitem) UnmarshalJSON(b []byte) error {
	var SortitemMap map[string]interface{}
	err := json.Unmarshal(b, &SortitemMap)
	if err != nil {
		return err
	}
	
	if Name, ok := SortitemMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Ascending, ok := SortitemMap["ascending"].(bool); ok {
		o.Ascending = &Ascending
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Sortitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
