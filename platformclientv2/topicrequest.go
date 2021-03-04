package platformclientv2
import (
	"encoding/json"
)

// Topicrequest
type Topicrequest struct { 
	// Name - The topic name
	Name *string `json:"name,omitempty"`


	// Description - The topic description
	Description *string `json:"description,omitempty"`


	// Strictness - The topic strictness, default value is 72
	Strictness *string `json:"strictness,omitempty"`


	// ProgramIds - The ids of programs associated to the topic
	ProgramIds *[]string `json:"programIds,omitempty"`


	// Tags - The topic tags
	Tags *[]string `json:"tags,omitempty"`


	// Dialect - The topic dialect
	Dialect *string `json:"dialect,omitempty"`


	// Participants - The topic participants, default value is All
	Participants *string `json:"participants,omitempty"`


	// Phrases - The topic phrases
	Phrases *[]Phrase `json:"phrases,omitempty"`

}

// String returns a JSON representation of the model
func (o *Topicrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
