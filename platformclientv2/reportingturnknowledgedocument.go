package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportingturnknowledgedocument
type Reportingturnknowledgedocument struct { 
	// Id - The ID of the knowledge document.
	Id *string `json:"id,omitempty"`


	// Question - The the question that was used to match against the search query.
	Question *string `json:"question,omitempty"`


	// Answer - The corresponding answer to the question.
	Answer *string `json:"answer,omitempty"`


	// Confidence - The confidence score of how well the question matched the search query.
	Confidence *float64 `json:"confidence,omitempty"`

}

func (o *Reportingturnknowledgedocument) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reportingturnknowledgedocument
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Question *string `json:"question,omitempty"`
		
		Answer *string `json:"answer,omitempty"`
		
		Confidence *float64 `json:"confidence,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Question: o.Question,
		
		Answer: o.Answer,
		
		Confidence: o.Confidence,
		Alias:    (*Alias)(o),
	})
}

func (o *Reportingturnknowledgedocument) UnmarshalJSON(b []byte) error {
	var ReportingturnknowledgedocumentMap map[string]interface{}
	err := json.Unmarshal(b, &ReportingturnknowledgedocumentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ReportingturnknowledgedocumentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Question, ok := ReportingturnknowledgedocumentMap["question"].(string); ok {
		o.Question = &Question
	}
    
	if Answer, ok := ReportingturnknowledgedocumentMap["answer"].(string); ok {
		o.Answer = &Answer
	}
    
	if Confidence, ok := ReportingturnknowledgedocumentMap["confidence"].(float64); ok {
		o.Confidence = &Confidence
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Reportingturnknowledgedocument) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
