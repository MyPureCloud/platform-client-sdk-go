package platformclientv2
import (
	"encoding/json"
)

// Attendancestatuslisting
type Attendancestatuslisting struct { 
	// Entities
	Entities *[]Attendancestatus `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Attendancestatuslisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
