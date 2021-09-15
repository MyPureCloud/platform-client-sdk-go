package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Outcomeachievement
type Outcomeachievement struct { 
	// Outcome - The outcome that was achieved.
	Outcome *Achievedoutcome `json:"outcome,omitempty"`


	// AchievedDate - Timestamp indicating when the outcome was achieved. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	AchievedDate *time.Time `json:"achievedDate,omitempty"`

}

func (o *Outcomeachievement) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outcomeachievement
	
	AchievedDate := new(string)
	if o.AchievedDate != nil {
		
		*AchievedDate = timeutil.Strftime(o.AchievedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		AchievedDate = nil
	}
	
	return json.Marshal(&struct { 
		Outcome *Achievedoutcome `json:"outcome,omitempty"`
		
		AchievedDate *string `json:"achievedDate,omitempty"`
		*Alias
	}{ 
		Outcome: o.Outcome,
		
		AchievedDate: AchievedDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Outcomeachievement) UnmarshalJSON(b []byte) error {
	var OutcomeachievementMap map[string]interface{}
	err := json.Unmarshal(b, &OutcomeachievementMap)
	if err != nil {
		return err
	}
	
	if Outcome, ok := OutcomeachievementMap["outcome"].(map[string]interface{}); ok {
		OutcomeString, _ := json.Marshal(Outcome)
		json.Unmarshal(OutcomeString, &o.Outcome)
	}
	
	if achievedDateString, ok := OutcomeachievementMap["achievedDate"].(string); ok {
		AchievedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", achievedDateString)
		o.AchievedDate = &AchievedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outcomeachievement) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
