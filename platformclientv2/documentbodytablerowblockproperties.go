package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentbodytablerowblockproperties
type Documentbodytablerowblockproperties struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// RowType - The type of the table row.
	RowType *string `json:"rowType,omitempty"`

	// Alignment - The alignment for the table row.
	Alignment *string `json:"alignment,omitempty"`

	// Height - The height for the table row.
	Height *float32 `json:"height,omitempty"`

	// BorderStyle - The border style for the table row.
	BorderStyle *string `json:"borderStyle,omitempty"`

	// BorderColor - The border color for the table row. For example black color - #000000
	BorderColor *string `json:"borderColor,omitempty"`

	// BackgroundColor - The background color for the table row. For example black color - #000000
	BackgroundColor *string `json:"backgroundColor,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Documentbodytablerowblockproperties) SetField(field string, fieldValue interface{}) {
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

func (o Documentbodytablerowblockproperties) MarshalJSON() ([]byte, error) {
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
	type Alias Documentbodytablerowblockproperties
	
	return json.Marshal(&struct { 
		RowType *string `json:"rowType,omitempty"`
		
		Alignment *string `json:"alignment,omitempty"`
		
		Height *float32 `json:"height,omitempty"`
		
		BorderStyle *string `json:"borderStyle,omitempty"`
		
		BorderColor *string `json:"borderColor,omitempty"`
		
		BackgroundColor *string `json:"backgroundColor,omitempty"`
		Alias
	}{ 
		RowType: o.RowType,
		
		Alignment: o.Alignment,
		
		Height: o.Height,
		
		BorderStyle: o.BorderStyle,
		
		BorderColor: o.BorderColor,
		
		BackgroundColor: o.BackgroundColor,
		Alias:    (Alias)(o),
	})
}

func (o *Documentbodytablerowblockproperties) UnmarshalJSON(b []byte) error {
	var DocumentbodytablerowblockpropertiesMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentbodytablerowblockpropertiesMap)
	if err != nil {
		return err
	}
	
	if RowType, ok := DocumentbodytablerowblockpropertiesMap["rowType"].(string); ok {
		o.RowType = &RowType
	}
    
	if Alignment, ok := DocumentbodytablerowblockpropertiesMap["alignment"].(string); ok {
		o.Alignment = &Alignment
	}
    
	if Height, ok := DocumentbodytablerowblockpropertiesMap["height"].(float64); ok {
		HeightFloat32 := float32(Height)
		o.Height = &HeightFloat32
	}
	
	if BorderStyle, ok := DocumentbodytablerowblockpropertiesMap["borderStyle"].(string); ok {
		o.BorderStyle = &BorderStyle
	}
    
	if BorderColor, ok := DocumentbodytablerowblockpropertiesMap["borderColor"].(string); ok {
		o.BorderColor = &BorderColor
	}
    
	if BackgroundColor, ok := DocumentbodytablerowblockpropertiesMap["backgroundColor"].(string); ok {
		o.BackgroundColor = &BackgroundColor
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Documentbodytablerowblockproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
