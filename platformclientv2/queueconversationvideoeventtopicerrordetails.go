package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationvideoeventtopicerrordetails
type Queueconversationvideoeventtopicerrordetails struct { 
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

func (o *Queueconversationvideoeventtopicerrordetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopicerrordetails
	
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

func (o *Queueconversationvideoeventtopicerrordetails) UnmarshalJSON(b []byte) error {
	var QueueconversationvideoeventtopicerrordetailsMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationvideoeventtopicerrordetailsMap)
	if err != nil {
		return err
	}
	
	if Status, ok := QueueconversationvideoeventtopicerrordetailsMap["status"].(float64); ok {
		StatusInt := int(Status)
		o.Status = &StatusInt
	}
	
	if Code, ok := QueueconversationvideoeventtopicerrordetailsMap["code"].(string); ok {
		o.Code = &Code
	}
	
	if Message, ok := QueueconversationvideoeventtopicerrordetailsMap["message"].(string); ok {
		o.Message = &Message
	}
	
	if MessageWithParams, ok := QueueconversationvideoeventtopicerrordetailsMap["messageWithParams"].(string); ok {
		o.MessageWithParams = &MessageWithParams
	}
	
	if MessageParams, ok := QueueconversationvideoeventtopicerrordetailsMap["messageParams"].(map[string]interface{}); ok {
		MessageParamsString, _ := json.Marshal(MessageParams)
		json.Unmarshal(MessageParamsString, &o.MessageParams)
	}
	
	if ContextId, ok := QueueconversationvideoeventtopicerrordetailsMap["contextId"].(string); ok {
		o.ContextId = &ContextId
	}
	
	if Uri, ok := QueueconversationvideoeventtopicerrordetailsMap["uri"].(string); ok {
		o.Uri = &Uri
	}
	
	if AdditionalProperties, ok := QueueconversationvideoeventtopicerrordetailsMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicerrordetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
