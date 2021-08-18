package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontactlistfilterconfigchangefilterpredicate
type Dialercontactlistfilterconfigchangefilterpredicate struct { 
	// Column
	Column *string `json:"column,omitempty"`


	// ColumnType
	ColumnType *string `json:"columnType,omitempty"`


	// Operator
	Operator *string `json:"operator,omitempty"`


	// Value
	Value *string `json:"value,omitempty"`


	// VarRange
	VarRange *Dialercontactlistfilterconfigchangerange `json:"range,omitempty"`


	// Inverted
	Inverted *bool `json:"inverted,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (u *Dialercontactlistfilterconfigchangefilterpredicate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercontactlistfilterconfigchangefilterpredicate

	

	return json.Marshal(&struct { 
		Column *string `json:"column,omitempty"`
		
		ColumnType *string `json:"columnType,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		VarRange *Dialercontactlistfilterconfigchangerange `json:"range,omitempty"`
		
		Inverted *bool `json:"inverted,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Column: u.Column,
		
		ColumnType: u.ColumnType,
		
		Operator: u.Operator,
		
		Value: u.Value,
		
		VarRange: u.VarRange,
		
		Inverted: u.Inverted,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialercontactlistfilterconfigchangefilterpredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
