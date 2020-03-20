package platformclientv2
import (
	"encoding/json"
)

// Architectpromptnotificationerrordetail
type Architectpromptnotificationerrordetail struct { 
	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// EntityId
	EntityId *string `json:"entityId,omitempty"`


	// EntityName
	EntityName *string `json:"entityName,omitempty"`


	// FieldName
	FieldName *string `json:"fieldName,omitempty"`

}

// String returns a JSON representation of the model
func (o *Architectpromptnotificationerrordetail) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
