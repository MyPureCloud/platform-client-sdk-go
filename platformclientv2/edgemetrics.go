package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Edgemetrics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgemetrics

	
	EventTime := new(string)
	if u.EventTime != nil {
		
		*EventTime = timeutil.Strftime(u.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventTime = nil
	}
	

	return json.Marshal(&struct { 
		Edge *Domainentityref `json:"edge,omitempty"`
		
		EventTime *string `json:"eventTime,omitempty"`
		
		UpTimeMsec *int `json:"upTimeMsec,omitempty"`
		
		Processors *[]Edgemetricsprocessor `json:"processors,omitempty"`
		
		Memory *[]Edgemetricsmemory `json:"memory,omitempty"`
		
		Disks *[]Edgemetricsdisk `json:"disks,omitempty"`
		
		Subsystems *[]Edgemetricssubsystem `json:"subsystems,omitempty"`
		
		Networks *[]Edgemetricsnetwork `json:"networks,omitempty"`
		*Alias
	}{ 
		Edge: u.Edge,
		
		EventTime: EventTime,
		
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
func (o *Edgemetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
