package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Domaincapabilities
type Domaincapabilities struct { 
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

func (o *Domaincapabilities) MarshalJSON() ([]byte, error) {
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
		*Alias
	}{ 
		Enabled: o.Enabled,
		
		Dhcp: o.Dhcp,
		
		Metric: o.Metric,
		
		AutoMetric: o.AutoMetric,
		
		SupportsMetric: o.SupportsMetric,
		
		PingEnabled: o.PingEnabled,
		Alias:    (*Alias)(o),
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
