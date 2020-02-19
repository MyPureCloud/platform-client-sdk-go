package platformclientv2
import (
	"encoding/json"
)

// Phonenumbercolumn
type Phonenumbercolumn struct { 
	// ColumnName
	ColumnName *string `json:"columnName,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Phonenumbercolumn) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
