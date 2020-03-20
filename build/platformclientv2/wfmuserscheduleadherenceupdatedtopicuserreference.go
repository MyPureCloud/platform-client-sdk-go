package platformclientv2
import (
	"encoding/json"
)

// Wfmuserscheduleadherenceupdatedtopicuserreference
type Wfmuserscheduleadherenceupdatedtopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmuserscheduleadherenceupdatedtopicuserreference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
