package platformclientv2
import (
	"encoding/json"
)

// Selectedcolumns
type Selectedcolumns struct { 
	// ColumnOrder - Indicates the order/position of the selected column
	ColumnOrder *int `json:"columnOrder,omitempty"`


	// ColumnName - Indicates enum name of the column from the export view
	ColumnName *string `json:"columnName,omitempty"`

}

// String returns a JSON representation of the model
func (o *Selectedcolumns) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
