package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Activitycodecontainer - Container for a map of ActivityCodeId to ActivityCode
type Activitycodecontainer struct { 
	// ActivityCodes - Map of activity code id to activity code
	ActivityCodes *map[string]Activitycode `json:"activityCodes,omitempty"`


	// Metadata - Version metadata for the associated management unit's list of activity codes
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

func (u *Activitycodecontainer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Activitycodecontainer

	

	return json.Marshal(&struct { 
		ActivityCodes *map[string]Activitycode `json:"activityCodes,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		ActivityCodes: u.ActivityCodes,
		
		Metadata: u.Metadata,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Activitycodecontainer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
