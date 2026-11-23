package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capabilitiesunavailablereason
type Capabilitiesunavailablereason struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CustomCondition
	CustomCondition *string `json:"customCondition,omitempty"`

	// Enabled
	Enabled *bool `json:"enabled,omitempty"`

	// ExceptionEncountered
	ExceptionEncountered *bool `json:"exceptionEncountered,omitempty"`

	// Feature - The capability isn't available because a feature is not available
	Feature *string `json:"feature,omitempty"`

	// FeaturesAllOf - The capability isn't available because all of a given set of features are not available
	FeaturesAllOf *[]string `json:"featuresAllOf,omitempty"`

	// FeaturesOneOf - The capability isn't available because a feature isn't available
	FeaturesOneOf *[]string `json:"featuresOneOf,omitempty"`

	// ProductsAllOf - The capability isn't available because all of a given set of products are not available
	ProductsAllOf *[]string `json:"productsAllOf,omitempty"`

	// ProductsOneOf - The capability isn't available because a product is not available
	ProductsOneOf *[]string `json:"productsOneOf,omitempty"`

	// UnavailableFeature
	UnavailableFeature *string `json:"unavailableFeature,omitempty"`

	// UnavailableProduct
	UnavailableProduct *string `json:"unavailableProduct,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capabilitiesunavailablereason) SetField(field string, fieldValue interface{}) {
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

func (o Capabilitiesunavailablereason) MarshalJSON() ([]byte, error) {
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
	type Alias Capabilitiesunavailablereason
	
	return json.Marshal(&struct { 
		CustomCondition *string `json:"customCondition,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		ExceptionEncountered *bool `json:"exceptionEncountered,omitempty"`
		
		Feature *string `json:"feature,omitempty"`
		
		FeaturesAllOf *[]string `json:"featuresAllOf,omitempty"`
		
		FeaturesOneOf *[]string `json:"featuresOneOf,omitempty"`
		
		ProductsAllOf *[]string `json:"productsAllOf,omitempty"`
		
		ProductsOneOf *[]string `json:"productsOneOf,omitempty"`
		
		UnavailableFeature *string `json:"unavailableFeature,omitempty"`
		
		UnavailableProduct *string `json:"unavailableProduct,omitempty"`
		Alias
	}{ 
		CustomCondition: o.CustomCondition,
		
		Enabled: o.Enabled,
		
		ExceptionEncountered: o.ExceptionEncountered,
		
		Feature: o.Feature,
		
		FeaturesAllOf: o.FeaturesAllOf,
		
		FeaturesOneOf: o.FeaturesOneOf,
		
		ProductsAllOf: o.ProductsAllOf,
		
		ProductsOneOf: o.ProductsOneOf,
		
		UnavailableFeature: o.UnavailableFeature,
		
		UnavailableProduct: o.UnavailableProduct,
		Alias:    (Alias)(o),
	})
}

func (o *Capabilitiesunavailablereason) UnmarshalJSON(b []byte) error {
	var CapabilitiesunavailablereasonMap map[string]interface{}
	err := json.Unmarshal(b, &CapabilitiesunavailablereasonMap)
	if err != nil {
		return err
	}
	
	if CustomCondition, ok := CapabilitiesunavailablereasonMap["customCondition"].(string); ok {
		o.CustomCondition = &CustomCondition
	}
    
	if Enabled, ok := CapabilitiesunavailablereasonMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if ExceptionEncountered, ok := CapabilitiesunavailablereasonMap["exceptionEncountered"].(bool); ok {
		o.ExceptionEncountered = &ExceptionEncountered
	}
    
	if Feature, ok := CapabilitiesunavailablereasonMap["feature"].(string); ok {
		o.Feature = &Feature
	}
    
	if FeaturesAllOf, ok := CapabilitiesunavailablereasonMap["featuresAllOf"].([]interface{}); ok {
		FeaturesAllOfString, _ := json.Marshal(FeaturesAllOf)
		json.Unmarshal(FeaturesAllOfString, &o.FeaturesAllOf)
	}
	
	if FeaturesOneOf, ok := CapabilitiesunavailablereasonMap["featuresOneOf"].([]interface{}); ok {
		FeaturesOneOfString, _ := json.Marshal(FeaturesOneOf)
		json.Unmarshal(FeaturesOneOfString, &o.FeaturesOneOf)
	}
	
	if ProductsAllOf, ok := CapabilitiesunavailablereasonMap["productsAllOf"].([]interface{}); ok {
		ProductsAllOfString, _ := json.Marshal(ProductsAllOf)
		json.Unmarshal(ProductsAllOfString, &o.ProductsAllOf)
	}
	
	if ProductsOneOf, ok := CapabilitiesunavailablereasonMap["productsOneOf"].([]interface{}); ok {
		ProductsOneOfString, _ := json.Marshal(ProductsOneOf)
		json.Unmarshal(ProductsOneOfString, &o.ProductsOneOf)
	}
	
	if UnavailableFeature, ok := CapabilitiesunavailablereasonMap["unavailableFeature"].(string); ok {
		o.UnavailableFeature = &UnavailableFeature
	}
    
	if UnavailableProduct, ok := CapabilitiesunavailablereasonMap["unavailableProduct"].(string); ok {
		o.UnavailableProduct = &UnavailableProduct
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Capabilitiesunavailablereason) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
