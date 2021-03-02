package platformclientv2
import (
	"encoding/json"
)

// Conversationeventtopicphonenumbercolumn
type Conversationeventtopicphonenumbercolumn struct { 
	// ColumnName
	ColumnName *string `json:"columnName,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationeventtopicphonenumbercolumn) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
