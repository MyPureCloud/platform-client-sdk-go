package platformclientv2
import (
	"encoding/json"
)

// Domainnetworkcommandresponse
type Domainnetworkcommandresponse struct { 
	// CorrelationId
	CorrelationId *string `json:"correlationId,omitempty"`


	// CommandName
	CommandName *string `json:"commandName,omitempty"`


	// Acknowledged
	Acknowledged *bool `json:"acknowledged,omitempty"`


	// ErrorInfo
	ErrorInfo **Errordetails `json:"errorInfo,omitempty"`

}

// String returns a JSON representation of the model
func (o *Domainnetworkcommandresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
