package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Nludetectionoutput
type Nludetectionoutput struct { 
	// Intents - The detected intents.
	Intents *[]Detectedintent `json:"intents,omitempty"`


	// DialogActs - The detected dialog acts.
	DialogActs *[]Detecteddialogact `json:"dialogActs,omitempty"`

}

// String returns a JSON representation of the model
func (o *Nludetectionoutput) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
