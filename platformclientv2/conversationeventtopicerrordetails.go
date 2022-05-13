package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopicerrordetails - Detailed information about an error response.
type Conversationeventtopicerrordetails struct { 
	// Status - The HTTP status code for this message (400, 401, 403, 404, 500, etc.
	Status *int `json:"status,omitempty"`


	// Code - A code unique to this error.
	Code *string `json:"code,omitempty"`


	// Message - Friendly description of this error.
	Message *string `json:"message,omitempty"`


	// MessageWithParams - This is the same as message except it uses template fields for variable replacement. For instance: 'User {username} was not found'
	MessageWithParams *string `json:"messageWithParams,omitempty"`


	// MessageParams - Used in conjunction with messageWithParams. These are the template parameters. For instance: UserParam.key = 'username', UserParam.value = 'john.doe'
	MessageParams *map[string]string `json:"messageParams,omitempty"`


	// ContextId - The correlation Id or context Id for this message. If left blank the Public API will look at the HTTP response header 'ININ-Correlation-Id' instead.
	ContextId *string `json:"contextId,omitempty"`


	// Uri
	Uri *string `json:"uri,omitempty"`

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
		*Alias
	}{ 
		Status: o.Status,
		
		Code: o.Code,
		
		Message: o.Message,
		
		MessageWithParams: o.MessageWithParams,
		
		MessageParams: o.MessageParams,
		
		ContextId: o.ContextId,
		
		Uri: o.Uri,
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
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventtopicerrordetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
