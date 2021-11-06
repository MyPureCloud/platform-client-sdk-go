package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventcobrowse - A CoBrowse event.
type Conversationeventcobrowse struct { 
	// VarType - Describes the type of CoBrowse event.
	VarType *string `json:"type,omitempty"`


	// SessionId - The CoBrowse session ID.
	SessionId *string `json:"sessionId,omitempty"`


	// SessionJoinToken - The CoBrowse session join token.
	SessionJoinToken *string `json:"sessionJoinToken,omitempty"`

}

func (o *Conversationeventcobrowse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationeventcobrowse
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		SessionJoinToken *string `json:"sessionJoinToken,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		SessionId: o.SessionId,
		
		SessionJoinToken: o.SessionJoinToken,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationeventcobrowse) UnmarshalJSON(b []byte) error {
	var ConversationeventcobrowseMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationeventcobrowseMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ConversationeventcobrowseMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if SessionId, ok := ConversationeventcobrowseMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
	
	if SessionJoinToken, ok := ConversationeventcobrowseMap["sessionJoinToken"].(string); ok {
		o.SessionJoinToken = &SessionJoinToken
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventcobrowse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
