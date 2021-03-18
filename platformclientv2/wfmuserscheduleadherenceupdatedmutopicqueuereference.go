package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmuserscheduleadherenceupdatedmutopicqueuereference
type Wfmuserscheduleadherenceupdatedmutopicqueuereference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmuserscheduleadherenceupdatedmutopicqueuereference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
