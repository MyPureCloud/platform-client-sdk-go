package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (o *Documentfaq) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentfaq
	
	return json.Marshal(&struct { 
		Question *string `json:"question,omitempty"`
		
		Answer *string `json:"answer,omitempty"`
		
		Alternatives *[]string `json:"alternatives,omitempty"`
		*Alias
	}{ 
		Question: o.Question,
		
		Answer: o.Answer,
		
		Alternatives: o.Alternatives,
		Alias:    (*Alias)(o),
	})
}

func (o *Documentfaq) UnmarshalJSON(b []byte) error {
	var DocumentfaqMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentfaqMap)
	if err != nil {
		return err
	}
	
	if Question, ok := DocumentfaqMap["question"].(string); ok {
		o.Question = &Question
	}
    
	if Answer, ok := DocumentfaqMap["answer"].(string); ok {
		o.Answer = &Answer
	}
    
	if Alternatives, ok := DocumentfaqMap["alternatives"].([]interface{}); ok {
		AlternativesString, _ := json.Marshal(Alternatives)
		json.Unmarshal(AlternativesString, &o.Alternatives)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentfaq) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
