package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Speechtextanalyticssettingsrequest
type Speechtextanalyticssettingsrequest struct { 
	// DefaultProgramId - Setting to choose name for the default program for topic detection
	DefaultProgramId *string `json:"defaultProgramId,omitempty"`


	// ExpectedDialects - Setting to choose expected dialects
	ExpectedDialects *[]string `json:"expectedDialects,omitempty"`

}

func (o *Speechtextanalyticssettingsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Speechtextanalyticssettingsrequest
	
	return json.Marshal(&struct { 
		DefaultProgramId *string `json:"defaultProgramId,omitempty"`
		
		ExpectedDialects *[]string `json:"expectedDialects,omitempty"`
		*Alias
	}{ 
		DefaultProgramId: o.DefaultProgramId,
		
		ExpectedDialects: o.ExpectedDialects,
		Alias:    (*Alias)(o),
	})
}

func (o *Speechtextanalyticssettingsrequest) UnmarshalJSON(b []byte) error {
	var SpeechtextanalyticssettingsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &SpeechtextanalyticssettingsrequestMap)
	if err != nil {
		return err
	}
	
	if DefaultProgramId, ok := SpeechtextanalyticssettingsrequestMap["defaultProgramId"].(string); ok {
		o.DefaultProgramId = &DefaultProgramId
	}
	
	if ExpectedDialects, ok := SpeechtextanalyticssettingsrequestMap["expectedDialects"].([]interface{}); ok {
		ExpectedDialectsString, _ := json.Marshal(ExpectedDialects)
		json.Unmarshal(ExpectedDialectsString, &o.ExpectedDialects)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Speechtextanalyticssettingsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
