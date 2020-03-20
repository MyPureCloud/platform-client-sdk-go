package platformclientv2
import (
	"encoding/json"
)

// Dialercontactlistconfigchangecontactphonenumbercolumn
type Dialercontactlistconfigchangecontactphonenumbercolumn struct { 
	// ColumnName
	ColumnName *string `json:"columnName,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// CallableTimeColumn
	CallableTimeColumn *string `json:"callableTimeColumn,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercontactlistconfigchangecontactphonenumbercolumn) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
