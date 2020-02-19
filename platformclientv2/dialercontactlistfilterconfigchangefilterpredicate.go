package platformclientv2
import (
	"encoding/json"
)

// Dialercontactlistfilterconfigchangefilterpredicate
type Dialercontactlistfilterconfigchangefilterpredicate struct { 
	// Column
	Column *string `json:"column,omitempty"`


	// ColumnType
	ColumnType *string `json:"columnType,omitempty"`


	// Operator
	Operator *string `json:"operator,omitempty"`


	// Value
	Value *string `json:"value,omitempty"`


	// VarRange
	VarRange *Dialercontactlistfilterconfigchangerange `json:"range,omitempty"`


	// Inverted
	Inverted *bool `json:"inverted,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercontactlistfilterconfigchangefilterpredicate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
