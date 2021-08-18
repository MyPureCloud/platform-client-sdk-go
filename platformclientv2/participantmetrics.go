package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Participantmetrics
type Participantmetrics struct { }

func (u *Participantmetrics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Participantmetrics

	

	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Participantmetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
