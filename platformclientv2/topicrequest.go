package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Topicrequest
type Topicrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The topic name
	Name *string `json:"name,omitempty"`

	// Description - The topic description
	Description *string `json:"description,omitempty"`

	// Strictness - The topic strictness, default value is 72
	Strictness *string `json:"strictness,omitempty"`

	// ProgramIds - The ids of programs associated to the topic
	ProgramIds *[]string `json:"programIds,omitempty"`

	// Tags - The topic tags
	Tags *[]string `json:"tags,omitempty"`

	// Dialect - The topic dialect
	Dialect *string `json:"dialect,omitempty"`

	// Participants - The topic participants, default value is All
	Participants *string `json:"participants,omitempty"`

	// Phrases - The topic phrases
	Phrases *[]Phrase `json:"phrases,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Topicrequest) SetField(field string, fieldValue interface{}) {
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

func (o Topicrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Topicrequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Strictness *string `json:"strictness,omitempty"`
		
		ProgramIds *[]string `json:"programIds,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		
		Dialect *string `json:"dialect,omitempty"`
		
		Participants *string `json:"participants,omitempty"`
		
		Phrases *[]Phrase `json:"phrases,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		Description: o.Description,
		
		Strictness: o.Strictness,
		
		ProgramIds: o.ProgramIds,
		
		Tags: o.Tags,
		
		Dialect: o.Dialect,
		
		Participants: o.Participants,
		
		Phrases: o.Phrases,
		Alias:    (Alias)(o),
	})
}

func (o *Topicrequest) UnmarshalJSON(b []byte) error {
	var TopicrequestMap map[string]interface{}
	err := json.Unmarshal(b, &TopicrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := TopicrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := TopicrequestMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Strictness, ok := TopicrequestMap["strictness"].(string); ok {
		o.Strictness = &Strictness
	}
    
	if ProgramIds, ok := TopicrequestMap["programIds"].([]interface{}); ok {
		ProgramIdsString, _ := json.Marshal(ProgramIds)
		json.Unmarshal(ProgramIdsString, &o.ProgramIds)
	}
	
	if Tags, ok := TopicrequestMap["tags"].([]interface{}); ok {
		TagsString, _ := json.Marshal(Tags)
		json.Unmarshal(TagsString, &o.Tags)
	}
	
	if Dialect, ok := TopicrequestMap["dialect"].(string); ok {
		o.Dialect = &Dialect
	}
    
	if Participants, ok := TopicrequestMap["participants"].(string); ok {
		o.Participants = &Participants
	}
    
	if Phrases, ok := TopicrequestMap["phrases"].([]interface{}); ok {
		PhrasesString, _ := json.Marshal(Phrases)
		json.Unmarshal(PhrasesString, &o.Phrases)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Topicrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
