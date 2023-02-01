package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Webchatroutingtarget
type Webchatroutingtarget struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// TargetType - The target type of the routing target, such as 'QUEUE'.
	TargetType *string `json:"targetType,omitempty"`

	// TargetAddress - The target of the route, in the format appropriate given the 'targetType'.
	TargetAddress *string `json:"targetAddress,omitempty"`

	// Skills - The list of skill names to use for routing.
	Skills *[]string `json:"skills,omitempty"`

	// Language - The language name to use for routing.
	Language *string `json:"language,omitempty"`

	// Priority - The priority to assign to the conversation for routing.
	Priority *int `json:"priority,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Webchatroutingtarget) SetField(field string, fieldValue interface{}) {
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

func (o Webchatroutingtarget) MarshalJSON() ([]byte, error) {
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
	type Alias Webchatroutingtarget
	
	return json.Marshal(&struct { 
		TargetType *string `json:"targetType,omitempty"`
		
		TargetAddress *string `json:"targetAddress,omitempty"`
		
		Skills *[]string `json:"skills,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		Alias
	}{ 
		TargetType: o.TargetType,
		
		TargetAddress: o.TargetAddress,
		
		Skills: o.Skills,
		
		Language: o.Language,
		
		Priority: o.Priority,
		Alias:    (Alias)(o),
	})
}

func (o *Webchatroutingtarget) UnmarshalJSON(b []byte) error {
	var WebchatroutingtargetMap map[string]interface{}
	err := json.Unmarshal(b, &WebchatroutingtargetMap)
	if err != nil {
		return err
	}
	
	if TargetType, ok := WebchatroutingtargetMap["targetType"].(string); ok {
		o.TargetType = &TargetType
	}
    
	if TargetAddress, ok := WebchatroutingtargetMap["targetAddress"].(string); ok {
		o.TargetAddress = &TargetAddress
	}
    
	if Skills, ok := WebchatroutingtargetMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if Language, ok := WebchatroutingtargetMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if Priority, ok := WebchatroutingtargetMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webchatroutingtarget) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
