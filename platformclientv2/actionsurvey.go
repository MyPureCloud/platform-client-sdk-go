package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Actionsurvey
type Actionsurvey struct { 
	// Questions - Questions shown to the user.
	Questions *[]Journeysurveyquestion `json:"questions,omitempty"`

}

func (u *Actionsurvey) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actionsurvey

	

	return json.Marshal(&struct { 
		Questions *[]Journeysurveyquestion `json:"questions,omitempty"`
		*Alias
	}{ 
		Questions: u.Questions,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Actionsurvey) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
