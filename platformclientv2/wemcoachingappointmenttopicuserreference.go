package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wemcoachingappointmenttopicuserreference
type Wemcoachingappointmenttopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Wemcoachingappointmenttopicuserreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wemcoachingappointmenttopicuserreference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Wemcoachingappointmenttopicuserreference) UnmarshalJSON(b []byte) error {
	var WemcoachingappointmenttopicuserreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &WemcoachingappointmenttopicuserreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WemcoachingappointmenttopicuserreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wemcoachingappointmenttopicuserreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
