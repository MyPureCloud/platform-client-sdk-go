package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webmessagingeventcobrowse - A Cobrowse event.
type Webmessagingeventcobrowse struct { 
	// VarType - Describes the type of Cobrowse event.
	VarType *string `json:"type,omitempty"`


	// SessionId - The Cobrowse session ID.
	SessionId *string `json:"sessionId,omitempty"`


	// SessionJoinToken - The Cobrowse session join token.
	SessionJoinToken *string `json:"sessionJoinToken,omitempty"`

}

func (o *Webmessagingeventcobrowse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webmessagingeventcobrowse
	
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

func (o *Webmessagingeventcobrowse) UnmarshalJSON(b []byte) error {
	var WebmessagingeventcobrowseMap map[string]interface{}
	err := json.Unmarshal(b, &WebmessagingeventcobrowseMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := WebmessagingeventcobrowseMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if SessionId, ok := WebmessagingeventcobrowseMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
	
	if SessionJoinToken, ok := WebmessagingeventcobrowseMap["sessionJoinToken"].(string); ok {
		o.SessionJoinToken = &SessionJoinToken
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webmessagingeventcobrowse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
