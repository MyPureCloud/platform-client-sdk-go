package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotflowturnrequest - Settings for a turn request to a bot flow.
type Textbotflowturnrequest struct { 
	// PreviousTurn - The reference to a previous turn if appropriate, used to avoid race conditions.
	PreviousTurn *Textbotturnreference `json:"previousTurn,omitempty"`


	// InputEventType - Indicates the type of input event being requested. If appropriate, fill out the matching user input object details on this request.
	InputEventType *string `json:"inputEventType,omitempty"`


	// InputEventUserInput - The data for the input event of this turn if it is a user input event. Only one inputEvent may be set.
	InputEventUserInput *Textbotuserinputevent `json:"inputEventUserInput,omitempty"`


	// InputEventError - The data for the input event of this turn if it is an error event. Only one inputEvent may be set.
	InputEventError *Textboterrorinputevent `json:"inputEventError,omitempty"`

}

func (u *Textbotflowturnrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotflowturnrequest

	

	return json.Marshal(&struct { 
		PreviousTurn *Textbotturnreference `json:"previousTurn,omitempty"`
		
		InputEventType *string `json:"inputEventType,omitempty"`
		
		InputEventUserInput *Textbotuserinputevent `json:"inputEventUserInput,omitempty"`
		
		InputEventError *Textboterrorinputevent `json:"inputEventError,omitempty"`
		*Alias
	}{ 
		PreviousTurn: u.PreviousTurn,
		
		InputEventType: u.InputEventType,
		
		InputEventUserInput: u.InputEventUserInput,
		
		InputEventError: u.InputEventError,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Textbotflowturnrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
