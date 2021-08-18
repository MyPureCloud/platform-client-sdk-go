package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Topicrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Topicrequest

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Strictness *string `json:"strictness,omitempty"`
		
		ProgramIds *[]string `json:"programIds,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		
		Dialect *string `json:"dialect,omitempty"`
		
		Participants *string `json:"participants,omitempty"`
		
		Phrases *[]Phrase `json:"phrases,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Description: u.Description,
		
		Strictness: u.Strictness,
		
		ProgramIds: u.ProgramIds,
		
		Tags: u.Tags,
		
		Dialect: u.Dialect,
		
		Participants: u.Participants,
		
		Phrases: u.Phrases,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Topicrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
