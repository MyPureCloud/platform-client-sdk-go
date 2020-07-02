package platformclientv2
import (
	"encoding/json"
)

// Wfmbushorttermforecastimportcompletetopicuserreference
type Wfmbushorttermforecastimportcompletetopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastimportcompletetopicuserreference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
