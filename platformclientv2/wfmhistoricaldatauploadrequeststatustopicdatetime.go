package platformclientv2
import (
	"encoding/json"
)

// Wfmhistoricaldatauploadrequeststatustopicdatetime
type Wfmhistoricaldatauploadrequeststatustopicdatetime struct { 
	// IMillis
	IMillis *int32 `json:"iMillis,omitempty"`


	// BeforeNow
	BeforeNow *bool `json:"beforeNow,omitempty"`


	// AfterNow
	AfterNow *bool `json:"afterNow,omitempty"`


	// EqualNow
	EqualNow *bool `json:"equalNow,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricaldatauploadrequeststatustopicdatetime) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
