package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyviewelementfilternumberpredicate - A numeric filter on an element within a journey view
type Journeyviewelementfilternumberpredicate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Dimension - the element's attribute being filtered on
	Dimension *string `json:"dimension,omitempty"`

	// Operator - Optional operator, default is Matches. Valid values: Matches
	Operator *string `json:"operator,omitempty"`

	// NoValue - set this to true if no specific value to be considered
	NoValue *bool `json:"noValue,omitempty"`

	// VarRange - the range of comparators to filter on
	VarRange *Journeyviewelementfilterrange `json:"range,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeyviewelementfilternumberpredicate) SetField(field string, fieldValue interface{}) {
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

func (o Journeyviewelementfilternumberpredicate) MarshalJSON() ([]byte, error) {
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
	type Alias Journeyviewelementfilternumberpredicate
	
	return json.Marshal(&struct { 
		Dimension *string `json:"dimension,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		NoValue *bool `json:"noValue,omitempty"`
		
		VarRange *Journeyviewelementfilterrange `json:"range,omitempty"`
		Alias
	}{ 
		Dimension: o.Dimension,
		
		Operator: o.Operator,
		
		NoValue: o.NoValue,
		
		VarRange: o.VarRange,
		Alias:    (Alias)(o),
	})
}

func (o *Journeyviewelementfilternumberpredicate) UnmarshalJSON(b []byte) error {
	var JourneyviewelementfilternumberpredicateMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyviewelementfilternumberpredicateMap)
	if err != nil {
		return err
	}
	
	if Dimension, ok := JourneyviewelementfilternumberpredicateMap["dimension"].(string); ok {
		o.Dimension = &Dimension
	}
    
	if Operator, ok := JourneyviewelementfilternumberpredicateMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if NoValue, ok := JourneyviewelementfilternumberpredicateMap["noValue"].(bool); ok {
		o.NoValue = &NoValue
	}
    
	if VarRange, ok := JourneyviewelementfilternumberpredicateMap["range"].(map[string]interface{}); ok {
		VarRangeString, _ := json.Marshal(VarRange)
		json.Unmarshal(VarRangeString, &o.VarRange)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyviewelementfilternumberpredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
