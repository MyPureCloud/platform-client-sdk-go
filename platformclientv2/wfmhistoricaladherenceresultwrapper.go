package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaladherenceresultwrapper
type Wfmhistoricaladherenceresultwrapper struct { 
	// EntityId - The operation ID of the historical adherence query
	EntityId *string `json:"entityId,omitempty"`


	// Data - The list of historical adherence query results
	Data *[]Historicaladherencequeryresult `json:"data,omitempty"`


	// LookupIdToSecondaryPresenceId - Map of secondary presence lookup ID to corresponding secondary presence ID
	LookupIdToSecondaryPresenceId *map[string]string `json:"lookupIdToSecondaryPresenceId,omitempty"`

}

func (o *Wfmhistoricaladherenceresultwrapper) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmhistoricaladherenceresultwrapper
	
	return json.Marshal(&struct { 
		EntityId *string `json:"entityId,omitempty"`
		
		Data *[]Historicaladherencequeryresult `json:"data,omitempty"`
		
		LookupIdToSecondaryPresenceId *map[string]string `json:"lookupIdToSecondaryPresenceId,omitempty"`
		*Alias
	}{ 
		EntityId: o.EntityId,
		
		Data: o.Data,
		
		LookupIdToSecondaryPresenceId: o.LookupIdToSecondaryPresenceId,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmhistoricaladherenceresultwrapper) UnmarshalJSON(b []byte) error {
	var WfmhistoricaladherenceresultwrapperMap map[string]interface{}
	err := json.Unmarshal(b, &WfmhistoricaladherenceresultwrapperMap)
	if err != nil {
		return err
	}
	
	if EntityId, ok := WfmhistoricaladherenceresultwrapperMap["entityId"].(string); ok {
		o.EntityId = &EntityId
	}
    
	if Data, ok := WfmhistoricaladherenceresultwrapperMap["data"].([]interface{}); ok {
		DataString, _ := json.Marshal(Data)
		json.Unmarshal(DataString, &o.Data)
	}
	
	if LookupIdToSecondaryPresenceId, ok := WfmhistoricaladherenceresultwrapperMap["lookupIdToSecondaryPresenceId"].(map[string]interface{}); ok {
		LookupIdToSecondaryPresenceIdString, _ := json.Marshal(LookupIdToSecondaryPresenceId)
		json.Unmarshal(LookupIdToSecondaryPresenceIdString, &o.LookupIdToSecondaryPresenceId)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherenceresultwrapper) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
