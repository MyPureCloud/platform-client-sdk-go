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

func (u *Speechtextanalyticssettingsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Speechtextanalyticssettingsrequest

	

	return json.Marshal(&struct { 
		DefaultProgramId *string `json:"defaultProgramId,omitempty"`
		
		ExpectedDialects *[]string `json:"expectedDialects,omitempty"`
		*Alias
	}{ 
		DefaultProgramId: u.DefaultProgramId,
		
		ExpectedDialects: u.ExpectedDialects,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Speechtextanalyticssettingsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
