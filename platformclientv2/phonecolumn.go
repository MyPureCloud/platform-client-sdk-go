package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Phonecolumn
type Phonecolumn struct { 
	// ColumnName - The name of the phone column.
	ColumnName *string `json:"columnName,omitempty"`


	// VarType - The type of the phone column. For example, 'cell' or 'home'.
	VarType *string `json:"type,omitempty"`

}

func (u *Phonecolumn) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Phonecolumn

	

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
func (o *Phonecolumn) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
