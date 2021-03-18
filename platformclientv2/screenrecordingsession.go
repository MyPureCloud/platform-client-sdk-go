package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Screenrecordingsession
type Screenrecordingsession struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// User
	User *User `json:"user,omitempty"`


	// CommunicationId - The id of the communication that is being recorded on the conversation
	CommunicationId *string `json:"communicationId,omitempty"`


	// Conversation
	Conversation *Conversation `json:"conversation,omitempty"`


	// StartTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartTime *time.Time `json:"startTime,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Screenrecordingsession) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
