package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Shifttradesettings
type Shifttradesettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

	// ExternalActivityRules - Rules that specify what to do with external activities that are part of a shift
	ExternalActivityRules *[]Shifttradeexternalactivityrule `json:"externalActivityRules,omitempty"`

	// ExternalActivityRuleIgnoresActivityCategoryRule - Whether external activity rules ignore activity category rules
	ExternalActivityRuleIgnoresActivityCategoryRule *bool `json:"externalActivityRuleIgnoresActivityCategoryRule,omitempty"`

	// MaxTradeSpanWeeks - The maximum number of weeks a shift trade can span
	MaxTradeSpanWeeks *int `json:"maxTradeSpanWeeks,omitempty"`

	// MaxTradesPerAgentPerWeek - The maximum number of shift trades an agent can submit per week
	MaxTradesPerAgentPerWeek *int `json:"maxTradesPerAgentPerWeek,omitempty"`

	// MinMinutesBetweenShifts - The minimum number of minutes between shifts
	MinMinutesBetweenShifts *int `json:"minMinutesBetweenShifts,omitempty"`

	// PlanningPeriodMinPaidViolations - How to handle shift trades which result in violations of planning period minimum paid time constraint
	PlanningPeriodMinPaidViolations *string `json:"planningPeriodMinPaidViolations,omitempty"`

	// PlanningPeriodMaxPaidViolations - How to handle shift trades which result in violations of planning period maximum paid time constraint
	PlanningPeriodMaxPaidViolations *string `json:"planningPeriodMaxPaidViolations,omitempty"`

	// MinMinutesBetweenShiftsViolations - How to handle shift trades which result in violations of minimum number of minutes between shifts constraint
	MinMinutesBetweenShiftsViolations *string `json:"minMinutesBetweenShiftsViolations,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Shifttradesettings) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Shifttradesettings) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

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
		
		ExternalActivityRules *[]Shifttradeexternalactivityrule `json:"externalActivityRules,omitempty"`
		
		ExternalActivityRuleIgnoresActivityCategoryRule *bool `json:"externalActivityRuleIgnoresActivityCategoryRule,omitempty"`
		
		MaxTradeSpanWeeks *int `json:"maxTradeSpanWeeks,omitempty"`
		
		MaxTradesPerAgentPerWeek *int `json:"maxTradesPerAgentPerWeek,omitempty"`
		
		MinMinutesBetweenShifts *int `json:"minMinutesBetweenShifts,omitempty"`
		
		PlanningPeriodMinPaidViolations *string `json:"planningPeriodMinPaidViolations,omitempty"`
		
		PlanningPeriodMaxPaidViolations *string `json:"planningPeriodMaxPaidViolations,omitempty"`
		
		MinMinutesBetweenShiftsViolations *string `json:"minMinutesBetweenShiftsViolations,omitempty"`
		Alias
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
		
		ExternalActivityRules: o.ExternalActivityRules,
		
		ExternalActivityRuleIgnoresActivityCategoryRule: o.ExternalActivityRuleIgnoresActivityCategoryRule,
		
		MaxTradeSpanWeeks: o.MaxTradeSpanWeeks,
		
		MaxTradesPerAgentPerWeek: o.MaxTradesPerAgentPerWeek,
		
		MinMinutesBetweenShifts: o.MinMinutesBetweenShifts,
		
		PlanningPeriodMinPaidViolations: o.PlanningPeriodMinPaidViolations,
		
		PlanningPeriodMaxPaidViolations: o.PlanningPeriodMaxPaidViolations,
		
		MinMinutesBetweenShiftsViolations: o.MinMinutesBetweenShiftsViolations,
		Alias:    (Alias)(o),
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
	
	if ExternalActivityRules, ok := ShifttradesettingsMap["externalActivityRules"].([]interface{}); ok {
		ExternalActivityRulesString, _ := json.Marshal(ExternalActivityRules)
		json.Unmarshal(ExternalActivityRulesString, &o.ExternalActivityRules)
	}
	
	if ExternalActivityRuleIgnoresActivityCategoryRule, ok := ShifttradesettingsMap["externalActivityRuleIgnoresActivityCategoryRule"].(bool); ok {
		o.ExternalActivityRuleIgnoresActivityCategoryRule = &ExternalActivityRuleIgnoresActivityCategoryRule
	}
    
	if MaxTradeSpanWeeks, ok := ShifttradesettingsMap["maxTradeSpanWeeks"].(float64); ok {
		MaxTradeSpanWeeksInt := int(MaxTradeSpanWeeks)
		o.MaxTradeSpanWeeks = &MaxTradeSpanWeeksInt
	}
	
	if MaxTradesPerAgentPerWeek, ok := ShifttradesettingsMap["maxTradesPerAgentPerWeek"].(float64); ok {
		MaxTradesPerAgentPerWeekInt := int(MaxTradesPerAgentPerWeek)
		o.MaxTradesPerAgentPerWeek = &MaxTradesPerAgentPerWeekInt
	}
	
	if MinMinutesBetweenShifts, ok := ShifttradesettingsMap["minMinutesBetweenShifts"].(float64); ok {
		MinMinutesBetweenShiftsInt := int(MinMinutesBetweenShifts)
		o.MinMinutesBetweenShifts = &MinMinutesBetweenShiftsInt
	}
	
	if PlanningPeriodMinPaidViolations, ok := ShifttradesettingsMap["planningPeriodMinPaidViolations"].(string); ok {
		o.PlanningPeriodMinPaidViolations = &PlanningPeriodMinPaidViolations
	}
    
	if PlanningPeriodMaxPaidViolations, ok := ShifttradesettingsMap["planningPeriodMaxPaidViolations"].(string); ok {
		o.PlanningPeriodMaxPaidViolations = &PlanningPeriodMaxPaidViolations
	}
    
	if MinMinutesBetweenShiftsViolations, ok := ShifttradesettingsMap["minMinutesBetweenShiftsViolations"].(string); ok {
		o.MinMinutesBetweenShiftsViolations = &MinMinutesBetweenShiftsViolations
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Shifttradesettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
