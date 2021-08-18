package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Createwebchatrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createwebchatrequest

	

	return json.Marshal(&struct { 
		QueueId *string `json:"queueId,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		SkillIds *[]string `json:"skillIds,omitempty"`
		
		LanguageId *string `json:"languageId,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		CustomerName *string `json:"customerName,omitempty"`
		*Alias
	}{ 
		QueueId: u.QueueId,
		
		Provider: u.Provider,
		
		SkillIds: u.SkillIds,
		
		LanguageId: u.LanguageId,
		
		Priority: u.Priority,
		
		Attributes: u.Attributes,
		
		CustomerName: u.CustomerName,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Createwebchatrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
