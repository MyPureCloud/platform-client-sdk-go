package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Errordetails
type Errordetails struct { 
	// Status
	Status *int `json:"status,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// MessageWithParams
	MessageWithParams *string `json:"messageWithParams,omitempty"`


	// MessageParams
	MessageParams *map[string]string `json:"messageParams,omitempty"`


	// Code
	Code *string `json:"code,omitempty"`


	// ContextId
	ContextId *string `json:"contextId,omitempty"`


	// Nested
	Nested **Errordetails `json:"nested,omitempty"`


	// Details
	Details *string `json:"details,omitempty"`

}

func (o *Errordetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Errordetails
	
	return json.Marshal(&struct { 
		Status *int `json:"status,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		MessageWithParams *string `json:"messageWithParams,omitempty"`
		
		MessageParams *map[string]string `json:"messageParams,omitempty"`
		
		Code *string `json:"code,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		Nested **Errordetails `json:"nested,omitempty"`
		
		Details *string `json:"details,omitempty"`
		*Alias
	}{ 
		Status: o.Status,
		
		Message: o.Message,
		
		MessageWithParams: o.MessageWithParams,
		
		MessageParams: o.MessageParams,
		
		Code: o.Code,
		
		ContextId: o.ContextId,
		
		Nested: o.Nested,
		
		Details: o.Details,
		Alias:    (*Alias)(o),
	})
}

func (o *Errordetails) UnmarshalJSON(b []byte) error {
	var ErrordetailsMap map[string]interface{}
	err := json.Unmarshal(b, &ErrordetailsMap)
	if err != nil {
		return err
	}
	
	if Status, ok := ErrordetailsMap["status"].(float64); ok {
		StatusInt := int(Status)
		o.Status = &StatusInt
	}
	
	if Message, ok := ErrordetailsMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if MessageWithParams, ok := ErrordetailsMap["messageWithParams"].(string); ok {
		o.MessageWithParams = &MessageWithParams
	}
    
	if MessageParams, ok := ErrordetailsMap["messageParams"].(map[string]interface{}); ok {
		MessageParamsString, _ := json.Marshal(MessageParams)
		json.Unmarshal(MessageParamsString, &o.MessageParams)
	}
	
	if Code, ok := ErrordetailsMap["code"].(string); ok {
		o.Code = &Code
	}
    
	if ContextId, ok := ErrordetailsMap["contextId"].(string); ok {
		o.ContextId = &ContextId
	}
    
	if Nested, ok := ErrordetailsMap["nested"].(map[string]interface{}); ok {
		NestedString, _ := json.Marshal(Nested)
		json.Unmarshal(NestedString, &o.Nested)
	}
	
	if Details, ok := ErrordetailsMap["details"].(string); ok {
		o.Details = &Details
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Errordetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
