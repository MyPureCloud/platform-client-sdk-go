package platformclientv2
import (
	"encoding/json"
)

// Wfmuserscheduleadherenceupdatedmutopicqueuereference
type Wfmuserscheduleadherenceupdatedmutopicqueuereference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmuserscheduleadherenceupdatedmutopicqueuereference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
