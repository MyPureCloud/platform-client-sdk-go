package platformclientv2
import (
	"encoding/json"
)

// Trunkmetricstopictrunkmetrics
type Trunkmetricstopictrunkmetrics struct { 
	// Calls
	Calls *Trunkmetricstopictrunkmetricscalls `json:"calls,omitempty"`


	// EventTime
	EventTime *Trunkmetricstopicoffsetdatetime `json:"eventTime,omitempty"`


	// Qos
	Qos *Trunkmetricstopictrunkmetricsqos `json:"qos,omitempty"`


	// Trunk
	Trunk *Trunkmetricstopicurireference `json:"trunk,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkmetricstopictrunkmetrics) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
