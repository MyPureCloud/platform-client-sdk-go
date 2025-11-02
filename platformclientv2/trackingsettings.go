package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Trackingsettings - Settings that control how tracking data is collected and filtered.
type Trackingsettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ExcludedQueryParameters - List of parameters to be excluded from the query string.
	ExcludedQueryParameters *[]string `json:"excludedQueryParameters,omitempty"`

	// ShouldKeepUrlFragment - Whether or not to keep the URL fragment.
	ShouldKeepUrlFragment *bool `json:"shouldKeepUrlFragment,omitempty"`

	// SearchQueryParameters - List of query parameters used for search (e.g. 'query').
	SearchQueryParameters *[]string `json:"searchQueryParameters,omitempty"`

	// IpFilters - IP address filtering configuration for tracking restrictions.
	IpFilters *[]Ipfilter `json:"ipFilters,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Trackingsettings) SetField(field string, fieldValue interface{}) {
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

func (o Trackingsettings) MarshalJSON() ([]byte, error) {
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
	type Alias Trackingsettings
	
	return json.Marshal(&struct { 
		ExcludedQueryParameters *[]string `json:"excludedQueryParameters,omitempty"`
		
		ShouldKeepUrlFragment *bool `json:"shouldKeepUrlFragment,omitempty"`
		
		SearchQueryParameters *[]string `json:"searchQueryParameters,omitempty"`
		
		IpFilters *[]Ipfilter `json:"ipFilters,omitempty"`
		Alias
	}{ 
		ExcludedQueryParameters: o.ExcludedQueryParameters,
		
		ShouldKeepUrlFragment: o.ShouldKeepUrlFragment,
		
		SearchQueryParameters: o.SearchQueryParameters,
		
		IpFilters: o.IpFilters,
		Alias:    (Alias)(o),
	})
}

func (o *Trackingsettings) UnmarshalJSON(b []byte) error {
	var TrackingsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &TrackingsettingsMap)
	if err != nil {
		return err
	}
	
	if ExcludedQueryParameters, ok := TrackingsettingsMap["excludedQueryParameters"].([]interface{}); ok {
		ExcludedQueryParametersString, _ := json.Marshal(ExcludedQueryParameters)
		json.Unmarshal(ExcludedQueryParametersString, &o.ExcludedQueryParameters)
	}
	
	if ShouldKeepUrlFragment, ok := TrackingsettingsMap["shouldKeepUrlFragment"].(bool); ok {
		o.ShouldKeepUrlFragment = &ShouldKeepUrlFragment
	}
    
	if SearchQueryParameters, ok := TrackingsettingsMap["searchQueryParameters"].([]interface{}); ok {
		SearchQueryParametersString, _ := json.Marshal(SearchQueryParameters)
		json.Unmarshal(SearchQueryParametersString, &o.SearchQueryParameters)
	}
	
	if IpFilters, ok := TrackingsettingsMap["ipFilters"].([]interface{}); ok {
		IpFiltersString, _ := json.Marshal(IpFilters)
		json.Unmarshal(IpFiltersString, &o.IpFilters)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trackingsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
