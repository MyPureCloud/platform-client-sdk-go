package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Actionsurvey
type Actionsurvey struct { 
	// Questions - Questions shown to the user.
	Questions *[]Journeysurveyquestion `json:"questions,omitempty"`

}

func (o *Actionsurvey) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actionsurvey
	
	return json.Marshal(&struct { 
		Questions *[]Journeysurveyquestion `json:"questions,omitempty"`
		*Alias
	}{ 
		Questions: o.Questions,
		Alias:    (*Alias)(o),
	})
}

func (o *Actionsurvey) UnmarshalJSON(b []byte) error {
	var ActionsurveyMap map[string]interface{}
	err := json.Unmarshal(b, &ActionsurveyMap)
	if err != nil {
		return err
	}
	
	if Questions, ok := ActionsurveyMap["questions"].([]interface{}); ok {
		QuestionsString, _ := json.Marshal(Questions)
		json.Unmarshal(QuestionsString, &o.Questions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Actionsurvey) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
