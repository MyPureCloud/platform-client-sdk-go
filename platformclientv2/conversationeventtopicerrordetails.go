package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopicerrordetails
type Conversationeventtopicerrordetails struct { 
	// Status
	Status *int `json:"status,omitempty"`


	// Code
	Code *string `json:"code,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// MessageWithParams
	MessageWithParams *string `json:"messageWithParams,omitempty"`


	// MessageParams
	MessageParams *map[string]string `json:"messageParams,omitempty"`


	// ContextId
	ContextId *string `json:"contextId,omitempty"`


	// Uri
	Uri *string `json:"uri,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Conversationeventtopicerrordetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationeventtopicerrordetails
	
	return json.Marshal(&struct { 
		Status *int `json:"status,omitempty"`
		
		Code *string `json:"code,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		MessageWithParams *string `json:"messageWithParams,omitempty"`
		
		MessageParams *map[string]string `json:"messageParams,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		Uri *string `json:"uri,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Status: o.Status,
		
		Code: o.Code,
		
		Message: o.Message,
		
		MessageWithParams: o.MessageWithParams,
		
		MessageParams: o.MessageParams,
		
		ContextId: o.ContextId,
		
		Uri: o.Uri,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationeventtopicerrordetails) UnmarshalJSON(b []byte) error {
	var ConversationeventtopicerrordetailsMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationeventtopicerrordetailsMap)
	if err != nil {
		return err
	}
	
	if Status, ok := ConversationeventtopicerrordetailsMap["status"].(float64); ok {
		StatusInt := int(Status)
		o.Status = &StatusInt
	}
	
	if Code, ok := ConversationeventtopicerrordetailsMap["code"].(string); ok {
		o.Code = &Code
	}
	
	if Message, ok := ConversationeventtopicerrordetailsMap["message"].(string); ok {
		o.Message = &Message
	}
	
	if MessageWithParams, ok := ConversationeventtopicerrordetailsMap["messageWithParams"].(string); ok {
		o.MessageWithParams = &MessageWithParams
	}
	
	if MessageParams, ok := ConversationeventtopicerrordetailsMap["messageParams"].(map[string]interface{}); ok {
		MessageParamsString, _ := json.Marshal(MessageParams)
		json.Unmarshal(MessageParamsString, &o.MessageParams)
	}
	
	if ContextId, ok := ConversationeventtopicerrordetailsMap["contextId"].(string); ok {
		o.ContextId = &ContextId
	}
	
	if Uri, ok := ConversationeventtopicerrordetailsMap["uri"].(string); ok {
		o.Uri = &Uri
	}
	
	if AdditionalProperties, ok := ConversationeventtopicerrordetailsMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventtopicerrordetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
