package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (o *Textbotflowturnresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotflowturnresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		PreviousTurn *Textbotturnreference `json:"previousTurn,omitempty"`
		
		Prompts *Textbotoutputprompts `json:"prompts,omitempty"`
		
		NextActionType *string `json:"nextActionType,omitempty"`
		
		NextActionDisconnect *Textbotdisconnectaction `json:"nextActionDisconnect,omitempty"`
		
		NextActionWaitForInput *Textbotwaitforinputaction `json:"nextActionWaitForInput,omitempty"`
		
		NextActionExit *Textbotexitaction `json:"nextActionExit,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		PreviousTurn: o.PreviousTurn,
		
		Prompts: o.Prompts,
		
		NextActionType: o.NextActionType,
		
		NextActionDisconnect: o.NextActionDisconnect,
		
		NextActionWaitForInput: o.NextActionWaitForInput,
		
		NextActionExit: o.NextActionExit,
		Alias:    (*Alias)(o),
	})
}

func (o *Textbotflowturnresponse) UnmarshalJSON(b []byte) error {
	var TextbotflowturnresponseMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotflowturnresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TextbotflowturnresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if PreviousTurn, ok := TextbotflowturnresponseMap["previousTurn"].(map[string]interface{}); ok {
		PreviousTurnString, _ := json.Marshal(PreviousTurn)
		json.Unmarshal(PreviousTurnString, &o.PreviousTurn)
	}
	
	if Prompts, ok := TextbotflowturnresponseMap["prompts"].(map[string]interface{}); ok {
		PromptsString, _ := json.Marshal(Prompts)
		json.Unmarshal(PromptsString, &o.Prompts)
	}
	
	if NextActionType, ok := TextbotflowturnresponseMap["nextActionType"].(string); ok {
		o.NextActionType = &NextActionType
	}
    
	if NextActionDisconnect, ok := TextbotflowturnresponseMap["nextActionDisconnect"].(map[string]interface{}); ok {
		NextActionDisconnectString, _ := json.Marshal(NextActionDisconnect)
		json.Unmarshal(NextActionDisconnectString, &o.NextActionDisconnect)
	}
	
	if NextActionWaitForInput, ok := TextbotflowturnresponseMap["nextActionWaitForInput"].(map[string]interface{}); ok {
		NextActionWaitForInputString, _ := json.Marshal(NextActionWaitForInput)
		json.Unmarshal(NextActionWaitForInputString, &o.NextActionWaitForInput)
	}
	
	if NextActionExit, ok := TextbotflowturnresponseMap["nextActionExit"].(map[string]interface{}); ok {
		NextActionExitString, _ := json.Marshal(NextActionExit)
		json.Unmarshal(NextActionExitString, &o.NextActionExit)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotflowturnresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
