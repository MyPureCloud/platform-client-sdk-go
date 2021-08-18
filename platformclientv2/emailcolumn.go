package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Emailcolumn) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Emailcolumn

	

	return json.Marshal(&struct { 
		ColumnName *string `json:"columnName,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		ColumnName: u.ColumnName,
		
		VarType: u.VarType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Emailcolumn) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
