package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createworkplan - Work plan information
type Createworkplan struct { 
	// Name - Name of this work plan
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


	// ConstrainPaidTimeGranularity - Whether paid time granularity should be constrained for this workplan
	ConstrainPaidTimeGranularity *bool `json:"constrainPaidTimeGranularity,omitempty"`


	// PaidTimeGranularityMinutes - Granularity in minutes allowed for shift paid time in this work plan. Used if constrainPaidTimeGranularity == true
	PaidTimeGranularityMinutes *int `json:"paidTimeGranularityMinutes,omitempty"`


	// ConstrainMinimumTimeBetweenShifts - Whether the minimum time between shifts constraint is enabled for this work plan
	ConstrainMinimumTimeBetweenShifts *bool `json:"constrainMinimumTimeBetweenShifts,omitempty"`


	// MinimumTimeBetweenShiftsMinutes - Minimum time between shifts in minutes defined in this work plan. Used if constrainMinimumTimeBetweenShifts == true
	MinimumTimeBetweenShiftsMinutes *int `json:"minimumTimeBetweenShiftsMinutes,omitempty"`


	// MaximumDays - Maximum number days in a week allowed to be scheduled for this work plan
	MaximumDays *int `json:"maximumDays,omitempty"`


	// MinimumConsecutiveNonWorkingMinutesPerWeek - Minimum amount of consecutive non working minutes per week that agents who are assigned this work plan are allowed to have off
	MinimumConsecutiveNonWorkingMinutesPerWeek *int `json:"minimumConsecutiveNonWorkingMinutesPerWeek,omitempty"`


	// ConstrainMaximumConsecutiveWorkingWeekends - Whether to constrain the maximum consecutive working weekends
	ConstrainMaximumConsecutiveWorkingWeekends *bool `json:"constrainMaximumConsecutiveWorkingWeekends,omitempty"`


	// MaximumConsecutiveWorkingWeekends - The maximum number of consecutive weekends that agents who are assigned to this work plan are allowed to work
	MaximumConsecutiveWorkingWeekends *int `json:"maximumConsecutiveWorkingWeekends,omitempty"`


	// MinimumWorkingDaysPerWeek - The minimum number of days that agents assigned to a work plan must work per week
	MinimumWorkingDaysPerWeek *int `json:"minimumWorkingDaysPerWeek,omitempty"`


	// ConstrainMaximumConsecutiveWorkingDays - Whether to constrain the maximum consecutive working days
	ConstrainMaximumConsecutiveWorkingDays *bool `json:"constrainMaximumConsecutiveWorkingDays,omitempty"`


	// MaximumConsecutiveWorkingDays - The maximum number of consecutive days that agents assigned to this work plan are allowed to work. Used if constrainMaximumConsecutiveWorkingDays == true
	MaximumConsecutiveWorkingDays *int `json:"maximumConsecutiveWorkingDays,omitempty"`


	// MinimumShiftStartDistanceMinutes - The time period in minutes for the duration between the start times of two consecutive working days
	MinimumShiftStartDistanceMinutes *int `json:"minimumShiftStartDistanceMinutes,omitempty"`


	// MinimumDaysOffPerPlanningPeriod - Minimum days off in the planning period
	MinimumDaysOffPerPlanningPeriod *int `json:"minimumDaysOffPerPlanningPeriod,omitempty"`


	// MaximumDaysOffPerPlanningPeriod - Maximum days off in the planning period
	MaximumDaysOffPerPlanningPeriod *int `json:"maximumDaysOffPerPlanningPeriod,omitempty"`


	// MinimumPaidMinutesPerPlanningPeriod - Minimum paid minutes in the planning period
	MinimumPaidMinutesPerPlanningPeriod *int `json:"minimumPaidMinutesPerPlanningPeriod,omitempty"`


	// MaximumPaidMinutesPerPlanningPeriod - Maximum paid minutes in the planning period
	MaximumPaidMinutesPerPlanningPeriod *int `json:"maximumPaidMinutesPerPlanningPeriod,omitempty"`


	// OptionalDays - Optional days to schedule for this work plan
	OptionalDays *Setwrapperdayofweek `json:"optionalDays,omitempty"`


	// ShiftStartVarianceType - This constraint ensures that an agent starts each workday within a user-defined time threshold
	ShiftStartVarianceType *string `json:"shiftStartVarianceType,omitempty"`


	// ShiftStartVariances - Variance in minutes among start times of shifts in this work plan
	ShiftStartVariances *Listwrappershiftstartvariance `json:"shiftStartVariances,omitempty"`


	// Shifts - Shifts in this work plan
	Shifts *[]Createworkplanshift `json:"shifts,omitempty"`


	// Agents - Agents in this work plan
	Agents *[]Userreference `json:"agents,omitempty"`

}

