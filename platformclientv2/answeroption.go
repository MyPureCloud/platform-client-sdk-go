package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Answeroption
type Answeroption struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// BuiltInType - The built-in type of this answer option. Only used for built-in answer options such as selection states for Multiple Select answer options. Possible values include: Selected, Unselected
	BuiltInType *string `json:"builtInType,omitempty"`

	// Text
	Text *string `json:"text,omitempty"`

	// Value
	Value *int `json:"value,omitempty"`

	// AssistanceConditions - List of assistance conditions which are combined together with a logical AND operator. Eg ( assistanceCondtion1 && assistanceCondition2 ) wherein assistanceCondition could be ( EXISTS topic1 || topic2 || ... ) or (NOTEXISTS topic3 || topic4 || ...).
	AssistanceConditions *[]Assistancecondition `json:"assistanceConditions,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Answeroption) SetField(field string, fieldValue interface{}) {
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

func (o Answeroption) MarshalJSON() ([]byte, error) {
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
	type Alias Answeroption
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		BuiltInType *string `json:"builtInType,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Value *int `json:"value,omitempty"`
		
		AssistanceConditions *[]Assistancecondition `json:"assistanceConditions,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		BuiltInType: o.BuiltInType,
		
		Text: o.Text,
		
		Value: o.Value,
		
		AssistanceConditions: o.AssistanceConditions,
		Alias:    (Alias)(o),
	})
}

func (o *Answeroption) UnmarshalJSON(b []byte) error {
	var AnsweroptionMap map[string]interface{}
	err := json.Unmarshal(b, &AnsweroptionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AnsweroptionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if BuiltInType, ok := AnsweroptionMap["builtInType"].(string); ok {
		o.BuiltInType = &BuiltInType
	}
    
	if Text, ok := AnsweroptionMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Value, ok := AnsweroptionMap["value"].(float64); ok {
		ValueInt := int(Value)
		o.Value = &ValueInt
	}
	
	if AssistanceConditions, ok := AnsweroptionMap["assistanceConditions"].([]interface{}); ok {
		AssistanceConditionsString, _ := json.Marshal(AssistanceConditions)
		json.Unmarshal(AssistanceConditionsString, &o.AssistanceConditions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Answeroption) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
