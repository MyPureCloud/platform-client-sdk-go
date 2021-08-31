package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotdisconnectaction - Settings for a next-action of disconnecting, including the reason code for the disconnect.
type Textbotdisconnectaction struct { 
	// Reason - The reason for the disconnect.
	Reason *string `json:"reason,omitempty"`


	// ReasonExtendedInfo - Extended information related to the reason, if available.
	ReasonExtendedInfo *string `json:"reasonExtendedInfo,omitempty"`


	// FlowLocation - Describes where in the Bot Flow the user was when the disconnect occurred.
	FlowLocation *Textbotflowlocation `json:"flowLocation,omitempty"`

}

func (o *Textbotdisconnectaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotdisconnectaction
	
	return json.Marshal(&struct { 
		Reason *string `json:"reason,omitempty"`
		
		ReasonExtendedInfo *string `json:"reasonExtendedInfo,omitempty"`
		
		FlowLocation *Textbotflowlocation `json:"flowLocation,omitempty"`
		*Alias
	}{ 
		Reason: o.Reason,
		
		ReasonExtendedInfo: o.ReasonExtendedInfo,
		
		FlowLocation: o.FlowLocation,
		Alias:    (*Alias)(o),
	})
}

func (o *Textbotdisconnectaction) UnmarshalJSON(b []byte) error {
	var TextbotdisconnectactionMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotdisconnectactionMap)
	if err != nil {
		return err
	}
	
	if Reason, ok := TextbotdisconnectactionMap["reason"].(string); ok {
		o.Reason = &Reason
	}
	
	if ReasonExtendedInfo, ok := TextbotdisconnectactionMap["reasonExtendedInfo"].(string); ok {
		o.ReasonExtendedInfo = &ReasonExtendedInfo
	}
	
	if FlowLocation, ok := TextbotdisconnectactionMap["flowLocation"].(map[string]interface{}); ok {
		FlowLocationString, _ := json.Marshal(FlowLocation)
		json.Unmarshal(FlowLocationString, &o.FlowLocation)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotdisconnectaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
