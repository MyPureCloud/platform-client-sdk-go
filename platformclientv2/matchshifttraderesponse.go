package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Matchshifttraderesponse
type Matchshifttraderesponse struct { 
	// Trade - The associated shift trade
	Trade *Shifttraderesponse `json:"trade,omitempty"`


	// Violations - Constraint violations which disallow this shift trade
	Violations *[]Shifttradematchviolation `json:"violations,omitempty"`


	// AdminReviewViolations - Constraint violations for this shift trade which require shift trade administrator review
	AdminReviewViolations *[]Shifttradematchviolation `json:"adminReviewViolations,omitempty"`

}

func (u *Matchshifttraderesponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Matchshifttraderesponse

	

	return json.Marshal(&struct { 
		Trade *Shifttraderesponse `json:"trade,omitempty"`
		
		Violations *[]Shifttradematchviolation `json:"violations,omitempty"`
		
		AdminReviewViolations *[]Shifttradematchviolation `json:"adminReviewViolations,omitempty"`
		*Alias
	}{ 
		Trade: u.Trade,
		
		Violations: u.Violations,
		
		AdminReviewViolations: u.AdminReviewViolations,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Matchshifttraderesponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
