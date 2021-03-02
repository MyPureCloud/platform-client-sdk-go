package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Shifttradematchreviewresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
