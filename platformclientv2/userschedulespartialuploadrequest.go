package platformclientv2
import (
	"encoding/json"
)

// Userschedulespartialuploadrequest - Request to upload partial set of user schedules
type Userschedulespartialuploadrequest struct { 
	// UserSchedules - User schedules that are part of partial request
	UserSchedules *map[string]Userschedule `json:"userSchedules,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userschedulespartialuploadrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
