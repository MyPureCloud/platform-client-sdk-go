package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
