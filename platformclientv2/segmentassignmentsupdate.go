package platformclientv2
import (
	"encoding/json"
)

// Segmentassignmentsupdate
type Segmentassignmentsupdate struct { 
	// SegmentId - ID of segment to be assigned/unnassigned
	SegmentId *string `json:"segmentId,omitempty"`


	// State - Segment assignment state
	State *string `json:"state,omitempty"`

}

// String returns a JSON representation of the model
func (o *Segmentassignmentsupdate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
