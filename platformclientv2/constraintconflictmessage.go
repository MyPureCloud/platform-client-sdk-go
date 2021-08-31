package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Constraintconflictmessage
type Constraintconflictmessage struct { 
	// Message - Message for how to resolve a set of conflicted work plan constraints
	Message *Workplanconstraintconflictmessage `json:"message,omitempty"`


	// ConflictedConstraintMessages - Messages for the set of conflicted work plan constraints. Each element indicates the message of a work plan constraint that is conflicted in the set
	ConflictedConstraintMessages *[]Workplanconstraintmessage `json:"conflictedConstraintMessages,omitempty"`

}

func (o *Constraintconflictmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Constraintconflictmessage
	
	return json.Marshal(&struct { 
		Message *Workplanconstraintconflictmessage `json:"message,omitempty"`
		
		ConflictedConstraintMessages *[]Workplanconstraintmessage `json:"conflictedConstraintMessages,omitempty"`
		*Alias
	}{ 
		Message: o.Message,
		
		ConflictedConstraintMessages: o.ConflictedConstraintMessages,
		Alias:    (*Alias)(o),
	})
}

func (o *Constraintconflictmessage) UnmarshalJSON(b []byte) error {
	var ConstraintconflictmessageMap map[string]interface{}
	err := json.Unmarshal(b, &ConstraintconflictmessageMap)
	if err != nil {
		return err
	}
	
	if Message, ok := ConstraintconflictmessageMap["message"].(map[string]interface{}); ok {
		MessageString, _ := json.Marshal(Message)
		json.Unmarshal(MessageString, &o.Message)
	}
	
	if ConflictedConstraintMessages, ok := ConstraintconflictmessageMap["conflictedConstraintMessages"].([]interface{}); ok {
		ConflictedConstraintMessagesString, _ := json.Marshal(ConflictedConstraintMessages)
		json.Unmarshal(ConflictedConstraintMessagesString, &o.ConflictedConstraintMessages)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Constraintconflictmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
