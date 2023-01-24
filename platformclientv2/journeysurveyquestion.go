package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeysurveyquestion
type Journeysurveyquestion struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - Type of survey question.
	VarType *string `json:"type,omitempty"`

	// Label - Label of question.
	Label *string `json:"label,omitempty"`

	// CustomerProperty - The customer property that the answer maps to.
	CustomerProperty *string `json:"customerProperty,omitempty"`

	// Choices - Choices available to user.
	Choices *[]string `json:"choices,omitempty"`

	// IsMandatory - Whether answering this question is mandatory.
	IsMandatory *bool `json:"isMandatory,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeysurveyquestion) SetField(field string, fieldValue interface{}) {
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

func (o Journeysurveyquestion) MarshalJSON() ([]byte, error) {
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
	type Alias Journeysurveyquestion
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Label *string `json:"label,omitempty"`
		
		CustomerProperty *string `json:"customerProperty,omitempty"`
		
		Choices *[]string `json:"choices,omitempty"`
		
		IsMandatory *bool `json:"isMandatory,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		Label: o.Label,
		
		CustomerProperty: o.CustomerProperty,
		
		Choices: o.Choices,
		
		IsMandatory: o.IsMandatory,
		Alias:    (Alias)(o),
	})
}

func (o *Journeysurveyquestion) UnmarshalJSON(b []byte) error {
	var JourneysurveyquestionMap map[string]interface{}
	err := json.Unmarshal(b, &JourneysurveyquestionMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := JourneysurveyquestionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Label, ok := JourneysurveyquestionMap["label"].(string); ok {
		o.Label = &Label
	}
    
	if CustomerProperty, ok := JourneysurveyquestionMap["customerProperty"].(string); ok {
		o.CustomerProperty = &CustomerProperty
	}
    
	if Choices, ok := JourneysurveyquestionMap["choices"].([]interface{}); ok {
		ChoicesString, _ := json.Marshal(Choices)
		json.Unmarshal(ChoicesString, &o.Choices)
	}
	
	if IsMandatory, ok := JourneysurveyquestionMap["isMandatory"].(bool); ok {
		o.IsMandatory = &IsMandatory
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeysurveyquestion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
