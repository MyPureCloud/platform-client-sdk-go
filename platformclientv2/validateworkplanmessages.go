package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Validateworkplanmessages
type Validateworkplanmessages struct { 
	// ViolationMessages - Messages for work plan violating some rules such as no shifts in a work plan
	ViolationMessages *[]Workplanconfigurationviolationmessage `json:"violationMessages,omitempty"`


	// ConstraintConflictMessage - This field is not null when there is a set of work plan constraints that conflict thus agent schedules cannot be generated
	ConstraintConflictMessage *Constraintconflictmessage `json:"constraintConflictMessage,omitempty"`

}

func (u *Validateworkplanmessages) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Validateworkplanmessages

	

	return json.Marshal(&struct { 
		ViolationMessages *[]Workplanconfigurationviolationmessage `json:"violationMessages,omitempty"`
		
		ConstraintConflictMessage *Constraintconflictmessage `json:"constraintConflictMessage,omitempty"`
		*Alias
	}{ 
		ViolationMessages: u.ViolationMessages,
		
		ConstraintConflictMessage: u.ConstraintConflictMessage,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Validateworkplanmessages) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
