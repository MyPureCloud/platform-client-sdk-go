package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Campaigninteractions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Campaigninteractions

	

	return json.Marshal(&struct { 
		Campaign *Domainentityref `json:"campaign,omitempty"`
		
		PendingInteractions *[]Campaigninteraction `json:"pendingInteractions,omitempty"`
		
		ProceedingInteractions *[]Campaigninteraction `json:"proceedingInteractions,omitempty"`
		
		PreviewingInteractions *[]Campaigninteraction `json:"previewingInteractions,omitempty"`
		
		InteractingInteractions *[]Campaigninteraction `json:"interactingInteractions,omitempty"`
		
		ScheduledInteractions *[]Campaigninteraction `json:"scheduledInteractions,omitempty"`
		*Alias
	}{ 
		Campaign: u.Campaign,
		
		PendingInteractions: u.PendingInteractions,
		
		ProceedingInteractions: u.ProceedingInteractions,
		
		PreviewingInteractions: u.PreviewingInteractions,
		
		InteractingInteractions: u.InteractingInteractions,
		
		ScheduledInteractions: u.ScheduledInteractions,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Campaigninteractions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
