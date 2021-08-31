package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmmoveagentscompletetopicwfmmoveagentdata
type Wfmmoveagentscompletetopicwfmmoveagentdata struct { 
	// User
	User *Wfmmoveagentscompletetopicuserreference `json:"user,omitempty"`


	// Result
	Result *string `json:"result,omitempty"`

}

func (o *Wfmmoveagentscompletetopicwfmmoveagentdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmmoveagentscompletetopicwfmmoveagentdata
	
	return json.Marshal(&struct { 
		User *Wfmmoveagentscompletetopicuserreference `json:"user,omitempty"`
		
		Result *string `json:"result,omitempty"`
		*Alias
	}{ 
		User: o.User,
		
		Result: o.Result,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmmoveagentscompletetopicwfmmoveagentdata) UnmarshalJSON(b []byte) error {
	var WfmmoveagentscompletetopicwfmmoveagentdataMap map[string]interface{}
	err := json.Unmarshal(b, &WfmmoveagentscompletetopicwfmmoveagentdataMap)
	if err != nil {
		return err
	}
	
	if User, ok := WfmmoveagentscompletetopicwfmmoveagentdataMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Result, ok := WfmmoveagentscompletetopicwfmmoveagentdataMap["result"].(string); ok {
		o.Result = &Result
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmmoveagentscompletetopicwfmmoveagentdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
