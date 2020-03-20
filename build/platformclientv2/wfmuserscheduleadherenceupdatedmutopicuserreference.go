package platformclientv2
import (
	"encoding/json"
)

// Wfmuserscheduleadherenceupdatedmutopicuserreference
type Wfmuserscheduleadherenceupdatedmutopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmuserscheduleadherenceupdatedmutopicuserreference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
