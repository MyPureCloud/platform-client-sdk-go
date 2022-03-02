package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmodulereassignsummary - Learning module reassign summary data
type Learningmodulereassignsummary struct { 
	// TotalReassigned - The total number of users for whom assignment is reassigned
	TotalReassigned *int `json:"totalReassigned,omitempty"`


	// CompletedCount - The total number of users who have the assignment in Completed state
	CompletedCount *int `json:"completedCount,omitempty"`


	// InProgressCount - The total number of users who have the assignment in InProgress state
	InProgressCount *int `json:"inProgressCount,omitempty"`


	// AssignedCount - The total number of users who have the assignment in Assigned state
	AssignedCount *int `json:"assignedCount,omitempty"`


	// NotCompletedCount - The total number of users who have their assignment overdue
	NotCompletedCount *int `json:"notCompletedCount,omitempty"`

}

func (o *Learningmodulereassignsummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningmodulereassignsummary
	
	return json.Marshal(&struct { 
		TotalReassigned *int `json:"totalReassigned,omitempty"`
		
		CompletedCount *int `json:"completedCount,omitempty"`
		
		InProgressCount *int `json:"inProgressCount,omitempty"`
		
		AssignedCount *int `json:"assignedCount,omitempty"`
		
		NotCompletedCount *int `json:"notCompletedCount,omitempty"`
		*Alias
	}{ 
		TotalReassigned: o.TotalReassigned,
		
		CompletedCount: o.CompletedCount,
		
		InProgressCount: o.InProgressCount,
		
		AssignedCount: o.AssignedCount,
		
		NotCompletedCount: o.NotCompletedCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningmodulereassignsummary) UnmarshalJSON(b []byte) error {
	var LearningmodulereassignsummaryMap map[string]interface{}
	err := json.Unmarshal(b, &LearningmodulereassignsummaryMap)
	if err != nil {
		return err
	}
	
	if TotalReassigned, ok := LearningmodulereassignsummaryMap["totalReassigned"].(float64); ok {
		TotalReassignedInt := int(TotalReassigned)
		o.TotalReassigned = &TotalReassignedInt
	}
	
	if CompletedCount, ok := LearningmodulereassignsummaryMap["completedCount"].(float64); ok {
		CompletedCountInt := int(CompletedCount)
		o.CompletedCount = &CompletedCountInt
	}
	
	if InProgressCount, ok := LearningmodulereassignsummaryMap["inProgressCount"].(float64); ok {
		InProgressCountInt := int(InProgressCount)
		o.InProgressCount = &InProgressCountInt
	}
	
	if AssignedCount, ok := LearningmodulereassignsummaryMap["assignedCount"].(float64); ok {
		AssignedCountInt := int(AssignedCount)
		o.AssignedCount = &AssignedCountInt
	}
	
	if NotCompletedCount, ok := LearningmodulereassignsummaryMap["notCompletedCount"].(float64); ok {
		NotCompletedCountInt := int(NotCompletedCount)
		o.NotCompletedCount = &NotCompletedCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningmodulereassignsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
