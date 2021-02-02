package platformclientv2
import (
	"time"
	"encoding/json"
)

// Edgemetrics
type Edgemetrics struct { 
	// Edge
	Edge *Domainentityref `json:"edge,omitempty"`


	// EventTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventTime *time.Time `json:"eventTime,omitempty"`


	// UpTimeMsec
	UpTimeMsec *int `json:"upTimeMsec,omitempty"`


	// Processors
	Processors *[]Edgemetricsprocessor `json:"processors,omitempty"`


	// Memory
	Memory *[]Edgemetricsmemory `json:"memory,omitempty"`


	// Disks
	Disks *[]Edgemetricsdisk `json:"disks,omitempty"`


	// Subsystems
	Subsystems *[]Edgemetricssubsystem `json:"subsystems,omitempty"`


	// Networks
	Networks *[]Edgemetricsnetwork `json:"networks,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgemetrics) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
