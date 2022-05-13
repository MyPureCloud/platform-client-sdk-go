package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queryrequestpredicate
type Queryrequestpredicate struct { 
	// Dimension - The dimension to be filtered
	Dimension *string `json:"dimension,omitempty"`


	// Value - The value to filter by
	Value *string `json:"value,omitempty"`

}

func (o *Queryrequestpredicate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queryrequestpredicate
	
	return json.Marshal(&struct { 
		Dimension *string `json:"dimension,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		Dimension: o.Dimension,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Queryrequestpredicate) UnmarshalJSON(b []byte) error {
	var QueryrequestpredicateMap map[string]interface{}
	err := json.Unmarshal(b, &QueryrequestpredicateMap)
	if err != nil {
		return err
	}
	
	if Dimension, ok := QueryrequestpredicateMap["dimension"].(string); ok {
		o.Dimension = &Dimension
	}
    
	if Value, ok := QueryrequestpredicateMap["value"].(string); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Queryrequestpredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
