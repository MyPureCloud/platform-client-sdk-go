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

func (u *Executerecordingjobsquery) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Executerecordingjobsquery

	

	return json.Marshal(&struct { 
		State *string `json:"state,omitempty"`
		*Alias
	}{ 
		State: u.State,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Executerecordingjobsquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
