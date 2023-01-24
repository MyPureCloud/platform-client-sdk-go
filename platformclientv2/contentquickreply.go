package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentquickreply - Quick reply object.
type Contentquickreply struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - A unique ID assigned to the quick reply (Deprecated).
	Id *string `json:"id,omitempty"`

	// Text - Text to show inside the quick reply. This is also used as the response text after clicking on the quick reply.
	Text *string `json:"text,omitempty"`

	// Payload - Content of the payload included in the quick reply response. Could be an ID identifying the quick reply response.
	Payload *string `json:"payload,omitempty"`

	// Image - URL of an image associated with the quick reply.
	Image *string `json:"image,omitempty"`

	// Action - Specifies the type of action that is triggered upon clicking the quick reply.
	Action *string `json:"action,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contentquickreply) SetField(field string, fieldValue interface{}) {
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

func (o Contentquickreply) MarshalJSON() ([]byte, error) {
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
	type Alias Contentquickreply
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Payload *string `json:"payload,omitempty"`
		
		Image *string `json:"image,omitempty"`
		
		Action *string `json:"action,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Text: o.Text,
		
		Payload: o.Payload,
		
		Image: o.Image,
		
		Action: o.Action,
		Alias:    (Alias)(o),
	})
}

func (o *Contentquickreply) UnmarshalJSON(b []byte) error {
	var ContentquickreplyMap map[string]interface{}
	err := json.Unmarshal(b, &ContentquickreplyMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ContentquickreplyMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Text, ok := ContentquickreplyMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Payload, ok := ContentquickreplyMap["payload"].(string); ok {
		o.Payload = &Payload
	}
    
	if Image, ok := ContentquickreplyMap["image"].(string); ok {
		o.Image = &Image
	}
    
	if Action, ok := ContentquickreplyMap["action"].(string); ok {
		o.Action = &Action
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contentquickreply) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
