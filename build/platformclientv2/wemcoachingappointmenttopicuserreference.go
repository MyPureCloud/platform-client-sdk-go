package platformclientv2
import (
	"encoding/json"
)

// Wemcoachingappointmenttopicuserreference
type Wemcoachingappointmenttopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wemcoachingappointmenttopicuserreference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
