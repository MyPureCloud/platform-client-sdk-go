package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Criteriagroup - A group of logical or a singular criteria used to create a query of executionData
type Criteriagroup struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// And - These criteriaItems will be AND'd together to find a match.
	And *[]Criteriaitem `json:"and,omitempty"`

	// Or - These criteriaItems will be OR'd together to find a match.
	Or *[]Criteriaitem `json:"or,omitempty"`

	// Not - These criteriaItems must all be false to find a match.
	Not *[]Criteriaitem `json:"not,omitempty"`

	// Criteria - A singular critieriaItem to match.
	Criteria *Criteriaitem `json:"criteria,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Criteriagroup) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Criteriagroup) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Criteriagroup
	
	return json.Marshal(&struct { 
		And *[]Criteriaitem `json:"and,omitempty"`
		
		Or *[]Criteriaitem `json:"or,omitempty"`
		
		Not *[]Criteriaitem `json:"not,omitempty"`
		
		Criteria *Criteriaitem `json:"criteria,omitempty"`
		Alias
	}{ 
		And: o.And,
		
		Or: o.Or,
		
		Not: o.Not,
		
		Criteria: o.Criteria,
		Alias:    (Alias)(o),
	})
}

func (o *Criteriagroup) UnmarshalJSON(b []byte) error {
	var CriteriagroupMap map[string]interface{}
	err := json.Unmarshal(b, &CriteriagroupMap)
	if err != nil {
		return err
	}
	
	if And, ok := CriteriagroupMap["and"].([]interface{}); ok {
		AndString, _ := json.Marshal(And)
		json.Unmarshal(AndString, &o.And)
	}
	
	if Or, ok := CriteriagroupMap["or"].([]interface{}); ok {
		OrString, _ := json.Marshal(Or)
		json.Unmarshal(OrString, &o.Or)
	}
	
	if Not, ok := CriteriagroupMap["not"].([]interface{}); ok {
		NotString, _ := json.Marshal(Not)
		json.Unmarshal(NotString, &o.Not)
	}
	
	if Criteria, ok := CriteriagroupMap["criteria"].(map[string]interface{}); ok {
		CriteriaString, _ := json.Marshal(Criteria)
		json.Unmarshal(CriteriaString, &o.Criteria)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Criteriagroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
