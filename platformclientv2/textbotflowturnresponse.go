package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotflowturnresponse - Information related to a success bot flow turn request.
type Textbotflowturnresponse struct { 
	// Id - The ID of the bot flow turn. If additional turns are needed, supply this ID as the previous turn in your next turn request.
	Id *string `json:"id,omitempty"`


	// PreviousTurn - The reference to a previous turn, if applicable.
	PreviousTurn *Textbotturnreference `json:"previousTurn,omitempty"`


	// Prompts - The output prompts for this turn.
	Prompts *Textbotoutputprompts `json:"prompts,omitempty"`


	// NextActionType - Indicates the suggested next action. If appropriate, the matching output event object includes additional information.
	NextActionType *string `json:"nextActionType,omitempty"`


	// NextActionDisconnect - The next action directive for this turn if it is a Disconnect type.
	NextActionDisconnect *Textbotdisconnectaction `json:"nextActionDisconnect,omitempty"`


	// NextActionWaitForInput - The next action directive for this turn if it is a WaitForInput type.
	NextActionWaitForInput *Textbotwaitforinputaction `json:"nextActionWaitForInput,omitempty"`


	// NextActionExit - The next action directive for this turn if it is an Exit type.
	NextActionExit *Textbotexitaction `json:"nextActionExit,omitempty"`

}

// String returns a JSON representation of the model
func (o *Textbotflowturnresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
