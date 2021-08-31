package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Campaigndiagnostics
type Campaigndiagnostics struct { 
	// CallableContacts - Campaign properties that can impact which contacts are callable
	CallableContacts *Callablecontactsdiagnostic `json:"callableContacts,omitempty"`


	// QueueUtilizationDiagnostic - Information regarding the campaign's queue
	QueueUtilizationDiagnostic *Queueutilizationdiagnostic `json:"queueUtilizationDiagnostic,omitempty"`


	// RuleSetDiagnostics - Information regarding the campaign's rule sets
	RuleSetDiagnostics *[]Rulesetdiagnostic `json:"ruleSetDiagnostics,omitempty"`


	// OutstandingInteractionsCount - Current number of outstanding interactions on the campaign
	OutstandingInteractionsCount *int `json:"outstandingInteractionsCount,omitempty"`


	// ScheduledInteractionsCount - Current number of scheduled interactions on the campaign
	ScheduledInteractionsCount *int `json:"scheduledInteractionsCount,omitempty"`

}

func (o *Campaigndiagnostics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Campaigndiagnostics
	
	return json.Marshal(&struct { 
		CallableContacts *Callablecontactsdiagnostic `json:"callableContacts,omitempty"`
		
		QueueUtilizationDiagnostic *Queueutilizationdiagnostic `json:"queueUtilizationDiagnostic,omitempty"`
		
		RuleSetDiagnostics *[]Rulesetdiagnostic `json:"ruleSetDiagnostics,omitempty"`
		
		OutstandingInteractionsCount *int `json:"outstandingInteractionsCount,omitempty"`
		
		ScheduledInteractionsCount *int `json:"scheduledInteractionsCount,omitempty"`
		*Alias
	}{ 
		CallableContacts: o.CallableContacts,
		
		QueueUtilizationDiagnostic: o.QueueUtilizationDiagnostic,
		
		RuleSetDiagnostics: o.RuleSetDiagnostics,
		
		OutstandingInteractionsCount: o.OutstandingInteractionsCount,
		
		ScheduledInteractionsCount: o.ScheduledInteractionsCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Campaigndiagnostics) UnmarshalJSON(b []byte) error {
	var CampaigndiagnosticsMap map[string]interface{}
	err := json.Unmarshal(b, &CampaigndiagnosticsMap)
	if err != nil {
		return err
	}
	
	if CallableContacts, ok := CampaigndiagnosticsMap["callableContacts"].(map[string]interface{}); ok {
		CallableContactsString, _ := json.Marshal(CallableContacts)
		json.Unmarshal(CallableContactsString, &o.CallableContacts)
	}
	
	if QueueUtilizationDiagnostic, ok := CampaigndiagnosticsMap["queueUtilizationDiagnostic"].(map[string]interface{}); ok {
		QueueUtilizationDiagnosticString, _ := json.Marshal(QueueUtilizationDiagnostic)
		json.Unmarshal(QueueUtilizationDiagnosticString, &o.QueueUtilizationDiagnostic)
	}
	
	if RuleSetDiagnostics, ok := CampaigndiagnosticsMap["ruleSetDiagnostics"].([]interface{}); ok {
		RuleSetDiagnosticsString, _ := json.Marshal(RuleSetDiagnostics)
		json.Unmarshal(RuleSetDiagnosticsString, &o.RuleSetDiagnostics)
	}
	
	if OutstandingInteractionsCount, ok := CampaigndiagnosticsMap["outstandingInteractionsCount"].(float64); ok {
		OutstandingInteractionsCountInt := int(OutstandingInteractionsCount)
		o.OutstandingInteractionsCount = &OutstandingInteractionsCountInt
	}
	
	if ScheduledInteractionsCount, ok := CampaigndiagnosticsMap["scheduledInteractionsCount"].(float64); ok {
		ScheduledInteractionsCountInt := int(ScheduledInteractionsCount)
		o.ScheduledInteractionsCount = &ScheduledInteractionsCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Campaigndiagnostics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
