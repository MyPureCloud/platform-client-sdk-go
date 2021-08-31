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

func (o *Selectedcolumns) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Selectedcolumns
	
	return json.Marshal(&struct { 
		ColumnOrder *int `json:"columnOrder,omitempty"`
		
		ColumnName *string `json:"columnName,omitempty"`
		*Alias
	}{ 
		ColumnOrder: o.ColumnOrder,
		
		ColumnName: o.ColumnName,
		Alias:    (*Alias)(o),
	})
}

func (o *Selectedcolumns) UnmarshalJSON(b []byte) error {
	var SelectedcolumnsMap map[string]interface{}
	err := json.Unmarshal(b, &SelectedcolumnsMap)
	if err != nil {
		return err
	}
	
	if ColumnOrder, ok := SelectedcolumnsMap["columnOrder"].(float64); ok {
		ColumnOrderInt := int(ColumnOrder)
		o.ColumnOrder = &ColumnOrderInt
	}
	
	if ColumnName, ok := SelectedcolumnsMap["columnName"].(string); ok {
		o.ColumnName = &ColumnName
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Selectedcolumns) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
