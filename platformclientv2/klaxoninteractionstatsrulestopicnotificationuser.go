package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Klaxoninteractionstatsrulestopicnotificationuser
type Klaxoninteractionstatsrulestopicnotificationuser struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// DisplayName
	DisplayName *string `json:"displayName,omitempty"`

}

func (u *Klaxoninteractionstatsrulestopicnotificationuser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Klaxoninteractionstatsrulestopicnotificationuser

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		DisplayName: u.DisplayName,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Klaxoninteractionstatsrulestopicnotificationuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
