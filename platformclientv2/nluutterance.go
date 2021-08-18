package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nluutterance
type Nluutterance struct { 
	// Segments - The list of segments that that constitute this utterance for the given intent.
	Segments *[]Nluutterancesegment `json:"segments,omitempty"`

}

func (u *Nluutterance) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nluutterance

	

	return json.Marshal(&struct { 
		Segments *[]Nluutterancesegment `json:"segments,omitempty"`
		*Alias
	}{ 
		Segments: u.Segments,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Nluutterance) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
