package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createworkplan
type Createworkplan struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createworkplan) SetField(field string, fieldValue interface{}) {
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

func (o Createworkplan) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
		Alias
	}{ 
		Name: o.Name,
		
		Enabled: o.Enabled,
		
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
		Alias:    (Alias)(o),
	})
}

func (o *Createworkplan) UnmarshalJSON(b []byte) error {
	var CreateworkplanMap map[string]interface{}
	err := json.Unmarshal(b, &CreateworkplanMap)
	if err != nil {
		return err
	}
	
	if Name, ok := CreateworkplanMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Enabled, ok := CreateworkplanMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if ConstrainWeeklyPaidTime, ok := CreateworkplanMap["constrainWeeklyPaidTime"].(bool); ok {
		o.ConstrainWeeklyPaidTime = &ConstrainWeeklyPaidTime
	}
    
	if FlexibleWeeklyPaidTime, ok := CreateworkplanMap["flexibleWeeklyPaidTime"].(bool); ok {
		o.FlexibleWeeklyPaidTime = &FlexibleWeeklyPaidTime
	}
    
	if WeeklyExactPaidMinutes, ok := CreateworkplanMap["weeklyExactPaidMinutes"].(float64); ok {
		WeeklyExactPaidMinutesInt := int(WeeklyExactPaidMinutes)
		o.WeeklyExactPaidMinutes = &WeeklyExactPaidMinutesInt
	}
	
	if WeeklyMinimumPaidMinutes, ok := CreateworkplanMap["weeklyMinimumPaidMinutes"].(float64); ok {
		WeeklyMinimumPaidMinutesInt := int(WeeklyMinimumPaidMinutes)
		o.WeeklyMinimumPaidMinutes = &WeeklyMinimumPaidMinutesInt
	}
	
	if WeeklyMaximumPaidMinutes, ok := CreateworkplanMap["weeklyMaximumPaidMinutes"].(float64); ok {
		WeeklyMaximumPaidMinutesInt := int(WeeklyMaximumPaidMinutes)
		o.WeeklyMaximumPaidMinutes = &WeeklyMaximumPaidMinutesInt
	}
	
	if ConstrainPaidTimeGranularity, ok := CreateworkplanMap["constrainPaidTimeGranularity"].(bool); ok {
		o.ConstrainPaidTimeGranularity = &ConstrainPaidTimeGranularity
	}
    
	if PaidTimeGranularityMinutes, ok := CreateworkplanMap["paidTimeGranularityMinutes"].(float64); ok {
		PaidTimeGranularityMinutesInt := int(PaidTimeGranularityMinutes)
		o.PaidTimeGranularityMinutes = &PaidTimeGranularityMinutesInt
	}
	
	if ConstrainMinimumTimeBetweenShifts, ok := CreateworkplanMap["constrainMinimumTimeBetweenShifts"].(bool); ok {
		o.ConstrainMinimumTimeBetweenShifts = &ConstrainMinimumTimeBetweenShifts
	}
    
	if MinimumTimeBetweenShiftsMinutes, ok := CreateworkplanMap["minimumTimeBetweenShiftsMinutes"].(float64); ok {
		MinimumTimeBetweenShiftsMinutesInt := int(MinimumTimeBetweenShiftsMinutes)
		o.MinimumTimeBetweenShiftsMinutes = &MinimumTimeBetweenShiftsMinutesInt
	}
	
	if MaximumDays, ok := CreateworkplanMap["maximumDays"].(float64); ok {
		MaximumDaysInt := int(MaximumDays)
		o.MaximumDays = &MaximumDaysInt
	}
	
	if MinimumConsecutiveNonWorkingMinutesPerWeek, ok := CreateworkplanMap["minimumConsecutiveNonWorkingMinutesPerWeek"].(float64); ok {
		MinimumConsecutiveNonWorkingMinutesPerWeekInt := int(MinimumConsecutiveNonWorkingMinutesPerWeek)
		o.MinimumConsecutiveNonWorkingMinutesPerWeek = &MinimumConsecutiveNonWorkingMinutesPerWeekInt
	}
	
	if ConstrainMaximumConsecutiveWorkingWeekends, ok := CreateworkplanMap["constrainMaximumConsecutiveWorkingWeekends"].(bool); ok {
		o.ConstrainMaximumConsecutiveWorkingWeekends = &ConstrainMaximumConsecutiveWorkingWeekends
	}
    
	if MaximumConsecutiveWorkingWeekends, ok := CreateworkplanMap["maximumConsecutiveWorkingWeekends"].(float64); ok {
		MaximumConsecutiveWorkingWeekendsInt := int(MaximumConsecutiveWorkingWeekends)
		o.MaximumConsecutiveWorkingWeekends = &MaximumConsecutiveWorkingWeekendsInt
	}
	
	if MinimumWorkingDaysPerWeek, ok := CreateworkplanMap["minimumWorkingDaysPerWeek"].(float64); ok {
		MinimumWorkingDaysPerWeekInt := int(MinimumWorkingDaysPerWeek)
		o.MinimumWorkingDaysPerWeek = &MinimumWorkingDaysPerWeekInt
	}
	
	if ConstrainMaximumConsecutiveWorkingDays, ok := CreateworkplanMap["constrainMaximumConsecutiveWorkingDays"].(bool); ok {
		o.ConstrainMaximumConsecutiveWorkingDays = &ConstrainMaximumConsecutiveWorkingDays
	}
    
	if MaximumConsecutiveWorkingDays, ok := CreateworkplanMap["maximumConsecutiveWorkingDays"].(float64); ok {
		MaximumConsecutiveWorkingDaysInt := int(MaximumConsecutiveWorkingDays)
		o.MaximumConsecutiveWorkingDays = &MaximumConsecutiveWorkingDaysInt
	}
	
	if MinimumShiftStartDistanceMinutes, ok := CreateworkplanMap["minimumShiftStartDistanceMinutes"].(float64); ok {
		MinimumShiftStartDistanceMinutesInt := int(MinimumShiftStartDistanceMinutes)
		o.MinimumShiftStartDistanceMinutes = &MinimumShiftStartDistanceMinutesInt
	}
	
	if MinimumDaysOffPerPlanningPeriod, ok := CreateworkplanMap["minimumDaysOffPerPlanningPeriod"].(float64); ok {
		MinimumDaysOffPerPlanningPeriodInt := int(MinimumDaysOffPerPlanningPeriod)
		o.MinimumDaysOffPerPlanningPeriod = &MinimumDaysOffPerPlanningPeriodInt
	}
	
	if MaximumDaysOffPerPlanningPeriod, ok := CreateworkplanMap["maximumDaysOffPerPlanningPeriod"].(float64); ok {
		MaximumDaysOffPerPlanningPeriodInt := int(MaximumDaysOffPerPlanningPeriod)
		o.MaximumDaysOffPerPlanningPeriod = &MaximumDaysOffPerPlanningPeriodInt
	}
	
	if MinimumPaidMinutesPerPlanningPeriod, ok := CreateworkplanMap["minimumPaidMinutesPerPlanningPeriod"].(float64); ok {
		MinimumPaidMinutesPerPlanningPeriodInt := int(MinimumPaidMinutesPerPlanningPeriod)
		o.MinimumPaidMinutesPerPlanningPeriod = &MinimumPaidMinutesPerPlanningPeriodInt
	}
	
	if MaximumPaidMinutesPerPlanningPeriod, ok := CreateworkplanMap["maximumPaidMinutesPerPlanningPeriod"].(float64); ok {
		MaximumPaidMinutesPerPlanningPeriodInt := int(MaximumPaidMinutesPerPlanningPeriod)
		o.MaximumPaidMinutesPerPlanningPeriod = &MaximumPaidMinutesPerPlanningPeriodInt
	}
	
	if OptionalDays, ok := CreateworkplanMap["optionalDays"].(map[string]interface{}); ok {
		OptionalDaysString, _ := json.Marshal(OptionalDays)
		json.Unmarshal(OptionalDaysString, &o.OptionalDays)
	}
	
	if ShiftStartVarianceType, ok := CreateworkplanMap["shiftStartVarianceType"].(string); ok {
		o.ShiftStartVarianceType = &ShiftStartVarianceType
	}
    
	if ShiftStartVariances, ok := CreateworkplanMap["shiftStartVariances"].(map[string]interface{}); ok {
		ShiftStartVariancesString, _ := json.Marshal(ShiftStartVariances)
		json.Unmarshal(ShiftStartVariancesString, &o.ShiftStartVariances)
	}
	
	if Shifts, ok := CreateworkplanMap["shifts"].([]interface{}); ok {
		ShiftsString, _ := json.Marshal(Shifts)
		json.Unmarshal(ShiftsString, &o.Shifts)
	}
	
	if Agents, ok := CreateworkplanMap["agents"].([]interface{}); ok {
		AgentsString, _ := json.Marshal(Agents)
		json.Unmarshal(AgentsString, &o.Agents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createworkplan) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
