package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeanddatesubconditionrange
type Timeanddatesubconditionrange struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Min - The minimum value of the range. Required for the operator BETWEEN. Format depends on type: timeOfDay: HH:mm, dayOfWeek: 1-7 (Monday-Sunday), dayOfMonth: 1-31, specificDate: yyyy-MM-dd (if includeYear=true) or MM-dd (if includeYear=false).
	Min *string `json:"min,omitempty"`

	// Max - The maximum value of the range. Required for the operator BETWEEN. Format follows the same rules as 'min'.
	Max *string `json:"max,omitempty"`

	// InSet - A set of values that the date/ time data should be in. Required for the IN operator. Format depends on type: dayOfWeek: 1-7 (Monday-Sunday), dayOfMonth: 1-31, and/ or LAST_DAY, ODD_DAY, EVEN_DAY,specificDate: yyyy-MM-dd (if includeYear=true) or MM-dd (if includeYear=false).
	InSet *[]string `json:"inSet,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Timeanddatesubconditionrange) SetField(field string, fieldValue interface{}) {
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

func (o Timeanddatesubconditionrange) MarshalJSON() ([]byte, error) {
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
	type Alias Timeanddatesubconditionrange
	
	return json.Marshal(&struct { 
		Min *string `json:"min,omitempty"`
		
		Max *string `json:"max,omitempty"`
		
		InSet *[]string `json:"inSet,omitempty"`
		Alias
	}{ 
		Min: o.Min,
		
		Max: o.Max,
		
		InSet: o.InSet,
		Alias:    (Alias)(o),
	})
}

func (o *Timeanddatesubconditionrange) UnmarshalJSON(b []byte) error {
	var TimeanddatesubconditionrangeMap map[string]interface{}
	err := json.Unmarshal(b, &TimeanddatesubconditionrangeMap)
	if err != nil {
		return err
	}
	
	if Min, ok := TimeanddatesubconditionrangeMap["min"].(string); ok {
		o.Min = &Min
	}
    
	if Max, ok := TimeanddatesubconditionrangeMap["max"].(string); ok {
		o.Max = &Max
	}
    
	if InSet, ok := TimeanddatesubconditionrangeMap["inSet"].([]interface{}); ok {
		InSetString, _ := json.Marshal(InSet)
		json.Unmarshal(InSetString, &o.InSet)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Timeanddatesubconditionrange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
