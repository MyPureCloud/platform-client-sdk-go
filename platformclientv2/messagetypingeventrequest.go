package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagetypingeventrequest
type Messagetypingeventrequest struct { 
	// Typing - Typing event
	Typing *Conversationeventtyping `json:"typing,omitempty"`


	// DateSent - The time when the message typing event was sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateSent *time.Time `json:"dateSent,omitempty"`

}

func (o *Messagetypingeventrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagetypingeventrequest
	
	DateSent := new(string)
	if o.DateSent != nil {
		
		*DateSent = timeutil.Strftime(o.DateSent, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateSent = nil
	}
	
	return json.Marshal(&struct { 
		Typing *Conversationeventtyping `json:"typing,omitempty"`
		
		DateSent *string `json:"dateSent,omitempty"`
		*Alias
	}{ 
		Typing: o.Typing,
		
		DateSent: DateSent,
		Alias:    (*Alias)(o),
	})
}

func (o *Messagetypingeventrequest) UnmarshalJSON(b []byte) error {
	var MessagetypingeventrequestMap map[string]interface{}
	err := json.Unmarshal(b, &MessagetypingeventrequestMap)
	if err != nil {
		return err
	}
	
	if Typing, ok := MessagetypingeventrequestMap["typing"].(map[string]interface{}); ok {
		TypingString, _ := json.Marshal(Typing)
		json.Unmarshal(TypingString, &o.Typing)
	}
	
	if dateSentString, ok := MessagetypingeventrequestMap["dateSent"].(string); ok {
		DateSent, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateSentString)
		o.DateSent = &DateSent
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messagetypingeventrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
