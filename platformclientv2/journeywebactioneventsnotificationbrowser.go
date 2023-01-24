package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebactioneventsnotificationbrowser
type Journeywebactioneventsnotificationbrowser struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Family
	Family *string `json:"family,omitempty"`

	// Version
	Version *string `json:"version,omitempty"`

	// Lang
	Lang *string `json:"lang,omitempty"`

	// Fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	// ViewHeight
	ViewHeight *int `json:"viewHeight,omitempty"`

	// ViewWidth
	ViewWidth *int `json:"viewWidth,omitempty"`

	// FeaturesFlash
	FeaturesFlash *bool `json:"featuresFlash,omitempty"`

	// FeaturesJava
	FeaturesJava *bool `json:"featuresJava,omitempty"`

	// FeaturesPdf
	FeaturesPdf *bool `json:"featuresPdf,omitempty"`

	// FeaturesWebrtc
	FeaturesWebrtc *bool `json:"featuresWebrtc,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeywebactioneventsnotificationbrowser) SetField(field string, fieldValue interface{}) {
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

func (o Journeywebactioneventsnotificationbrowser) MarshalJSON() ([]byte, error) {
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
	type Alias Journeywebactioneventsnotificationbrowser
	
	return json.Marshal(&struct { 
		Family *string `json:"family,omitempty"`
		
		Version *string `json:"version,omitempty"`
		
		Lang *string `json:"lang,omitempty"`
		
		Fingerprint *string `json:"fingerprint,omitempty"`
		
		ViewHeight *int `json:"viewHeight,omitempty"`
		
		ViewWidth *int `json:"viewWidth,omitempty"`
		
		FeaturesFlash *bool `json:"featuresFlash,omitempty"`
		
		FeaturesJava *bool `json:"featuresJava,omitempty"`
		
		FeaturesPdf *bool `json:"featuresPdf,omitempty"`
		
		FeaturesWebrtc *bool `json:"featuresWebrtc,omitempty"`
		Alias
	}{ 
		Family: o.Family,
		
		Version: o.Version,
		
		Lang: o.Lang,
		
		Fingerprint: o.Fingerprint,
		
		ViewHeight: o.ViewHeight,
		
		ViewWidth: o.ViewWidth,
		
		FeaturesFlash: o.FeaturesFlash,
		
		FeaturesJava: o.FeaturesJava,
		
		FeaturesPdf: o.FeaturesPdf,
		
		FeaturesWebrtc: o.FeaturesWebrtc,
		Alias:    (Alias)(o),
	})
}

func (o *Journeywebactioneventsnotificationbrowser) UnmarshalJSON(b []byte) error {
	var JourneywebactioneventsnotificationbrowserMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebactioneventsnotificationbrowserMap)
	if err != nil {
		return err
	}
	
	if Family, ok := JourneywebactioneventsnotificationbrowserMap["family"].(string); ok {
		o.Family = &Family
	}
    
	if Version, ok := JourneywebactioneventsnotificationbrowserMap["version"].(string); ok {
		o.Version = &Version
	}
    
	if Lang, ok := JourneywebactioneventsnotificationbrowserMap["lang"].(string); ok {
		o.Lang = &Lang
	}
    
	if Fingerprint, ok := JourneywebactioneventsnotificationbrowserMap["fingerprint"].(string); ok {
		o.Fingerprint = &Fingerprint
	}
    
	if ViewHeight, ok := JourneywebactioneventsnotificationbrowserMap["viewHeight"].(float64); ok {
		ViewHeightInt := int(ViewHeight)
		o.ViewHeight = &ViewHeightInt
	}
	
	if ViewWidth, ok := JourneywebactioneventsnotificationbrowserMap["viewWidth"].(float64); ok {
		ViewWidthInt := int(ViewWidth)
		o.ViewWidth = &ViewWidthInt
	}
	
	if FeaturesFlash, ok := JourneywebactioneventsnotificationbrowserMap["featuresFlash"].(bool); ok {
		o.FeaturesFlash = &FeaturesFlash
	}
    
	if FeaturesJava, ok := JourneywebactioneventsnotificationbrowserMap["featuresJava"].(bool); ok {
		o.FeaturesJava = &FeaturesJava
	}
    
	if FeaturesPdf, ok := JourneywebactioneventsnotificationbrowserMap["featuresPdf"].(bool); ok {
		o.FeaturesPdf = &FeaturesPdf
	}
    
	if FeaturesWebrtc, ok := JourneywebactioneventsnotificationbrowserMap["featuresWebrtc"].(bool); ok {
		o.FeaturesWebrtc = &FeaturesWebrtc
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebactioneventsnotificationbrowser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
