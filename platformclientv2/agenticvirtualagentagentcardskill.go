package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agenticvirtualagentagentcardskill - A2A agent card skill.
type Agenticvirtualagentagentcardskill struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - Unique identifier for the skill.
	Id *string `json:"id,omitempty"`

	// Name - Human-readable name of the skill.
	Name *string `json:"name,omitempty"`

	// Description - Detailed explanation of what the skill does.
	Description *string `json:"description,omitempty"`

	// Tags - Keywords for categorization and discovery.
	Tags *[]string `json:"tags,omitempty"`

	// Examples - Sample prompts or use cases.
	Examples *[]string `json:"examples,omitempty"`

	// InputModes - Supported input media types.
	InputModes *[]string `json:"inputModes,omitempty"`

	// OutputModes - Supported output media types.
	OutputModes *[]string `json:"outputModes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agenticvirtualagentagentcardskill) SetField(field string, fieldValue interface{}) {
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

func (o Agenticvirtualagentagentcardskill) MarshalJSON() ([]byte, error) {
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
	type Alias Agenticvirtualagentagentcardskill
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		
		Examples *[]string `json:"examples,omitempty"`
		
		InputModes *[]string `json:"inputModes,omitempty"`
		
		OutputModes *[]string `json:"outputModes,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Tags: o.Tags,
		
		Examples: o.Examples,
		
		InputModes: o.InputModes,
		
		OutputModes: o.OutputModes,
		Alias:    (Alias)(o),
	})
}

func (o *Agenticvirtualagentagentcardskill) UnmarshalJSON(b []byte) error {
	var AgenticvirtualagentagentcardskillMap map[string]interface{}
	err := json.Unmarshal(b, &AgenticvirtualagentagentcardskillMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AgenticvirtualagentagentcardskillMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := AgenticvirtualagentagentcardskillMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := AgenticvirtualagentagentcardskillMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Tags, ok := AgenticvirtualagentagentcardskillMap["tags"].([]interface{}); ok {
		TagsString, _ := json.Marshal(Tags)
		json.Unmarshal(TagsString, &o.Tags)
	}
	
	if Examples, ok := AgenticvirtualagentagentcardskillMap["examples"].([]interface{}); ok {
		ExamplesString, _ := json.Marshal(Examples)
		json.Unmarshal(ExamplesString, &o.Examples)
	}
	
	if InputModes, ok := AgenticvirtualagentagentcardskillMap["inputModes"].([]interface{}); ok {
		InputModesString, _ := json.Marshal(InputModes)
		json.Unmarshal(InputModesString, &o.InputModes)
	}
	
	if OutputModes, ok := AgenticvirtualagentagentcardskillMap["outputModes"].([]interface{}); ok {
		OutputModesString, _ := json.Marshal(OutputModes)
		json.Unmarshal(OutputModesString, &o.OutputModes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentagentcardskill) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
