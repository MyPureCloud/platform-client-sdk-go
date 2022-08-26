package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Eventcobrowse - A CoBrowse event.
type Eventcobrowse struct { 
	// VarType - Describes the type of CoBrowse event.
	VarType *string `json:"type,omitempty"`


	// SessionId - The CoBrowse session ID.
	SessionId *string `json:"sessionId,omitempty"`


	// SessionJoinToken - The CoBrowse session join token.
	SessionJoinToken *string `json:"sessionJoinToken,omitempty"`

}

func (o *Eventcobrowse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Eventcobrowse
	
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

func (o *Eventcobrowse) UnmarshalJSON(b []byte) error {
	var EventcobrowseMap map[string]interface{}
	err := json.Unmarshal(b, &EventcobrowseMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := EventcobrowseMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if SessionId, ok := EventcobrowseMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if SessionJoinToken, ok := EventcobrowseMap["sessionJoinToken"].(string); ok {
		o.SessionJoinToken = &SessionJoinToken
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Eventcobrowse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
