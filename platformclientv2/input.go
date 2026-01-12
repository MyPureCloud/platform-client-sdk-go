package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Input - Input component configuration
type Input struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - Unique identifier for the input field
	Id *string `json:"id,omitempty"`

	// Title - Title of the input field
	Title *string `json:"title,omitempty"`

	// Subtitle - Subtitle of the input field
	Subtitle *string `json:"subtitle,omitempty"`

	// PlaceholderText - Placeholder text for the input
	PlaceholderText *string `json:"placeholderText,omitempty"`

	// IsMultipleLine - Whether the input supports multiple lines
	IsMultipleLine *bool `json:"isMultipleLine,omitempty"`

	// IsRequired - Whether the input is required
	IsRequired *bool `json:"isRequired,omitempty"`

	// KeyboardType - Type of keyboard to be shown
	KeyboardType *string `json:"keyboardType,omitempty"`

	// AutoCompleteType - A string value representing the keyboard and system information about the expected semantic meaning for the content that users enter
	AutoCompleteType *string `json:"autoCompleteType,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Input) SetField(field string, fieldValue interface{}) {
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

func (o Input) MarshalJSON() ([]byte, error) {
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
	type Alias Input
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Subtitle *string `json:"subtitle,omitempty"`
		
		PlaceholderText *string `json:"placeholderText,omitempty"`
		
		IsMultipleLine *bool `json:"isMultipleLine,omitempty"`
		
		IsRequired *bool `json:"isRequired,omitempty"`
		
		KeyboardType *string `json:"keyboardType,omitempty"`
		
		AutoCompleteType *string `json:"autoCompleteType,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Title: o.Title,
		
		Subtitle: o.Subtitle,
		
		PlaceholderText: o.PlaceholderText,
		
		IsMultipleLine: o.IsMultipleLine,
		
		IsRequired: o.IsRequired,
		
		KeyboardType: o.KeyboardType,
		
		AutoCompleteType: o.AutoCompleteType,
		Alias:    (Alias)(o),
	})
}

func (o *Input) UnmarshalJSON(b []byte) error {
	var InputMap map[string]interface{}
	err := json.Unmarshal(b, &InputMap)
	if err != nil {
		return err
	}
	
	if Id, ok := InputMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Title, ok := InputMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Subtitle, ok := InputMap["subtitle"].(string); ok {
		o.Subtitle = &Subtitle
	}
    
	if PlaceholderText, ok := InputMap["placeholderText"].(string); ok {
		o.PlaceholderText = &PlaceholderText
	}
    
	if IsMultipleLine, ok := InputMap["isMultipleLine"].(bool); ok {
		o.IsMultipleLine = &IsMultipleLine
	}
    
	if IsRequired, ok := InputMap["isRequired"].(bool); ok {
		o.IsRequired = &IsRequired
	}
    
	if KeyboardType, ok := InputMap["keyboardType"].(string); ok {
		o.KeyboardType = &KeyboardType
	}
    
	if AutoCompleteType, ok := InputMap["autoCompleteType"].(string); ok {
		o.AutoCompleteType = &AutoCompleteType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Input) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
