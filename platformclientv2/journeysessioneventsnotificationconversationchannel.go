package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeysessioneventsnotificationconversationchannel
type Journeysessioneventsnotificationconversationchannel struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// MessageType
	MessageType *string `json:"messageType,omitempty"`


	// Platform
	Platform *string `json:"platform,omitempty"`

}

func (o *Journeysessioneventsnotificationconversationchannel) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeysessioneventsnotificationconversationchannel
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		MessageType *string `json:"messageType,omitempty"`
		
		Platform *string `json:"platform,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		MessageType: o.MessageType,
		
		Platform: o.Platform,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeysessioneventsnotificationconversationchannel) UnmarshalJSON(b []byte) error {
	var JourneysessioneventsnotificationconversationchannelMap map[string]interface{}
	err := json.Unmarshal(b, &JourneysessioneventsnotificationconversationchannelMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := JourneysessioneventsnotificationconversationchannelMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if MessageType, ok := JourneysessioneventsnotificationconversationchannelMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
    
	if Platform, ok := JourneysessioneventsnotificationconversationchannelMap["platform"].(string); ok {
		o.Platform = &Platform
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeysessioneventsnotificationconversationchannel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
