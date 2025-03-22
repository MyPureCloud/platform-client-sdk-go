package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentbodytablecellblockproperties
type Documentbodytablecellblockproperties struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CellType - The type of the table cell.
	CellType *string `json:"cellType,omitempty"`

	// Width - The width of the table cell converted to em unit.
	Width *float32 `json:"width,omitempty"`

	// WidthWithUnit - The width of the table cell in the specified unit.
	WidthWithUnit *Documentelementlength `json:"widthWithUnit,omitempty"`

	// Height - The height for the table cell.
	Height *float32 `json:"height,omitempty"`

	// HorizontalAlign - The horizontal alignment for the table cell.
	HorizontalAlign *string `json:"horizontalAlign,omitempty"`

	// VerticalAlign - The vertical alignment for the table cell.
	VerticalAlign *string `json:"verticalAlign,omitempty"`

	// BorderWidth - The border width for the table cell. The valid values in 'em'
	BorderWidth *float32 `json:"borderWidth,omitempty"`

	// BorderStyle - The border style for the table cell.
	BorderStyle *string `json:"borderStyle,omitempty"`

	// BorderColor - The border color for the table cell. For example black color - #000000
	BorderColor *string `json:"borderColor,omitempty"`

	// BackgroundColor - The background color for the table cell. For example black color - #000000
	BackgroundColor *string `json:"backgroundColor,omitempty"`

	// Scope - The scope for the table cell.
	Scope *string `json:"scope,omitempty"`

	// ColSpan - The colSpan for the table cell.
	ColSpan *int `json:"colSpan,omitempty"`

	// RowSpan - The rowSpan for the table cell.
	RowSpan *int `json:"rowSpan,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Documentbodytablecellblockproperties) SetField(field string, fieldValue interface{}) {
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

func (o Documentbodytablecellblockproperties) MarshalJSON() ([]byte, error) {
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
	type Alias Documentbodytablecellblockproperties
	
	return json.Marshal(&struct { 
		CellType *string `json:"cellType,omitempty"`
		
		Width *float32 `json:"width,omitempty"`
		
		WidthWithUnit *Documentelementlength `json:"widthWithUnit,omitempty"`
		
		Height *float32 `json:"height,omitempty"`
		
		HorizontalAlign *string `json:"horizontalAlign,omitempty"`
		
		VerticalAlign *string `json:"verticalAlign,omitempty"`
		
		BorderWidth *float32 `json:"borderWidth,omitempty"`
		
		BorderStyle *string `json:"borderStyle,omitempty"`
		
		BorderColor *string `json:"borderColor,omitempty"`
		
		BackgroundColor *string `json:"backgroundColor,omitempty"`
		
		Scope *string `json:"scope,omitempty"`
		
		ColSpan *int `json:"colSpan,omitempty"`
		
		RowSpan *int `json:"rowSpan,omitempty"`
		Alias
	}{ 
		CellType: o.CellType,
		
		Width: o.Width,
		
		WidthWithUnit: o.WidthWithUnit,
		
		Height: o.Height,
		
		HorizontalAlign: o.HorizontalAlign,
		
		VerticalAlign: o.VerticalAlign,
		
		BorderWidth: o.BorderWidth,
		
		BorderStyle: o.BorderStyle,
		
		BorderColor: o.BorderColor,
		
		BackgroundColor: o.BackgroundColor,
		
		Scope: o.Scope,
		
		ColSpan: o.ColSpan,
		
		RowSpan: o.RowSpan,
		Alias:    (Alias)(o),
	})
}

func (o *Documentbodytablecellblockproperties) UnmarshalJSON(b []byte) error {
	var DocumentbodytablecellblockpropertiesMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentbodytablecellblockpropertiesMap)
	if err != nil {
		return err
	}
	
	if CellType, ok := DocumentbodytablecellblockpropertiesMap["cellType"].(string); ok {
		o.CellType = &CellType
	}
    
	if Width, ok := DocumentbodytablecellblockpropertiesMap["width"].(float64); ok {
		WidthFloat32 := float32(Width)
		o.Width = &WidthFloat32
	}
	
	if WidthWithUnit, ok := DocumentbodytablecellblockpropertiesMap["widthWithUnit"].(map[string]interface{}); ok {
		WidthWithUnitString, _ := json.Marshal(WidthWithUnit)
		json.Unmarshal(WidthWithUnitString, &o.WidthWithUnit)
	}
	
	if Height, ok := DocumentbodytablecellblockpropertiesMap["height"].(float64); ok {
		HeightFloat32 := float32(Height)
		o.Height = &HeightFloat32
	}
	
	if HorizontalAlign, ok := DocumentbodytablecellblockpropertiesMap["horizontalAlign"].(string); ok {
		o.HorizontalAlign = &HorizontalAlign
	}
    
	if VerticalAlign, ok := DocumentbodytablecellblockpropertiesMap["verticalAlign"].(string); ok {
		o.VerticalAlign = &VerticalAlign
	}
    
	if BorderWidth, ok := DocumentbodytablecellblockpropertiesMap["borderWidth"].(float64); ok {
		BorderWidthFloat32 := float32(BorderWidth)
		o.BorderWidth = &BorderWidthFloat32
	}
	
	if BorderStyle, ok := DocumentbodytablecellblockpropertiesMap["borderStyle"].(string); ok {
		o.BorderStyle = &BorderStyle
	}
    
	if BorderColor, ok := DocumentbodytablecellblockpropertiesMap["borderColor"].(string); ok {
		o.BorderColor = &BorderColor
	}
    
	if BackgroundColor, ok := DocumentbodytablecellblockpropertiesMap["backgroundColor"].(string); ok {
		o.BackgroundColor = &BackgroundColor
	}
    
	if Scope, ok := DocumentbodytablecellblockpropertiesMap["scope"].(string); ok {
		o.Scope = &Scope
	}
    
	if ColSpan, ok := DocumentbodytablecellblockpropertiesMap["colSpan"].(float64); ok {
		ColSpanInt := int(ColSpan)
		o.ColSpan = &ColSpanInt
	}
	
	if RowSpan, ok := DocumentbodytablecellblockpropertiesMap["rowSpan"].(float64); ok {
		RowSpanInt := int(RowSpan)
		o.RowSpan = &RowSpanInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentbodytablecellblockproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
