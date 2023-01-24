package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Gkndocumentationresult
type Gkndocumentationresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Content - The text or html content for the documentation entity. Will be returned in responses for certain entities.
	Content *string `json:"content,omitempty"`

	// Link - URL link for the documentation entity. Will be returned in responses for certain entities.
	Link *string `json:"link,omitempty"`

	// Title - The title of the documentation entity. Will be returned in responses for certain entities.
	Title *string `json:"title,omitempty"`

	// VarType - The search type. Will be returned in responses for certain entities.
	VarType *string `json:"_type,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Gkndocumentationresult) SetField(field string, fieldValue interface{}) {
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

func (o Gkndocumentationresult) MarshalJSON() ([]byte, error) {
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
	type Alias Gkndocumentationresult
	
	return json.Marshal(&struct { 
		Content *string `json:"content,omitempty"`
		
		Link *string `json:"link,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		VarType *string `json:"_type,omitempty"`
		Alias
	}{ 
		Content: o.Content,
		
		Link: o.Link,
		
		Title: o.Title,
		
		VarType: o.VarType,
		Alias:    (Alias)(o),
	})
}

func (o *Gkndocumentationresult) UnmarshalJSON(b []byte) error {
	var GkndocumentationresultMap map[string]interface{}
	err := json.Unmarshal(b, &GkndocumentationresultMap)
	if err != nil {
		return err
	}
	
	if Content, ok := GkndocumentationresultMap["content"].(string); ok {
		o.Content = &Content
	}
    
	if Link, ok := GkndocumentationresultMap["link"].(string); ok {
		o.Link = &Link
	}
    
	if Title, ok := GkndocumentationresultMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if VarType, ok := GkndocumentationresultMap["_type"].(string); ok {
		o.VarType = &VarType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Gkndocumentationresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
