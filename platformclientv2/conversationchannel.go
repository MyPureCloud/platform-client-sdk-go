package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationchannel
type Conversationchannel struct { 
	// VarType - The type or category of this channel.
	VarType *string `json:"type,omitempty"`


	// Platform - The platform used to deliver media for the conversation for a given channel (e.g. Twitter, Messenger, PureCloud Edge).
	Platform *string `json:"platform,omitempty"`

}

func (o *Conversationchannel) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationchannel
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Platform *string `json:"platform,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Platform: o.Platform,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationchannel) UnmarshalJSON(b []byte) error {
	var ConversationchannelMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationchannelMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ConversationchannelMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Platform, ok := ConversationchannelMap["platform"].(string); ok {
		o.Platform = &Platform
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationchannel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
