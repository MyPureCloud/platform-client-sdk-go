package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentactivitychangedtopicorganizationpresence
type Agentactivitychangedtopicorganizationpresence struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// SystemPresence
	SystemPresence *string `json:"systemPresence,omitempty"`

}

func (o *Agentactivitychangedtopicorganizationpresence) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Agentactivitychangedtopicorganizationpresence
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SystemPresence *string `json:"systemPresence,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		SystemPresence: o.SystemPresence,
		Alias:    (*Alias)(o),
	})
}

func (o *Agentactivitychangedtopicorganizationpresence) UnmarshalJSON(b []byte) error {
	var AgentactivitychangedtopicorganizationpresenceMap map[string]interface{}
	err := json.Unmarshal(b, &AgentactivitychangedtopicorganizationpresenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AgentactivitychangedtopicorganizationpresenceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if SystemPresence, ok := AgentactivitychangedtopicorganizationpresenceMap["systemPresence"].(string); ok {
		o.SystemPresence = &SystemPresence
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Agentactivitychangedtopicorganizationpresence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
