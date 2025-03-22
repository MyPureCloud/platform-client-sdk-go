package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentbodytableproperties
type Documentbodytableproperties struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Width - The width of the table converted to em unit.
	Width *float32 `json:"width,omitempty"`

	// WidthWithUnit - The width of the table in the specified unit.
	WidthWithUnit *Documentelementlength `json:"widthWithUnit,omitempty"`

	// Height - The height for the table.
	Height *float32 `json:"height,omitempty"`

	// CellSpacing - The cell spacing for the table. The valid values in 'em'.
	CellSpacing *float32 `json:"cellSpacing,omitempty"`

	// CellPadding - The cell padding for the table. The valid values in 'em'.
	CellPadding *float32 `json:"cellPadding,omitempty"`

	// BorderWidth - The border width for the table. The valid values in 'em'
	BorderWidth *float32 `json:"borderWidth,omitempty"`

	// Alignment - The alignment for the table.
	Alignment *string `json:"alignment,omitempty"`

	// BorderStyle - The border style for the table.
	BorderStyle *string `json:"borderStyle,omitempty"`

	// BorderColor - The border color for the table. The valid values in hex color code representation. For example black color - #000000
	BorderColor *string `json:"borderColor,omitempty"`

	// BackgroundColor - The background color for the table. The valid values in hex color code representation. For example black color - #000000
	BackgroundColor *string `json:"backgroundColor,omitempty"`

	// Caption - The caption for the table. The valid values in hex color code representation. For example black color - #000000
	Caption *Documentbodytablecaptionblock `json:"caption,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Documentbodytableproperties) SetField(field string, fieldValue interface{}) {
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

func (o Documentbodytableproperties) MarshalJSON() ([]byte, error) {
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
	type Alias Documentbodytableproperties
	
	return json.Marshal(&struct { 
		Width *float32 `json:"width,omitempty"`
		
		WidthWithUnit *Documentelementlength `json:"widthWithUnit,omitempty"`
		
		Height *float32 `json:"height,omitempty"`
		
		CellSpacing *float32 `json:"cellSpacing,omitempty"`
		
		CellPadding *float32 `json:"cellPadding,omitempty"`
		
		BorderWidth *float32 `json:"borderWidth,omitempty"`
		
		Alignment *string `json:"alignment,omitempty"`
		
		BorderStyle *string `json:"borderStyle,omitempty"`
		
		BorderColor *string `json:"borderColor,omitempty"`
		
		BackgroundColor *string `json:"backgroundColor,omitempty"`
		
		Caption *Documentbodytablecaptionblock `json:"caption,omitempty"`
		Alias
	}{ 
		Width: o.Width,
		
		WidthWithUnit: o.WidthWithUnit,
		
		Height: o.Height,
		
		CellSpacing: o.CellSpacing,
		
		CellPadding: o.CellPadding,
		
		BorderWidth: o.BorderWidth,
		
		Alignment: o.Alignment,
		
		BorderStyle: o.BorderStyle,
		
		BorderColor: o.BorderColor,
		
		BackgroundColor: o.BackgroundColor,
		
		Caption: o.Caption,
		Alias:    (Alias)(o),
	})
}

func (o *Documentbodytableproperties) UnmarshalJSON(b []byte) error {
	var DocumentbodytablepropertiesMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentbodytablepropertiesMap)
	if err != nil {
		return err
	}
	
	if Width, ok := DocumentbodytablepropertiesMap["width"].(float64); ok {
		WidthFloat32 := float32(Width)
		o.Width = &WidthFloat32
	}
	
	if WidthWithUnit, ok := DocumentbodytablepropertiesMap["widthWithUnit"].(map[string]interface{}); ok {
		WidthWithUnitString, _ := json.Marshal(WidthWithUnit)
		json.Unmarshal(WidthWithUnitString, &o.WidthWithUnit)
	}
	
	if Height, ok := DocumentbodytablepropertiesMap["height"].(float64); ok {
		HeightFloat32 := float32(Height)
		o.Height = &HeightFloat32
	}
	
	if CellSpacing, ok := DocumentbodytablepropertiesMap["cellSpacing"].(float64); ok {
		CellSpacingFloat32 := float32(CellSpacing)
		o.CellSpacing = &CellSpacingFloat32
	}
	
	if CellPadding, ok := DocumentbodytablepropertiesMap["cellPadding"].(float64); ok {
		CellPaddingFloat32 := float32(CellPadding)
		o.CellPadding = &CellPaddingFloat32
	}
	
	if BorderWidth, ok := DocumentbodytablepropertiesMap["borderWidth"].(float64); ok {
		BorderWidthFloat32 := float32(BorderWidth)
		o.BorderWidth = &BorderWidthFloat32
	}
	
	if Alignment, ok := DocumentbodytablepropertiesMap["alignment"].(string); ok {
		o.Alignment = &Alignment
	}
    
	if BorderStyle, ok := DocumentbodytablepropertiesMap["borderStyle"].(string); ok {
		o.BorderStyle = &BorderStyle
	}
    
	if BorderColor, ok := DocumentbodytablepropertiesMap["borderColor"].(string); ok {
		o.BorderColor = &BorderColor
	}
    
	if BackgroundColor, ok := DocumentbodytablepropertiesMap["backgroundColor"].(string); ok {
		o.BackgroundColor = &BackgroundColor
	}
    
	if Caption, ok := DocumentbodytablepropertiesMap["caption"].(map[string]interface{}); ok {
		CaptionString, _ := json.Marshal(Caption)
		json.Unmarshal(CaptionString, &o.Caption)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentbodytableproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
