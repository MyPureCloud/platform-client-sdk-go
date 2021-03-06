package platformclientv2
import (
	"encoding/json"
)

// Activitycodecontainer - Container for a map of ActivityCodeId to ActivityCode
type Activitycodecontainer struct { 
	// ActivityCodes - Map of activity code id to activity code
	ActivityCodes *map[string]Activitycode `json:"activityCodes,omitempty"`


	// Metadata - Version metadata for the associated management unit's list of activity codes
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

// String returns a JSON representation of the model
func (o *Activitycodecontainer) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
