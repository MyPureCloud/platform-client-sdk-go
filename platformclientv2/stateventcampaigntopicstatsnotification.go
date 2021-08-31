package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventcampaigntopicstatsnotification
type Stateventcampaigntopicstatsnotification struct { 
	// Group
	Group *map[string]string `json:"group,omitempty"`


	// Data
	Data *[]Stateventcampaigntopicdatum `json:"data,omitempty"`

}

func (o *Stateventcampaigntopicstatsnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Stateventcampaigntopicstatsnotification
	
	return json.Marshal(&struct { 
		Group *map[string]string `json:"group,omitempty"`
		
		Data *[]Stateventcampaigntopicdatum `json:"data,omitempty"`
		*Alias
	}{ 
		Group: o.Group,
		
		Data: o.Data,
		Alias:    (*Alias)(o),
	})
}

func (o *Stateventcampaigntopicstatsnotification) UnmarshalJSON(b []byte) error {
	var StateventcampaigntopicstatsnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &StateventcampaigntopicstatsnotificationMap)
	if err != nil {
		return err
	}
	
	if Group, ok := StateventcampaigntopicstatsnotificationMap["group"].(map[string]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	
	if Data, ok := StateventcampaigntopicstatsnotificationMap["data"].([]interface{}); ok {
		DataString, _ := json.Marshal(Data)
		json.Unmarshal(DataString, &o.Data)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Stateventcampaigntopicstatsnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
