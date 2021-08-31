package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmuserscheduleadherenceupdatedteamtopicqueuereference
type Wfmuserscheduleadherenceupdatedteamtopicqueuereference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Wfmuserscheduleadherenceupdatedteamtopicqueuereference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmuserscheduleadherenceupdatedteamtopicqueuereference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmuserscheduleadherenceupdatedteamtopicqueuereference) UnmarshalJSON(b []byte) error {
	var WfmuserscheduleadherenceupdatedteamtopicqueuereferenceMap map[string]interface{}
	err := json.Unmarshal(b, &WfmuserscheduleadherenceupdatedteamtopicqueuereferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmuserscheduleadherenceupdatedteamtopicqueuereferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmuserscheduleadherenceupdatedteamtopicqueuereference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
