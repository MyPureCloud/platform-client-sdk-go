package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentlistpickeritem - Represents a selectable item in a list picker.
type Conversationcontentlistpickeritem struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - Unique identifier for the list picker item
	Id *string `json:"id,omitempty"`

	// Title - The main text displayed for the item.
	Title *string `json:"title,omitempty"`

	// Subtitle - Additional text providing more details about the item.
	Subtitle *string `json:"subtitle,omitempty"`

	// ImageUrl - URL of an image to be displayed alongside the item.
	ImageUrl *string `json:"imageUrl,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationcontentlistpickeritem) SetField(field string, fieldValue interface{}) {
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

func (o Conversationcontentlistpickeritem) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationcontentlistpickeritem
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Subtitle *string `json:"subtitle,omitempty"`
		
		ImageUrl *string `json:"imageUrl,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Title: o.Title,
		
		Subtitle: o.Subtitle,
		
		ImageUrl: o.ImageUrl,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationcontentlistpickeritem) UnmarshalJSON(b []byte) error {
	var ConversationcontentlistpickeritemMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcontentlistpickeritemMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationcontentlistpickeritemMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Title, ok := ConversationcontentlistpickeritemMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Subtitle, ok := ConversationcontentlistpickeritemMap["subtitle"].(string); ok {
		o.Subtitle = &Subtitle
	}
    
	if ImageUrl, ok := ConversationcontentlistpickeritemMap["imageUrl"].(string); ok {
		o.ImageUrl = &ImageUrl
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontentlistpickeritem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
