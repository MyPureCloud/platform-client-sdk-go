package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgemetricstopicedgemetrics
type Edgemetricstopicedgemetrics struct { 
	// Edge
	Edge *Edgemetricstopicurireference `json:"edge,omitempty"`


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

func (u *Edgemetricstopicedgemetrics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgemetricstopicedgemetrics

	

	return json.Marshal(&struct { 
		Edge *Edgemetricstopicurireference `json:"edge,omitempty"`
		
		UpTimeMsec *int `json:"upTimeMsec,omitempty"`
		
		Processors *[]Edgemetricstopicedgemetricprocessor `json:"processors,omitempty"`
		
		Memory *[]Edgemetricstopicedgemetricmemory `json:"memory,omitempty"`
		
		Disks *[]Edgemetricstopicedgemetricdisk `json:"disks,omitempty"`
		
		Subsystems *[]Edgemetricstopicedgemetricsubsystem `json:"subsystems,omitempty"`
		
		Networks *[]Edgemetricstopicedgemetricnetworks `json:"networks,omitempty"`
		*Alias
	}{ 
		Edge: u.Edge,
		
		UpTimeMsec: u.UpTimeMsec,
		
		Processors: u.Processors,
		
		Memory: u.Memory,
		
		Disks: u.Disks,
		
		Subsystems: u.Subsystems,
		
		Networks: u.Networks,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Edgemetricstopicedgemetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
