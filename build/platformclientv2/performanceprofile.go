package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Performanceprofile
type Performanceprofile struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - A name for this performance profile
	Name *string `json:"name,omitempty"`


	// Description - A description about this performance profile
	Description *string `json:"description,omitempty"`


	// MetricOrders - Order of the associated metrics. The list should contain valid ids for metrics
	MetricOrders *[]string `json:"metricOrders,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Performanceprofile) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
