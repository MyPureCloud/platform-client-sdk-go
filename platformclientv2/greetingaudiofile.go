package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Greetingaudiofile
type Greetingaudiofile struct { 
	// DurationMilliseconds
	DurationMilliseconds *int `json:"durationMilliseconds,omitempty"`


	// SizeBytes
	SizeBytes *int `json:"sizeBytes,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Greetingaudiofile) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Greetingaudiofile

	

	return json.Marshal(&struct { 
		DurationMilliseconds *int `json:"durationMilliseconds,omitempty"`
		
		SizeBytes *int `json:"sizeBytes,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		DurationMilliseconds: u.DurationMilliseconds,
		
		SizeBytes: u.SizeBytes,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Greetingaudiofile) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
