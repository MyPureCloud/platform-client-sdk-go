package platformclientv2
import (
	"encoding/json"
)

// Queueconversationcobrowseeventtopicdetail
type Queueconversationcobrowseeventtopicdetail struct { 
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
func (o *Queueconversationcobrowseeventtopicdetail) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
