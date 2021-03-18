package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
