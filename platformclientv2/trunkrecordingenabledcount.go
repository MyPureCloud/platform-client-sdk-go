package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkrecordingenabledcount
type Trunkrecordingenabledcount struct { 
	// EnabledCount - The amount of trunks that have recording enabled
	EnabledCount *int `json:"enabledCount,omitempty"`


	// DisabledCount - The amount of trunks that do not have recording enabled
	DisabledCount *int `json:"disabledCount,omitempty"`

}

func (u *Trunkrecordingenabledcount) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkrecordingenabledcount

	

	return json.Marshal(&struct { 
		EnabledCount *int `json:"enabledCount,omitempty"`
		
		DisabledCount *int `json:"disabledCount,omitempty"`
		*Alias
	}{ 
		EnabledCount: u.EnabledCount,
		
		DisabledCount: u.DisabledCount,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Trunkrecordingenabledcount) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
