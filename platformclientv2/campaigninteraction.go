package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Campaigninteraction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Campaigninteraction

	
	LastActivePreviewWrapupTime := new(string)
	if u.LastActivePreviewWrapupTime != nil {
		
		*LastActivePreviewWrapupTime = timeutil.Strftime(u.LastActivePreviewWrapupTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		LastActivePreviewWrapupTime = nil
	}
	
	CreationTime := new(string)
	if u.CreationTime != nil {
		
		*CreationTime = timeutil.Strftime(u.CreationTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreationTime = nil
	}
	
	CallPlacedTime := new(string)
	if u.CallPlacedTime != nil {
		
		*CallPlacedTime = timeutil.Strftime(u.CallPlacedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CallPlacedTime = nil
	}
	
	CallRoutedTime := new(string)
	if u.CallRoutedTime != nil {
		
		*CallRoutedTime = timeutil.Strftime(u.CallRoutedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CallRoutedTime = nil
	}
	
	PreviewConnectedTime := new(string)
	if u.PreviewConnectedTime != nil {
		
		*PreviewConnectedTime = timeutil.Strftime(u.PreviewConnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		PreviewConnectedTime = nil
	}
	
	PreviewPopDeliveredTime := new(string)
	if u.PreviewPopDeliveredTime != nil {
		
		*PreviewPopDeliveredTime = timeutil.Strftime(u.PreviewPopDeliveredTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		PreviewPopDeliveredTime = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Campaign *Domainentityref `json:"campaign,omitempty"`
		
		Agent *Domainentityref `json:"agent,omitempty"`
		
		Contact *Domainentityref `json:"contact,omitempty"`
		
		DestinationAddress *string `json:"destinationAddress,omitempty"`
		
		ActivePreviewCall *bool `json:"activePreviewCall,omitempty"`
		
		LastActivePreviewWrapupTime *string `json:"lastActivePreviewWrapupTime,omitempty"`
		
		CreationTime *string `json:"creationTime,omitempty"`
		
		CallPlacedTime *string `json:"callPlacedTime,omitempty"`
		
		CallRoutedTime *string `json:"callRoutedTime,omitempty"`
		
		PreviewConnectedTime *string `json:"previewConnectedTime,omitempty"`
		
		Queue *Domainentityref `json:"queue,omitempty"`
		
		Script *Domainentityref `json:"script,omitempty"`
		
		Disposition *string `json:"disposition,omitempty"`
		
		CallerName *string `json:"callerName,omitempty"`
		
		CallerAddress *string `json:"callerAddress,omitempty"`
		
		PreviewPopDeliveredTime *string `json:"previewPopDeliveredTime,omitempty"`
		
		Conversation *Conversationbasic `json:"conversation,omitempty"`
		
		DialerSystemParticipantId *string `json:"dialerSystemParticipantId,omitempty"`
		
		DialingMode *string `json:"dialingMode,omitempty"`
		
		Skills *[]Domainentityref `json:"skills,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Campaign: u.Campaign,
		
		Agent: u.Agent,
		
		Contact: u.Contact,
		
		DestinationAddress: u.DestinationAddress,
		
		ActivePreviewCall: u.ActivePreviewCall,
		
		LastActivePreviewWrapupTime: LastActivePreviewWrapupTime,
		
		CreationTime: CreationTime,
		
		CallPlacedTime: CallPlacedTime,
		
		CallRoutedTime: CallRoutedTime,
		
		PreviewConnectedTime: PreviewConnectedTime,
		
		Queue: u.Queue,
		
		Script: u.Script,
		
		Disposition: u.Disposition,
		
		CallerName: u.CallerName,
		
		CallerAddress: u.CallerAddress,
		
		PreviewPopDeliveredTime: PreviewPopDeliveredTime,
		
		Conversation: u.Conversation,
		
		DialerSystemParticipantId: u.DialerSystemParticipantId,
		
		DialingMode: u.DialingMode,
		
		Skills: u.Skills,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Campaigninteraction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
