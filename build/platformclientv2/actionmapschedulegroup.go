package platformclientv2
import (
	"encoding/json"
)

// Actionmapschedulegroup
type Actionmapschedulegroup struct { 
	// Id - The ID of the action maps's associated schedule group.
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Actionmapschedulegroup) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
