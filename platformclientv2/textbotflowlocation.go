package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotflowlocation - Describes a flow location.
type Textbotflowlocation struct { 
	// ActionName - The name of the action that was active when the event of interest happened.
	ActionName *string `json:"actionName,omitempty"`


	// ActionNumber - The number of the action that was active when the event of interest happened.
	ActionNumber *int `json:"actionNumber,omitempty"`


	// SequenceName - The name of the state or task which was active when the event of interest happened.
	SequenceName *string `json:"sequenceName,omitempty"`

}

func (u *Textbotflowlocation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotflowlocation

	

	return json.Marshal(&struct { 
		ActionName *string `json:"actionName,omitempty"`
		
		ActionNumber *int `json:"actionNumber,omitempty"`
		
		SequenceName *string `json:"sequenceName,omitempty"`
		*Alias
	}{ 
		ActionName: u.ActionName,
		
		ActionNumber: u.ActionNumber,
		
		SequenceName: u.SequenceName,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Textbotflowlocation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
