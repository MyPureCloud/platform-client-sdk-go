package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Nluutterance
type Nluutterance struct { 
	// Segments - The list of segments that that constitute this utterance for the given intent.
	Segments *[]Nluutterancesegment `json:"segments,omitempty"`

}

// String returns a JSON representation of the model
func (o *Nluutterance) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
