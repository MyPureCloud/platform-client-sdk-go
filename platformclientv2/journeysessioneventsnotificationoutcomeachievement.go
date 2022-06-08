package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeysessioneventsnotificationoutcomeachievement
type Journeysessioneventsnotificationoutcomeachievement struct { 
	// Outcome
	Outcome *Journeysessioneventsnotificationoutcome `json:"outcome,omitempty"`


	// AchievedDate
	AchievedDate *time.Time `json:"achievedDate,omitempty"`

}

func (o *Journeysessioneventsnotificationoutcomeachievement) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeysessioneventsnotificationoutcomeachievement
	
	AchievedDate := new(string)
	if o.AchievedDate != nil {
		
		*AchievedDate = timeutil.Strftime(o.AchievedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		AchievedDate = nil
	}
	
	return json.Marshal(&struct { 
		Outcome *Journeysessioneventsnotificationoutcome `json:"outcome,omitempty"`
		
		AchievedDate *string `json:"achievedDate,omitempty"`
		*Alias
	}{ 
		Outcome: o.Outcome,
		
		AchievedDate: AchievedDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeysessioneventsnotificationoutcomeachievement) UnmarshalJSON(b []byte) error {
	var JourneysessioneventsnotificationoutcomeachievementMap map[string]interface{}
	err := json.Unmarshal(b, &JourneysessioneventsnotificationoutcomeachievementMap)
	if err != nil {
		return err
	}
	
	if Outcome, ok := JourneysessioneventsnotificationoutcomeachievementMap["outcome"].(map[string]interface{}); ok {
		OutcomeString, _ := json.Marshal(Outcome)
		json.Unmarshal(OutcomeString, &o.Outcome)
	}
	
	if achievedDateString, ok := JourneysessioneventsnotificationoutcomeachievementMap["achievedDate"].(string); ok {
		AchievedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", achievedDateString)
		o.AchievedDate = &AchievedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeysessioneventsnotificationoutcomeachievement) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
