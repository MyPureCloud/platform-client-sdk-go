package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkmetrics
type Trunkmetrics struct { 
	// EventTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventTime *time.Time `json:"eventTime,omitempty"`


	// LogicalInterface
	LogicalInterface *Domainentityref `json:"logicalInterface,omitempty"`


	// Trunk
	Trunk *Domainentityref `json:"trunk,omitempty"`


	// Calls
	Calls *Trunkmetricscalls `json:"calls,omitempty"`


	// Qos
	Qos *Trunkmetricsqos `json:"qos,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkmetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
