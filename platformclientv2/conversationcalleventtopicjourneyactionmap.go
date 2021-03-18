package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcalleventtopicjourneyactionmap
type Conversationcalleventtopicjourneyactionmap struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationcalleventtopicjourneyactionmap) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
