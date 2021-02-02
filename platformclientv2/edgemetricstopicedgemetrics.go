package platformclientv2
import (
	"encoding/json"
)

// Edgemetricstopicedgemetrics
type Edgemetricstopicedgemetrics struct { 
	// Edge
	Edge *Edgemetricstopicurireference `json:"edge,omitempty"`


	// EventTime
	EventTime *Edgemetricstopicoffsetdatetime `json:"eventTime,omitempty"`


	// UpTimeMsec
	UpTimeMsec *int `json:"upTimeMsec,omitempty"`


	// Processors
	Processors *[]Edgemetricstopicedgemetricprocessor `json:"processors,omitempty"`


	// Memory
	Memory *[]Edgemetricstopicedgemetricmemory `json:"memory,omitempty"`


	// Disks
	Disks *[]Edgemetricstopicedgemetricdisk `json:"disks,omitempty"`


	// Subsystems
	Subsystems *[]Edgemetricstopicedgemetricsubsystem `json:"subsystems,omitempty"`


	// Networks
	Networks *[]Edgemetricstopicedgemetricnetworks `json:"networks,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgemetricstopicedgemetrics) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
