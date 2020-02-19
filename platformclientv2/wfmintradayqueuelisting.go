package platformclientv2
import (
	"encoding/json"
)

// Wfmintradayqueuelisting - A list of IntradayQueue objects
type Wfmintradayqueuelisting struct { 
	// Entities
	Entities *[]Intradayqueue `json:"entities,omitempty"`


	// NoDataReason
	NoDataReason *string `json:"noDataReason,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmintradayqueuelisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
