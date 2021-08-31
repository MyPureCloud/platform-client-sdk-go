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

func (o *Shifttradesettings) MarshalJSON() ([]byte, error) {
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
		Enabled: o.Enabled,
		
		AutoReview: o.AutoReview,
		
		AllowDirectTrades: o.AllowDirectTrades,
		
		MinHoursInFuture: o.MinHoursInFuture,
		
		UnequalPaid: o.UnequalPaid,
		
		OneSided: o.OneSided,
		
		WeeklyMinPaidViolations: o.WeeklyMinPaidViolations,
		
		WeeklyMaxPaidViolations: o.WeeklyMaxPaidViolations,
		
		RequiresMatchingQueues: o.RequiresMatchingQueues,
		
		RequiresMatchingLanguages: o.RequiresMatchingLanguages,
		
		RequiresMatchingSkills: o.RequiresMatchingSkills,
		
		RequiresMatchingPlanningGroups: o.RequiresMatchingPlanningGroups,
		
		ActivityCategoryRules: o.ActivityCategoryRules,
		Alias:    (*Alias)(o),
	})
}

func (o *Shifttradesettings) UnmarshalJSON(b []byte) error {
	var ShifttradesettingsMap map[string]interface{}
	err := json.Unmarshal(b, &ShifttradesettingsMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := ShifttradesettingsMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
	
	if AutoReview, ok := ShifttradesettingsMap["autoReview"].(bool); ok {
		o.AutoReview = &AutoReview
	}
	
	if AllowDirectTrades, ok := ShifttradesettingsMap["allowDirectTrades"].(bool); ok {
		o.AllowDirectTrades = &AllowDirectTrades
	}
	
	if MinHoursInFuture, ok := ShifttradesettingsMap["minHoursInFuture"].(float64); ok {
		MinHoursInFutureInt := int(MinHoursInFuture)
		o.MinHoursInFuture = &MinHoursInFutureInt
	}
	
	if UnequalPaid, ok := ShifttradesettingsMap["unequalPaid"].(string); ok {
		o.UnequalPaid = &UnequalPaid
	}
	
	if OneSided, ok := ShifttradesettingsMap["oneSided"].(string); ok {
		o.OneSided = &OneSided
	}
	
	if WeeklyMinPaidViolations, ok := ShifttradesettingsMap["weeklyMinPaidViolations"].(string); ok {
		o.WeeklyMinPaidViolations = &WeeklyMinPaidViolations
	}
	
	if WeeklyMaxPaidViolations, ok := ShifttradesettingsMap["weeklyMaxPaidViolations"].(string); ok {
		o.WeeklyMaxPaidViolations = &WeeklyMaxPaidViolations
	}
	
	if RequiresMatchingQueues, ok := ShifttradesettingsMap["requiresMatchingQueues"].(bool); ok {
		o.RequiresMatchingQueues = &RequiresMatchingQueues
	}
	
	if RequiresMatchingLanguages, ok := ShifttradesettingsMap["requiresMatchingLanguages"].(bool); ok {
		o.RequiresMatchingLanguages = &RequiresMatchingLanguages
	}
	
	if RequiresMatchingSkills, ok := ShifttradesettingsMap["requiresMatchingSkills"].(bool); ok {
		o.RequiresMatchingSkills = &RequiresMatchingSkills
	}
	
	if RequiresMatchingPlanningGroups, ok := ShifttradesettingsMap["requiresMatchingPlanningGroups"].(bool); ok {
		o.RequiresMatchingPlanningGroups = &RequiresMatchingPlanningGroups
	}
	
	if ActivityCategoryRules, ok := ShifttradesettingsMap["activityCategoryRules"].([]interface{}); ok {
		ActivityCategoryRulesString, _ := json.Marshal(ActivityCategoryRules)
		json.Unmarshal(ActivityCategoryRulesString, &o.ActivityCategoryRules)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Shifttradesettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
