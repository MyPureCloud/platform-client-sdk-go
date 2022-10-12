package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentqueryinterval
type Documentqueryinterval struct { 
	// Field - Specifies the date field to be used for date and time range.
	Field *string `json:"field,omitempty"`


	// Value - Specifies the date and time range for filtering the documents. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Value *string `json:"value,omitempty"`

}

func (o *Documentqueryinterval) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentqueryinterval
	
	return json.Marshal(&struct { 
		Field *string `json:"field,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		Field: o.Field,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Documentqueryinterval) UnmarshalJSON(b []byte) error {
	var DocumentqueryintervalMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentqueryintervalMap)
	if err != nil {
		return err
	}
	
	if Field, ok := DocumentqueryintervalMap["field"].(string); ok {
		o.Field = &Field
	}
    
	if Value, ok := DocumentqueryintervalMap["value"].(string); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Documentqueryinterval) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
