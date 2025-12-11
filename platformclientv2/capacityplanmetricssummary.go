package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capacityplanmetricssummary
type Capacityplanmetricssummary struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// RequiredStaffFullTimeEquivalentCount - The total staff requirements for all planning groups in the capacity plan, aggregated by the selected time granularity
	RequiredStaffFullTimeEquivalentCount *[]float64 `json:"requiredStaffFullTimeEquivalentCount,omitempty"`

	// PlannedFullTimeEquivalentCount - The planned full time equivalent for all staffing groups in the capacity plan, aggregated by the selected time granularity
	PlannedFullTimeEquivalentCount *[]float64 `json:"plannedFullTimeEquivalentCount,omitempty"`

	// StaffingOverUnderFullTimeEquivalentCount - The difference between the summary of planning group required full time equivalent and planned full time equivalent, aggregated by the selected time granularity
	StaffingOverUnderFullTimeEquivalentCount *[]float64 `json:"staffingOverUnderFullTimeEquivalentCount,omitempty"`

	// StartingFullTimeEquivalentCount - The total starting full time equivalent for all staffing groups in the capacity plan, aggregated by the selected time granularity
	StartingFullTimeEquivalentCount *[]float64 `json:"startingFullTimeEquivalentCount,omitempty"`

	// AttritionFullTimeEquivalentCount - The sum of all staffing group attrition full time equivalent in the capacity plan, aggregated by the selected time granularity
	AttritionFullTimeEquivalentCount *[]float64 `json:"attritionFullTimeEquivalentCount,omitempty"`

	// AttritionPercentage - The total attrition percentage of staffing groups in the capacity plan in the scale of 0.0-100.0, aggregated by the selected time granularity
	AttritionPercentage *[]float64 `json:"attritionPercentage,omitempty"`

	// NewHireFullTimeEquivalentCount - The sum of all staffing group new hire full time equivalent in the capacity plan, aggregated by the selected time granularity
	NewHireFullTimeEquivalentCount *[]float64 `json:"newHireFullTimeEquivalentCount,omitempty"`

	// TransfersFullTimeEquivalentCount - The sum of all staffing group projected transfers of agents into or out of this capacity plan, aggregated by the selected time granularity
	TransfersFullTimeEquivalentCount *[]float64 `json:"transfersFullTimeEquivalentCount,omitempty"`

	// ExtraTimeUnderTimeFullTimeEquivalentCount - The sum of all staffing group extra or under time full time equivalent count in the capacity plan, aggregated by the selected time granularity
	ExtraTimeUnderTimeFullTimeEquivalentCount *[]float64 `json:"extraTimeUnderTimeFullTimeEquivalentCount,omitempty"`

	// ShrinkageFullTimeEquivalentCount - The sum of all staffing groups shrinkage full time equivalent, aggregated by the selected time granularity
	ShrinkageFullTimeEquivalentCount *[]float64 `json:"shrinkageFullTimeEquivalentCount,omitempty"`

	// ShrinkagePercentage - The total shrinkage percentage of all staffing groups in the capacity plan in the scale of 0.0-100.0, aggregated by the selected time granularity
	ShrinkagePercentage *[]float64 `json:"shrinkagePercentage,omitempty"`

	// EndOfMonthPlannedFullTimeEquivalentCount - The total sum of all staffing group end of month planned full time equivalent for this capacity plan, aggregated by the selected time granularity
	EndOfMonthPlannedFullTimeEquivalentCount *[]float64 `json:"endOfMonthPlannedFullTimeEquivalentCount,omitempty"`

	// NetFullTimeEquivalentCount - The sum of all staffing groups net full time equivalent in the capacity plan, aggregated by the selected time granularity
	NetFullTimeEquivalentCount *[]float64 `json:"netFullTimeEquivalentCount,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capacityplanmetricssummary) SetField(field string, fieldValue interface{}) {
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

func (o Capacityplanmetricssummary) MarshalJSON() ([]byte, error) {
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
	type Alias Capacityplanmetricssummary
	
	return json.Marshal(&struct { 
		RequiredStaffFullTimeEquivalentCount *[]float64 `json:"requiredStaffFullTimeEquivalentCount,omitempty"`
		
		PlannedFullTimeEquivalentCount *[]float64 `json:"plannedFullTimeEquivalentCount,omitempty"`
		
		StaffingOverUnderFullTimeEquivalentCount *[]float64 `json:"staffingOverUnderFullTimeEquivalentCount,omitempty"`
		
		StartingFullTimeEquivalentCount *[]float64 `json:"startingFullTimeEquivalentCount,omitempty"`
		
		AttritionFullTimeEquivalentCount *[]float64 `json:"attritionFullTimeEquivalentCount,omitempty"`
		
		AttritionPercentage *[]float64 `json:"attritionPercentage,omitempty"`
		
		NewHireFullTimeEquivalentCount *[]float64 `json:"newHireFullTimeEquivalentCount,omitempty"`
		
		TransfersFullTimeEquivalentCount *[]float64 `json:"transfersFullTimeEquivalentCount,omitempty"`
		
		ExtraTimeUnderTimeFullTimeEquivalentCount *[]float64 `json:"extraTimeUnderTimeFullTimeEquivalentCount,omitempty"`
		
		ShrinkageFullTimeEquivalentCount *[]float64 `json:"shrinkageFullTimeEquivalentCount,omitempty"`
		
		ShrinkagePercentage *[]float64 `json:"shrinkagePercentage,omitempty"`
		
		EndOfMonthPlannedFullTimeEquivalentCount *[]float64 `json:"endOfMonthPlannedFullTimeEquivalentCount,omitempty"`
		
		NetFullTimeEquivalentCount *[]float64 `json:"netFullTimeEquivalentCount,omitempty"`
		Alias
	}{ 
		RequiredStaffFullTimeEquivalentCount: o.RequiredStaffFullTimeEquivalentCount,
		
		PlannedFullTimeEquivalentCount: o.PlannedFullTimeEquivalentCount,
		
		StaffingOverUnderFullTimeEquivalentCount: o.StaffingOverUnderFullTimeEquivalentCount,
		
		StartingFullTimeEquivalentCount: o.StartingFullTimeEquivalentCount,
		
		AttritionFullTimeEquivalentCount: o.AttritionFullTimeEquivalentCount,
		
		AttritionPercentage: o.AttritionPercentage,
		
		NewHireFullTimeEquivalentCount: o.NewHireFullTimeEquivalentCount,
		
		TransfersFullTimeEquivalentCount: o.TransfersFullTimeEquivalentCount,
		
		ExtraTimeUnderTimeFullTimeEquivalentCount: o.ExtraTimeUnderTimeFullTimeEquivalentCount,
		
		ShrinkageFullTimeEquivalentCount: o.ShrinkageFullTimeEquivalentCount,
		
		ShrinkagePercentage: o.ShrinkagePercentage,
		
		EndOfMonthPlannedFullTimeEquivalentCount: o.EndOfMonthPlannedFullTimeEquivalentCount,
		
		NetFullTimeEquivalentCount: o.NetFullTimeEquivalentCount,
		Alias:    (Alias)(o),
	})
}

func (o *Capacityplanmetricssummary) UnmarshalJSON(b []byte) error {
	var CapacityplanmetricssummaryMap map[string]interface{}
	err := json.Unmarshal(b, &CapacityplanmetricssummaryMap)
	if err != nil {
		return err
	}
	
	if RequiredStaffFullTimeEquivalentCount, ok := CapacityplanmetricssummaryMap["requiredStaffFullTimeEquivalentCount"].([]interface{}); ok {
		RequiredStaffFullTimeEquivalentCountString, _ := json.Marshal(RequiredStaffFullTimeEquivalentCount)
		json.Unmarshal(RequiredStaffFullTimeEquivalentCountString, &o.RequiredStaffFullTimeEquivalentCount)
	}
	
	if PlannedFullTimeEquivalentCount, ok := CapacityplanmetricssummaryMap["plannedFullTimeEquivalentCount"].([]interface{}); ok {
		PlannedFullTimeEquivalentCountString, _ := json.Marshal(PlannedFullTimeEquivalentCount)
		json.Unmarshal(PlannedFullTimeEquivalentCountString, &o.PlannedFullTimeEquivalentCount)
	}
	
	if StaffingOverUnderFullTimeEquivalentCount, ok := CapacityplanmetricssummaryMap["staffingOverUnderFullTimeEquivalentCount"].([]interface{}); ok {
		StaffingOverUnderFullTimeEquivalentCountString, _ := json.Marshal(StaffingOverUnderFullTimeEquivalentCount)
		json.Unmarshal(StaffingOverUnderFullTimeEquivalentCountString, &o.StaffingOverUnderFullTimeEquivalentCount)
	}
	
	if StartingFullTimeEquivalentCount, ok := CapacityplanmetricssummaryMap["startingFullTimeEquivalentCount"].([]interface{}); ok {
		StartingFullTimeEquivalentCountString, _ := json.Marshal(StartingFullTimeEquivalentCount)
		json.Unmarshal(StartingFullTimeEquivalentCountString, &o.StartingFullTimeEquivalentCount)
	}
	
	if AttritionFullTimeEquivalentCount, ok := CapacityplanmetricssummaryMap["attritionFullTimeEquivalentCount"].([]interface{}); ok {
		AttritionFullTimeEquivalentCountString, _ := json.Marshal(AttritionFullTimeEquivalentCount)
		json.Unmarshal(AttritionFullTimeEquivalentCountString, &o.AttritionFullTimeEquivalentCount)
	}
	
	if AttritionPercentage, ok := CapacityplanmetricssummaryMap["attritionPercentage"].([]interface{}); ok {
		AttritionPercentageString, _ := json.Marshal(AttritionPercentage)
		json.Unmarshal(AttritionPercentageString, &o.AttritionPercentage)
	}
	
	if NewHireFullTimeEquivalentCount, ok := CapacityplanmetricssummaryMap["newHireFullTimeEquivalentCount"].([]interface{}); ok {
		NewHireFullTimeEquivalentCountString, _ := json.Marshal(NewHireFullTimeEquivalentCount)
		json.Unmarshal(NewHireFullTimeEquivalentCountString, &o.NewHireFullTimeEquivalentCount)
	}
	
	if TransfersFullTimeEquivalentCount, ok := CapacityplanmetricssummaryMap["transfersFullTimeEquivalentCount"].([]interface{}); ok {
		TransfersFullTimeEquivalentCountString, _ := json.Marshal(TransfersFullTimeEquivalentCount)
		json.Unmarshal(TransfersFullTimeEquivalentCountString, &o.TransfersFullTimeEquivalentCount)
	}
	
	if ExtraTimeUnderTimeFullTimeEquivalentCount, ok := CapacityplanmetricssummaryMap["extraTimeUnderTimeFullTimeEquivalentCount"].([]interface{}); ok {
		ExtraTimeUnderTimeFullTimeEquivalentCountString, _ := json.Marshal(ExtraTimeUnderTimeFullTimeEquivalentCount)
		json.Unmarshal(ExtraTimeUnderTimeFullTimeEquivalentCountString, &o.ExtraTimeUnderTimeFullTimeEquivalentCount)
	}
	
	if ShrinkageFullTimeEquivalentCount, ok := CapacityplanmetricssummaryMap["shrinkageFullTimeEquivalentCount"].([]interface{}); ok {
		ShrinkageFullTimeEquivalentCountString, _ := json.Marshal(ShrinkageFullTimeEquivalentCount)
		json.Unmarshal(ShrinkageFullTimeEquivalentCountString, &o.ShrinkageFullTimeEquivalentCount)
	}
	
	if ShrinkagePercentage, ok := CapacityplanmetricssummaryMap["shrinkagePercentage"].([]interface{}); ok {
		ShrinkagePercentageString, _ := json.Marshal(ShrinkagePercentage)
		json.Unmarshal(ShrinkagePercentageString, &o.ShrinkagePercentage)
	}
	
	if EndOfMonthPlannedFullTimeEquivalentCount, ok := CapacityplanmetricssummaryMap["endOfMonthPlannedFullTimeEquivalentCount"].([]interface{}); ok {
		EndOfMonthPlannedFullTimeEquivalentCountString, _ := json.Marshal(EndOfMonthPlannedFullTimeEquivalentCount)
		json.Unmarshal(EndOfMonthPlannedFullTimeEquivalentCountString, &o.EndOfMonthPlannedFullTimeEquivalentCount)
	}
	
	if NetFullTimeEquivalentCount, ok := CapacityplanmetricssummaryMap["netFullTimeEquivalentCount"].([]interface{}); ok {
		NetFullTimeEquivalentCountString, _ := json.Marshal(NetFullTimeEquivalentCount)
		json.Unmarshal(NetFullTimeEquivalentCountString, &o.NetFullTimeEquivalentCount)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Capacityplanmetricssummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
