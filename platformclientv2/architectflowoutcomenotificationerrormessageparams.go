package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflowoutcomenotificationerrormessageparams
type Architectflowoutcomenotificationerrormessageparams struct { 
	// AdditionalProperties
	AdditionalProperties *map[string]string `json:"additionalProperties,omitempty"`

}

func (u *Architectflowoutcomenotificationerrormessageparams) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectflowoutcomenotificationerrormessageparams

	

	return json.Marshal(&struct { 
		AdditionalProperties *map[string]string `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Architectflowoutcomenotificationerrormessageparams) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
