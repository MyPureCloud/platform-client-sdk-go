package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkerrordetail
type Bulkerrordetail struct { 
	// FieldName
	FieldName *string `json:"fieldName,omitempty"`


	// Value
	Value *string `json:"value,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`

}

func (o *Bulkerrordetail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkerrordetail
	
	return json.Marshal(&struct { 
		FieldName *string `json:"fieldName,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		Message *string `json:"message,omitempty"`
		*Alias
	}{ 
		FieldName: o.FieldName,
		
		Value: o.Value,
		
		Message: o.Message,
		Alias:    (*Alias)(o),
	})
}

func (o *Bulkerrordetail) UnmarshalJSON(b []byte) error {
	var BulkerrordetailMap map[string]interface{}
	err := json.Unmarshal(b, &BulkerrordetailMap)
	if err != nil {
		return err
	}
	
	if FieldName, ok := BulkerrordetailMap["fieldName"].(string); ok {
		o.FieldName = &FieldName
	}
    
	if Value, ok := BulkerrordetailMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if Message, ok := BulkerrordetailMap["message"].(string); ok {
		o.Message = &Message
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkerrordetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
