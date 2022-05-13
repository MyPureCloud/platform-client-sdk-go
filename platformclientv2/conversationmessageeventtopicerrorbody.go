package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessageeventtopicerrorbody
type Conversationmessageeventtopicerrorbody struct { 
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
	Details *[]Conversationmessageeventtopicdetail `json:"details,omitempty"`


	// Errors
	Errors *[]Conversationmessageeventtopicerrorbody `json:"errors,omitempty"`

}

func (o *Conversationmessageeventtopicerrorbody) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessageeventtopicerrorbody
	
	return json.Marshal(&struct { 
		Message *string `json:"message,omitempty"`
		
		Code *string `json:"code,omitempty"`
		
		Status *int `json:"status,omitempty"`
		
		EntityId *string `json:"entityId,omitempty"`
		
		EntityName *string `json:"entityName,omitempty"`
		
		MessageWithParams *string `json:"messageWithParams,omitempty"`
		
		MessageParams *map[string]string `json:"messageParams,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		Details *[]Conversationmessageeventtopicdetail `json:"details,omitempty"`
		
		Errors *[]Conversationmessageeventtopicerrorbody `json:"errors,omitempty"`
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

func (o *Conversationmessageeventtopicerrorbody) UnmarshalJSON(b []byte) error {
	var ConversationmessageeventtopicerrorbodyMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmessageeventtopicerrorbodyMap)
	if err != nil {
		return err
	}
	
	if Message, ok := ConversationmessageeventtopicerrorbodyMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if Code, ok := ConversationmessageeventtopicerrorbodyMap["code"].(string); ok {
		o.Code = &Code
	}
    
	if Status, ok := ConversationmessageeventtopicerrorbodyMap["status"].(float64); ok {
		StatusInt := int(Status)
		o.Status = &StatusInt
	}
	
	if EntityId, ok := ConversationmessageeventtopicerrorbodyMap["entityId"].(string); ok {
		o.EntityId = &EntityId
	}
    
	if EntityName, ok := ConversationmessageeventtopicerrorbodyMap["entityName"].(string); ok {
		o.EntityName = &EntityName
	}
    
	if MessageWithParams, ok := ConversationmessageeventtopicerrorbodyMap["messageWithParams"].(string); ok {
		o.MessageWithParams = &MessageWithParams
	}
    
	if MessageParams, ok := ConversationmessageeventtopicerrorbodyMap["messageParams"].(map[string]interface{}); ok {
		MessageParamsString, _ := json.Marshal(MessageParams)
		json.Unmarshal(MessageParamsString, &o.MessageParams)
	}
	
	if ContextId, ok := ConversationmessageeventtopicerrorbodyMap["contextId"].(string); ok {
		o.ContextId = &ContextId
	}
    
	if Details, ok := ConversationmessageeventtopicerrorbodyMap["details"].([]interface{}); ok {
		DetailsString, _ := json.Marshal(Details)
		json.Unmarshal(DetailsString, &o.Details)
	}
	
	if Errors, ok := ConversationmessageeventtopicerrorbodyMap["errors"].([]interface{}); ok {
		ErrorsString, _ := json.Marshal(Errors)
		json.Unmarshal(ErrorsString, &o.Errors)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessageeventtopicerrorbody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
