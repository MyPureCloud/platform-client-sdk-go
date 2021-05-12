package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Documentfaq
type Documentfaq struct { 
	// Question - The question for this FAQ
	Question *string `json:"question,omitempty"`


	// Answer - The answer for this FAQ
	Answer *string `json:"answer,omitempty"`


	// Alternatives - List of Alternative questions related to the answer which helps in improving the likelihood of a match to user query
	Alternatives *[]string `json:"alternatives,omitempty"`

}

// String returns a JSON representation of the model
func (o *Documentfaq) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
