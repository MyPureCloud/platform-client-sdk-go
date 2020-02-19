package platformclientv2
import (
	"encoding/json"
)

// Phonecolumn
type Phonecolumn struct { 
	// ColumnName - The name of the phone column.
	ColumnName *string `json:"columnName,omitempty"`


	// VarType - The type of the phone column. For example, 'cell' or 'home'.
	VarType *string `json:"type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Phonecolumn) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
