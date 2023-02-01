package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Testexecutionresult
type Testexecutionresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Operations - Execution operations performed as part of the test
	Operations *[]Testexecutionoperationresult `json:"operations,omitempty"`

	// VarError - The final error encountered during the test that resulted in test failure
	VarError *Errorbody `json:"error,omitempty"`

	// FinalResult - The final result of the test. This is the response that would be returned during normal action execution
	FinalResult *interface{} `json:"finalResult,omitempty"`

	// Success - Indicates whether or not the test was a success
	Success *bool `json:"success,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Testexecutionresult) SetField(field string, fieldValue interface{}) {
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

func (o Testexecutionresult) MarshalJSON() ([]byte, error) {
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
	type Alias Testexecutionresult
	
	return json.Marshal(&struct { 
		Operations *[]Testexecutionoperationresult `json:"operations,omitempty"`
		
		VarError *Errorbody `json:"error,omitempty"`
		
		FinalResult *interface{} `json:"finalResult,omitempty"`
		
		Success *bool `json:"success,omitempty"`
		Alias
	}{ 
		Operations: o.Operations,
		
		VarError: o.VarError,
		
		FinalResult: o.FinalResult,
		
		Success: o.Success,
		Alias:    (Alias)(o),
	})
}

func (o *Testexecutionresult) UnmarshalJSON(b []byte) error {
	var TestexecutionresultMap map[string]interface{}
	err := json.Unmarshal(b, &TestexecutionresultMap)
	if err != nil {
		return err
	}
	
	if Operations, ok := TestexecutionresultMap["operations"].([]interface{}); ok {
		OperationsString, _ := json.Marshal(Operations)
		json.Unmarshal(OperationsString, &o.Operations)
	}
	
	if VarError, ok := TestexecutionresultMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	
	if FinalResult, ok := TestexecutionresultMap["finalResult"].(map[string]interface{}); ok {
		FinalResultString, _ := json.Marshal(FinalResult)
		json.Unmarshal(FinalResultString, &o.FinalResult)
	}
	
	if Success, ok := TestexecutionresultMap["success"].(bool); ok {
		o.Success = &Success
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Testexecutionresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
