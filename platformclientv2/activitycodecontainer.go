package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Activitycodecontainer
type Activitycodecontainer struct { 
	// ActivityCodes - Map of activity code id to activity code
	ActivityCodes *map[string]Activitycode `json:"activityCodes,omitempty"`


	// Metadata - Version metadata for the associated management unit's list of activity codes
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

func (o *Activitycodecontainer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Activitycodecontainer
	
	return json.Marshal(&struct { 
		ActivityCodes *map[string]Activitycode `json:"activityCodes,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		ActivityCodes: o.ActivityCodes,
		
		Metadata: o.Metadata,
		Alias:    (*Alias)(o),
	})
}

func (o *Activitycodecontainer) UnmarshalJSON(b []byte) error {
	var ActivitycodecontainerMap map[string]interface{}
	err := json.Unmarshal(b, &ActivitycodecontainerMap)
	if err != nil {
		return err
	}
	
	if ActivityCodes, ok := ActivitycodecontainerMap["activityCodes"].(map[string]interface{}); ok {
		ActivityCodesString, _ := json.Marshal(ActivityCodes)
		json.Unmarshal(ActivityCodesString, &o.ActivityCodes)
	}
	
	if Metadata, ok := ActivitycodecontainerMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Activitycodecontainer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
