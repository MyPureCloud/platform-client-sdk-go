package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Groupgreetingeventgreetingaudiofile
type Groupgreetingeventgreetingaudiofile struct { 
	// DurationMilliseconds
	DurationMilliseconds *int `json:"durationMilliseconds,omitempty"`


	// SizeBytes
	SizeBytes *int `json:"sizeBytes,omitempty"`

}

func (u *Groupgreetingeventgreetingaudiofile) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Groupgreetingeventgreetingaudiofile

	

	return json.Marshal(&struct { 
		DurationMilliseconds *int `json:"durationMilliseconds,omitempty"`
		
		SizeBytes *int `json:"sizeBytes,omitempty"`
		*Alias
	}{ 
		DurationMilliseconds: u.DurationMilliseconds,
		
		SizeBytes: u.SizeBytes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Groupgreetingeventgreetingaudiofile) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
