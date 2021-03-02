package platformclientv2
import (
	"encoding/json"
)

// Speechtextanalyticssettingsrequest
type Speechtextanalyticssettingsrequest struct { 
	// DefaultProgramId - Setting to choose name for the default program for topic detection
	DefaultProgramId *string `json:"defaultProgramId,omitempty"`


	// ExpectedDialects - Setting to choose expected dialects
	ExpectedDialects *[]string `json:"expectedDialects,omitempty"`

}

// String returns a JSON representation of the model
func (o *Speechtextanalyticssettingsrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
