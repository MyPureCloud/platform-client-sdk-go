package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Predictedanswer
type Predictedanswer struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// AnswerId - The unique identifier of the suggested predicted answer.
	AnswerId *string `json:"answerId,omitempty"`

	// Explanation - An explanation providing the reasoning behind the suggested answer.
	Explanation *string `json:"explanation,omitempty"`

	// FailureType - Describes the type of error associated with the predicted answer.
	FailureType *string `json:"failureType,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Predictedanswer) SetField(field string, fieldValue interface{}) {
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

func (o Predictedanswer) MarshalJSON() ([]byte, error) {
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
	type Alias Predictedanswer
	
	return json.Marshal(&struct { 
		AnswerId *string `json:"answerId,omitempty"`
		
		Explanation *string `json:"explanation,omitempty"`
		
		FailureType *string `json:"failureType,omitempty"`
		Alias
	}{ 
		AnswerId: o.AnswerId,
		
		Explanation: o.Explanation,
		
		FailureType: o.FailureType,
		Alias:    (Alias)(o),
	})
}

func (o *Predictedanswer) UnmarshalJSON(b []byte) error {
	var PredictedanswerMap map[string]interface{}
	err := json.Unmarshal(b, &PredictedanswerMap)
	if err != nil {
		return err
	}
	
	if AnswerId, ok := PredictedanswerMap["answerId"].(string); ok {
		o.AnswerId = &AnswerId
	}
    
	if Explanation, ok := PredictedanswerMap["explanation"].(string); ok {
		o.Explanation = &Explanation
	}
    
	if FailureType, ok := PredictedanswerMap["failureType"].(string); ok {
		o.FailureType = &FailureType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Predictedanswer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
