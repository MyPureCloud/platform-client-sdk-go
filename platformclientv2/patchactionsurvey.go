package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchactionsurvey
type Patchactionsurvey struct { 
	// Questions - Questions shown to the user.
	Questions *[]Patchsurveyquestion `json:"questions,omitempty"`

}

func (o *Patchactionsurvey) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchactionsurvey
	
	return json.Marshal(&struct { 
		Questions *[]Patchsurveyquestion `json:"questions,omitempty"`
		*Alias
	}{ 
		Questions: o.Questions,
		Alias:    (*Alias)(o),
	})
}

func (o *Patchactionsurvey) UnmarshalJSON(b []byte) error {
	var PatchactionsurveyMap map[string]interface{}
	err := json.Unmarshal(b, &PatchactionsurveyMap)
	if err != nil {
		return err
	}
	
	if Questions, ok := PatchactionsurveyMap["questions"].([]interface{}); ok {
		QuestionsString, _ := json.Marshal(Questions)
		json.Unmarshal(QuestionsString, &o.Questions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Patchactionsurvey) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
