package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Querycriteriagroup - A group of logical items for library queries
type Querycriteriagroup struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// And - Items that will be AND'd together
	And *[]Querycriteriaitem `json:"and,omitempty"`

	// Or - Items that will be OR'd together
	Or *[]Querycriteriaitem `json:"or,omitempty"`

	// Not - Items that must all be false
	Not *[]Querycriteriaitem `json:"not,omitempty"`

	// Criteria - A single item
	Criteria *Querycriteriaitem `json:"criteria,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Querycriteriagroup) SetField(field string, fieldValue interface{}) {
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

func (o Querycriteriagroup) MarshalJSON() ([]byte, error) {
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
	type Alias Querycriteriagroup
	
	return json.Marshal(&struct { 
		And *[]Querycriteriaitem `json:"and,omitempty"`
		
		Or *[]Querycriteriaitem `json:"or,omitempty"`
		
		Not *[]Querycriteriaitem `json:"not,omitempty"`
		
		Criteria *Querycriteriaitem `json:"criteria,omitempty"`
		Alias
	}{ 
		And: o.And,
		
		Or: o.Or,
		
		Not: o.Not,
		
		Criteria: o.Criteria,
		Alias:    (Alias)(o),
	})
}

func (o *Querycriteriagroup) UnmarshalJSON(b []byte) error {
	var QuerycriteriagroupMap map[string]interface{}
	err := json.Unmarshal(b, &QuerycriteriagroupMap)
	if err != nil {
		return err
	}
	
	if And, ok := QuerycriteriagroupMap["and"].([]interface{}); ok {
		AndString, _ := json.Marshal(And)
		json.Unmarshal(AndString, &o.And)
	}
	
	if Or, ok := QuerycriteriagroupMap["or"].([]interface{}); ok {
		OrString, _ := json.Marshal(Or)
		json.Unmarshal(OrString, &o.Or)
	}
	
	if Not, ok := QuerycriteriagroupMap["not"].([]interface{}); ok {
		NotString, _ := json.Marshal(Not)
		json.Unmarshal(NotString, &o.Not)
	}
	
	if Criteria, ok := QuerycriteriagroupMap["criteria"].(map[string]interface{}); ok {
		CriteriaString, _ := json.Marshal(Criteria)
		json.Unmarshal(CriteriaString, &o.Criteria)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Querycriteriagroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
