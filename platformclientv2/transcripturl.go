package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Transcripturl
type Transcripturl struct { 
	// Url - The pre-signed S3 URL of the transcript
	Url *string `json:"url,omitempty"`

}

func (u *Transcripturl) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Transcripturl

	

	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		*Alias
	}{ 
		Url: u.Url,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Transcripturl) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
