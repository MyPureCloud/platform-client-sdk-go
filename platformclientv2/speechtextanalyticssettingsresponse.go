package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Speechtextanalyticssettingsresponse
type Speechtextanalyticssettingsresponse struct { 
	// DefaultProgram - Setting to choose name for the default program for topic detection
	DefaultProgram *Addressableentityref `json:"defaultProgram,omitempty"`


	// ExpectedDialects - Setting to choose expected dialects
	ExpectedDialects *[]string `json:"expectedDialects,omitempty"`

}

func (o *Speechtextanalyticssettingsresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Speechtextanalyticssettingsresponse
	
	return json.Marshal(&struct { 
		DefaultProgram *Addressableentityref `json:"defaultProgram,omitempty"`
		
		ExpectedDialects *[]string `json:"expectedDialects,omitempty"`
		*Alias
	}{ 
		DefaultProgram: o.DefaultProgram,
		
		ExpectedDialects: o.ExpectedDialects,
		Alias:    (*Alias)(o),
	})
}

func (o *Speechtextanalyticssettingsresponse) UnmarshalJSON(b []byte) error {
	var SpeechtextanalyticssettingsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &SpeechtextanalyticssettingsresponseMap)
	if err != nil {
		return err
	}
	
	if DefaultProgram, ok := SpeechtextanalyticssettingsresponseMap["defaultProgram"].(map[string]interface{}); ok {
		DefaultProgramString, _ := json.Marshal(DefaultProgram)
		json.Unmarshal(DefaultProgramString, &o.DefaultProgram)
	}
	
	if ExpectedDialects, ok := SpeechtextanalyticssettingsresponseMap["expectedDialects"].([]interface{}); ok {
		ExpectedDialectsString, _ := json.Marshal(ExpectedDialects)
		json.Unmarshal(ExpectedDialectsString, &o.ExpectedDialects)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Speechtextanalyticssettingsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
