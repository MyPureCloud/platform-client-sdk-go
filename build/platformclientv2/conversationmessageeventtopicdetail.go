package platformclientv2
import (
	"encoding/json"
)

// Conversationmessageeventtopicdetail
type Conversationmessageeventtopicdetail struct { 
	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// FieldName
	FieldName *string `json:"fieldName,omitempty"`


	// EntityId
	EntityId *string `json:"entityId,omitempty"`


	// EntityName
	EntityName *string `json:"entityName,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationmessageeventtopicdetail) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
