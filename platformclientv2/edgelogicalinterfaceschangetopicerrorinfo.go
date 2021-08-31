package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgelogicalinterfaceschangetopicerrorinfo
type Edgelogicalinterfaceschangetopicerrorinfo struct { 
	// Message
	Message *string `json:"message,omitempty"`


	// MessageWithParams
	MessageWithParams *string `json:"messageWithParams,omitempty"`


	// MessageParams
	MessageParams *map[string]string `json:"messageParams,omitempty"`


	// Code
	Code *string `json:"code,omitempty"`

}

func (o *Edgelogicalinterfaceschangetopicerrorinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgelogicalinterfaceschangetopicerrorinfo
	
	return json.Marshal(&struct { 
		Message *string `json:"message,omitempty"`
		
		MessageWithParams *string `json:"messageWithParams,omitempty"`
		
		MessageParams *map[string]string `json:"messageParams,omitempty"`
		
		Code *string `json:"code,omitempty"`
		*Alias
	}{ 
		Message: o.Message,
		
		MessageWithParams: o.MessageWithParams,
		
		MessageParams: o.MessageParams,
		
		Code: o.Code,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgelogicalinterfaceschangetopicerrorinfo) UnmarshalJSON(b []byte) error {
	var EdgelogicalinterfaceschangetopicerrorinfoMap map[string]interface{}
	err := json.Unmarshal(b, &EdgelogicalinterfaceschangetopicerrorinfoMap)
	if err != nil {
		return err
	}
	
	if Message, ok := EdgelogicalinterfaceschangetopicerrorinfoMap["message"].(string); ok {
		o.Message = &Message
	}
	
	if MessageWithParams, ok := EdgelogicalinterfaceschangetopicerrorinfoMap["messageWithParams"].(string); ok {
		o.MessageWithParams = &MessageWithParams
	}
	
	if MessageParams, ok := EdgelogicalinterfaceschangetopicerrorinfoMap["messageParams"].(map[string]interface{}); ok {
		MessageParamsString, _ := json.Marshal(MessageParams)
		json.Unmarshal(MessageParamsString, &o.MessageParams)
	}
	
	if Code, ok := EdgelogicalinterfaceschangetopicerrorinfoMap["code"].(string); ok {
		o.Code = &Code
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgelogicalinterfaceschangetopicerrorinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
