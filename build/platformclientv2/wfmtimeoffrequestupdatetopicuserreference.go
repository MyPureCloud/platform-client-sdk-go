package platformclientv2
import (
	"encoding/json"
)

// Wfmtimeoffrequestupdatetopicuserreference
type Wfmtimeoffrequestupdatetopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmtimeoffrequestupdatetopicuserreference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
