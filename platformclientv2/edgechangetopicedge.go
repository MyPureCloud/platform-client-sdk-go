package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgechangetopicedge
type Edgechangetopicedge struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// OnlineStatus
	OnlineStatus *string `json:"onlineStatus,omitempty"`

}

func (o *Edgechangetopicedge) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgechangetopicedge
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		OnlineStatus *string `json:"onlineStatus,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		OnlineStatus: o.OnlineStatus,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgechangetopicedge) UnmarshalJSON(b []byte) error {
	var EdgechangetopicedgeMap map[string]interface{}
	err := json.Unmarshal(b, &EdgechangetopicedgeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EdgechangetopicedgeMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if OnlineStatus, ok := EdgechangetopicedgeMap["onlineStatus"].(string); ok {
		o.OnlineStatus = &OnlineStatus
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Edgechangetopicedge) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
