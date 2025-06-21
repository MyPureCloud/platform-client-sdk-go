package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeanddatesubcondition
type Timeanddatesubcondition struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - The type of time/date sub-condition.
	VarType *string `json:"type,omitempty"`

	// Operator - The operator to use for comparison.
	Operator *string `json:"operator,omitempty"`

	// Inverted - If true, inverts the result of evaluating this sub-condition. Default is false.
	Inverted *bool `json:"inverted,omitempty"`

	// IncludeYear - If true, includes year in date comparison for specificDate type. When false, only month and day are compared. Default is true. Only applicable for specificDate type.
	IncludeYear *bool `json:"includeYear,omitempty"`

	// ThresholdValue - Threshold value for BEFORE or AFTER operators. Format depends on type: timeOfDay: HH:mm, dayOfWeek: 1-7 (Monday-Sunday), dayOfMonth: 1-31 and/ or LAST_DAY, ODD_DAY, EVEN_DAY, specificDate: yyyy-MM-dd (if includeYear=true) or MM-dd (if includeYear=false). For single-value comparison, use a list with one element.
	ThresholdValue *string `json:"thresholdValue,omitempty"`

	// VarRange - A range of values for BETWEEN and IN operators. Format follows the same rules as 'thresholdValue'.
	VarRange *Timeanddatesubconditionrange `json:"range,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Timeanddatesubcondition) SetField(field string, fieldValue interface{}) {
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

func (o Timeanddatesubcondition) MarshalJSON() ([]byte, error) {
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
	type Alias Timeanddatesubcondition
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Inverted *bool `json:"inverted,omitempty"`
		
		IncludeYear *bool `json:"includeYear,omitempty"`
		
		ThresholdValue *string `json:"thresholdValue,omitempty"`
		
		VarRange *Timeanddatesubconditionrange `json:"range,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		Operator: o.Operator,
		
		Inverted: o.Inverted,
		
		IncludeYear: o.IncludeYear,
		
		ThresholdValue: o.ThresholdValue,
		
		VarRange: o.VarRange,
		Alias:    (Alias)(o),
	})
}

func (o *Timeanddatesubcondition) UnmarshalJSON(b []byte) error {
	var TimeanddatesubconditionMap map[string]interface{}
	err := json.Unmarshal(b, &TimeanddatesubconditionMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := TimeanddatesubconditionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Operator, ok := TimeanddatesubconditionMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Inverted, ok := TimeanddatesubconditionMap["inverted"].(bool); ok {
		o.Inverted = &Inverted
	}
    
	if IncludeYear, ok := TimeanddatesubconditionMap["includeYear"].(bool); ok {
		o.IncludeYear = &IncludeYear
	}
    
	if ThresholdValue, ok := TimeanddatesubconditionMap["thresholdValue"].(string); ok {
		o.ThresholdValue = &ThresholdValue
	}
    
	if VarRange, ok := TimeanddatesubconditionMap["range"].(map[string]interface{}); ok {
		VarRangeString, _ := json.Marshal(VarRange)
		json.Unmarshal(VarRangeString, &o.VarRange)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Timeanddatesubcondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
