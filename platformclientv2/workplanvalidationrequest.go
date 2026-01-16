package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanvalidationrequest
type Workplanvalidationrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

	// ConstrainMaximumWorkingWeekendsPerPlanningPeriod - Whether to constrain the maximum working weekends in the planning period
	ConstrainMaximumWorkingWeekendsPerPlanningPeriod *bool `json:"constrainMaximumWorkingWeekendsPerPlanningPeriod,omitempty"`

	// MaximumWorkingWeekendsPerPlanningPeriod - Maximum working weekends in the planning period
	MaximumWorkingWeekendsPerPlanningPeriod *int `json:"maximumWorkingWeekendsPerPlanningPeriod,omitempty"`

	// OptionalDays - Optional days to schedule for this work plan
	OptionalDays *Setwrapperdayofweek `json:"optionalDays,omitempty"`

	// ShiftStartVarianceType - This constraint ensures that an agent starts each workday within a user-defined time threshold
	ShiftStartVarianceType *string `json:"shiftStartVarianceType,omitempty"`

	// ShiftStartVariancePeriod - The length of the period over which the maximum shift start time variance is applied
	ShiftStartVariancePeriod *string `json:"shiftStartVariancePeriod,omitempty"`

	// ShiftStartVariances - Variance in minutes among start times of shifts in this work plan
	ShiftStartVariances *Listwrappershiftstartvariance `json:"shiftStartVariances,omitempty"`

	// Shifts - Shifts in this work plan
	Shifts *[]Workplanshift `json:"shifts,omitempty"`

	// Agents - Agents in this work plan
	Agents *[]Deletableuserreference `json:"agents,omitempty"`

	// AgentCount - Number of agents in this work plan
	AgentCount *int `json:"agentCount,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workplanvalidationrequest) SetField(field string, fieldValue interface{}) {
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

func (o Workplanvalidationrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Workplanvalidationrequest
	
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
		
		ConstrainMaximumWorkingWeekendsPerPlanningPeriod *bool `json:"constrainMaximumWorkingWeekendsPerPlanningPeriod,omitempty"`
		
		MaximumWorkingWeekendsPerPlanningPeriod *int `json:"maximumWorkingWeekendsPerPlanningPeriod,omitempty"`
		
		OptionalDays *Setwrapperdayofweek `json:"optionalDays,omitempty"`
		
		ShiftStartVarianceType *string `json:"shiftStartVarianceType,omitempty"`
		
		ShiftStartVariancePeriod *string `json:"shiftStartVariancePeriod,omitempty"`
		
		ShiftStartVariances *Listwrappershiftstartvariance `json:"shiftStartVariances,omitempty"`
		
		Shifts *[]Workplanshift `json:"shifts,omitempty"`
		
		Agents *[]Deletableuserreference `json:"agents,omitempty"`
		
		AgentCount *int `json:"agentCount,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
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
		
		ConstrainMaximumWorkingWeekendsPerPlanningPeriod: o.ConstrainMaximumWorkingWeekendsPerPlanningPeriod,
		
		MaximumWorkingWeekendsPerPlanningPeriod: o.MaximumWorkingWeekendsPerPlanningPeriod,
		
		OptionalDays: o.OptionalDays,
		
		ShiftStartVarianceType: o.ShiftStartVarianceType,
		
		ShiftStartVariancePeriod: o.ShiftStartVariancePeriod,
		
		ShiftStartVariances: o.ShiftStartVariances,
		
		Shifts: o.Shifts,
		
		Agents: o.Agents,
		
		AgentCount: o.AgentCount,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Workplanvalidationrequest) UnmarshalJSON(b []byte) error {
	var WorkplanvalidationrequestMap map[string]interface{}
	err := json.Unmarshal(b, &WorkplanvalidationrequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WorkplanvalidationrequestMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WorkplanvalidationrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Enabled, ok := WorkplanvalidationrequestMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if Valid, ok := WorkplanvalidationrequestMap["valid"].(bool); ok {
		o.Valid = &Valid
	}
    
	if ConstrainWeeklyPaidTime, ok := WorkplanvalidationrequestMap["constrainWeeklyPaidTime"].(bool); ok {
		o.ConstrainWeeklyPaidTime = &ConstrainWeeklyPaidTime
	}
    
	if FlexibleWeeklyPaidTime, ok := WorkplanvalidationrequestMap["flexibleWeeklyPaidTime"].(bool); ok {
		o.FlexibleWeeklyPaidTime = &FlexibleWeeklyPaidTime
	}
    
	if WeeklyExactPaidMinutes, ok := WorkplanvalidationrequestMap["weeklyExactPaidMinutes"].(float64); ok {
		WeeklyExactPaidMinutesInt := int(WeeklyExactPaidMinutes)
		o.WeeklyExactPaidMinutes = &WeeklyExactPaidMinutesInt
	}
	
	if WeeklyMinimumPaidMinutes, ok := WorkplanvalidationrequestMap["weeklyMinimumPaidMinutes"].(float64); ok {
		WeeklyMinimumPaidMinutesInt := int(WeeklyMinimumPaidMinutes)
		o.WeeklyMinimumPaidMinutes = &WeeklyMinimumPaidMinutesInt
	}
	
	if WeeklyMaximumPaidMinutes, ok := WorkplanvalidationrequestMap["weeklyMaximumPaidMinutes"].(float64); ok {
		WeeklyMaximumPaidMinutesInt := int(WeeklyMaximumPaidMinutes)
		o.WeeklyMaximumPaidMinutes = &WeeklyMaximumPaidMinutesInt
	}
	
	if ConstrainPaidTimeGranularity, ok := WorkplanvalidationrequestMap["constrainPaidTimeGranularity"].(bool); ok {
		o.ConstrainPaidTimeGranularity = &ConstrainPaidTimeGranularity
	}
    
	if PaidTimeGranularityMinutes, ok := WorkplanvalidationrequestMap["paidTimeGranularityMinutes"].(float64); ok {
		PaidTimeGranularityMinutesInt := int(PaidTimeGranularityMinutes)
		o.PaidTimeGranularityMinutes = &PaidTimeGranularityMinutesInt
	}
	
	if ConstrainMinimumTimeBetweenShifts, ok := WorkplanvalidationrequestMap["constrainMinimumTimeBetweenShifts"].(bool); ok {
		o.ConstrainMinimumTimeBetweenShifts = &ConstrainMinimumTimeBetweenShifts
	}
    
	if MinimumTimeBetweenShiftsMinutes, ok := WorkplanvalidationrequestMap["minimumTimeBetweenShiftsMinutes"].(float64); ok {
		MinimumTimeBetweenShiftsMinutesInt := int(MinimumTimeBetweenShiftsMinutes)
		o.MinimumTimeBetweenShiftsMinutes = &MinimumTimeBetweenShiftsMinutesInt
	}
	
	if MaximumDays, ok := WorkplanvalidationrequestMap["maximumDays"].(float64); ok {
		MaximumDaysInt := int(MaximumDays)
		o.MaximumDays = &MaximumDaysInt
	}
	
	if MinimumConsecutiveNonWorkingMinutesPerWeek, ok := WorkplanvalidationrequestMap["minimumConsecutiveNonWorkingMinutesPerWeek"].(float64); ok {
		MinimumConsecutiveNonWorkingMinutesPerWeekInt := int(MinimumConsecutiveNonWorkingMinutesPerWeek)
		o.MinimumConsecutiveNonWorkingMinutesPerWeek = &MinimumConsecutiveNonWorkingMinutesPerWeekInt
	}
	
	if ConstrainMaximumConsecutiveWorkingWeekends, ok := WorkplanvalidationrequestMap["constrainMaximumConsecutiveWorkingWeekends"].(bool); ok {
		o.ConstrainMaximumConsecutiveWorkingWeekends = &ConstrainMaximumConsecutiveWorkingWeekends
	}
    
	if MaximumConsecutiveWorkingWeekends, ok := WorkplanvalidationrequestMap["maximumConsecutiveWorkingWeekends"].(float64); ok {
		MaximumConsecutiveWorkingWeekendsInt := int(MaximumConsecutiveWorkingWeekends)
		o.MaximumConsecutiveWorkingWeekends = &MaximumConsecutiveWorkingWeekendsInt
	}
	
	if MinimumWorkingDaysPerWeek, ok := WorkplanvalidationrequestMap["minimumWorkingDaysPerWeek"].(float64); ok {
		MinimumWorkingDaysPerWeekInt := int(MinimumWorkingDaysPerWeek)
		o.MinimumWorkingDaysPerWeek = &MinimumWorkingDaysPerWeekInt
	}
	
	if ConstrainMaximumConsecutiveWorkingDays, ok := WorkplanvalidationrequestMap["constrainMaximumConsecutiveWorkingDays"].(bool); ok {
		o.ConstrainMaximumConsecutiveWorkingDays = &ConstrainMaximumConsecutiveWorkingDays
	}
    
	if MaximumConsecutiveWorkingDays, ok := WorkplanvalidationrequestMap["maximumConsecutiveWorkingDays"].(float64); ok {
		MaximumConsecutiveWorkingDaysInt := int(MaximumConsecutiveWorkingDays)
		o.MaximumConsecutiveWorkingDays = &MaximumConsecutiveWorkingDaysInt
	}
	
	if MinimumShiftStartDistanceMinutes, ok := WorkplanvalidationrequestMap["minimumShiftStartDistanceMinutes"].(float64); ok {
		MinimumShiftStartDistanceMinutesInt := int(MinimumShiftStartDistanceMinutes)
		o.MinimumShiftStartDistanceMinutes = &MinimumShiftStartDistanceMinutesInt
	}
	
	if MinimumDaysOffPerPlanningPeriod, ok := WorkplanvalidationrequestMap["minimumDaysOffPerPlanningPeriod"].(float64); ok {
		MinimumDaysOffPerPlanningPeriodInt := int(MinimumDaysOffPerPlanningPeriod)
		o.MinimumDaysOffPerPlanningPeriod = &MinimumDaysOffPerPlanningPeriodInt
	}
	
	if MaximumDaysOffPerPlanningPeriod, ok := WorkplanvalidationrequestMap["maximumDaysOffPerPlanningPeriod"].(float64); ok {
		MaximumDaysOffPerPlanningPeriodInt := int(MaximumDaysOffPerPlanningPeriod)
		o.MaximumDaysOffPerPlanningPeriod = &MaximumDaysOffPerPlanningPeriodInt
	}
	
	if MinimumPaidMinutesPerPlanningPeriod, ok := WorkplanvalidationrequestMap["minimumPaidMinutesPerPlanningPeriod"].(float64); ok {
		MinimumPaidMinutesPerPlanningPeriodInt := int(MinimumPaidMinutesPerPlanningPeriod)
		o.MinimumPaidMinutesPerPlanningPeriod = &MinimumPaidMinutesPerPlanningPeriodInt
	}
	
	if MaximumPaidMinutesPerPlanningPeriod, ok := WorkplanvalidationrequestMap["maximumPaidMinutesPerPlanningPeriod"].(float64); ok {
		MaximumPaidMinutesPerPlanningPeriodInt := int(MaximumPaidMinutesPerPlanningPeriod)
		o.MaximumPaidMinutesPerPlanningPeriod = &MaximumPaidMinutesPerPlanningPeriodInt
	}
	
	if ConstrainMaximumWorkingWeekendsPerPlanningPeriod, ok := WorkplanvalidationrequestMap["constrainMaximumWorkingWeekendsPerPlanningPeriod"].(bool); ok {
		o.ConstrainMaximumWorkingWeekendsPerPlanningPeriod = &ConstrainMaximumWorkingWeekendsPerPlanningPeriod
	}
    
	if MaximumWorkingWeekendsPerPlanningPeriod, ok := WorkplanvalidationrequestMap["maximumWorkingWeekendsPerPlanningPeriod"].(float64); ok {
		MaximumWorkingWeekendsPerPlanningPeriodInt := int(MaximumWorkingWeekendsPerPlanningPeriod)
		o.MaximumWorkingWeekendsPerPlanningPeriod = &MaximumWorkingWeekendsPerPlanningPeriodInt
	}
	
	if OptionalDays, ok := WorkplanvalidationrequestMap["optionalDays"].(map[string]interface{}); ok {
		OptionalDaysString, _ := json.Marshal(OptionalDays)
		json.Unmarshal(OptionalDaysString, &o.OptionalDays)
	}
	
	if ShiftStartVarianceType, ok := WorkplanvalidationrequestMap["shiftStartVarianceType"].(string); ok {
		o.ShiftStartVarianceType = &ShiftStartVarianceType
	}
    
	if ShiftStartVariancePeriod, ok := WorkplanvalidationrequestMap["shiftStartVariancePeriod"].(string); ok {
		o.ShiftStartVariancePeriod = &ShiftStartVariancePeriod
	}
    
	if ShiftStartVariances, ok := WorkplanvalidationrequestMap["shiftStartVariances"].(map[string]interface{}); ok {
		ShiftStartVariancesString, _ := json.Marshal(ShiftStartVariances)
		json.Unmarshal(ShiftStartVariancesString, &o.ShiftStartVariances)
	}
	
	if Shifts, ok := WorkplanvalidationrequestMap["shifts"].([]interface{}); ok {
		ShiftsString, _ := json.Marshal(Shifts)
		json.Unmarshal(ShiftsString, &o.Shifts)
	}
	
	if Agents, ok := WorkplanvalidationrequestMap["agents"].([]interface{}); ok {
		AgentsString, _ := json.Marshal(Agents)
		json.Unmarshal(AgentsString, &o.Agents)
	}
	
	if AgentCount, ok := WorkplanvalidationrequestMap["agentCount"].(float64); ok {
		AgentCountInt := int(AgentCount)
		o.AgentCount = &AgentCountInt
	}
	
	if SelfUri, ok := WorkplanvalidationrequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Workplanvalidationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
