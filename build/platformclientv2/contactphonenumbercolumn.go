package platformclientv2
import (
	"encoding/json"
)

// Contactphonenumbercolumn
type Contactphonenumbercolumn struct { 
	// ColumnName - The name of the phone column.
	ColumnName *string `json:"columnName,omitempty"`


	// VarType - Indicates the type of the phone column. For example, 'cell' or 'home'.
	VarType *string `json:"type,omitempty"`


	// CallableTimeColumn - A column that indicates the timezone to use for a given contact when checking callable times. Not allowed if 'automaticTimeZoneMapping' is set to true.
	CallableTimeColumn *string `json:"callableTimeColumn,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contactphonenumbercolumn) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
