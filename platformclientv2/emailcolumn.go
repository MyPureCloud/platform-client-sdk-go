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


	// ContactableTimeColumn - A column that indicates the timezone to use for a given contact when checking contactable times.
	ContactableTimeColumn *string `json:"contactableTimeColumn,omitempty"`

}

func (o *Emailcolumn) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Emailcolumn
	
	return json.Marshal(&struct { 
		ColumnName *string `json:"columnName,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		ContactableTimeColumn *string `json:"contactableTimeColumn,omitempty"`
		*Alias
	}{ 
		ColumnName: o.ColumnName,
		
		VarType: o.VarType,
		
		ContactableTimeColumn: o.ContactableTimeColumn,
		Alias:    (*Alias)(o),
	})
}

func (o *Emailcolumn) UnmarshalJSON(b []byte) error {
	var EmailcolumnMap map[string]interface{}
	err := json.Unmarshal(b, &EmailcolumnMap)
	if err != nil {
		return err
	}
	
	if ColumnName, ok := EmailcolumnMap["columnName"].(string); ok {
		o.ColumnName = &ColumnName
	}
    
	if VarType, ok := EmailcolumnMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if ContactableTimeColumn, ok := EmailcolumnMap["contactableTimeColumn"].(string); ok {
		o.ContactableTimeColumn = &ContactableTimeColumn
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Emailcolumn) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
