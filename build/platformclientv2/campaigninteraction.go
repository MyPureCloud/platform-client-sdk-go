package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Campaigninteraction
type Campaigninteraction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Campaign
	Campaign *Domainentityref `json:"campaign,omitempty"`


	// Agent
	Agent *Domainentityref `json:"agent,omitempty"`


	// Contact
	Contact *Domainentityref `json:"contact,omitempty"`


	// DestinationAddress
	DestinationAddress *string `json:"destinationAddress,omitempty"`


	// ActivePreviewCall - Boolean value if there is an active preview call on the interaction
	ActivePreviewCall *bool `json:"activePreviewCall,omitempty"`


	// LastActivePreviewWrapupTime - The time when the last preview of the interaction was wrapped up. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	LastActivePreviewWrapupTime *time.Time `json:"lastActivePreviewWrapupTime,omitempty"`


	// CreationTime - The time when dialer created the interaction. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreationTime *time.Time `json:"creationTime,omitempty"`


	// CallPlacedTime - The time when the agent or system places the call. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CallPlacedTime *time.Time `json:"callPlacedTime,omitempty"`


	// CallRoutedTime - The time when the agent was connected to the call. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CallRoutedTime *time.Time `json:"callRoutedTime,omitempty"`


	// PreviewConnectedTime - The time when the customer and routing participant are connected. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	PreviewConnectedTime *time.Time `json:"previewConnectedTime,omitempty"`


	// Queue
	Queue *Domainentityref `json:"queue,omitempty"`


	// Script
	Script *Domainentityref `json:"script,omitempty"`


	// Disposition - Describes what happened with call analysis for instance: disposition.classification.callable.person, disposition.classification.callable.noanswer
	Disposition *string `json:"disposition,omitempty"`


	// CallerName
	CallerName *string `json:"callerName,omitempty"`


	// CallerAddress
	CallerAddress *string `json:"callerAddress,omitempty"`


	// PreviewPopDeliveredTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	PreviewPopDeliveredTime *time.Time `json:"previewPopDeliveredTime,omitempty"`


	// Conversation
	Conversation *Conversationbasic `json:"conversation,omitempty"`


	// DialerSystemParticipantId - conversation participant id that is the dialer system participant to monitor the call from dialer perspective
	DialerSystemParticipantId *string `json:"dialerSystemParticipantId,omitempty"`


	// DialingMode
	DialingMode *string `json:"dialingMode,omitempty"`


	// Skills - Any skills that are attached to the call for routing
	Skills *[]Domainentityref `json:"skills,omitempty"`

}

// String returns a JSON representation of the model
func (o *Campaigninteraction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
