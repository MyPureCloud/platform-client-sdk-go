package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Domaincapabilities
type Domaincapabilities struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Enabled - True if this address family on the interface is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Dhcp - True if this address family on the interface is using DHCP.
	Dhcp *bool `json:"dhcp,omitempty"`

	// Metric - The metric being used for the address family on this interface. Lower values will have a higher priority. If autoMetric is true, this value will be the automatically calculated metric. To set this value be sure autoMetric is false. If no value is returned, metric configuration is not supported on this Edge.
	Metric *int `json:"metric,omitempty"`

	// AutoMetric - True if the metric is being calculated automatically for the address family on this interface.
	AutoMetric *bool `json:"autoMetric,omitempty"`

	// SupportsMetric - True if metric configuration is supported.
	SupportsMetric *bool `json:"supportsMetric,omitempty"`

	// PingEnabled - Set to true to enable this address family on this interface to respond to ping requests.
	PingEnabled *bool `json:"pingEnabled,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Domaincapabilities) SetField(field string, fieldValue interface{}) {
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

func (o Domaincapabilities) MarshalJSON() ([]byte, error) {
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
	type Alias Domaincapabilities
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		Dhcp *bool `json:"dhcp,omitempty"`
		
		Metric *int `json:"metric,omitempty"`
		
		AutoMetric *bool `json:"autoMetric,omitempty"`
		
		SupportsMetric *bool `json:"supportsMetric,omitempty"`
		
		PingEnabled *bool `json:"pingEnabled,omitempty"`
		Alias
	}{ 
		Enabled: o.Enabled,
		
		Dhcp: o.Dhcp,
		
		Metric: o.Metric,
		
		AutoMetric: o.AutoMetric,
		
		SupportsMetric: o.SupportsMetric,
		
		PingEnabled: o.PingEnabled,
		Alias:    (Alias)(o),
	})
}

func (o *Domaincapabilities) UnmarshalJSON(b []byte) error {
	var DomaincapabilitiesMap map[string]interface{}
	err := json.Unmarshal(b, &DomaincapabilitiesMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := DomaincapabilitiesMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if Dhcp, ok := DomaincapabilitiesMap["dhcp"].(bool); ok {
		o.Dhcp = &Dhcp
	}
    
	if Metric, ok := DomaincapabilitiesMap["metric"].(float64); ok {
		MetricInt := int(Metric)
		o.Metric = &MetricInt
	}
	
	if AutoMetric, ok := DomaincapabilitiesMap["autoMetric"].(bool); ok {
		o.AutoMetric = &AutoMetric
	}
    
	if SupportsMetric, ok := DomaincapabilitiesMap["supportsMetric"].(bool); ok {
		o.SupportsMetric = &SupportsMetric
	}
    
	if PingEnabled, ok := DomaincapabilitiesMap["pingEnabled"].(bool); ok {
		o.PingEnabled = &PingEnabled
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Domaincapabilities) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
