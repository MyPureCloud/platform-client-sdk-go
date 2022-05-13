package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmenttopicuserreference
type Learningassignmenttopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Learningassignmenttopicuserreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmenttopicuserreference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningassignmenttopicuserreference) UnmarshalJSON(b []byte) error {
	var LearningassignmenttopicuserreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmenttopicuserreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LearningassignmenttopicuserreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmenttopicuserreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
