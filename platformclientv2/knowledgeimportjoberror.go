package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeimportjoberror
type Knowledgeimportjoberror struct { 
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
	Details *[]Detail `json:"details,omitempty"`


	// Errors
	Errors *[]Errorbody `json:"errors,omitempty"`


	// DocumentIndex - Index of the faulty document.
	DocumentIndex *int `json:"documentIndex,omitempty"`

}

func (o *Knowledgeimportjoberror) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgeimportjoberror
	
	return json.Marshal(&struct { 
		Message *string `json:"message,omitempty"`
		
		Code *string `json:"code,omitempty"`
		
		Status *int `json:"status,omitempty"`
		
		EntityId *string `json:"entityId,omitempty"`
		
		EntityName *string `json:"entityName,omitempty"`
		
		MessageWithParams *string `json:"messageWithParams,omitempty"`
		
		MessageParams *map[string]string `json:"messageParams,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		Details *[]Detail `json:"details,omitempty"`
		
		Errors *[]Errorbody `json:"errors,omitempty"`
		
		DocumentIndex *int `json:"documentIndex,omitempty"`
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
		
		DocumentIndex: o.DocumentIndex,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgeimportjoberror) UnmarshalJSON(b []byte) error {
	var KnowledgeimportjoberrorMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeimportjoberrorMap)
	if err != nil {
		return err
	}
	
	if Message, ok := KnowledgeimportjoberrorMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if Code, ok := KnowledgeimportjoberrorMap["code"].(string); ok {
		o.Code = &Code
	}
    
	if Status, ok := KnowledgeimportjoberrorMap["status"].(float64); ok {
		StatusInt := int(Status)
		o.Status = &StatusInt
	}
	
	if EntityId, ok := KnowledgeimportjoberrorMap["entityId"].(string); ok {
		o.EntityId = &EntityId
	}
    
	if EntityName, ok := KnowledgeimportjoberrorMap["entityName"].(string); ok {
		o.EntityName = &EntityName
	}
    
	if MessageWithParams, ok := KnowledgeimportjoberrorMap["messageWithParams"].(string); ok {
		o.MessageWithParams = &MessageWithParams
	}
    
	if MessageParams, ok := KnowledgeimportjoberrorMap["messageParams"].(map[string]interface{}); ok {
		MessageParamsString, _ := json.Marshal(MessageParams)
		json.Unmarshal(MessageParamsString, &o.MessageParams)
	}
	
	if ContextId, ok := KnowledgeimportjoberrorMap["contextId"].(string); ok {
		o.ContextId = &ContextId
	}
    
	if Details, ok := KnowledgeimportjoberrorMap["details"].([]interface{}); ok {
		DetailsString, _ := json.Marshal(Details)
		json.Unmarshal(DetailsString, &o.Details)
	}
	
	if Errors, ok := KnowledgeimportjoberrorMap["errors"].([]interface{}); ok {
		ErrorsString, _ := json.Marshal(Errors)
		json.Unmarshal(ErrorsString, &o.Errors)
	}
	
	if DocumentIndex, ok := KnowledgeimportjoberrorMap["documentIndex"].(float64); ok {
		DocumentIndexInt := int(DocumentIndex)
		o.DocumentIndex = &DocumentIndexInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeimportjoberror) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
