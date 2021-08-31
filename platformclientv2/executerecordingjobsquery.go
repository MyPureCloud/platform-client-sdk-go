package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Executerecordingjobsquery
type Executerecordingjobsquery struct { 
	// State - The desired state for the job to be set to.
	State *string `json:"state,omitempty"`

}

func (o *Executerecordingjobsquery) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Executerecordingjobsquery
	
	return json.Marshal(&struct { 
		State *string `json:"state,omitempty"`
		*Alias
	}{ 
		State: o.State,
		Alias:    (*Alias)(o),
	})
}

func (o *Executerecordingjobsquery) UnmarshalJSON(b []byte) error {
	var ExecuterecordingjobsqueryMap map[string]interface{}
	err := json.Unmarshal(b, &ExecuterecordingjobsqueryMap)
	if err != nil {
		return err
	}
	
	if State, ok := ExecuterecordingjobsqueryMap["state"].(string); ok {
		o.State = &State
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Executerecordingjobsquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
