package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Shifttradematchreviewresponse
type Shifttradematchreviewresponse struct { 
	// InitiatingUser - Details for the initiatingUser side of the shift trade
	InitiatingUser *Shifttradematchreviewuserresponse `json:"initiatingUser,omitempty"`


	// ReceivingUser - Details for the receivingUser side of the shift trade
	ReceivingUser *Shifttradematchreviewuserresponse `json:"receivingUser,omitempty"`


	// Violations - Constraint violations introduced after being matched that would normally disallow a trade, but which can still be overridden by the shift trade administrator
	Violations *[]Shifttradematchviolation `json:"violations,omitempty"`


	// AdminReviewViolations - Constraint violations associated with this shift trade which require shift trade administrator review
	AdminReviewViolations *[]Shifttradematchviolation `json:"adminReviewViolations,omitempty"`

}

func (u *Shifttradematchreviewresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shifttradematchreviewresponse

	

	return json.Marshal(&struct { 
		InitiatingUser *Shifttradematchreviewuserresponse `json:"initiatingUser,omitempty"`
		
		ReceivingUser *Shifttradematchreviewuserresponse `json:"receivingUser,omitempty"`
		
		Violations *[]Shifttradematchviolation `json:"violations,omitempty"`
		
		AdminReviewViolations *[]Shifttradematchviolation `json:"adminReviewViolations,omitempty"`
		*Alias
	}{ 
		InitiatingUser: u.InitiatingUser,
		
		ReceivingUser: u.ReceivingUser,
		
		Violations: u.Violations,
		
		AdminReviewViolations: u.AdminReviewViolations,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Shifttradematchreviewresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
