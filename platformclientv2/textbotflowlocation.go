package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Textbotflowlocation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
