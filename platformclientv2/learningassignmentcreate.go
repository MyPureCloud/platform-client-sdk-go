package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentcreate
type Learningassignmentcreate struct { 
	// ModuleId - The Learning module Id associated with this assignment
	ModuleId *string `json:"moduleId,omitempty"`


	// UserId - The User for whom the assignment is assigned
	UserId *string `json:"userId,omitempty"`


	// RecommendedCompletionDate - The recommended completion date of assignment. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	RecommendedCompletionDate *time.Time `json:"recommendedCompletionDate,omitempty"`

}

func (o *Learningassignmentcreate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentcreate
	
	RecommendedCompletionDate := new(string)
	if o.RecommendedCompletionDate != nil {
		
		*RecommendedCompletionDate = timeutil.Strftime(o.RecommendedCompletionDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		RecommendedCompletionDate = nil
	}
	
	return json.Marshal(&struct { 
		ModuleId *string `json:"moduleId,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		RecommendedCompletionDate *string `json:"recommendedCompletionDate,omitempty"`
		*Alias
	}{ 
		ModuleId: o.ModuleId,
		
		UserId: o.UserId,
		
		RecommendedCompletionDate: RecommendedCompletionDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningassignmentcreate) UnmarshalJSON(b []byte) error {
	var LearningassignmentcreateMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmentcreateMap)
	if err != nil {
		return err
	}
	
	if ModuleId, ok := LearningassignmentcreateMap["moduleId"].(string); ok {
		o.ModuleId = &ModuleId
	}
    
	if UserId, ok := LearningassignmentcreateMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if recommendedCompletionDateString, ok := LearningassignmentcreateMap["recommendedCompletionDate"].(string); ok {
		RecommendedCompletionDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", recommendedCompletionDateString)
		o.RecommendedCompletionDate = &RecommendedCompletionDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmentcreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
