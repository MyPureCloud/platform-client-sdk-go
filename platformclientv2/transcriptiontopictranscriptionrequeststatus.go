package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Transcriptiontopictranscriptionrequeststatus
type Transcriptiontopictranscriptionrequeststatus struct { 
	// OffsetMs
	OffsetMs *int `json:"offsetMs,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`

}

// String returns a JSON representation of the model
func (o *Transcriptiontopictranscriptionrequeststatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
