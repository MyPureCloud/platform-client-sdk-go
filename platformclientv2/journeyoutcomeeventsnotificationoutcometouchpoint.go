package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyoutcomeeventsnotificationoutcometouchpoint
type Journeyoutcomeeventsnotificationoutcometouchpoint struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Channels
	Channels *[]Journeyoutcomeeventsnotificationoutcometouchpointchannel `json:"channels,omitempty"`


	// CreatedDate
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// ActionMap
	ActionMap *Journeyoutcomeeventsnotificationactionmap `json:"actionMap,omitempty"`

}

func (o *Journeyoutcomeeventsnotificationoutcometouchpoint) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeyoutcomeeventsnotificationoutcometouchpoint
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Channels *[]Journeyoutcomeeventsnotificationoutcometouchpointchannel `json:"channels,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		ActionMap *Journeyoutcomeeventsnotificationactionmap `json:"actionMap,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Channels: o.Channels,
		
		CreatedDate: CreatedDate,
		
		ActionMap: o.ActionMap,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeyoutcomeeventsnotificationoutcometouchpoint) UnmarshalJSON(b []byte) error {
	var JourneyoutcomeeventsnotificationoutcometouchpointMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyoutcomeeventsnotificationoutcometouchpointMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JourneyoutcomeeventsnotificationoutcometouchpointMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Channels, ok := JourneyoutcomeeventsnotificationoutcometouchpointMap["channels"].([]interface{}); ok {
		ChannelsString, _ := json.Marshal(Channels)
		json.Unmarshal(ChannelsString, &o.Channels)
	}
	
	if createdDateString, ok := JourneyoutcomeeventsnotificationoutcometouchpointMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if ActionMap, ok := JourneyoutcomeeventsnotificationoutcometouchpointMap["actionMap"].(map[string]interface{}); ok {
		ActionMapString, _ := json.Marshal(ActionMap)
		json.Unmarshal(ActionMapString, &o.ActionMap)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyoutcomeeventsnotificationoutcometouchpoint) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
