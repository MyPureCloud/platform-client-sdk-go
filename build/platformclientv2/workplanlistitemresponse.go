package platformclientv2
import (
	"encoding/json"
)

// Workplanlistitemresponse - Work plan information
type Workplanlistitemresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Enabled - Whether the work plan is enabled for scheduling
	Enabled *bool `json:"enabled,omitempty"`


	// ConstrainWeeklyPaidTime - Whether the weekly paid time constraint is enabled for this work plan
	ConstrainWeeklyPaidTime *bool `json:"constrainWeeklyPaidTime,omitempty"`


	// FlexibleWeeklyPaidTime - Whether the weekly paid time constraint is flexible for this work plan
	FlexibleWeeklyPaidTime *bool `json:"flexibleWeeklyPaidTime,omitempty"`


	// WeeklyExactPaidMinutes - Exact weekly paid time in minutes for this work plan. Used if flexibleWeeklyPaidTime == false
	WeeklyExactPaidMinutes *int `json:"weeklyExactPaidMinutes,omitempty"`


	// WeeklyMinimumPaidMinutes - Minimum weekly paid time in minutes for this work plan. Used if flexibleWeeklyPaidTime == true
	WeeklyMinimumPaidMinutes *int `json:"weeklyMinimumPaidMinutes,omitempty"`


	// WeeklyMaximumPaidMinutes - Maximum weekly paid time in minutes for this work plan. Used if flexibleWeeklyPaidTime == true
	WeeklyMaximumPaidMinutes *int `json:"weeklyMaximumPaidMinutes,omitempty"`


	// ConstrainPaidTimeGranularity - Whether paid time granularity is constrained for this workplan
	ConstrainPaidTimeGranularity *bool `json:"constrainPaidTimeGranularity,omitempty"`


	// PaidTimeGranularityMinutes - Granularity in minutes allowed for shift paid time in this work plan. Used if constrainPaidTimeGranularity == true
	PaidTimeGranularityMinutes *int `json:"paidTimeGranularityMinutes,omitempty"`


	// ConstrainMinimumTimeBetweenShifts - Whether the minimum time between shifts constraint is enabled for this work plan
	ConstrainMinimumTimeBetweenShifts *bool `json:"constrainMinimumTimeBetweenShifts,omitempty"`


	// MinimumTimeBetweenShiftsMinutes - Minimum time between shifts in minutes defined in this work plan. Used if constrainMinimumTimeBetweenShifts == true
	MinimumTimeBetweenShiftsMinutes *int `json:"minimumTimeBetweenShiftsMinutes,omitempty"`


	// MaximumDays - Maximum number days in a week allowed to be scheduled for this work plan
	MaximumDays *int `json:"maximumDays,omitempty"`


	// MinimumWorkingDaysPerWeek - The minimum number of days that agents assigned to a work plan must work per week
	MinimumWorkingDaysPerWeek *int `json:"minimumWorkingDaysPerWeek,omitempty"`


	// OptionalDays - Optional days to schedule for this work plan. Populate with expand=details
	OptionalDays *Setwrapperdayofweek `json:"optionalDays,omitempty"`


	// ShiftStartVariances - Variance in minutes among start times of shifts in this work plan. Populate with expand=details
	ShiftStartVariances *Listwrappershiftstartvariance `json:"shiftStartVariances,omitempty"`


	// Shifts - Shifts in this work plan. Populate with expand=details (defaults to empty list)
	Shifts *[]Workplanshift `json:"shifts,omitempty"`


	// Agents - Agents in this work plan. Populate with expand=details (defaults to empty list)
	Agents *[]Deletableuserreference `json:"agents,omitempty"`


	// Metadata - Version metadata for this work plan
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// AgentCount - Number of agents in this work plan.  Populate with expand=agentCount
	AgentCount *int `json:"agentCount,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workplanlistitemresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
