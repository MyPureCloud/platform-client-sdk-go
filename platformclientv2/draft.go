package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Draft
type Draft struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - Draft name
	Name *string `json:"name,omitempty"`


	// Miner - Miner to which the draft belongs.
	Miner *Miner `json:"miner,omitempty"`


	// Intents - Draft intent object.
	Intents *[]Draftintents `json:"intents,omitempty"`


	// DateCreated - Date when the draft was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date when the draft was updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Draft) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
