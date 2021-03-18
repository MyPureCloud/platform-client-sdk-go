package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Webdeploymentsdeploymenttopicwebmessagingdeploymentchangeeventbody
type Webdeploymentsdeploymenttopicwebmessagingdeploymentchangeeventbody struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Configuration
	Configuration *Webdeploymentsdeploymenttopicwebmessagingconfigchangeeventbody `json:"configuration,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`

}

// String returns a JSON representation of the model
func (o *Webdeploymentsdeploymenttopicwebmessagingdeploymentchangeeventbody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
