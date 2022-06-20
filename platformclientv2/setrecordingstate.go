package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Setrecordingstate
type Setrecordingstate struct { 
	// RecordingState - The value of the recordingState to set.
	RecordingState *string `json:"recordingState,omitempty"`

}

func (o *Setrecordingstate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Setrecordingstate
	
	return json.Marshal(&struct { 
		RecordingState *string `json:"recordingState,omitempty"`
		*Alias
	}{ 
		RecordingState: o.RecordingState,
		Alias:    (*Alias)(o),
	})
}

func (o *Setrecordingstate) UnmarshalJSON(b []byte) error {
	var SetrecordingstateMap map[string]interface{}
	err := json.Unmarshal(b, &SetrecordingstateMap)
	if err != nil {
		return err
	}
	
	if RecordingState, ok := SetrecordingstateMap["recordingState"].(string); ok {
		o.RecordingState = &RecordingState
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Setrecordingstate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
