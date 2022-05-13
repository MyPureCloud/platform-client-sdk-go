package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationsocialexpressioneventtopicdetail
type Conversationsocialexpressioneventtopicdetail struct { 
	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// FieldName
	FieldName *string `json:"fieldName,omitempty"`


	// EntityId
	EntityId *string `json:"entityId,omitempty"`


	// EntityName
	EntityName *string `json:"entityName,omitempty"`

}

func (o *Conversationsocialexpressioneventtopicdetail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationsocialexpressioneventtopicdetail
	
	return json.Marshal(&struct { 
		ErrorCode *string `json:"errorCode,omitempty"`
		
		FieldName *string `json:"fieldName,omitempty"`
		
		EntityId *string `json:"entityId,omitempty"`
		
		EntityName *string `json:"entityName,omitempty"`
		*Alias
	}{ 
		ErrorCode: o.ErrorCode,
		
		FieldName: o.FieldName,
		
		EntityId: o.EntityId,
		
		EntityName: o.EntityName,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationsocialexpressioneventtopicdetail) UnmarshalJSON(b []byte) error {
	var ConversationsocialexpressioneventtopicdetailMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationsocialexpressioneventtopicdetailMap)
	if err != nil {
		return err
	}
	
	if ErrorCode, ok := ConversationsocialexpressioneventtopicdetailMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if FieldName, ok := ConversationsocialexpressioneventtopicdetailMap["fieldName"].(string); ok {
		o.FieldName = &FieldName
	}
    
	if EntityId, ok := ConversationsocialexpressioneventtopicdetailMap["entityId"].(string); ok {
		o.EntityId = &EntityId
	}
    
	if EntityName, ok := ConversationsocialexpressioneventtopicdetailMap["entityName"].(string); ok {
		o.EntityName = &EntityName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationsocialexpressioneventtopicdetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
