package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotexitaction - Settings for a next-action of exiting the bot flow. Any output variables are available in the details.
type Textbotexitaction struct { 
	// Reason - The reason for the exit.
	Reason *string `json:"reason,omitempty"`


	// ReasonExtendedInfo - Extended information related to the reason, if available.
	ReasonExtendedInfo *string `json:"reasonExtendedInfo,omitempty"`


	// ActiveIntent - The active intent at the time of the exit.
	ActiveIntent *string `json:"activeIntent,omitempty"`


	// FlowLocation - Describes where in the Bot Flow the user was when the exit occurred.
	FlowLocation *Textbotflowlocation `json:"flowLocation,omitempty"`


	// OutputData - The output data for the bot flow.
	OutputData *Textbotinputoutputdata `json:"outputData,omitempty"`

}

func (u *Textbotexitaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotexitaction

	

	return json.Marshal(&struct { 
		Reason *string `json:"reason,omitempty"`
		
		ReasonExtendedInfo *string `json:"reasonExtendedInfo,omitempty"`
		
		ActiveIntent *string `json:"activeIntent,omitempty"`
		
		FlowLocation *Textbotflowlocation `json:"flowLocation,omitempty"`
		
		OutputData *Textbotinputoutputdata `json:"outputData,omitempty"`
		*Alias
	}{ 
		Reason: u.Reason,
		
		ReasonExtendedInfo: u.ReasonExtendedInfo,
		
		ActiveIntent: u.ActiveIntent,
		
		FlowLocation: u.FlowLocation,
		
		OutputData: u.OutputData,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Textbotexitaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
