package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Validateworkplanmessages) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
