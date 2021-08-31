package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Facebookappcredentials
type Facebookappcredentials struct { 
	// Id - Genesys Cloud Facebook App Id
	Id *string `json:"id,omitempty"`

}

func (o *Facebookappcredentials) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Facebookappcredentials
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Facebookappcredentials) UnmarshalJSON(b []byte) error {
	var FacebookappcredentialsMap map[string]interface{}
	err := json.Unmarshal(b, &FacebookappcredentialsMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FacebookappcredentialsMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Facebookappcredentials) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
