package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmodulesummary - Learning module summary data
type Learningmodulesummary struct { 
	// AssignedCount - The total number of assignments assigned for a learning module
	AssignedCount *int `json:"assignedCount,omitempty"`


	// CompletedCount - The number of assignments completed for a learning module
	CompletedCount *int `json:"completedCount,omitempty"`


	// PassedCount - The number of assignments passed for a learning module
	PassedCount *int `json:"passedCount,omitempty"`


	// CompletedSum - The sum of assignment scores for a learning module
	CompletedSum *float32 `json:"completedSum,omitempty"`

}

func (o *Learningmodulesummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningmodulesummary
	
	return json.Marshal(&struct { 
		AssignedCount *int `json:"assignedCount,omitempty"`
		
		CompletedCount *int `json:"completedCount,omitempty"`
		
		PassedCount *int `json:"passedCount,omitempty"`
		
		CompletedSum *float32 `json:"completedSum,omitempty"`
		*Alias
	}{ 
		AssignedCount: o.AssignedCount,
		
		CompletedCount: o.CompletedCount,
		
		PassedCount: o.PassedCount,
		
		CompletedSum: o.CompletedSum,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningmodulesummary) UnmarshalJSON(b []byte) error {
	var LearningmodulesummaryMap map[string]interface{}
	err := json.Unmarshal(b, &LearningmodulesummaryMap)
	if err != nil {
		return err
	}
	
	if AssignedCount, ok := LearningmodulesummaryMap["assignedCount"].(float64); ok {
		AssignedCountInt := int(AssignedCount)
		o.AssignedCount = &AssignedCountInt
	}
	
	if CompletedCount, ok := LearningmodulesummaryMap["completedCount"].(float64); ok {
		CompletedCountInt := int(CompletedCount)
		o.CompletedCount = &CompletedCountInt
	}
	
	if PassedCount, ok := LearningmodulesummaryMap["passedCount"].(float64); ok {
		PassedCountInt := int(PassedCount)
		o.PassedCount = &PassedCountInt
	}
	
	if CompletedSum, ok := LearningmodulesummaryMap["completedSum"].(float64); ok {
		CompletedSumFloat32 := float32(CompletedSum)
		o.CompletedSum = &CompletedSumFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningmodulesummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
