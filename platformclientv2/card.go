package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Card
type Card struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Title - Text to show in the title.
	Title *string `json:"title,omitempty"`

	// Description - Text to show in the description.
	Description *string `json:"description,omitempty"`

	// Url - URL of an image.
	Url *string `json:"url,omitempty"`

	// DefaultAction - The default action to be taken.
	DefaultAction *Cardaction `json:"defaultAction,omitempty"`

	// Actions - List of possible action objects.
	Actions *[]Cardaction `json:"actions,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Card) SetField(field string, fieldValue interface{}) {
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

func (o Card) MarshalJSON() ([]byte, error) {
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
	type Alias Card
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		DefaultAction *Cardaction `json:"defaultAction,omitempty"`
		
		Actions *[]Cardaction `json:"actions,omitempty"`
		Alias
	}{ 
		Title: o.Title,
		
		Description: o.Description,
		
		Url: o.Url,
		
		DefaultAction: o.DefaultAction,
		
		Actions: o.Actions,
		Alias:    (Alias)(o),
	})
}

func (o *Card) UnmarshalJSON(b []byte) error {
	var CardMap map[string]interface{}
	err := json.Unmarshal(b, &CardMap)
	if err != nil {
		return err
	}
	
	if Title, ok := CardMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Description, ok := CardMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Url, ok := CardMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if DefaultAction, ok := CardMap["defaultAction"].(map[string]interface{}); ok {
		DefaultActionString, _ := json.Marshal(DefaultAction)
		json.Unmarshal(DefaultActionString, &o.DefaultAction)
	}
	
	if Actions, ok := CardMap["actions"].([]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Card) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
