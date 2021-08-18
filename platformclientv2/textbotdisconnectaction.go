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

func (u *Textbotdisconnectaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotdisconnectaction

	

	return json.Marshal(&struct { 
		Reason *string `json:"reason,omitempty"`
		
		ReasonExtendedInfo *string `json:"reasonExtendedInfo,omitempty"`
		
		FlowLocation *Textbotflowlocation `json:"flowLocation,omitempty"`
		*Alias
	}{ 
		Reason: u.Reason,
		
		ReasonExtendedInfo: u.ReasonExtendedInfo,
		
		FlowLocation: u.FlowLocation,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Textbotdisconnectaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
