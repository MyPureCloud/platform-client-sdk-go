package platformclientv2
import (
	"encoding/json"
)

// Wfmmoveagentscompletetopicwfmmoveagentdata
type Wfmmoveagentscompletetopicwfmmoveagentdata struct { 
	// User
	User *Wfmmoveagentscompletetopicuserreference `json:"user,omitempty"`


	// Result
	Result *string `json:"result,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmmoveagentscompletetopicwfmmoveagentdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
