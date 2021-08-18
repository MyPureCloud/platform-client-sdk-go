package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactphonenumbercolumn
type Contactphonenumbercolumn struct { 
	// ColumnName - The name of the phone column.
	ColumnName *string `json:"columnName,omitempty"`


	// VarType - Indicates the type of the phone column. For example, 'cell' or 'home'.
	VarType *string `json:"type,omitempty"`


	// CallableTimeColumn - A column that indicates the timezone to use for a given contact when checking callable times. Not allowed if 'automaticTimeZoneMapping' is set to true.
	CallableTimeColumn *string `json:"callableTimeColumn,omitempty"`

}

func (u *Contactphonenumbercolumn) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactphonenumbercolumn

	

	return json.Marshal(&struct { 
		ColumnName *string `json:"columnName,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		CallableTimeColumn *string `json:"callableTimeColumn,omitempty"`
		*Alias
	}{ 
		ColumnName: u.ColumnName,
		
		VarType: u.VarType,
		
		CallableTimeColumn: u.CallableTimeColumn,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Contactphonenumbercolumn) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
