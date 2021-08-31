package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationemaileventtopicdetail
type Queueconversationemaileventtopicdetail struct { 
	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// FieldName
	FieldName *string `json:"fieldName,omitempty"`


	// EntityId
	EntityId *string `json:"entityId,omitempty"`


	// EntityName
	EntityName *string `json:"entityName,omitempty"`

}

func (o *Queueconversationemaileventtopicdetail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationemaileventtopicdetail
	
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

func (o *Queueconversationemaileventtopicdetail) UnmarshalJSON(b []byte) error {
	var QueueconversationemaileventtopicdetailMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationemaileventtopicdetailMap)
	if err != nil {
		return err
	}
	
	if ErrorCode, ok := QueueconversationemaileventtopicdetailMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
	
	if FieldName, ok := QueueconversationemaileventtopicdetailMap["fieldName"].(string); ok {
		o.FieldName = &FieldName
	}
	
	if EntityId, ok := QueueconversationemaileventtopicdetailMap["entityId"].(string); ok {
		o.EntityId = &EntityId
	}
	
	if EntityName, ok := QueueconversationemaileventtopicdetailMap["entityName"].(string); ok {
		o.EntityName = &EntityName
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationemaileventtopicdetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
