package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Greetingmediainfo
type Greetingmediainfo struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// MediaFileUri
	MediaFileUri *string `json:"mediaFileUri,omitempty"`

	// MediaImageUri
	MediaImageUri *string `json:"mediaImageUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Greetingmediainfo) SetField(field string, fieldValue interface{}) {
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

func (o Greetingmediainfo) MarshalJSON() ([]byte, error) {
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
	type Alias Greetingmediainfo
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		MediaFileUri *string `json:"mediaFileUri,omitempty"`
		
		MediaImageUri *string `json:"mediaImageUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		MediaFileUri: o.MediaFileUri,
		
		MediaImageUri: o.MediaImageUri,
		Alias:    (Alias)(o),
	})
}

func (o *Greetingmediainfo) UnmarshalJSON(b []byte) error {
	var GreetingmediainfoMap map[string]interface{}
	err := json.Unmarshal(b, &GreetingmediainfoMap)
	if err != nil {
		return err
	}
	
	if Id, ok := GreetingmediainfoMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if MediaFileUri, ok := GreetingmediainfoMap["mediaFileUri"].(string); ok {
		o.MediaFileUri = &MediaFileUri
	}
    
	if MediaImageUri, ok := GreetingmediainfoMap["mediaImageUri"].(string); ok {
		o.MediaImageUri = &MediaImageUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Greetingmediainfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
