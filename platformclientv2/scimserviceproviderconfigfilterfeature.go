package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimserviceproviderconfigfilterfeature - Defines a \"filter\" request in the SCIM service provider's configuration.
type Scimserviceproviderconfigfilterfeature struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Supported - Indicates whether configuration options are supported.
	Supported *bool `json:"supported,omitempty"`

	// MaxResults - The maximum number of results returned from a filtered query.
	MaxResults *int `json:"maxResults,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Scimserviceproviderconfigfilterfeature) SetField(field string, fieldValue interface{}) {
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

func (o Scimserviceproviderconfigfilterfeature) MarshalJSON() ([]byte, error) {
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
	type Alias Scimserviceproviderconfigfilterfeature
	
	return json.Marshal(&struct { 
		Supported *bool `json:"supported,omitempty"`
		
		MaxResults *int `json:"maxResults,omitempty"`
		Alias
	}{ 
		Supported: o.Supported,
		
		MaxResults: o.MaxResults,
		Alias:    (Alias)(o),
	})
}

func (o *Scimserviceproviderconfigfilterfeature) UnmarshalJSON(b []byte) error {
	var ScimserviceproviderconfigfilterfeatureMap map[string]interface{}
	err := json.Unmarshal(b, &ScimserviceproviderconfigfilterfeatureMap)
	if err != nil {
		return err
	}
	
	if Supported, ok := ScimserviceproviderconfigfilterfeatureMap["supported"].(bool); ok {
		o.Supported = &Supported
	}
    
	if MaxResults, ok := ScimserviceproviderconfigfilterfeatureMap["maxResults"].(float64); ok {
		MaxResultsInt := int(MaxResults)
		o.MaxResults = &MaxResultsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Scimserviceproviderconfigfilterfeature) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
