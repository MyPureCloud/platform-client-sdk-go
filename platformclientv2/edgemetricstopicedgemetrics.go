package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgemetricstopicedgemetrics
type Edgemetricstopicedgemetrics struct { 
	// Edge
	Edge *Edgemetricstopicurireference `json:"edge,omitempty"`


	// EventTime
	EventTime *time.Time `json:"eventTime,omitempty"`


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

func (o *Edgemetricstopicedgemetrics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgemetricstopicedgemetrics
	
	EventTime := new(string)
	if o.EventTime != nil {
		
		*EventTime = timeutil.Strftime(o.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventTime = nil
	}
	
	return json.Marshal(&struct { 
		Edge *Edgemetricstopicurireference `json:"edge,omitempty"`
		
		EventTime *string `json:"eventTime,omitempty"`
		
		UpTimeMsec *int `json:"upTimeMsec,omitempty"`
		
		Processors *[]Edgemetricstopicedgemetricprocessor `json:"processors,omitempty"`
		
		Memory *[]Edgemetricstopicedgemetricmemory `json:"memory,omitempty"`
		
		Disks *[]Edgemetricstopicedgemetricdisk `json:"disks,omitempty"`
		
		Subsystems *[]Edgemetricstopicedgemetricsubsystem `json:"subsystems,omitempty"`
		
		Networks *[]Edgemetricstopicedgemetricnetworks `json:"networks,omitempty"`
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

func (o *Edgemetricstopicedgemetrics) UnmarshalJSON(b []byte) error {
	var EdgemetricstopicedgemetricsMap map[string]interface{}
	err := json.Unmarshal(b, &EdgemetricstopicedgemetricsMap)
	if err != nil {
		return err
	}
	
	if Edge, ok := EdgemetricstopicedgemetricsMap["edge"].(map[string]interface{}); ok {
		EdgeString, _ := json.Marshal(Edge)
		json.Unmarshal(EdgeString, &o.Edge)
	}
	
	if eventTimeString, ok := EdgemetricstopicedgemetricsMap["eventTime"].(string); ok {
		EventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventTimeString)
		o.EventTime = &EventTime
	}
	
	if UpTimeMsec, ok := EdgemetricstopicedgemetricsMap["upTimeMsec"].(float64); ok {
		UpTimeMsecInt := int(UpTimeMsec)
		o.UpTimeMsec = &UpTimeMsecInt
	}
	
	if Processors, ok := EdgemetricstopicedgemetricsMap["processors"].([]interface{}); ok {
		ProcessorsString, _ := json.Marshal(Processors)
		json.Unmarshal(ProcessorsString, &o.Processors)
	}
	
	if Memory, ok := EdgemetricstopicedgemetricsMap["memory"].([]interface{}); ok {
		MemoryString, _ := json.Marshal(Memory)
		json.Unmarshal(MemoryString, &o.Memory)
	}
	
	if Disks, ok := EdgemetricstopicedgemetricsMap["disks"].([]interface{}); ok {
		DisksString, _ := json.Marshal(Disks)
		json.Unmarshal(DisksString, &o.Disks)
	}
	
	if Subsystems, ok := EdgemetricstopicedgemetricsMap["subsystems"].([]interface{}); ok {
		SubsystemsString, _ := json.Marshal(Subsystems)
		json.Unmarshal(SubsystemsString, &o.Subsystems)
	}
	
	if Networks, ok := EdgemetricstopicedgemetricsMap["networks"].([]interface{}); ok {
		NetworksString, _ := json.Marshal(Networks)
		json.Unmarshal(NetworksString, &o.Networks)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgemetricstopicedgemetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
