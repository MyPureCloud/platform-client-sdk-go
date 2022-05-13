package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanlistitemresponse - Work plan information
type Workplanlistitemresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Enabled - Whether the work plan is enabled for scheduling
	Enabled *bool `json:"enabled,omitempty"`


	// Valid - Whether the work plan is valid or not
	Valid *bool `json:"valid,omitempty"`


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


	// ConstrainPaidTimeGranularity - Whether paid time granularity is constrained for this work plan
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


	// OptionalDays - Optional days to schedule for this work plan. Populate with expand=details
	OptionalDays *Setwrapperdayofweek `json:"optionalDays,omitempty"`


	// ShiftStartVarianceType - This constraint ensures that an agent starts each workday within a user-defined time threshold
	ShiftStartVarianceType *string `json:"shiftStartVarianceType,omitempty"`


	// ShiftStartVariances - Variance in minutes among start times of shifts in this work plan. Populate with expand=details
	ShiftStartVariances *Listwrappershiftstartvariance `json:"shiftStartVariances,omitempty"`


	// Shifts - Shifts in this work plan. Populate with expand=details (defaults to empty list)
	Shifts *[]Workplanshift `json:"shifts,omitempty"`


	// Agents - Agents in this work plan. Populate with expand=details (defaults to empty list)
	Agents *[]Deletableuserreference `json:"agents,omitempty"`


	// AgentCount - Number of agents in this work plan.  Populate with expand=agentCount
	AgentCount *int `json:"agentCount,omitempty"`


	// Metadata - Version metadata for this work plan
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Workplanlistitemresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workplanlistitemresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		Valid *bool `json:"valid,omitempty"`
		
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
		
		Shifts *[]Workplanshift `json:"shifts,omitempty"`
		
		Agents *[]Deletableuserreference `json:"agents,omitempty"`
		
		AgentCount *int `json:"agentCount,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Enabled: o.Enabled,
		
		Valid: o.Valid,
		
		ConstrainWeeklyPaidTime: o.ConstrainWeeklyPaidTime,
		
		FlexibleWeeklyPaidTime: o.FlexibleWeeklyPaidTime,
		
		WeeklyExactPaidMinutes: o.WeeklyExactPaidMinutes,
		
		WeeklyMinimumPaidMinutes: o.WeeklyMinimumPaidMinutes,
		
		WeeklyMaximumPaidMinutes: o.WeeklyMaximumPaidMinutes,
		
		ConstrainPaidTimeGranularity: o.ConstrainPaidTimeGranularity,
		
		PaidTimeGranularityMinutes: o.PaidTimeGranularityMinutes,
		
		ConstrainMinimumTimeBetweenShifts: o.ConstrainMinimumTimeBetweenShifts,
		
		MinimumTimeBetweenShiftsMinutes: o.MinimumTimeBetweenShiftsMinutes,
		
		MaximumDays: o.MaximumDays,
		
		MinimumConsecutiveNonWorkingMinutesPerWeek: o.MinimumConsecutiveNonWorkingMinutesPerWeek,
		
		ConstrainMaximumConsecutiveWorkingWeekends: o.ConstrainMaximumConsecutiveWorkingWeekends,
		
		MaximumConsecutiveWorkingWeekends: o.MaximumConsecutiveWorkingWeekends,
		
		MinimumWorkingDaysPerWeek: o.MinimumWorkingDaysPerWeek,
		
		ConstrainMaximumConsecutiveWorkingDays: o.ConstrainMaximumConsecutiveWorkingDays,
		
		MaximumConsecutiveWorkingDays: o.MaximumConsecutiveWorkingDays,
		
		MinimumShiftStartDistanceMinutes: o.MinimumShiftStartDistanceMinutes,
		
		MinimumDaysOffPerPlanningPeriod: o.MinimumDaysOffPerPlanningPeriod,
		
		MaximumDaysOffPerPlanningPeriod: o.MaximumDaysOffPerPlanningPeriod,
		
		MinimumPaidMinutesPerPlanningPeriod: o.MinimumPaidMinutesPerPlanningPeriod,
		
		MaximumPaidMinutesPerPlanningPeriod: o.MaximumPaidMinutesPerPlanningPeriod,
		
		OptionalDays: o.OptionalDays,
		
		ShiftStartVarianceType: o.ShiftStartVarianceType,
		
		ShiftStartVariances: o.ShiftStartVariances,
		
		Shifts: o.Shifts,
		
		Agents: o.Agents,
		
		AgentCount: o.AgentCount,
		
		Metadata: o.Metadata,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Workplanlistitemresponse) UnmarshalJSON(b []byte) error {
	var WorkplanlistitemresponseMap map[string]interface{}
	err := json.Unmarshal(b, &WorkplanlistitemresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WorkplanlistitemresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WorkplanlistitemresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Enabled, ok := WorkplanlistitemresponseMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if Valid, ok := WorkplanlistitemresponseMap["valid"].(bool); ok {
		o.Valid = &Valid
	}
    
	if ConstrainWeeklyPaidTime, ok := WorkplanlistitemresponseMap["constrainWeeklyPaidTime"].(bool); ok {
		o.ConstrainWeeklyPaidTime = &ConstrainWeeklyPaidTime
	}
    
	if FlexibleWeeklyPaidTime, ok := WorkplanlistitemresponseMap["flexibleWeeklyPaidTime"].(bool); ok {
		o.FlexibleWeeklyPaidTime = &FlexibleWeeklyPaidTime
	}
    
	if WeeklyExactPaidMinutes, ok := WorkplanlistitemresponseMap["weeklyExactPaidMinutes"].(float64); ok {
		WeeklyExactPaidMinutesInt := int(WeeklyExactPaidMinutes)
		o.WeeklyExactPaidMinutes = &WeeklyExactPaidMinutesInt
	}
	
	if WeeklyMinimumPaidMinutes, ok := WorkplanlistitemresponseMap["weeklyMinimumPaidMinutes"].(float64); ok {
		WeeklyMinimumPaidMinutesInt := int(WeeklyMinimumPaidMinutes)
		o.WeeklyMinimumPaidMinutes = &WeeklyMinimumPaidMinutesInt
	}
	
	if WeeklyMaximumPaidMinutes, ok := WorkplanlistitemresponseMap["weeklyMaximumPaidMinutes"].(float64); ok {
		WeeklyMaximumPaidMinutesInt := int(WeeklyMaximumPaidMinutes)
		o.WeeklyMaximumPaidMinutes = &WeeklyMaximumPaidMinutesInt
	}
	
	if ConstrainPaidTimeGranularity, ok := WorkplanlistitemresponseMap["constrainPaidTimeGranularity"].(bool); ok {
		o.ConstrainPaidTimeGranularity = &ConstrainPaidTimeGranularity
	}
    
	if PaidTimeGranularityMinutes, ok := WorkplanlistitemresponseMap["paidTimeGranularityMinutes"].(float64); ok {
		PaidTimeGranularityMinutesInt := int(PaidTimeGranularityMinutes)
		o.PaidTimeGranularityMinutes = &PaidTimeGranularityMinutesInt
	}
	
	if ConstrainMinimumTimeBetweenShifts, ok := WorkplanlistitemresponseMap["constrainMinimumTimeBetweenShifts"].(bool); ok {
		o.ConstrainMinimumTimeBetweenShifts = &ConstrainMinimumTimeBetweenShifts
	}
    
	if MinimumTimeBetweenShiftsMinutes, ok := WorkplanlistitemresponseMap["minimumTimeBetweenShiftsMinutes"].(float64); ok {
		MinimumTimeBetweenShiftsMinutesInt := int(MinimumTimeBetweenShiftsMinutes)
		o.MinimumTimeBetweenShiftsMinutes = &MinimumTimeBetweenShiftsMinutesInt
	}
	
	if MaximumDays, ok := WorkplanlistitemresponseMap["maximumDays"].(float64); ok {
		MaximumDaysInt := int(MaximumDays)
		o.MaximumDays = &MaximumDaysInt
	}
	
	if MinimumConsecutiveNonWorkingMinutesPerWeek, ok := WorkplanlistitemresponseMap["minimumConsecutiveNonWorkingMinutesPerWeek"].(float64); ok {
		MinimumConsecutiveNonWorkingMinutesPerWeekInt := int(MinimumConsecutiveNonWorkingMinutesPerWeek)
		o.MinimumConsecutiveNonWorkingMinutesPerWeek = &MinimumConsecutiveNonWorkingMinutesPerWeekInt
	}
	
	if ConstrainMaximumConsecutiveWorkingWeekends, ok := WorkplanlistitemresponseMap["constrainMaximumConsecutiveWorkingWeekends"].(bool); ok {
		o.ConstrainMaximumConsecutiveWorkingWeekends = &ConstrainMaximumConsecutiveWorkingWeekends
	}
    
	if MaximumConsecutiveWorkingWeekends, ok := WorkplanlistitemresponseMap["maximumConsecutiveWorkingWeekends"].(float64); ok {
		MaximumConsecutiveWorkingWeekendsInt := int(MaximumConsecutiveWorkingWeekends)
		o.MaximumConsecutiveWorkingWeekends = &MaximumConsecutiveWorkingWeekendsInt
	}
	
	if MinimumWorkingDaysPerWeek, ok := WorkplanlistitemresponseMap["minimumWorkingDaysPerWeek"].(float64); ok {
		MinimumWorkingDaysPerWeekInt := int(MinimumWorkingDaysPerWeek)
		o.MinimumWorkingDaysPerWeek = &MinimumWorkingDaysPerWeekInt
	}
	
	if ConstrainMaximumConsecutiveWorkingDays, ok := WorkplanlistitemresponseMap["constrainMaximumConsecutiveWorkingDays"].(bool); ok {
		o.ConstrainMaximumConsecutiveWorkingDays = &ConstrainMaximumConsecutiveWorkingDays
	}
    
	if MaximumConsecutiveWorkingDays, ok := WorkplanlistitemresponseMap["maximumConsecutiveWorkingDays"].(float64); ok {
		MaximumConsecutiveWorkingDaysInt := int(MaximumConsecutiveWorkingDays)
		o.MaximumConsecutiveWorkingDays = &MaximumConsecutiveWorkingDaysInt
	}
	
	if MinimumShiftStartDistanceMinutes, ok := WorkplanlistitemresponseMap["minimumShiftStartDistanceMinutes"].(float64); ok {
		MinimumShiftStartDistanceMinutesInt := int(MinimumShiftStartDistanceMinutes)
		o.MinimumShiftStartDistanceMinutes = &MinimumShiftStartDistanceMinutesInt
	}
	
	if MinimumDaysOffPerPlanningPeriod, ok := WorkplanlistitemresponseMap["minimumDaysOffPerPlanningPeriod"].(float64); ok {
		MinimumDaysOffPerPlanningPeriodInt := int(MinimumDaysOffPerPlanningPeriod)
		o.MinimumDaysOffPerPlanningPeriod = &MinimumDaysOffPerPlanningPeriodInt
	}
	
	if MaximumDaysOffPerPlanningPeriod, ok := WorkplanlistitemresponseMap["maximumDaysOffPerPlanningPeriod"].(float64); ok {
		MaximumDaysOffPerPlanningPeriodInt := int(MaximumDaysOffPerPlanningPeriod)
		o.MaximumDaysOffPerPlanningPeriod = &MaximumDaysOffPerPlanningPeriodInt
	}
	
	if MinimumPaidMinutesPerPlanningPeriod, ok := WorkplanlistitemresponseMap["minimumPaidMinutesPerPlanningPeriod"].(float64); ok {
		MinimumPaidMinutesPerPlanningPeriodInt := int(MinimumPaidMinutesPerPlanningPeriod)
		o.MinimumPaidMinutesPerPlanningPeriod = &MinimumPaidMinutesPerPlanningPeriodInt
	}
	
	if MaximumPaidMinutesPerPlanningPeriod, ok := WorkplanlistitemresponseMap["maximumPaidMinutesPerPlanningPeriod"].(float64); ok {
		MaximumPaidMinutesPerPlanningPeriodInt := int(MaximumPaidMinutesPerPlanningPeriod)
		o.MaximumPaidMinutesPerPlanningPeriod = &MaximumPaidMinutesPerPlanningPeriodInt
	}
	
	if OptionalDays, ok := WorkplanlistitemresponseMap["optionalDays"].(map[string]interface{}); ok {
		OptionalDaysString, _ := json.Marshal(OptionalDays)
		json.Unmarshal(OptionalDaysString, &o.OptionalDays)
	}
	
	if ShiftStartVarianceType, ok := WorkplanlistitemresponseMap["shiftStartVarianceType"].(string); ok {
		o.ShiftStartVarianceType = &ShiftStartVarianceType
	}
    
	if ShiftStartVariances, ok := WorkplanlistitemresponseMap["shiftStartVariances"].(map[string]interface{}); ok {
		ShiftStartVariancesString, _ := json.Marshal(ShiftStartVariances)
		json.Unmarshal(ShiftStartVariancesString, &o.ShiftStartVariances)
	}
	
	if Shifts, ok := WorkplanlistitemresponseMap["shifts"].([]interface{}); ok {
		ShiftsString, _ := json.Marshal(Shifts)
		json.Unmarshal(ShiftsString, &o.Shifts)
	}
	
	if Agents, ok := WorkplanlistitemresponseMap["agents"].([]interface{}); ok {
		AgentsString, _ := json.Marshal(Agents)
		json.Unmarshal(AgentsString, &o.Agents)
	}
	
	if AgentCount, ok := WorkplanlistitemresponseMap["agentCount"].(float64); ok {
		AgentCountInt := int(AgentCount)
		o.AgentCount = &AgentCountInt
	}
	
	if Metadata, ok := WorkplanlistitemresponseMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if SelfUri, ok := WorkplanlistitemresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Workplanlistitemresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
