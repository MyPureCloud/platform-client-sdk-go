package platformclientv2
import (
	"encoding/json"
)

// Wfmuserscheduleadherenceupdatedteamtopicqueuereference
type Wfmuserscheduleadherenceupdatedteamtopicqueuereference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmuserscheduleadherenceupdatedteamtopicqueuereference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
