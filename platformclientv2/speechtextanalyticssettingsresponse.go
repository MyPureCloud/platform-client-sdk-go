package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Speechtextanalyticssettingsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
