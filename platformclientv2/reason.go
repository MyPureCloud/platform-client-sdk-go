package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Reason - Reasons for a failed message receipt.
type Reason struct { 
	// Code - The reason code for the failed message receipt.
	Code *string `json:"code,omitempty"`


	// Message - Description of the reason for the failed message receipt.
	Message *string `json:"message,omitempty"`

}

func (o *Reason) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reason
	
	return json.Marshal(&struct { 
		Code *string `json:"code,omitempty"`
		
		Message *string `json:"message,omitempty"`
		*Alias
	}{ 
		Code: o.Code,
		
		Message: o.Message,
		Alias:    (*Alias)(o),
	})
}

func (o *Reason) UnmarshalJSON(b []byte) error {
	var ReasonMap map[string]interface{}
	err := json.Unmarshal(b, &ReasonMap)
	if err != nil {
		return err
	}
	
	if Code, ok := ReasonMap["code"].(string); ok {
		o.Code = &Code
	}
	
	if Message, ok := ReasonMap["message"].(string); ok {
		o.Message = &Message
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Reason) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
