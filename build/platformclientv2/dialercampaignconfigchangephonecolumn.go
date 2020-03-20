package platformclientv2
import (
	"encoding/json"
)

// Dialercampaignconfigchangephonecolumn
type Dialercampaignconfigchangephonecolumn struct { 
	// ColumnName
	ColumnName *string `json:"columnName,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercampaignconfigchangephonecolumn) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
