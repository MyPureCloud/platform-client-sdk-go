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

func (o *Agentactivitychangedtopicpresence) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Agentactivitychangedtopicpresence
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	return json.Marshal(&struct { 
		PresenceDefinition *Agentactivitychangedtopicorganizationpresence `json:"presenceDefinition,omitempty"`
		
		PresenceMessage *string `json:"presenceMessage,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		*Alias
	}{ 
		PresenceDefinition: o.PresenceDefinition,
		
		PresenceMessage: o.PresenceMessage,
		
		ModifiedDate: ModifiedDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Agentactivitychangedtopicpresence) UnmarshalJSON(b []byte) error {
	var AgentactivitychangedtopicpresenceMap map[string]interface{}
	err := json.Unmarshal(b, &AgentactivitychangedtopicpresenceMap)
	if err != nil {
		return err
	}
	
	if PresenceDefinition, ok := AgentactivitychangedtopicpresenceMap["presenceDefinition"].(map[string]interface{}); ok {
		PresenceDefinitionString, _ := json.Marshal(PresenceDefinition)
		json.Unmarshal(PresenceDefinitionString, &o.PresenceDefinition)
	}
	
	if PresenceMessage, ok := AgentactivitychangedtopicpresenceMap["presenceMessage"].(string); ok {
		o.PresenceMessage = &PresenceMessage
	}
	
	if modifiedDateString, ok := AgentactivitychangedtopicpresenceMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Agentactivitychangedtopicpresence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
