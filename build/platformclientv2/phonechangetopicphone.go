package platformclientv2
import (
	"encoding/json"
)

// Phonechangetopicphone
type Phonechangetopicphone struct { 
	// UserAgentInfo
	UserAgentInfo *Phonechangetopicuseragentinfo `json:"userAgentInfo,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// Status
	Status *Phonechangetopicphonestatus `json:"status,omitempty"`


	// SecondaryStatus
	SecondaryStatus *Phonechangetopicphonestatus `json:"secondaryStatus,omitempty"`

}

// String returns a JSON representation of the model
func (o *Phonechangetopicphone) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
