package platformclientv2
import (
	"encoding/json"
)

// Wfmforecastmodificationattributes
type Wfmforecastmodificationattributes struct { 
	// Queues - The queues to which to apply a modification
	Queues *[]Queuereference `json:"queues,omitempty"`


	// MediaTypes - The media types to which to apply a modification
	MediaTypes *[]string `json:"mediaTypes,omitempty"`


	// Languages - The languages to which to apply a modification
	Languages *[]Languagereference `json:"languages,omitempty"`


	// SkillSets - The skill sets to which to apply a modification
	SkillSets *[][]Routingskillreference `json:"skillSets,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmforecastmodificationattributes) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
