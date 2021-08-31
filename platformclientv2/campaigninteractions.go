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

func (o *Campaigninteractions) MarshalJSON() ([]byte, error) {
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
		Campaign: o.Campaign,
		
		PendingInteractions: o.PendingInteractions,
		
		ProceedingInteractions: o.ProceedingInteractions,
		
		PreviewingInteractions: o.PreviewingInteractions,
		
		InteractingInteractions: o.InteractingInteractions,
		
		ScheduledInteractions: o.ScheduledInteractions,
		Alias:    (*Alias)(o),
	})
}

func (o *Campaigninteractions) UnmarshalJSON(b []byte) error {
	var CampaigninteractionsMap map[string]interface{}
	err := json.Unmarshal(b, &CampaigninteractionsMap)
	if err != nil {
		return err
	}
	
	if Campaign, ok := CampaigninteractionsMap["campaign"].(map[string]interface{}); ok {
		CampaignString, _ := json.Marshal(Campaign)
		json.Unmarshal(CampaignString, &o.Campaign)
	}
	
	if PendingInteractions, ok := CampaigninteractionsMap["pendingInteractions"].([]interface{}); ok {
		PendingInteractionsString, _ := json.Marshal(PendingInteractions)
		json.Unmarshal(PendingInteractionsString, &o.PendingInteractions)
	}
	
	if ProceedingInteractions, ok := CampaigninteractionsMap["proceedingInteractions"].([]interface{}); ok {
		ProceedingInteractionsString, _ := json.Marshal(ProceedingInteractions)
		json.Unmarshal(ProceedingInteractionsString, &o.ProceedingInteractions)
	}
	
	if PreviewingInteractions, ok := CampaigninteractionsMap["previewingInteractions"].([]interface{}); ok {
		PreviewingInteractionsString, _ := json.Marshal(PreviewingInteractions)
		json.Unmarshal(PreviewingInteractionsString, &o.PreviewingInteractions)
	}
	
	if InteractingInteractions, ok := CampaigninteractionsMap["interactingInteractions"].([]interface{}); ok {
		InteractingInteractionsString, _ := json.Marshal(InteractingInteractions)
		json.Unmarshal(InteractingInteractionsString, &o.InteractingInteractions)
	}
	
	if ScheduledInteractions, ok := CampaigninteractionsMap["scheduledInteractions"].([]interface{}); ok {
		ScheduledInteractionsString, _ := json.Marshal(ScheduledInteractions)
		json.Unmarshal(ScheduledInteractionsString, &o.ScheduledInteractions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Campaigninteractions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
