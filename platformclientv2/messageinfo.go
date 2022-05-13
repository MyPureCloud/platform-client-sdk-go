package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messageinfo
type Messageinfo struct { 
	// LocalizableMessageCode - Key that can be used to localize the message.
	LocalizableMessageCode *string `json:"localizableMessageCode,omitempty"`


	// Message - Description of the message.
	Message *string `json:"message,omitempty"`


	// MessageWithParams - Message with template fields for variable replacement.
	MessageWithParams *string `json:"messageWithParams,omitempty"`


	// MessageParams - Map with fields for variable replacement.
	MessageParams *map[string]string `json:"messageParams,omitempty"`

}

func (o *Messageinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messageinfo
	
	return json.Marshal(&struct { 
		LocalizableMessageCode *string `json:"localizableMessageCode,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		MessageWithParams *string `json:"messageWithParams,omitempty"`
		
		MessageParams *map[string]string `json:"messageParams,omitempty"`
		*Alias
	}{ 
		LocalizableMessageCode: o.LocalizableMessageCode,
		
		Message: o.Message,
		
		MessageWithParams: o.MessageWithParams,
		
		MessageParams: o.MessageParams,
		Alias:    (*Alias)(o),
	})
}

func (o *Messageinfo) UnmarshalJSON(b []byte) error {
	var MessageinfoMap map[string]interface{}
	err := json.Unmarshal(b, &MessageinfoMap)
	if err != nil {
		return err
	}
	
	if LocalizableMessageCode, ok := MessageinfoMap["localizableMessageCode"].(string); ok {
		o.LocalizableMessageCode = &LocalizableMessageCode
	}
    
	if Message, ok := MessageinfoMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if MessageWithParams, ok := MessageinfoMap["messageWithParams"].(string); ok {
		o.MessageWithParams = &MessageWithParams
	}
    
	if MessageParams, ok := MessageinfoMap["messageParams"].(map[string]interface{}); ok {
		MessageParamsString, _ := json.Marshal(MessageParams)
		json.Unmarshal(MessageParamsString, &o.MessageParams)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messageinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
