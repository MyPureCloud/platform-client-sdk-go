package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Addconversationresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Addconversationresponse

	

	return json.Marshal(&struct { 
		Conversation *Conversationreference `json:"conversation,omitempty"`
		
		Appointment *Coachingappointmentreference `json:"appointment,omitempty"`
		*Alias
	}{ 
		Conversation: u.Conversation,
		
		Appointment: u.Appointment,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Addconversationresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
