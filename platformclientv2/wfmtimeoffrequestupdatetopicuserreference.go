package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmtimeoffrequestupdatetopicuserreference
type Wfmtimeoffrequestupdatetopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Wfmtimeoffrequestupdatetopicuserreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmtimeoffrequestupdatetopicuserreference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmtimeoffrequestupdatetopicuserreference) UnmarshalJSON(b []byte) error {
	var WfmtimeoffrequestupdatetopicuserreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &WfmtimeoffrequestupdatetopicuserreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmtimeoffrequestupdatetopicuserreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmtimeoffrequestupdatetopicuserreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
