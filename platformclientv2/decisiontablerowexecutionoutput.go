package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Decisiontablerowexecutionoutput
type Decisiontablerowexecutionoutput struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// RowId - Unique rule identifier.
	RowId *string `json:"rowId,omitempty"`

	// RowIndex - Unique rule identifier.
	RowIndex *int `json:"rowIndex,omitempty"`

	// Outputs - The JSON output produced by this rule. Valid according to the execution output contract.
	Outputs *map[string]interface{} `json:"outputs,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Decisiontablerowexecutionoutput) SetField(field string, fieldValue interface{}) {
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

func (o Decisiontablerowexecutionoutput) MarshalJSON() ([]byte, error) {
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
	type Alias Decisiontablerowexecutionoutput
	
	return json.Marshal(&struct { 
		RowId *string `json:"rowId,omitempty"`
		
		RowIndex *int `json:"rowIndex,omitempty"`
		
		Outputs *map[string]interface{} `json:"outputs,omitempty"`
		Alias
	}{ 
		RowId: o.RowId,
		
		RowIndex: o.RowIndex,
		
		Outputs: o.Outputs,
		Alias:    (Alias)(o),
	})
}

func (o *Decisiontablerowexecutionoutput) UnmarshalJSON(b []byte) error {
	var DecisiontablerowexecutionoutputMap map[string]interface{}
	err := json.Unmarshal(b, &DecisiontablerowexecutionoutputMap)
	if err != nil {
		return err
	}
	
	if RowId, ok := DecisiontablerowexecutionoutputMap["rowId"].(string); ok {
		o.RowId = &RowId
	}
    
	if RowIndex, ok := DecisiontablerowexecutionoutputMap["rowIndex"].(float64); ok {
		RowIndexInt := int(RowIndex)
		o.RowIndex = &RowIndexInt
	}
	
	if Outputs, ok := DecisiontablerowexecutionoutputMap["outputs"].(map[string]interface{}); ok {
		OutputsString, _ := json.Marshal(Outputs)
		json.Unmarshal(OutputsString, &o.Outputs)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Decisiontablerowexecutionoutput) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
