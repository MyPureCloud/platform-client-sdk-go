package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Keyperformanceindicatorassessment
type Keyperformanceindicatorassessment struct { 
	// Kpi - Name of the key performance indicator assessed.
	Kpi *string `json:"kpi,omitempty"`


	// AssessmentResult - The overall result of the assessment for a key performance indicator.
	AssessmentResult *string `json:"assessmentResult,omitempty"`


	// Checks - Set of checks executed as part of an assessment.
	Checks *[]Check `json:"checks,omitempty"`

}

// String returns a JSON representation of the model
func (o *Keyperformanceindicatorassessment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
