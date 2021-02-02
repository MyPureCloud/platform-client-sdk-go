package platformclientv2
import (
	"encoding/json"
)

// Createwebchatrequest
type Createwebchatrequest struct { 
	// QueueId - The ID of the queue to use for routing the chat conversation.
	QueueId *string `json:"queueId,omitempty"`


	// Provider - The name of the provider that is sourcing the web chat.
	Provider *string `json:"provider,omitempty"`


	// SkillIds - The list of skill ID's to use for routing.
	SkillIds *[]string `json:"skillIds,omitempty"`


	// LanguageId - The ID of the langauge to use for routing.
	LanguageId *string `json:"languageId,omitempty"`


	// Priority - The priority to assign to the conversation for routing.
	Priority *int `json:"priority,omitempty"`


	// Attributes - The list of attributes to associate with the customer participant.
	Attributes *map[string]string `json:"attributes,omitempty"`


	// CustomerName - The name of the customer participating in the web chat.
	CustomerName *string `json:"customerName,omitempty"`

}

// String returns a JSON representation of the model
func (o *Createwebchatrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
