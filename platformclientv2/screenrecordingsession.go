package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (o *Screenrecordingsession) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Screenrecordingsession
	
	StartTime := new(string)
	if o.StartTime != nil {
		
		*StartTime = timeutil.Strftime(o.StartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		User *User `json:"user,omitempty"`
		
		CommunicationId *string `json:"communicationId,omitempty"`
		
		Conversation *Conversation `json:"conversation,omitempty"`
		
		StartTime *string `json:"startTime,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		User: o.User,
		
		CommunicationId: o.CommunicationId,
		
		Conversation: o.Conversation,
		
		StartTime: StartTime,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Screenrecordingsession) UnmarshalJSON(b []byte) error {
	var ScreenrecordingsessionMap map[string]interface{}
	err := json.Unmarshal(b, &ScreenrecordingsessionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ScreenrecordingsessionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := ScreenrecordingsessionMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if User, ok := ScreenrecordingsessionMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if CommunicationId, ok := ScreenrecordingsessionMap["communicationId"].(string); ok {
		o.CommunicationId = &CommunicationId
	}
	
	if Conversation, ok := ScreenrecordingsessionMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if startTimeString, ok := ScreenrecordingsessionMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	
	if SelfUri, ok := ScreenrecordingsessionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Screenrecordingsession) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
