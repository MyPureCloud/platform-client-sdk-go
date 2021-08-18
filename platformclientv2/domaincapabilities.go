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

func (u *Domaincapabilities) MarshalJSON() ([]byte, error) {
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
		Enabled: u.Enabled,
		
		Dhcp: u.Dhcp,
		
		Metric: u.Metric,
		
		AutoMetric: u.AutoMetric,
		
		SupportsMetric: u.SupportsMetric,
		
		PingEnabled: u.PingEnabled,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Domaincapabilities) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
