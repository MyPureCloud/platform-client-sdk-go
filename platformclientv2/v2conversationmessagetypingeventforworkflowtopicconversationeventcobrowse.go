package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforworkflowtopicconversationeventcobrowse
type V2conversationmessagetypingeventforworkflowtopicconversationeventcobrowse struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// SessionId
	SessionId *string `json:"sessionId,omitempty"`


	// SessionJoinToken
	SessionJoinToken *string `json:"sessionJoinToken,omitempty"`

}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationeventcobrowse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforworkflowtopicconversationeventcobrowse
	
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

func (o *V2conversationmessagetypingeventforworkflowtopicconversationeventcobrowse) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforworkflowtopicconversationeventcobrowseMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforworkflowtopicconversationeventcobrowseMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := V2conversationmessagetypingeventforworkflowtopicconversationeventcobrowseMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if SessionId, ok := V2conversationmessagetypingeventforworkflowtopicconversationeventcobrowseMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
	
	if SessionJoinToken, ok := V2conversationmessagetypingeventforworkflowtopicconversationeventcobrowseMap["sessionJoinToken"].(string); ok {
		o.SessionJoinToken = &SessionJoinToken
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforworkflowtopicconversationeventcobrowse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
