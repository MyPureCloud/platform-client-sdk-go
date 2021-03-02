package platformclientv2
import (
	"encoding/json"
)

// Queueconversationvideoeventtopicphonenumbercolumn
type Queueconversationvideoeventtopicphonenumbercolumn struct { 
	// ColumnName
	ColumnName *string `json:"columnName,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicphonenumbercolumn) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
