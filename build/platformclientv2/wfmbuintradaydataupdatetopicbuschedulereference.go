package platformclientv2
import (
	"encoding/json"
)

// Wfmbuintradaydataupdatetopicbuschedulereference
type Wfmbuintradaydataupdatetopicbuschedulereference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// WeekDate
	WeekDate *Wfmbuintradaydataupdatetopiclocaldate `json:"weekDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicbuschedulereference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
