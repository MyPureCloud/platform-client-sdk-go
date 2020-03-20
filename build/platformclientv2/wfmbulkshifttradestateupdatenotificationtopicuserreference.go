package platformclientv2
import (
	"encoding/json"
)

// Wfmbulkshifttradestateupdatenotificationtopicuserreference
type Wfmbulkshifttradestateupdatenotificationtopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbulkshifttradestateupdatenotificationtopicuserreference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
