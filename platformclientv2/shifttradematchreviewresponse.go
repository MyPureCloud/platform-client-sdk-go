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

func (o *Shifttradematchreviewresponse) MarshalJSON() ([]byte, error) {
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
		InitiatingUser: o.InitiatingUser,
		
		ReceivingUser: o.ReceivingUser,
		
		Violations: o.Violations,
		
		AdminReviewViolations: o.AdminReviewViolations,
		Alias:    (*Alias)(o),
	})
}

func (o *Shifttradematchreviewresponse) UnmarshalJSON(b []byte) error {
	var ShifttradematchreviewresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ShifttradematchreviewresponseMap)
	if err != nil {
		return err
	}
	
	if InitiatingUser, ok := ShifttradematchreviewresponseMap["initiatingUser"].(map[string]interface{}); ok {
		InitiatingUserString, _ := json.Marshal(InitiatingUser)
		json.Unmarshal(InitiatingUserString, &o.InitiatingUser)
	}
	
	if ReceivingUser, ok := ShifttradematchreviewresponseMap["receivingUser"].(map[string]interface{}); ok {
		ReceivingUserString, _ := json.Marshal(ReceivingUser)
		json.Unmarshal(ReceivingUserString, &o.ReceivingUser)
	}
	
	if Violations, ok := ShifttradematchreviewresponseMap["violations"].([]interface{}); ok {
		ViolationsString, _ := json.Marshal(Violations)
		json.Unmarshal(ViolationsString, &o.Violations)
	}
	
	if AdminReviewViolations, ok := ShifttradematchreviewresponseMap["adminReviewViolations"].([]interface{}); ok {
		AdminReviewViolationsString, _ := json.Marshal(AdminReviewViolations)
		json.Unmarshal(AdminReviewViolationsString, &o.AdminReviewViolations)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Shifttradematchreviewresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
