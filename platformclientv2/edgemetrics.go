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

func (o *Edgemetrics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgemetrics
	
	EventTime := new(string)
	if o.EventTime != nil {
		
		*EventTime = timeutil.Strftime(o.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Edge: o.Edge,
		
		EventTime: EventTime,
		
		UpTimeMsec: o.UpTimeMsec,
		
		Processors: o.Processors,
		
		Memory: o.Memory,
		
		Disks: o.Disks,
		
		Subsystems: o.Subsystems,
		
		Networks: o.Networks,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgemetrics) UnmarshalJSON(b []byte) error {
	var EdgemetricsMap map[string]interface{}
	err := json.Unmarshal(b, &EdgemetricsMap)
	if err != nil {
		return err
	}
	
	if Edge, ok := EdgemetricsMap["edge"].(map[string]interface{}); ok {
		EdgeString, _ := json.Marshal(Edge)
		json.Unmarshal(EdgeString, &o.Edge)
	}
	
	if eventTimeString, ok := EdgemetricsMap["eventTime"].(string); ok {
		EventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventTimeString)
		o.EventTime = &EventTime
	}
	
	if UpTimeMsec, ok := EdgemetricsMap["upTimeMsec"].(float64); ok {
		UpTimeMsecInt := int(UpTimeMsec)
		o.UpTimeMsec = &UpTimeMsecInt
	}
	
	if Processors, ok := EdgemetricsMap["processors"].([]interface{}); ok {
		ProcessorsString, _ := json.Marshal(Processors)
		json.Unmarshal(ProcessorsString, &o.Processors)
	}
	
	if Memory, ok := EdgemetricsMap["memory"].([]interface{}); ok {
		MemoryString, _ := json.Marshal(Memory)
		json.Unmarshal(MemoryString, &o.Memory)
	}
	
	if Disks, ok := EdgemetricsMap["disks"].([]interface{}); ok {
		DisksString, _ := json.Marshal(Disks)
		json.Unmarshal(DisksString, &o.Disks)
	}
	
	if Subsystems, ok := EdgemetricsMap["subsystems"].([]interface{}); ok {
		SubsystemsString, _ := json.Marshal(Subsystems)
		json.Unmarshal(SubsystemsString, &o.Subsystems)
	}
	
	if Networks, ok := EdgemetricsMap["networks"].([]interface{}); ok {
		NetworksString, _ := json.Marshal(Networks)
		json.Unmarshal(NetworksString, &o.Networks)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgemetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
