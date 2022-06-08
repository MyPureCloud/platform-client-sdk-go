package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeysessioneventsnotificationconnectedqueue
type Journeysessioneventsnotificationconnectedqueue struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Journeysessioneventsnotificationconnectedqueue) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeysessioneventsnotificationconnectedqueue
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeysessioneventsnotificationconnectedqueue) UnmarshalJSON(b []byte) error {
	var JourneysessioneventsnotificationconnectedqueueMap map[string]interface{}
	err := json.Unmarshal(b, &JourneysessioneventsnotificationconnectedqueueMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JourneysessioneventsnotificationconnectedqueueMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if SelfUri, ok := JourneysessioneventsnotificationconnectedqueueMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeysessioneventsnotificationconnectedqueue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
