package platformclientv2
import (
	"encoding/json"
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
	return string(j)
}
