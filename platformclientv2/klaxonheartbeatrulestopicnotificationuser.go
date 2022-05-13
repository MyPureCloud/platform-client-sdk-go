package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Klaxonheartbeatrulestopicnotificationuser
type Klaxonheartbeatrulestopicnotificationuser struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// DisplayName
	DisplayName *string `json:"displayName,omitempty"`

}

func (o *Klaxonheartbeatrulestopicnotificationuser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Klaxonheartbeatrulestopicnotificationuser
	
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

func (o *Klaxonheartbeatrulestopicnotificationuser) UnmarshalJSON(b []byte) error {
	var KlaxonheartbeatrulestopicnotificationuserMap map[string]interface{}
	err := json.Unmarshal(b, &KlaxonheartbeatrulestopicnotificationuserMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KlaxonheartbeatrulestopicnotificationuserMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if DisplayName, ok := KlaxonheartbeatrulestopicnotificationuserMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Klaxonheartbeatrulestopicnotificationuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
