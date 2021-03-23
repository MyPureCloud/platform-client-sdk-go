package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Emailcolumn
type Emailcolumn struct { 
	// ColumnName - The name of the email column.
	ColumnName *string `json:"columnName,omitempty"`


	// VarType - Indicates the type of the email column. For example, 'work' or 'personal'.
	VarType *string `json:"type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Emailcolumn) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
