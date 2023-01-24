package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentoffer
type Contentoffer struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ImageUrl - URL for image displayed to the customer when displaying content offer.
	ImageUrl *string `json:"imageUrl,omitempty"`

	// DisplayMode - The display mode of Genesys Widgets when displaying content offer.
	DisplayMode *string `json:"displayMode,omitempty"`

	// LayoutMode - The layout mode of the text shown to the user when displaying content offer.
	LayoutMode *string `json:"layoutMode,omitempty"`

	// Title - Title used in the header of the content offer.
	Title *string `json:"title,omitempty"`

	// Headline - Headline displayed above the body text of the content offer.
	Headline *string `json:"headline,omitempty"`

	// Body - Body text of the content offer.
	Body *string `json:"body,omitempty"`

	// CallToAction - Properties customizing the call to action button on the content offer.
	CallToAction *Calltoaction `json:"callToAction,omitempty"`

	// Style - Properties customizing the styling of the content offer.
	Style *Contentofferstylingconfiguration `json:"style,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contentoffer) SetField(field string, fieldValue interface{}) {
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

func (o Contentoffer) MarshalJSON() ([]byte, error) {
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
	type Alias Contentoffer
	
	return json.Marshal(&struct { 
		ImageUrl *string `json:"imageUrl,omitempty"`
		
		DisplayMode *string `json:"displayMode,omitempty"`
		
		LayoutMode *string `json:"layoutMode,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Headline *string `json:"headline,omitempty"`
		
		Body *string `json:"body,omitempty"`
		
		CallToAction *Calltoaction `json:"callToAction,omitempty"`
		
		Style *Contentofferstylingconfiguration `json:"style,omitempty"`
		Alias
	}{ 
		ImageUrl: o.ImageUrl,
		
		DisplayMode: o.DisplayMode,
		
		LayoutMode: o.LayoutMode,
		
		Title: o.Title,
		
		Headline: o.Headline,
		
		Body: o.Body,
		
		CallToAction: o.CallToAction,
		
		Style: o.Style,
		Alias:    (Alias)(o),
	})
}

func (o *Contentoffer) UnmarshalJSON(b []byte) error {
	var ContentofferMap map[string]interface{}
	err := json.Unmarshal(b, &ContentofferMap)
	if err != nil {
		return err
	}
	
	if ImageUrl, ok := ContentofferMap["imageUrl"].(string); ok {
		o.ImageUrl = &ImageUrl
	}
    
	if DisplayMode, ok := ContentofferMap["displayMode"].(string); ok {
		o.DisplayMode = &DisplayMode
	}
    
	if LayoutMode, ok := ContentofferMap["layoutMode"].(string); ok {
		o.LayoutMode = &LayoutMode
	}
    
	if Title, ok := ContentofferMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Headline, ok := ContentofferMap["headline"].(string); ok {
		o.Headline = &Headline
	}
    
	if Body, ok := ContentofferMap["body"].(string); ok {
		o.Body = &Body
	}
    
	if CallToAction, ok := ContentofferMap["callToAction"].(map[string]interface{}); ok {
		CallToActionString, _ := json.Marshal(CallToAction)
		json.Unmarshal(CallToActionString, &o.CallToAction)
	}
	
	if Style, ok := ContentofferMap["style"].(map[string]interface{}); ok {
		StyleString, _ := json.Marshal(Style)
		json.Unmarshal(StyleString, &o.Style)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentoffer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
