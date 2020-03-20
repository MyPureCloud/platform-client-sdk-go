package platformclientv2
import (
	"encoding/json"
)

// Shifttradesettings
type Shifttradesettings struct { 
	// Enabled - Whether shift trading is enabled for this management unit
	Enabled *bool `json:"enabled,omitempty"`


	// AutoReview - Whether automatic shift trade review is enabled according to the rules defined in for this management unit
	AutoReview *bool `json:"autoReview,omitempty"`


	// AllowDirectTrades - Whether direct shift trades between agents are allowed
	AllowDirectTrades *bool `json:"allowDirectTrades,omitempty"`


	// MinHoursInFuture - The minimum number of hours in the future shift trades are allowed
	MinHoursInFuture *int32 `json:"minHoursInFuture,omitempty"`


	// UnequalPaid - How to handle shift trades which involve unequal paid times
	UnequalPaid *string `json:"unequalPaid,omitempty"`


	// OneSided - How to handle one-sided shift trades
	OneSided *string `json:"oneSided,omitempty"`


	// WeeklyMinPaidViolations - How to handle shift trades which result in violations of weekly minimum paid time constraint
	WeeklyMinPaidViolations *string `json:"weeklyMinPaidViolations,omitempty"`


	// WeeklyMaxPaidViolations - How to handle shift trades which result in violations of weekly maximum paid time constraint
	WeeklyMaxPaidViolations *string `json:"weeklyMaxPaidViolations,omitempty"`


	// RequiresMatchingQueues - Whether to constrain shift trades to agents with matching queues
	RequiresMatchingQueues *bool `json:"requiresMatchingQueues,omitempty"`


	// RequiresMatchingLanguages - Whether to constrain shift trades to agents with matching languages
	RequiresMatchingLanguages *bool `json:"requiresMatchingLanguages,omitempty"`


	// RequiresMatchingSkills - Whether to constrain shift trades to agents with matching skills
	RequiresMatchingSkills *bool `json:"requiresMatchingSkills,omitempty"`


	// RequiresMatchingPlanningGroups - Whether to constrain shift trades to agents with matching planning groups
	RequiresMatchingPlanningGroups *bool `json:"requiresMatchingPlanningGroups,omitempty"`


	// ActivityCategoryRules - Rules that specify what to do with activity categories that are part of a shift defined in a trade
	ActivityCategoryRules *[]Shifttradeactivityrule `json:"activityCategoryRules,omitempty"`

}

// String returns a JSON representation of the model
func (o *Shifttradesettings) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
