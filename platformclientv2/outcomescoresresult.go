package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Outcomescoresresult
type Outcomescoresresult struct { 
	// OutcomeScores - List of scored outcomes in the session.
	OutcomeScores *[]Outcomeeventscore `json:"outcomeScores,omitempty"`


	// ModifiedDate - Timestamp indicating the last time that the event was scored. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

}

func (o *Outcomescoresresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outcomescoresresult
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	return json.Marshal(&struct { 
		OutcomeScores *[]Outcomeeventscore `json:"outcomeScores,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		*Alias
	}{ 
		OutcomeScores: o.OutcomeScores,
		
		ModifiedDate: ModifiedDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Outcomescoresresult) UnmarshalJSON(b []byte) error {
	var OutcomescoresresultMap map[string]interface{}
	err := json.Unmarshal(b, &OutcomescoresresultMap)
	if err != nil {
		return err
	}
	
	if OutcomeScores, ok := OutcomescoresresultMap["outcomeScores"].([]interface{}); ok {
		OutcomeScoresString, _ := json.Marshal(OutcomeScores)
		json.Unmarshal(OutcomeScoresString, &o.OutcomeScores)
	}
	
	if modifiedDateString, ok := OutcomescoresresultMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outcomescoresresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
