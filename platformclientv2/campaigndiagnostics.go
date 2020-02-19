package platformclientv2
import (
	"encoding/json"
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
	OutstandingInteractionsCount *int32 `json:"outstandingInteractionsCount,omitempty"`


	// ScheduledInteractionsCount - Current number of scheduled interactions on the campaign
	ScheduledInteractionsCount *int32 `json:"scheduledInteractionsCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Campaigndiagnostics) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
