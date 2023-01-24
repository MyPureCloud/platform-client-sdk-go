package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Testmatcheseventoperation - Results from evaluating matching criteria against test input
type Testmatcheseventoperation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the processing step
	Name *string `json:"name,omitempty"`

	// Step - The number of the processing step
	Step *int `json:"step,omitempty"`

	// MatchedTriggers - Triggers that matched
	MatchedTriggers *[]Testmodetrigger `json:"matchedTriggers,omitempty"`

	// UnmatchedTriggers - Triggers that did not match
	UnmatchedTriggers *[]Testmodetrigger `json:"unmatchedTriggers,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Testmatcheseventoperation) SetField(field string, fieldValue interface{}) {
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

func (o Testmatcheseventoperation) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Testmatcheseventoperation
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Step *int `json:"step,omitempty"`
		
		MatchedTriggers *[]Testmodetrigger `json:"matchedTriggers,omitempty"`
		
		UnmatchedTriggers *[]Testmodetrigger `json:"unmatchedTriggers,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		Step: o.Step,
		
		MatchedTriggers: o.MatchedTriggers,
		
		UnmatchedTriggers: o.UnmatchedTriggers,
		Alias:    (Alias)(o),
	})
}

func (o *Testmatcheseventoperation) UnmarshalJSON(b []byte) error {
	var TestmatcheseventoperationMap map[string]interface{}
	err := json.Unmarshal(b, &TestmatcheseventoperationMap)
	if err != nil {
		return err
	}
	
	if Name, ok := TestmatcheseventoperationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Step, ok := TestmatcheseventoperationMap["step"].(float64); ok {
		StepInt := int(Step)
		o.Step = &StepInt
	}
	
	if MatchedTriggers, ok := TestmatcheseventoperationMap["matchedTriggers"].([]interface{}); ok {
		MatchedTriggersString, _ := json.Marshal(MatchedTriggers)
		json.Unmarshal(MatchedTriggersString, &o.MatchedTriggers)
	}
	
	if UnmatchedTriggers, ok := TestmatcheseventoperationMap["unmatchedTriggers"].([]interface{}); ok {
		UnmatchedTriggersString, _ := json.Marshal(UnmatchedTriggers)
		json.Unmarshal(UnmatchedTriggersString, &o.UnmatchedTriggers)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Testmatcheseventoperation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
