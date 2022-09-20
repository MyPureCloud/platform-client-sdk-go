package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentreassign
type Learningassignmentreassign struct { 
	// RecommendedCompletionDate - The recommended completion date of assignment. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	RecommendedCompletionDate *time.Time `json:"recommendedCompletionDate,omitempty"`


	// LengthInMinutes - The length in minutes of assignment
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`

}

func (o *Learningassignmentreassign) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentreassign
	
	RecommendedCompletionDate := new(string)
	if o.RecommendedCompletionDate != nil {
		
		*RecommendedCompletionDate = timeutil.Strftime(o.RecommendedCompletionDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		RecommendedCompletionDate = nil
	}
	
	return json.Marshal(&struct { 
		RecommendedCompletionDate *string `json:"recommendedCompletionDate,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		*Alias
	}{ 
		RecommendedCompletionDate: RecommendedCompletionDate,
		
		LengthInMinutes: o.LengthInMinutes,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningassignmentreassign) UnmarshalJSON(b []byte) error {
	var LearningassignmentreassignMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmentreassignMap)
	if err != nil {
		return err
	}
	
	if recommendedCompletionDateString, ok := LearningassignmentreassignMap["recommendedCompletionDate"].(string); ok {
		RecommendedCompletionDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", recommendedCompletionDateString)
		o.RecommendedCompletionDate = &RecommendedCompletionDate
	}
	
	if LengthInMinutes, ok := LearningassignmentreassignMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmentreassign) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
