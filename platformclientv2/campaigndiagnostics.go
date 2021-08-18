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

func (u *Campaigndiagnostics) MarshalJSON() ([]byte, error) {
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
		CallableContacts: u.CallableContacts,
		
		QueueUtilizationDiagnostic: u.QueueUtilizationDiagnostic,
		
		RuleSetDiagnostics: u.RuleSetDiagnostics,
		
		OutstandingInteractionsCount: u.OutstandingInteractionsCount,
		
		ScheduledInteractionsCount: u.ScheduledInteractionsCount,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Campaigndiagnostics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
