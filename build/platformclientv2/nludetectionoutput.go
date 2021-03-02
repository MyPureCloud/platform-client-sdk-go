package platformclientv2
import (
	"encoding/json"
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
	return string(j)
}
