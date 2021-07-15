package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Minerintent
type Minerintent struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - Intent name.
	Name *string `json:"name,omitempty"`


	// Miner - The miner to which the intent belongs.
	Miner *Miner `json:"miner,omitempty"`


	// Utterances - The utterances that are extracted for an Intent.
	Utterances *[]Utterance `json:"utterances,omitempty"`


	// AnalyticVolumePercent - Percentage of conversations belonging to the intent.
	AnalyticVolumePercent *float64 `json:"analyticVolumePercent,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Minerintent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
