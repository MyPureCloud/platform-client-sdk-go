package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Textbotdisconnectaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
