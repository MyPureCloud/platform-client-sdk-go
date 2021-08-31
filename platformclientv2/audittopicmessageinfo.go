package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Audittopicmessageinfo
type Audittopicmessageinfo struct { 
	// LocalizableMessageCode
	LocalizableMessageCode *string `json:"localizableMessageCode,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// MessageWithParams
	MessageWithParams *string `json:"messageWithParams,omitempty"`


	// MessageParams
	MessageParams *map[string]string `json:"messageParams,omitempty"`

}

func (o *Audittopicmessageinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Audittopicmessageinfo
	
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

func (o *Audittopicmessageinfo) UnmarshalJSON(b []byte) error {
	var AudittopicmessageinfoMap map[string]interface{}
	err := json.Unmarshal(b, &AudittopicmessageinfoMap)
	if err != nil {
		return err
	}
	
	if LocalizableMessageCode, ok := AudittopicmessageinfoMap["localizableMessageCode"].(string); ok {
		o.LocalizableMessageCode = &LocalizableMessageCode
	}
	
	if Message, ok := AudittopicmessageinfoMap["message"].(string); ok {
		o.Message = &Message
	}
	
	if MessageWithParams, ok := AudittopicmessageinfoMap["messageWithParams"].(string); ok {
		o.MessageWithParams = &MessageWithParams
	}
	
	if MessageParams, ok := AudittopicmessageinfoMap["messageParams"].(map[string]interface{}); ok {
		MessageParamsString, _ := json.Marshal(MessageParams)
		json.Unmarshal(MessageParamsString, &o.MessageParams)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Audittopicmessageinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
