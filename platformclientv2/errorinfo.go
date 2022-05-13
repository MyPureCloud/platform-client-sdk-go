package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Errorinfo
type Errorinfo struct { 
	// Message
	Message *string `json:"message,omitempty"`


	// Code
	Code *string `json:"code,omitempty"`

}

func (o *Errorinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Errorinfo
	
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

func (o *Errorinfo) UnmarshalJSON(b []byte) error {
	var ErrorinfoMap map[string]interface{}
	err := json.Unmarshal(b, &ErrorinfoMap)
	if err != nil {
		return err
	}
	
	if Message, ok := ErrorinfoMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if Code, ok := ErrorinfoMap["code"].(string); ok {
		o.Code = &Code
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Errorinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
