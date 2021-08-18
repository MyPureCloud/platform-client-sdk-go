package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Programrequest
type Programrequest struct { 
	// Name - The program name
	Name *string `json:"name,omitempty"`


	// Description - The program description
	Description *string `json:"description,omitempty"`


	// TopicIds - The ids of topics associated to the program
	TopicIds *[]string `json:"topicIds,omitempty"`


	// Tags - The program tags
	Tags *[]string `json:"tags,omitempty"`

}

func (u *Programrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Programrequest

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		TopicIds *[]string `json:"topicIds,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Description: u.Description,
		
		TopicIds: u.TopicIds,
		
		Tags: u.Tags,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Programrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
