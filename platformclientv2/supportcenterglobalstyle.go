package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Supportcenterglobalstyle
type Supportcenterglobalstyle struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// BackgroundColor - Global background color, in hexadecimal format, eg #ffffff
	BackgroundColor *string `json:"backgroundColor,omitempty"`

	// PrimaryColor - Global primary color, in hexadecimal format, eg #ffffff
	PrimaryColor *string `json:"primaryColor,omitempty"`

	// PrimaryColorDark - Global dark primary color, in hexadecimal format, eg #ffffff
	PrimaryColorDark *string `json:"primaryColorDark,omitempty"`

	// PrimaryColorLight - Global light primary color, in hexadecimal format, eg #ffffff
	PrimaryColorLight *string `json:"primaryColorLight,omitempty"`

	// TextColor - Global text color, in hexadecimal format, eg #ffffff
	TextColor *string `json:"textColor,omitempty"`

	// FontFamily - Global font family
	FontFamily *string `json:"fontFamily,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Supportcenterglobalstyle) SetField(field string, fieldValue interface{}) {
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

func (o Supportcenterglobalstyle) MarshalJSON() ([]byte, error) {
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
	type Alias Supportcenterglobalstyle
	
	return json.Marshal(&struct { 
		BackgroundColor *string `json:"backgroundColor,omitempty"`
		
		PrimaryColor *string `json:"primaryColor,omitempty"`
		
		PrimaryColorDark *string `json:"primaryColorDark,omitempty"`
		
		PrimaryColorLight *string `json:"primaryColorLight,omitempty"`
		
		TextColor *string `json:"textColor,omitempty"`
		
		FontFamily *string `json:"fontFamily,omitempty"`
		Alias
	}{ 
		BackgroundColor: o.BackgroundColor,
		
		PrimaryColor: o.PrimaryColor,
		
		PrimaryColorDark: o.PrimaryColorDark,
		
		PrimaryColorLight: o.PrimaryColorLight,
		
		TextColor: o.TextColor,
		
		FontFamily: o.FontFamily,
		Alias:    (Alias)(o),
	})
}

func (o *Supportcenterglobalstyle) UnmarshalJSON(b []byte) error {
	var SupportcenterglobalstyleMap map[string]interface{}
	err := json.Unmarshal(b, &SupportcenterglobalstyleMap)
	if err != nil {
		return err
	}
	
	if BackgroundColor, ok := SupportcenterglobalstyleMap["backgroundColor"].(string); ok {
		o.BackgroundColor = &BackgroundColor
	}
    
	if PrimaryColor, ok := SupportcenterglobalstyleMap["primaryColor"].(string); ok {
		o.PrimaryColor = &PrimaryColor
	}
    
	if PrimaryColorDark, ok := SupportcenterglobalstyleMap["primaryColorDark"].(string); ok {
		o.PrimaryColorDark = &PrimaryColorDark
	}
    
	if PrimaryColorLight, ok := SupportcenterglobalstyleMap["primaryColorLight"].(string); ok {
		o.PrimaryColorLight = &PrimaryColorLight
	}
    
	if TextColor, ok := SupportcenterglobalstyleMap["textColor"].(string); ok {
		o.TextColor = &TextColor
	}
    
	if FontFamily, ok := SupportcenterglobalstyleMap["fontFamily"].(string); ok {
		o.FontFamily = &FontFamily
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Supportcenterglobalstyle) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
