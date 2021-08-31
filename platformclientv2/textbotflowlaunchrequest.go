package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (o *Textbotflowlaunchrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotflowlaunchrequest
	
	return json.Marshal(&struct { 
		Flow *Textbotflow `json:"flow,omitempty"`
		
		ExternalSessionId *string `json:"externalSessionId,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		InputData *Textbotinputoutputdata `json:"inputData,omitempty"`
		
		Channel *Textbotchannel `json:"channel,omitempty"`
		*Alias
	}{ 
		Flow: o.Flow,
		
		ExternalSessionId: o.ExternalSessionId,
		
		ConversationId: o.ConversationId,
		
		InputData: o.InputData,
		
		Channel: o.Channel,
		Alias:    (*Alias)(o),
	})
}

func (o *Textbotflowlaunchrequest) UnmarshalJSON(b []byte) error {
	var TextbotflowlaunchrequestMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotflowlaunchrequestMap)
	if err != nil {
		return err
	}
	
	if Flow, ok := TextbotflowlaunchrequestMap["flow"].(map[string]interface{}); ok {
		FlowString, _ := json.Marshal(Flow)
		json.Unmarshal(FlowString, &o.Flow)
	}
	
	if ExternalSessionId, ok := TextbotflowlaunchrequestMap["externalSessionId"].(string); ok {
		o.ExternalSessionId = &ExternalSessionId
	}
	
	if ConversationId, ok := TextbotflowlaunchrequestMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
	
	if InputData, ok := TextbotflowlaunchrequestMap["inputData"].(map[string]interface{}); ok {
		InputDataString, _ := json.Marshal(InputData)
		json.Unmarshal(InputDataString, &o.InputData)
	}
	
	if Channel, ok := TextbotflowlaunchrequestMap["channel"].(map[string]interface{}); ok {
		ChannelString, _ := json.Marshal(Channel)
		json.Unmarshal(ChannelString, &o.Channel)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotflowlaunchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
