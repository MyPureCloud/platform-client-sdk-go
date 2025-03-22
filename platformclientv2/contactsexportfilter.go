package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactsexportfilter
type Contactsexportfilter struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Eq - Filtered field should have the same value
	Eq *Contactsexportfieldfilter `json:"eq,omitempty"`

	// In - Filtered field should match one of the listed values
	In *Contactsexportfieldlistfilter `json:"in,omitempty"`

	// And - Boolean AND combination of filters
	And *[]Contactsexportfilter `json:"and,omitempty"`

	// Or - Boolean OR combination of filters
	Or *[]Contactsexportfilter `json:"or,omitempty"`

	// Not - Boolean negation of filters
	Not **Contactsexportfilter `json:"not,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contactsexportfilter) SetField(field string, fieldValue interface{}) {
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

func (o Contactsexportfilter) MarshalJSON() ([]byte, error) {
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
	type Alias Contactsexportfilter
	
	return json.Marshal(&struct { 
		Eq *Contactsexportfieldfilter `json:"eq,omitempty"`
		
		In *Contactsexportfieldlistfilter `json:"in,omitempty"`
		
		And *[]Contactsexportfilter `json:"and,omitempty"`
		
		Or *[]Contactsexportfilter `json:"or,omitempty"`
		
		Not **Contactsexportfilter `json:"not,omitempty"`
		Alias
	}{ 
		Eq: o.Eq,
		
		In: o.In,
		
		And: o.And,
		
		Or: o.Or,
		
		Not: o.Not,
		Alias:    (Alias)(o),
	})
}

func (o *Contactsexportfilter) UnmarshalJSON(b []byte) error {
	var ContactsexportfilterMap map[string]interface{}
	err := json.Unmarshal(b, &ContactsexportfilterMap)
	if err != nil {
		return err
	}
	
	if Eq, ok := ContactsexportfilterMap["eq"].(map[string]interface{}); ok {
		EqString, _ := json.Marshal(Eq)
		json.Unmarshal(EqString, &o.Eq)
	}
	
	if In, ok := ContactsexportfilterMap["in"].(map[string]interface{}); ok {
		InString, _ := json.Marshal(In)
		json.Unmarshal(InString, &o.In)
	}
	
	if And, ok := ContactsexportfilterMap["and"].([]interface{}); ok {
		AndString, _ := json.Marshal(And)
		json.Unmarshal(AndString, &o.And)
	}
	
	if Or, ok := ContactsexportfilterMap["or"].([]interface{}); ok {
		OrString, _ := json.Marshal(Or)
		json.Unmarshal(OrString, &o.Or)
	}
	
	if Not, ok := ContactsexportfilterMap["not"].(map[string]interface{}); ok {
		NotString, _ := json.Marshal(Not)
		json.Unmarshal(NotString, &o.Not)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contactsexportfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
