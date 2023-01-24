package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Skillgroupcondition
type Skillgroupcondition struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// RoutingSkillConditions - Routing skill conditions that will be used for building the query
	RoutingSkillConditions *[]Skillgrouproutingcondition `json:"routingSkillConditions,omitempty"`

	// LanguageSkillConditions - Routing skill conditions that will be used for building the query
	LanguageSkillConditions *[]Skillgrouplanguagecondition `json:"languageSkillConditions,omitempty"`

	// Operation - Operator that will be applied to the conditions
	Operation *string `json:"operation,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Skillgroupcondition) SetField(field string, fieldValue interface{}) {
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

func (o Skillgroupcondition) MarshalJSON() ([]byte, error) {
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
	type Alias Skillgroupcondition
	
	return json.Marshal(&struct { 
		RoutingSkillConditions *[]Skillgrouproutingcondition `json:"routingSkillConditions,omitempty"`
		
		LanguageSkillConditions *[]Skillgrouplanguagecondition `json:"languageSkillConditions,omitempty"`
		
		Operation *string `json:"operation,omitempty"`
		Alias
	}{ 
		RoutingSkillConditions: o.RoutingSkillConditions,
		
		LanguageSkillConditions: o.LanguageSkillConditions,
		
		Operation: o.Operation,
		Alias:    (Alias)(o),
	})
}

func (o *Skillgroupcondition) UnmarshalJSON(b []byte) error {
	var SkillgroupconditionMap map[string]interface{}
	err := json.Unmarshal(b, &SkillgroupconditionMap)
	if err != nil {
		return err
	}
	
	if RoutingSkillConditions, ok := SkillgroupconditionMap["routingSkillConditions"].([]interface{}); ok {
		RoutingSkillConditionsString, _ := json.Marshal(RoutingSkillConditions)
		json.Unmarshal(RoutingSkillConditionsString, &o.RoutingSkillConditions)
	}
	
	if LanguageSkillConditions, ok := SkillgroupconditionMap["languageSkillConditions"].([]interface{}); ok {
		LanguageSkillConditionsString, _ := json.Marshal(LanguageSkillConditions)
		json.Unmarshal(LanguageSkillConditionsString, &o.LanguageSkillConditions)
	}
	
	if Operation, ok := SkillgroupconditionMap["operation"].(string); ok {
		o.Operation = &Operation
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Skillgroupcondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
