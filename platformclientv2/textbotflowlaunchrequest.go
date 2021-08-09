package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotflowlaunchrequest - Settings for launching an instance of a bot flow.
type Textbotflowlaunchrequest struct { 
	// Flow - Specifies which Bot Flow to launch.
	Flow *Textbotflow `json:"flow,omitempty"`


	// ExternalSessionId - The ID of the external session that is associated with the bot flow.
	ExternalSessionId *string `json:"externalSessionId,omitempty"`


	// ConversationId - A conversation ID to associate with the bot flow, if available.
	ConversationId *string `json:"conversationId,omitempty"`


	// InputData - Input values to the flow. Valid values are defined by the flow's input JSON schema.
	InputData *Textbotinputoutputdata `json:"inputData,omitempty"`


	// Channel - Channel information relevant to the bot flow.
	Channel *Textbotchannel `json:"channel,omitempty"`

}

// String returns a JSON representation of the model
func (o *Textbotflowlaunchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
