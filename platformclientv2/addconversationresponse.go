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

func (o *Addconversationresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Addconversationresponse
	
	return json.Marshal(&struct { 
		Conversation *Conversationreference `json:"conversation,omitempty"`
		
		Appointment *Coachingappointmentreference `json:"appointment,omitempty"`
		*Alias
	}{ 
		Conversation: o.Conversation,
		
		Appointment: o.Appointment,
		Alias:    (*Alias)(o),
	})
}

func (o *Addconversationresponse) UnmarshalJSON(b []byte) error {
	var AddconversationresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AddconversationresponseMap)
	if err != nil {
		return err
	}
	
	if Conversation, ok := AddconversationresponseMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if Appointment, ok := AddconversationresponseMap["appointment"].(map[string]interface{}); ok {
		AppointmentString, _ := json.Marshal(Appointment)
		json.Unmarshal(AppointmentString, &o.Appointment)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Addconversationresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
