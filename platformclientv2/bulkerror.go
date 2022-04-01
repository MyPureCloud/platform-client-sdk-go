package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkerror
type Bulkerror struct { 
	// Message - Error message of the bulk operation result.
	Message *string `json:"message,omitempty"`


	// Code - Error code of the bulk operation result.
	Code *string `json:"code,omitempty"`

}

func (o *Bulkerror) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkerror
	
	return json.Marshal(&struct { 
		Message *string `json:"message,omitempty"`
		
		Code *string `json:"code,omitempty"`
		*Alias
	}{ 
		Message: o.Message,
		
		Code: o.Code,
		Alias:    (*Alias)(o),
	})
}

func (o *Bulkerror) UnmarshalJSON(b []byte) error {
	var BulkerrorMap map[string]interface{}
	err := json.Unmarshal(b, &BulkerrorMap)
	if err != nil {
		return err
	}
	
	if Message, ok := BulkerrorMap["message"].(string); ok {
		o.Message = &Message
	}
	
	if Code, ok := BulkerrorMap["code"].(string); ok {
		o.Code = &Code
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkerror) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
