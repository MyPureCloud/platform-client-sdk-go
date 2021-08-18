package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchactionsurvey
type Patchactionsurvey struct { 
	// Questions - Questions shown to the user.
	Questions *[]Patchsurveyquestion `json:"questions,omitempty"`

}

func (u *Patchactionsurvey) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchactionsurvey

	

	return json.Marshal(&struct { 
		Questions *[]Patchsurveyquestion `json:"questions,omitempty"`
		*Alias
	}{ 
		Questions: u.Questions,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Patchactionsurvey) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
