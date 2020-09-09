package platformclientv2
import (
	"encoding/json"
)

// Wfmuserscheduleadherenceupdatedteamtopicuserreference
type Wfmuserscheduleadherenceupdatedteamtopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmuserscheduleadherenceupdatedteamtopicuserreference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
