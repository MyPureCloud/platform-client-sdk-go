package platformclientv2
import (
	"encoding/json"
)

// Architectflowoutcomenotificationerrordetail
type Architectflowoutcomenotificationerrordetail struct { 
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
func (o *Architectflowoutcomenotificationerrordetail) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
