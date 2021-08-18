package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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
	MinHoursInFuture *int `json:"minHoursInFuture,omitempty"`


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

func (u *Shifttradesettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shifttradesettings

	

	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		AutoReview *bool `json:"autoReview,omitempty"`
		
		AllowDirectTrades *bool `json:"allowDirectTrades,omitempty"`
		
		MinHoursInFuture *int `json:"minHoursInFuture,omitempty"`
		
		UnequalPaid *string `json:"unequalPaid,omitempty"`
		
		OneSided *string `json:"oneSided,omitempty"`
		
		WeeklyMinPaidViolations *string `json:"weeklyMinPaidViolations,omitempty"`
		
		WeeklyMaxPaidViolations *string `json:"weeklyMaxPaidViolations,omitempty"`
		
		RequiresMatchingQueues *bool `json:"requiresMatchingQueues,omitempty"`
		
		RequiresMatchingLanguages *bool `json:"requiresMatchingLanguages,omitempty"`
		
		RequiresMatchingSkills *bool `json:"requiresMatchingSkills,omitempty"`
		
		RequiresMatchingPlanningGroups *bool `json:"requiresMatchingPlanningGroups,omitempty"`
		
		ActivityCategoryRules *[]Shifttradeactivityrule `json:"activityCategoryRules,omitempty"`
		*Alias
	}{ 
		Enabled: u.Enabled,
		
		AutoReview: u.AutoReview,
		
		AllowDirectTrades: u.AllowDirectTrades,
		
		MinHoursInFuture: u.MinHoursInFuture,
		
		UnequalPaid: u.UnequalPaid,
		
		OneSided: u.OneSided,
		
		WeeklyMinPaidViolations: u.WeeklyMinPaidViolations,
		
		WeeklyMaxPaidViolations: u.WeeklyMaxPaidViolations,
		
		RequiresMatchingQueues: u.RequiresMatchingQueues,
		
		RequiresMatchingLanguages: u.RequiresMatchingLanguages,
		
		RequiresMatchingSkills: u.RequiresMatchingSkills,
		
		RequiresMatchingPlanningGroups: u.RequiresMatchingPlanningGroups,
		
		ActivityCategoryRules: u.ActivityCategoryRules,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Shifttradesettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
