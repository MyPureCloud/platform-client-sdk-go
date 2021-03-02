package platformclientv2
import (
	"encoding/json"
)

// Wfmbushorttermforecastupdatecompletetopicuserreference
type Wfmbushorttermforecastupdatecompletetopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastupdatecompletetopicuserreference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
