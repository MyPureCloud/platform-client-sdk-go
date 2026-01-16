package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Updateguideversion - Request body for updating a guide version
type Updateguideversion struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Instruction - The instruction given to this version of the guide, for how it should behave when interacting with a User.
	Instruction *string `json:"instruction,omitempty"`

	// Variables - The variables associated with this version of the guide. Includes input variables (provided) and output variables (captured during execution).
	Variables *[]Variable `json:"variables,omitempty"`

	// Resources - The resources associated with this version of the guide.
	Resources *Guideversionresources `json:"resources,omitempty"`

	// KnowledgeSettings - The knowledge settings associated with this version of the guide.
	KnowledgeSettings *Authoringknowledgesettings `json:"knowledgeSettings,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Updateguideversion) SetField(field string, fieldValue interface{}) {
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

func (o Updateguideversion) MarshalJSON() ([]byte, error) {
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
	type Alias Updateguideversion
	
	return json.Marshal(&struct { 
		Instruction *string `json:"instruction,omitempty"`
		
		Variables *[]Variable `json:"variables,omitempty"`
		
		Resources *Guideversionresources `json:"resources,omitempty"`
		
		KnowledgeSettings *Authoringknowledgesettings `json:"knowledgeSettings,omitempty"`
		Alias
	}{ 
		Instruction: o.Instruction,
		
		Variables: o.Variables,
		
		Resources: o.Resources,
		
		KnowledgeSettings: o.KnowledgeSettings,
		Alias:    (Alias)(o),
	})
}

func (o *Updateguideversion) UnmarshalJSON(b []byte) error {
	var UpdateguideversionMap map[string]interface{}
	err := json.Unmarshal(b, &UpdateguideversionMap)
	if err != nil {
		return err
	}
	
	if Instruction, ok := UpdateguideversionMap["instruction"].(string); ok {
		o.Instruction = &Instruction
	}
    
	if Variables, ok := UpdateguideversionMap["variables"].([]interface{}); ok {
		VariablesString, _ := json.Marshal(Variables)
		json.Unmarshal(VariablesString, &o.Variables)
	}
	
	if Resources, ok := UpdateguideversionMap["resources"].(map[string]interface{}); ok {
		ResourcesString, _ := json.Marshal(Resources)
		json.Unmarshal(ResourcesString, &o.Resources)
	}
	
	if KnowledgeSettings, ok := UpdateguideversionMap["knowledgeSettings"].(map[string]interface{}); ok {
		KnowledgeSettingsString, _ := json.Marshal(KnowledgeSettings)
		json.Unmarshal(KnowledgeSettingsString, &o.KnowledgeSettings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updateguideversion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
