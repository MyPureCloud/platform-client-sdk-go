package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactlistfilterpredicate
type Contactlistfilterpredicate struct { 
	// Column - Contact list column from the ContactListFilter's contactList.
	Column *string `json:"column,omitempty"`


	// ColumnType - The type of data in the contact column.
	ColumnType *string `json:"columnType,omitempty"`


	// Operator - The operator for this ContactListFilterPredicate.
	Operator *string `json:"operator,omitempty"`


	// Value - Value with which to compare the contact's data. This could be text, a number, or a relative time. A value for relative time should follow the format PxxDTyyHzzM, where xx, yy, and zz specify the days, hours and minutes. For example, a value of P01DT08H30M corresponds to 1 day, 8 hours, and 30 minutes from now. To specify a time in the past, include a negative sign before each numeric value. For example, a value of P-01DT-08H-30M corresponds to 1 day, 8 hours, and 30 minutes in the past. You can also do things like P01DT00H-30M, which would correspond to 23 hours and 30 minutes from now (1 day - 30 minutes).
	Value *string `json:"value,omitempty"`


	// VarRange - A range of values. Required for operators BETWEEN and IN.
	VarRange *Contactlistfilterrange `json:"range,omitempty"`


	// Inverted - Inverts the result of the predicate (i.e., if the predicate returns true, inverting it will return false).
	Inverted *bool `json:"inverted,omitempty"`

}

func (o *Contactlistfilterpredicate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactlistfilterpredicate
	
	return json.Marshal(&struct { 
		Column *string `json:"column,omitempty"`
		
		ColumnType *string `json:"columnType,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		VarRange *Contactlistfilterrange `json:"range,omitempty"`
		
		Inverted *bool `json:"inverted,omitempty"`
		*Alias
	}{ 
		Column: o.Column,
		
		ColumnType: o.ColumnType,
		
		Operator: o.Operator,
		
		Value: o.Value,
		
		VarRange: o.VarRange,
		
		Inverted: o.Inverted,
		Alias:    (*Alias)(o),
	})
}

func (o *Contactlistfilterpredicate) UnmarshalJSON(b []byte) error {
	var ContactlistfilterpredicateMap map[string]interface{}
	err := json.Unmarshal(b, &ContactlistfilterpredicateMap)
	if err != nil {
		return err
	}
	
	if Column, ok := ContactlistfilterpredicateMap["column"].(string); ok {
		o.Column = &Column
	}
    
	if ColumnType, ok := ContactlistfilterpredicateMap["columnType"].(string); ok {
		o.ColumnType = &ColumnType
	}
    
	if Operator, ok := ContactlistfilterpredicateMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Value, ok := ContactlistfilterpredicateMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if VarRange, ok := ContactlistfilterpredicateMap["range"].(map[string]interface{}); ok {
		VarRangeString, _ := json.Marshal(VarRange)
		json.Unmarshal(VarRangeString, &o.VarRange)
	}
	
	if Inverted, ok := ContactlistfilterpredicateMap["inverted"].(bool); ok {
		o.Inverted = &Inverted
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contactlistfilterpredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
