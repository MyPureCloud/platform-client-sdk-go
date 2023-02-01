package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentofferstylingconfiguration
type Contentofferstylingconfiguration struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Position - Properties for customizing the positioning of the content offer.
	Position *Contentpositionproperties `json:"position,omitempty"`

	// Offer - Properties for customizing the appearance of the content offer.
	Offer *Contentofferstyleproperties `json:"offer,omitempty"`

	// CloseButton - Properties for customizing the appearance of the close button.
	CloseButton *Closebuttonstyleproperties `json:"closeButton,omitempty"`

	// CtaButton - Properties for customizing the appearance of the CTA button.
	CtaButton *Ctabuttonstyleproperties `json:"ctaButton,omitempty"`

	// Title - Properties for customizing the appearance of the title text.
	Title *Textstyleproperties `json:"title,omitempty"`

	// Headline - Properties for customizing the appearance of the headline text.
	Headline *Textstyleproperties `json:"headline,omitempty"`

	// Body - Properties for customizing the appearance of the body text.
	Body *Textstyleproperties `json:"body,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contentofferstylingconfiguration) SetField(field string, fieldValue interface{}) {
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

func (o Contentofferstylingconfiguration) MarshalJSON() ([]byte, error) {
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
	type Alias Contentofferstylingconfiguration
	
	return json.Marshal(&struct { 
		Position *Contentpositionproperties `json:"position,omitempty"`
		
		Offer *Contentofferstyleproperties `json:"offer,omitempty"`
		
		CloseButton *Closebuttonstyleproperties `json:"closeButton,omitempty"`
		
		CtaButton *Ctabuttonstyleproperties `json:"ctaButton,omitempty"`
		
		Title *Textstyleproperties `json:"title,omitempty"`
		
		Headline *Textstyleproperties `json:"headline,omitempty"`
		
		Body *Textstyleproperties `json:"body,omitempty"`
		Alias
	}{ 
		Position: o.Position,
		
		Offer: o.Offer,
		
		CloseButton: o.CloseButton,
		
		CtaButton: o.CtaButton,
		
		Title: o.Title,
		
		Headline: o.Headline,
		
		Body: o.Body,
		Alias:    (Alias)(o),
	})
}

func (o *Contentofferstylingconfiguration) UnmarshalJSON(b []byte) error {
	var ContentofferstylingconfigurationMap map[string]interface{}
	err := json.Unmarshal(b, &ContentofferstylingconfigurationMap)
	if err != nil {
		return err
	}
	
	if Position, ok := ContentofferstylingconfigurationMap["position"].(map[string]interface{}); ok {
		PositionString, _ := json.Marshal(Position)
		json.Unmarshal(PositionString, &o.Position)
	}
	
	if Offer, ok := ContentofferstylingconfigurationMap["offer"].(map[string]interface{}); ok {
		OfferString, _ := json.Marshal(Offer)
		json.Unmarshal(OfferString, &o.Offer)
	}
	
	if CloseButton, ok := ContentofferstylingconfigurationMap["closeButton"].(map[string]interface{}); ok {
		CloseButtonString, _ := json.Marshal(CloseButton)
		json.Unmarshal(CloseButtonString, &o.CloseButton)
	}
	
	if CtaButton, ok := ContentofferstylingconfigurationMap["ctaButton"].(map[string]interface{}); ok {
		CtaButtonString, _ := json.Marshal(CtaButton)
		json.Unmarshal(CtaButtonString, &o.CtaButton)
	}
	
	if Title, ok := ContentofferstylingconfigurationMap["title"].(map[string]interface{}); ok {
		TitleString, _ := json.Marshal(Title)
		json.Unmarshal(TitleString, &o.Title)
	}
	
	if Headline, ok := ContentofferstylingconfigurationMap["headline"].(map[string]interface{}); ok {
		HeadlineString, _ := json.Marshal(Headline)
		json.Unmarshal(HeadlineString, &o.Headline)
	}
	
	if Body, ok := ContentofferstylingconfigurationMap["body"].(map[string]interface{}); ok {
		BodyString, _ := json.Marshal(Body)
		json.Unmarshal(BodyString, &o.Body)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentofferstylingconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
