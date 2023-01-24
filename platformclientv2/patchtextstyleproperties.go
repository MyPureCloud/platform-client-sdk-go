package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchtextstyleproperties
type Patchtextstyleproperties struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Color - Color of the text. (eg. #FFFFFF)
	Color *string `json:"color,omitempty"`

	// Font - Font of the text. (eg. Helvetica)
	Font *string `json:"font,omitempty"`

	// FontSize - Font size of the text. (eg. '12')
	FontSize *string `json:"fontSize,omitempty"`

	// TextAlign - Text alignment.
	TextAlign *string `json:"textAlign,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Patchtextstyleproperties) SetField(field string, fieldValue interface{}) {
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

func (o Patchtextstyleproperties) MarshalJSON() ([]byte, error) {
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
	type Alias Patchtextstyleproperties
	
	return json.Marshal(&struct { 
		Color *string `json:"color,omitempty"`
		
		Font *string `json:"font,omitempty"`
		
		FontSize *string `json:"fontSize,omitempty"`
		
		TextAlign *string `json:"textAlign,omitempty"`
		Alias
	}{ 
		Color: o.Color,
		
		Font: o.Font,
		
		FontSize: o.FontSize,
		
		TextAlign: o.TextAlign,
		Alias:    (Alias)(o),
	})
}

func (o *Patchtextstyleproperties) UnmarshalJSON(b []byte) error {
	var PatchtextstylepropertiesMap map[string]interface{}
	err := json.Unmarshal(b, &PatchtextstylepropertiesMap)
	if err != nil {
		return err
	}
	
	if Color, ok := PatchtextstylepropertiesMap["color"].(string); ok {
		o.Color = &Color
	}
    
	if Font, ok := PatchtextstylepropertiesMap["font"].(string); ok {
		o.Font = &Font
	}
    
	if FontSize, ok := PatchtextstylepropertiesMap["fontSize"].(string); ok {
		o.FontSize = &FontSize
	}
    
	if TextAlign, ok := PatchtextstylepropertiesMap["textAlign"].(string); ok {
		o.TextAlign = &TextAlign
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Patchtextstyleproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
