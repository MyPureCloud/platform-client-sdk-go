package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcallbackeventtopicerrorbody
type Conversationcallbackeventtopicerrorbody struct { 
	// Message
	Message *string `json:"message,omitempty"`


	// Code
	Code *string `json:"code,omitempty"`


	// Status
	Status *int `json:"status,omitempty"`


	// EntityId
	EntityId *string `json:"entityId,omitempty"`


	// EntityName
	EntityName *string `json:"entityName,omitempty"`


	// MessageWithParams
	MessageWithParams *string `json:"messageWithParams,omitempty"`


	// MessageParams
	MessageParams *map[string]string `json:"messageParams,omitempty"`


	// ContextId
	ContextId *string `json:"contextId,omitempty"`


	// Details
	Details *[]Conversationcallbackeventtopicdetail `json:"details,omitempty"`


	// Errors
	Errors *[]Conversationcallbackeventtopicerrorbody `json:"errors,omitempty"`

}

func (o *Conversationcallbackeventtopicerrorbody) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcallbackeventtopicerrorbody
	
	return json.Marshal(&struct { 
		Message *string `json:"message,omitempty"`
		
		Code *string `json:"code,omitempty"`
		
		Status *int `json:"status,omitempty"`
		
		EntityId *string `json:"entityId,omitempty"`
		
		EntityName *string `json:"entityName,omitempty"`
		
		MessageWithParams *string `json:"messageWithParams,omitempty"`
		
		MessageParams *map[string]string `json:"messageParams,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		Details *[]Conversationcallbackeventtopicdetail `json:"details,omitempty"`
		
		Errors *[]Conversationcallbackeventtopicerrorbody `json:"errors,omitempty"`
		*Alias
	}{ 
		Message: o.Message,
		
		Code: o.Code,
		
		Status: o.Status,
		
		EntityId: o.EntityId,
		
		EntityName: o.EntityName,
		
		MessageWithParams: o.MessageWithParams,
		
		MessageParams: o.MessageParams,
		
		ContextId: o.ContextId,
		
		Details: o.Details,
		
		Errors: o.Errors,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationcallbackeventtopicerrorbody) UnmarshalJSON(b []byte) error {
	var ConversationcallbackeventtopicerrorbodyMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcallbackeventtopicerrorbodyMap)
	if err != nil {
		return err
	}
	
	if Message, ok := ConversationcallbackeventtopicerrorbodyMap["message"].(string); ok {
		o.Message = &Message
	}
	
	if Code, ok := ConversationcallbackeventtopicerrorbodyMap["code"].(string); ok {
		o.Code = &Code
	}
	
	if Status, ok := ConversationcallbackeventtopicerrorbodyMap["status"].(float64); ok {
		StatusInt := int(Status)
		o.Status = &StatusInt
	}
	
	if EntityId, ok := ConversationcallbackeventtopicerrorbodyMap["entityId"].(string); ok {
		o.EntityId = &EntityId
	}
	
	if EntityName, ok := ConversationcallbackeventtopicerrorbodyMap["entityName"].(string); ok {
		o.EntityName = &EntityName
	}
	
	if MessageWithParams, ok := ConversationcallbackeventtopicerrorbodyMap["messageWithParams"].(string); ok {
		o.MessageWithParams = &MessageWithParams
	}
	
	if MessageParams, ok := ConversationcallbackeventtopicerrorbodyMap["messageParams"].(map[string]interface{}); ok {
		MessageParamsString, _ := json.Marshal(MessageParams)
		json.Unmarshal(MessageParamsString, &o.MessageParams)
	}
	
	if ContextId, ok := ConversationcallbackeventtopicerrorbodyMap["contextId"].(string); ok {
		o.ContextId = &ContextId
	}
	
	if Details, ok := ConversationcallbackeventtopicerrorbodyMap["details"].([]interface{}); ok {
		DetailsString, _ := json.Marshal(Details)
		json.Unmarshal(DetailsString, &o.Details)
	}
	
	if Errors, ok := ConversationcallbackeventtopicerrorbodyMap["errors"].([]interface{}); ok {
		ErrorsString, _ := json.Marshal(Errors)
		json.Unmarshal(ErrorsString, &o.Errors)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcallbackeventtopicerrorbody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
