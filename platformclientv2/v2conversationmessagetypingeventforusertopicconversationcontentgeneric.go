package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationcontentgeneric
type V2conversationmessagetypingeventforusertopicconversationcontentgeneric struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Title
	Title *string `json:"title,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`

	// Image
	Image *string `json:"image,omitempty"`

	// Video
	Video *string `json:"video,omitempty"`

	// Actions
	Actions *V2conversationmessagetypingeventforusertopicconversationcontentactions `json:"actions,omitempty"`

	// Components
	Components *[]V2conversationmessagetypingeventforusertopicconversationbuttoncomponent `json:"components,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2conversationmessagetypingeventforusertopicconversationcontentgeneric) SetField(field string, fieldValue interface{}) {
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

func (o V2conversationmessagetypingeventforusertopicconversationcontentgeneric) MarshalJSON() ([]byte, error) {
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
	type Alias V2conversationmessagetypingeventforusertopicconversationcontentgeneric
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Image *string `json:"image,omitempty"`
		
		Video *string `json:"video,omitempty"`
		
		Actions *V2conversationmessagetypingeventforusertopicconversationcontentactions `json:"actions,omitempty"`
		
		Components *[]V2conversationmessagetypingeventforusertopicconversationbuttoncomponent `json:"components,omitempty"`
		Alias
	}{ 
		Title: o.Title,
		
		Description: o.Description,
		
		Image: o.Image,
		
		Video: o.Video,
		
		Actions: o.Actions,
		
		Components: o.Components,
		Alias:    (Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforusertopicconversationcontentgeneric) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationcontentgenericMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationcontentgenericMap)
	if err != nil {
		return err
	}
	
	if Title, ok := V2conversationmessagetypingeventforusertopicconversationcontentgenericMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Description, ok := V2conversationmessagetypingeventforusertopicconversationcontentgenericMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Image, ok := V2conversationmessagetypingeventforusertopicconversationcontentgenericMap["image"].(string); ok {
		o.Image = &Image
	}
    
	if Video, ok := V2conversationmessagetypingeventforusertopicconversationcontentgenericMap["video"].(string); ok {
		o.Video = &Video
	}
    
	if Actions, ok := V2conversationmessagetypingeventforusertopicconversationcontentgenericMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	
	if Components, ok := V2conversationmessagetypingeventforusertopicconversationcontentgenericMap["components"].([]interface{}); ok {
		ComponentsString, _ := json.Marshal(Components)
		json.Unmarshal(ComponentsString, &o.Components)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationcontentgeneric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
