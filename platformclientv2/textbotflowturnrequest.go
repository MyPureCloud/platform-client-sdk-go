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

func (o *Textbotflowturnrequest) MarshalJSON() ([]byte, error) {
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
		PreviousTurn: o.PreviousTurn,
		
		InputEventType: o.InputEventType,
		
		InputEventUserInput: o.InputEventUserInput,
		
		InputEventError: o.InputEventError,
		Alias:    (*Alias)(o),
	})
}

func (o *Textbotflowturnrequest) UnmarshalJSON(b []byte) error {
	var TextbotflowturnrequestMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotflowturnrequestMap)
	if err != nil {
		return err
	}
	
	if PreviousTurn, ok := TextbotflowturnrequestMap["previousTurn"].(map[string]interface{}); ok {
		PreviousTurnString, _ := json.Marshal(PreviousTurn)
		json.Unmarshal(PreviousTurnString, &o.PreviousTurn)
	}
	
	if InputEventType, ok := TextbotflowturnrequestMap["inputEventType"].(string); ok {
		o.InputEventType = &InputEventType
	}
	
	if InputEventUserInput, ok := TextbotflowturnrequestMap["inputEventUserInput"].(map[string]interface{}); ok {
		InputEventUserInputString, _ := json.Marshal(InputEventUserInput)
		json.Unmarshal(InputEventUserInputString, &o.InputEventUserInput)
	}
	
	if InputEventError, ok := TextbotflowturnrequestMap["inputEventError"].(map[string]interface{}); ok {
		InputEventErrorString, _ := json.Marshal(InputEventError)
		json.Unmarshal(InputEventErrorString, &o.InputEventError)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotflowturnrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
