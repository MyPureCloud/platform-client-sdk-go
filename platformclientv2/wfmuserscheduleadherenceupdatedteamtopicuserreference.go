package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmuserscheduleadherenceupdatedteamtopicuserreference
type Wfmuserscheduleadherenceupdatedteamtopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Wfmuserscheduleadherenceupdatedteamtopicuserreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmuserscheduleadherenceupdatedteamtopicuserreference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmuserscheduleadherenceupdatedteamtopicuserreference) UnmarshalJSON(b []byte) error {
	var WfmuserscheduleadherenceupdatedteamtopicuserreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &WfmuserscheduleadherenceupdatedteamtopicuserreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmuserscheduleadherenceupdatedteamtopicuserreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmuserscheduleadherenceupdatedteamtopicuserreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
