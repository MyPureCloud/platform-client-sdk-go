package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentbodyparagraphproperties
type Documentbodyparagraphproperties struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// FontSize - The font size for the paragraph. The valid values in 'em'.
	FontSize *string `json:"fontSize,omitempty"`

	// FontType - The font type for the paragraph.
	FontType *string `json:"fontType,omitempty"`

	// TextColor - The text color for the paragraph. The valid values in hex color code representation. For example black color - #000000
	TextColor *string `json:"textColor,omitempty"`

	// BackgroundColor - The background color for the paragraph. The valid values in hex color code representation. For example black color - #000000
	BackgroundColor *string `json:"backgroundColor,omitempty"`

	// Align - The align type for the paragraph.
	Align *string `json:"align,omitempty"`

	// Indentation - The indentation color for the paragraph. The valid values in 'em'.
	Indentation *float32 `json:"indentation,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Documentbodyparagraphproperties) SetField(field string, fieldValue interface{}) {
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

func (o Documentbodyparagraphproperties) MarshalJSON() ([]byte, error) {
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
	type Alias Documentbodyparagraphproperties
	
	return json.Marshal(&struct { 
		FontSize *string `json:"fontSize,omitempty"`
		
		FontType *string `json:"fontType,omitempty"`
		
		TextColor *string `json:"textColor,omitempty"`
		
		BackgroundColor *string `json:"backgroundColor,omitempty"`
		
		Align *string `json:"align,omitempty"`
		
		Indentation *float32 `json:"indentation,omitempty"`
		Alias
	}{ 
		FontSize: o.FontSize,
		
		FontType: o.FontType,
		
		TextColor: o.TextColor,
		
		BackgroundColor: o.BackgroundColor,
		
		Align: o.Align,
		
		Indentation: o.Indentation,
		Alias:    (Alias)(o),
	})
}

func (o *Documentbodyparagraphproperties) UnmarshalJSON(b []byte) error {
	var DocumentbodyparagraphpropertiesMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentbodyparagraphpropertiesMap)
	if err != nil {
		return err
	}
	
	if FontSize, ok := DocumentbodyparagraphpropertiesMap["fontSize"].(string); ok {
		o.FontSize = &FontSize
	}
    
	if FontType, ok := DocumentbodyparagraphpropertiesMap["fontType"].(string); ok {
		o.FontType = &FontType
	}
    
	if TextColor, ok := DocumentbodyparagraphpropertiesMap["textColor"].(string); ok {
		o.TextColor = &TextColor
	}
    
	if BackgroundColor, ok := DocumentbodyparagraphpropertiesMap["backgroundColor"].(string); ok {
		o.BackgroundColor = &BackgroundColor
	}
    
	if Align, ok := DocumentbodyparagraphpropertiesMap["align"].(string); ok {
		o.Align = &Align
	}
    
	if Indentation, ok := DocumentbodyparagraphpropertiesMap["indentation"].(float64); ok {
		IndentationFloat32 := float32(Indentation)
		o.Indentation = &IndentationFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentbodyparagraphproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
