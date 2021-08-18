package platformclientv2
import (
	"github.com/leekchan/timeutil"
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
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (u *Dialercontactlistconfigchangecontactphonenumbercolumn) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercontactlistconfigchangecontactphonenumbercolumn

	

	return json.Marshal(&struct { 
		ColumnName *string `json:"columnName,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		CallableTimeColumn *string `json:"callableTimeColumn,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		ColumnName: u.ColumnName,
		
		VarType: u.VarType,
		
		CallableTimeColumn: u.CallableTimeColumn,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialercontactlistconfigchangecontactphonenumbercolumn) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
