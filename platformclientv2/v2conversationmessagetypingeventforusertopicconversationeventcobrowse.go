package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationeventcobrowse
type V2conversationmessagetypingeventforusertopicconversationeventcobrowse struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// SessionId
	SessionId *string `json:"sessionId,omitempty"`


	// SessionJoinToken
	SessionJoinToken *string `json:"sessionJoinToken,omitempty"`

}

func (o *V2conversationmessagetypingeventforusertopicconversationeventcobrowse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforusertopicconversationeventcobrowse
	
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

func (o *V2conversationmessagetypingeventforusertopicconversationeventcobrowse) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationeventcobrowseMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationeventcobrowseMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := V2conversationmessagetypingeventforusertopicconversationeventcobrowseMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if SessionId, ok := V2conversationmessagetypingeventforusertopicconversationeventcobrowseMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
	
	if SessionJoinToken, ok := V2conversationmessagetypingeventforusertopicconversationeventcobrowseMap["sessionJoinToken"].(string); ok {
		o.SessionJoinToken = &SessionJoinToken
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationeventcobrowse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
