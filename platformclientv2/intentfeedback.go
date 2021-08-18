package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Intentfeedback
type Intentfeedback struct { 
	// Name - The name of the detected intent.
	Name *string `json:"name,omitempty"`


	// Probability - The probability of the detected intent.
	Probability *float64 `json:"probability,omitempty"`


	// Entities - The collection of named entities detected.
	Entities *[]Detectednamedentity `json:"entities,omitempty"`


	// Assessment - The assessment on the detection for feedback text.
	Assessment *string `json:"assessment,omitempty"`

}

func (u *Intentfeedback) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Intentfeedback

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Probability *float64 `json:"probability,omitempty"`
		
		Entities *[]Detectednamedentity `json:"entities,omitempty"`
		
		Assessment *string `json:"assessment,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Probability: u.Probability,
		
		Entities: u.Entities,
		
		Assessment: u.Assessment,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Intentfeedback) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
