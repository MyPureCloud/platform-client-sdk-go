package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Facetentry
type Facetentry struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Attribute
	Attribute *Termattribute `json:"attribute,omitempty"`

	// Statistics
	Statistics *Facetstatistics `json:"statistics,omitempty"`

	// Other
	Other *int `json:"other,omitempty"`

	// Total
	Total *int `json:"total,omitempty"`

	// Missing
	Missing *int `json:"missing,omitempty"`

	// TermCount
	TermCount *int `json:"termCount,omitempty"`

	// TermType
	TermType *string `json:"termType,omitempty"`

	// Terms
	Terms *[]Facetterm `json:"terms,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Facetentry) SetField(field string, fieldValue interface{}) {
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

func (o Facetentry) MarshalJSON() ([]byte, error) {
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
	type Alias Facetentry
	
	return json.Marshal(&struct { 
		Attribute *Termattribute `json:"attribute,omitempty"`
		
		Statistics *Facetstatistics `json:"statistics,omitempty"`
		
		Other *int `json:"other,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		Missing *int `json:"missing,omitempty"`
		
		TermCount *int `json:"termCount,omitempty"`
		
		TermType *string `json:"termType,omitempty"`
		
		Terms *[]Facetterm `json:"terms,omitempty"`
		Alias
	}{ 
		Attribute: o.Attribute,
		
		Statistics: o.Statistics,
		
		Other: o.Other,
		
		Total: o.Total,
		
		Missing: o.Missing,
		
		TermCount: o.TermCount,
		
		TermType: o.TermType,
		
		Terms: o.Terms,
		Alias:    (Alias)(o),
	})
}

func (o *Facetentry) UnmarshalJSON(b []byte) error {
	var FacetentryMap map[string]interface{}
	err := json.Unmarshal(b, &FacetentryMap)
	if err != nil {
		return err
	}
	
	if Attribute, ok := FacetentryMap["attribute"].(map[string]interface{}); ok {
		AttributeString, _ := json.Marshal(Attribute)
		json.Unmarshal(AttributeString, &o.Attribute)
	}
	
	if Statistics, ok := FacetentryMap["statistics"].(map[string]interface{}); ok {
		StatisticsString, _ := json.Marshal(Statistics)
		json.Unmarshal(StatisticsString, &o.Statistics)
	}
	
	if Other, ok := FacetentryMap["other"].(float64); ok {
		OtherInt := int(Other)
		o.Other = &OtherInt
	}
	
	if Total, ok := FacetentryMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if Missing, ok := FacetentryMap["missing"].(float64); ok {
		MissingInt := int(Missing)
		o.Missing = &MissingInt
	}
	
	if TermCount, ok := FacetentryMap["termCount"].(float64); ok {
		TermCountInt := int(TermCount)
		o.TermCount = &TermCountInt
	}
	
	if TermType, ok := FacetentryMap["termType"].(string); ok {
		o.TermType = &TermType
	}
    
	if Terms, ok := FacetentryMap["terms"].([]interface{}); ok {
		TermsString, _ := json.Marshal(Terms)
		json.Unmarshal(TermsString, &o.Terms)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Facetentry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
