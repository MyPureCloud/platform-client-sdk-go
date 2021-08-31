package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Eventmessage
type Eventmessage struct { 
	// Code
	Code *string `json:"code,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// MessageWithParams
	MessageWithParams *string `json:"messageWithParams,omitempty"`


	// MessageParams
	MessageParams *map[string]interface{} `json:"messageParams,omitempty"`


	// DocumentationUri
	DocumentationUri *string `json:"documentationUri,omitempty"`


	// ResourceURIs
	ResourceURIs *[]string `json:"resourceURIs,omitempty"`

}

func (o *Eventmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Eventmessage
	
	return json.Marshal(&struct { 
		Code *string `json:"code,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		MessageWithParams *string `json:"messageWithParams,omitempty"`
		
		MessageParams *map[string]interface{} `json:"messageParams,omitempty"`
		
		DocumentationUri *string `json:"documentationUri,omitempty"`
		
		ResourceURIs *[]string `json:"resourceURIs,omitempty"`
		*Alias
	}{ 
		Code: o.Code,
		
		Message: o.Message,
		
		MessageWithParams: o.MessageWithParams,
		
		MessageParams: o.MessageParams,
		
		DocumentationUri: o.DocumentationUri,
		
		ResourceURIs: o.ResourceURIs,
		Alias:    (*Alias)(o),
	})
}

func (o *Eventmessage) UnmarshalJSON(b []byte) error {
	var EventmessageMap map[string]interface{}
	err := json.Unmarshal(b, &EventmessageMap)
	if err != nil {
		return err
	}
	
	if Code, ok := EventmessageMap["code"].(string); ok {
		o.Code = &Code
	}
	
	if Message, ok := EventmessageMap["message"].(string); ok {
		o.Message = &Message
	}
	
	if MessageWithParams, ok := EventmessageMap["messageWithParams"].(string); ok {
		o.MessageWithParams = &MessageWithParams
	}
	
	if MessageParams, ok := EventmessageMap["messageParams"].(map[string]interface{}); ok {
		MessageParamsString, _ := json.Marshal(MessageParams)
		json.Unmarshal(MessageParamsString, &o.MessageParams)
	}
	
	if DocumentationUri, ok := EventmessageMap["documentationUri"].(string); ok {
		o.DocumentationUri = &DocumentationUri
	}
	
	if ResourceURIs, ok := EventmessageMap["resourceURIs"].([]interface{}); ok {
		ResourceURIsString, _ := json.Marshal(ResourceURIs)
		json.Unmarshal(ResourceURIsString, &o.ResourceURIs)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Eventmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
