package platformclientv2
import (
	"encoding/json"
)

// Wfmusernotificationtopicuserreference
type Wfmusernotificationtopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmusernotificationtopicuserreference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
