package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Campaigninteractions
type Campaigninteractions struct { 
	// Campaign
	Campaign *Domainentityref `json:"campaign,omitempty"`


	// PendingInteractions
	PendingInteractions *[]Campaigninteraction `json:"pendingInteractions,omitempty"`


	// ProceedingInteractions
	ProceedingInteractions *[]Campaigninteraction `json:"proceedingInteractions,omitempty"`


	// PreviewingInteractions
	PreviewingInteractions *[]Campaigninteraction `json:"previewingInteractions,omitempty"`


	// InteractingInteractions
	InteractingInteractions *[]Campaigninteraction `json:"interactingInteractions,omitempty"`


	// ScheduledInteractions
	ScheduledInteractions *[]Campaigninteraction `json:"scheduledInteractions,omitempty"`

}

// String returns a JSON representation of the model
func (o *Campaigninteractions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
