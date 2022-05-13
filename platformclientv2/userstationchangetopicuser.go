package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userstationchangetopicuser
type Userstationchangetopicuser struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Userstationchangetopicuser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userstationchangetopicuser
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Userstationchangetopicuser) UnmarshalJSON(b []byte) error {
	var UserstationchangetopicuserMap map[string]interface{}
	err := json.Unmarshal(b, &UserstationchangetopicuserMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UserstationchangetopicuserMap["id"].(string); ok {
		o.Id = &Id
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userstationchangetopicuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
