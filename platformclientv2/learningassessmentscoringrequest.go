package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassessmentscoringrequest
type Learningassessmentscoringrequest struct { 
	// AssessmentForm - The assessment form to score against
	AssessmentForm *Assessmentform `json:"assessmentForm,omitempty"`


	// Answers - The answers to score
	Answers *Assessmentscoringset `json:"answers,omitempty"`

}

// String returns a JSON representation of the model
func (o *Learningassessmentscoringrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
