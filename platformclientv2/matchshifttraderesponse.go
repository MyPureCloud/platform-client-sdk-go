package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Matchshifttraderesponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
