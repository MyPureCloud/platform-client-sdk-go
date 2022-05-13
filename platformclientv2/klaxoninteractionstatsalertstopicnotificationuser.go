package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Klaxoninteractionstatsalertstopicnotificationuser
type Klaxoninteractionstatsalertstopicnotificationuser struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// DisplayName
	DisplayName *string `json:"displayName,omitempty"`

}

func (o *Klaxoninteractionstatsalertstopicnotificationuser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Klaxoninteractionstatsalertstopicnotificationuser
	
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

func (o *Klaxoninteractionstatsalertstopicnotificationuser) UnmarshalJSON(b []byte) error {
	var KlaxoninteractionstatsalertstopicnotificationuserMap map[string]interface{}
	err := json.Unmarshal(b, &KlaxoninteractionstatsalertstopicnotificationuserMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KlaxoninteractionstatsalertstopicnotificationuserMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if DisplayName, ok := KlaxoninteractionstatsalertstopicnotificationuserMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Klaxoninteractionstatsalertstopicnotificationuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
