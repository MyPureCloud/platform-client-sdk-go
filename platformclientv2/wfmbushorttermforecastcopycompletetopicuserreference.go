package platformclientv2
import (
	"encoding/json"
)

// Wfmbushorttermforecastcopycompletetopicuserreference
type Wfmbushorttermforecastcopycompletetopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastcopycompletetopicuserreference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
