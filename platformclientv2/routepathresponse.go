package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Routepathresponse
type Routepathresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Queue - The ID of the queue associated with the route path
	Queue *Queuereference `json:"queue,omitempty"`

	// MediaType - The media type of the given queue associated with the route path
	MediaType *string `json:"mediaType,omitempty"`

	// Language - The ID of the language associated with the route path
	Language *Languagereference `json:"language,omitempty"`

	// Skills - The set of skills associated with the route path
	Skills *[]Routingskillreference `json:"skills,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Routepathresponse) SetField(field string, fieldValue interface{}) {
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

func (o Routepathresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Routepathresponse
	
	return json.Marshal(&struct { 
		Queue *Queuereference `json:"queue,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		Language *Languagereference `json:"language,omitempty"`
		
		Skills *[]Routingskillreference `json:"skills,omitempty"`
		Alias
	}{ 
		Queue: o.Queue,
		
		MediaType: o.MediaType,
		
		Language: o.Language,
		
		Skills: o.Skills,
		Alias:    (Alias)(o),
	})
}

func (o *Routepathresponse) UnmarshalJSON(b []byte) error {
	var RoutepathresponseMap map[string]interface{}
	err := json.Unmarshal(b, &RoutepathresponseMap)
	if err != nil {
		return err
	}
	
	if Queue, ok := RoutepathresponseMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if MediaType, ok := RoutepathresponseMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Language, ok := RoutepathresponseMap["language"].(map[string]interface{}); ok {
		LanguageString, _ := json.Marshal(Language)
		json.Unmarshal(LanguageString, &o.Language)
	}
	
	if Skills, ok := RoutepathresponseMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Routepathresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
