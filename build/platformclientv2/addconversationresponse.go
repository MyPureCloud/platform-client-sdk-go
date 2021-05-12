package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Addconversationresponse
type Addconversationresponse struct { 
	// Conversation - The conversation reference
	Conversation *Conversationreference `json:"conversation,omitempty"`


	// Appointment - The appointment reference
	Appointment *Coachingappointmentreference `json:"appointment,omitempty"`

}

// String returns a JSON representation of the model
func (o *Addconversationresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