func (u *Createworkplan) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createworkplan

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		ConstrainWeeklyPaidTime *bool `json:"constrainWeeklyPaidTime,omitempty"`
		
		FlexibleWeeklyPaidTime *bool `json:"flexibleWeeklyPaidTime,omitempty"`
		
		WeeklyExactPaidMinutes *int `json:"weeklyExactPaidMinutes,omitempty"`
		
		WeeklyMinimumPaidMinutes *int `json:"weeklyMinimumPaidMinutes,omitempty"`
		
		WeeklyMaximumPaidMinutes *int `json:"weeklyMaximumPaidMinutes,omitempty"`
		
		ConstrainPaidTimeGranularity *bool `json:"constrainPaidTimeGranularity,omitempty"`
		
		PaidTimeGranularityMinutes *int `json:"paidTimeGranularityMinutes,omitempty"`
		
		ConstrainMinimumTimeBetweenShifts *bool `json:"constrainMinimumTimeBetweenShifts,omitempty"`
		
		MinimumTimeBetweenShiftsMinutes *int `json:"minimumTimeBetweenShiftsMinutes,omitempty"`
		
		MaximumDays *int `json:"maximumDays,omitempty"`
		
		MinimumConsecutiveNonWorkingMinutesPerWeek *int `json:"minimumConsecutiveNonWorkingMinutesPerWeek,omitempty"`
		
		ConstrainMaximumConsecutiveWorkingWeekends *bool `json:"constrainMaximumConsecutiveWorkingWeekends,omitempty"`
		
		MaximumConsecutiveWorkingWeekends *int `json:"maximumConsecutiveWorkingWeekends,omitempty"`
		
		MinimumWorkingDaysPerWeek *int `json:"minimumWorkingDaysPerWeek,omitempty"`
		
		ConstrainMaximumConsecutiveWorkingDays *bool `json:"constrainMaximumConsecutiveWorkingDays,omitempty"`
		
		MaximumConsecutiveWorkingDays *int `json:"maximumConsecutiveWorkingDays,omitempty"`
		
		MinimumShiftStartDistanceMinutes *int `json:"minimumShiftStartDistanceMinutes,omitempty"`
		
		MinimumDaysOffPerPlanningPeriod *int `json:"minimumDaysOffPerPlanningPeriod,omitempty"`
		
		MaximumDaysOffPerPlanningPeriod *int `json:"maximumDaysOffPerPlanningPeriod,omitempty"`
		
		MinimumPaidMinutesPerPlanningPeriod *int `json:"minimumPaidMinutesPerPlanningPeriod,omitempty"`
		
		MaximumPaidMinutesPerPlanningPeriod *int `json:"maximumPaidMinutesPerPlanningPeriod,omitempty"`
		
		OptionalDays *Setwrapperdayofweek `json:"optionalDays,omitempty"`
		
		ShiftStartVarianceType *string `json:"shiftStartVarianceType,omitempty"`
		
		ShiftStartVariances *Listwrappershiftstartvariance `json:"shiftStartVariances,omitempty"`
		
		Shifts *[]Createworkplanshift `json:"shifts,omitempty"`
		
		Agents *[]Userreference `json:"agents,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Enabled: u.Enabled,
		
		ConstrainWeeklyPaidTime: u.ConstrainWeeklyPaidTime,
		
		FlexibleWeeklyPaidTime: u.FlexibleWeeklyPaidTime,
		
		WeeklyExactPaidMinutes: u.WeeklyExactPaidMinutes,
		
		WeeklyMinimumPaidMinutes: u.WeeklyMinimumPaidMinutes,
		
		WeeklyMaximumPaidMinutes: u.WeeklyMaximumPaidMinutes,
		
		ConstrainPaidTimeGranularity: u.ConstrainPaidTimeGranularity,
		
		PaidTimeGranularityMinutes: u.PaidTimeGranularityMinutes,
		
		ConstrainMinimumTimeBetweenShifts: u.ConstrainMinimumTimeBetweenShifts,
		
		MinimumTimeBetweenShiftsMinutes: u.MinimumTimeBetweenShiftsMinutes,
		
		MaximumDays: u.MaximumDays,
		
		MinimumConsecutiveNonWorkingMinutesPerWeek: u.MinimumConsecutiveNonWorkingMinutesPerWeek,
		
		ConstrainMaximumConsecutiveWorkingWeekends: u.ConstrainMaximumConsecutiveWorkingWeekends,
		
		MaximumConsecutiveWorkingWeekends: u.MaximumConsecutiveWorkingWeekends,
		
		MinimumWorkingDaysPerWeek: u.MinimumWorkingDaysPerWeek,
		
		ConstrainMaximumConsecutiveWorkingDays: u.ConstrainMaximumConsecutiveWorkingDays,
		
		MaximumConsecutiveWorkingDays: u.MaximumConsecutiveWorkingDays,
		
		MinimumShiftStartDistanceMinutes: u.MinimumShiftStartDistanceMinutes,
		
		MinimumDaysOffPerPlanningPeriod: u.MinimumDaysOffPerPlanningPeriod,
		
		MaximumDaysOffPerPlanningPeriod: u.MaximumDaysOffPerPlanningPeriod,
		
		MinimumPaidMinutesPerPlanningPeriod: u.MinimumPaidMinutesPerPlanningPeriod,
		
		MaximumPaidMinutesPerPlanningPeriod: u.MaximumPaidMinutesPerPlanningPeriod,
		
		OptionalDays: u.OptionalDays,
		
		ShiftStartVarianceType: u.ShiftStartVarianceType,
		
		ShiftStartVariances: u.ShiftStartVariances,
		
		Shifts: u.Shifts,
		
		Agents: u.Agents,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Createworkplan) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
