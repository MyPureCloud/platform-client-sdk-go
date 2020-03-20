package platformclientv2
import (
	"encoding/json"
)

// Keywordset
type Keywordset struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Queues
	Queues *[]Queue `json:"queues,omitempty"`


	// Language - Language code, such as 'en-US'
	Language *string `json:"language,omitempty"`


	// Agents
	Agents *[]User `json:"agents,omitempty"`


	// Keywords - The list of keywords to be used for keyword spotting.
	Keywords *[]Keyword `json:"keywords,omitempty"`


	// ParticipantPurposes - The types of participants to use keyword spotting on.
	ParticipantPurposes *[]string `json:"participantPurposes,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Keywordset) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
