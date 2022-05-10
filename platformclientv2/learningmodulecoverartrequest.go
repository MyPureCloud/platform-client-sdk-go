package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmodulecoverartrequest
type Learningmodulecoverartrequest struct { 
	// Id - The key identifier for the cover art
	Id *string `json:"id,omitempty"`

}

func (o *Learningmodulecoverartrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningmodulecoverartrequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningmodulecoverartrequest) UnmarshalJSON(b []byte) error {
	var LearningmodulecoverartrequestMap map[string]interface{}
	err := json.Unmarshal(b, &LearningmodulecoverartrequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LearningmodulecoverartrequestMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningmodulecoverartrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
