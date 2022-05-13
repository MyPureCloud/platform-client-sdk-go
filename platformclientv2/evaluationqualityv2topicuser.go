package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationqualityv2topicuser
type Evaluationqualityv2topicuser struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// DisplayName
	DisplayName *string `json:"displayName,omitempty"`

}

func (o *Evaluationqualityv2topicuser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationqualityv2topicuser
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		DisplayName: o.DisplayName,
		Alias:    (*Alias)(o),
	})
}

func (o *Evaluationqualityv2topicuser) UnmarshalJSON(b []byte) error {
	var Evaluationqualityv2topicuserMap map[string]interface{}
	err := json.Unmarshal(b, &Evaluationqualityv2topicuserMap)
	if err != nil {
		return err
	}
	
	if Id, ok := Evaluationqualityv2topicuserMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if DisplayName, ok := Evaluationqualityv2topicuserMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationqualityv2topicuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
