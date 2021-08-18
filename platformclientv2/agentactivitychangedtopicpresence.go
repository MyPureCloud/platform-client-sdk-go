package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentactivitychangedtopicpresence
type Agentactivitychangedtopicpresence struct { 
	// PresenceDefinition
	PresenceDefinition *Agentactivitychangedtopicorganizationpresence `json:"presenceDefinition,omitempty"`


	// PresenceMessage
	PresenceMessage *string `json:"presenceMessage,omitempty"`


	// ModifiedDate
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

}

func (u *Agentactivitychangedtopicpresence) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Agentactivitychangedtopicpresence

	
	ModifiedDate := new(string)
	if u.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(u.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	

	return json.Marshal(&struct { 
		PresenceDefinition *Agentactivitychangedtopicorganizationpresence `json:"presenceDefinition,omitempty"`
		
		PresenceMessage *string `json:"presenceMessage,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		*Alias
	}{ 
		PresenceDefinition: u.PresenceDefinition,
		
		PresenceMessage: u.PresenceMessage,
		
		ModifiedDate: ModifiedDate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Agentactivitychangedtopicpresence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
