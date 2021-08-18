package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Selectedcolumns
type Selectedcolumns struct { 
	// ColumnOrder - Indicates the order/position of the selected column
	ColumnOrder *int `json:"columnOrder,omitempty"`


	// ColumnName - Indicates enum name of the column from the export view
	ColumnName *string `json:"columnName,omitempty"`

}

func (u *Selectedcolumns) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Selectedcolumns

	

	return json.Marshal(&struct { 
		ColumnOrder *int `json:"columnOrder,omitempty"`
		
		ColumnName *string `json:"columnName,omitempty"`
		*Alias
	}{ 
		ColumnOrder: u.ColumnOrder,
		
		ColumnName: u.ColumnName,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Selectedcolumns) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
