package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userstationchangetopicuserstations
type Userstationchangetopicuserstations struct { 
	// AssociatedStation
	AssociatedStation *Userstationchangetopicuserstation `json:"associatedStation,omitempty"`

}

func (o *Userstationchangetopicuserstations) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userstationchangetopicuserstations
	
	return json.Marshal(&struct { 
		AssociatedStation *Userstationchangetopicuserstation `json:"associatedStation,omitempty"`
		*Alias
	}{ 
		AssociatedStation: o.AssociatedStation,
		Alias:    (*Alias)(o),
	})
}

func (o *Userstationchangetopicuserstations) UnmarshalJSON(b []byte) error {
	var UserstationchangetopicuserstationsMap map[string]interface{}
	err := json.Unmarshal(b, &UserstationchangetopicuserstationsMap)
	if err != nil {
		return err
	}
	
	if AssociatedStation, ok := UserstationchangetopicuserstationsMap["associatedStation"].(map[string]interface{}); ok {
		AssociatedStationString, _ := json.Marshal(AssociatedStation)
		json.Unmarshal(AssociatedStationString, &o.AssociatedStation)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userstationchangetopicuserstations) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
