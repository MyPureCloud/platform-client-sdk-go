package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Testexecutionoperationresult
type Testexecutionoperationresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Step - The step number to indicate the order in which the operation was performed
	Step *int `json:"step,omitempty"`

	// Name - Name of the operation performed
	Name *string `json:"name,omitempty"`

	// Success - Indicated whether or not the operation was successful
	Success *bool `json:"success,omitempty"`

	// Result - The result of the operation
	Result *interface{} `json:"result,omitempty"`

	// VarError - Error that occurred during the operation
	VarError *Errorbody `json:"error,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Testexecutionoperationresult) SetField(field string, fieldValue interface{}) {
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

func (o Testexecutionoperationresult) MarshalJSON() ([]byte, error) {
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
	type Alias Testexecutionoperationresult
	
	return json.Marshal(&struct { 
		Step *int `json:"step,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Success *bool `json:"success,omitempty"`
		
		Result *interface{} `json:"result,omitempty"`
		
		VarError *Errorbody `json:"error,omitempty"`
		Alias
	}{ 
		Step: o.Step,
		
		Name: o.Name,
		
		Success: o.Success,
		
		Result: o.Result,
		
		VarError: o.VarError,
		Alias:    (Alias)(o),
	})
}

func (o *Testexecutionoperationresult) UnmarshalJSON(b []byte) error {
	var TestexecutionoperationresultMap map[string]interface{}
	err := json.Unmarshal(b, &TestexecutionoperationresultMap)
	if err != nil {
		return err
	}
	
	if Step, ok := TestexecutionoperationresultMap["step"].(float64); ok {
		StepInt := int(Step)
		o.Step = &StepInt
	}
	
	if Name, ok := TestexecutionoperationresultMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Success, ok := TestexecutionoperationresultMap["success"].(bool); ok {
		o.Success = &Success
	}
    
	if Result, ok := TestexecutionoperationresultMap["result"].(map[string]interface{}); ok {
		ResultString, _ := json.Marshal(Result)
		json.Unmarshal(ResultString, &o.Result)
	}
	
	if VarError, ok := TestexecutionoperationresultMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Testexecutionoperationresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
