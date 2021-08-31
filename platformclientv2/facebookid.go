package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Facebookid - User information for a Facebook user interacting with a page or app
type Facebookid struct { 
	// Ids - The set of scopedIds that this person has. Each scopedId is specific to a page or app that the user interacts with.
	Ids *[]Facebookscopedid `json:"ids,omitempty"`


	// DisplayName - The displayName of this person's Facebook account. Roughly translates to user.first_name + ' ' + user.last_name in the Facebook API.
	DisplayName *string `json:"displayName,omitempty"`

}

func (o *Facebookid) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Facebookid
	
	return json.Marshal(&struct { 
		Ids *[]Facebookscopedid `json:"ids,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		*Alias
	}{ 
		Ids: o.Ids,
		
		DisplayName: o.DisplayName,
		Alias:    (*Alias)(o),
	})
}

func (o *Facebookid) UnmarshalJSON(b []byte) error {
	var FacebookidMap map[string]interface{}
	err := json.Unmarshal(b, &FacebookidMap)
	if err != nil {
		return err
	}
	
	if Ids, ok := FacebookidMap["ids"].([]interface{}); ok {
		IdsString, _ := json.Marshal(Ids)
		json.Unmarshal(IdsString, &o.Ids)
	}
	
	if DisplayName, ok := FacebookidMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Facebookid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
