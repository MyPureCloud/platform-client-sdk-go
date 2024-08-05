package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Updateworkplanbid - Update work plan bid model
type Updateworkplanbid struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the work plan bid
	Name *string `json:"name,omitempty"`

	// Forecast - The selected forecast in this work plan bid
	Forecast *Bushorttermforecastweekreference `json:"forecast,omitempty"`

	// BidWindowStartDate - The bid start date where agents start participate in work plan bidding in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	BidWindowStartDate *time.Time `json:"bidWindowStartDate,omitempty"`

	// BidWindowEndDate - The bid end date in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	BidWindowEndDate *time.Time `json:"bidWindowEndDate,omitempty"`

	// EffectiveDate - The date when agents will be assigned to the new work plan in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	EffectiveDate *time.Time `json:"effectiveDate,omitempty"`

	// AgentRankingType - The type of agent ranking selected for this bid
	AgentRankingType *string `json:"agentRankingType,omitempty"`

	// RankingTiebreakerType - Ranking tiebreaker
	RankingTiebreakerType *string `json:"rankingTiebreakerType,omitempty"`

	// WorkPlanFieldsVisibleToAgents - The work plan fields visible to agents whenever work plan preferences are made
	WorkPlanFieldsVisibleToAgents *Listwrapperagentworkplanfield `json:"workPlanFieldsVisibleToAgents,omitempty"`

	// Status - The state of the bid
	Status *string `json:"status,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Updateworkplanbid) SetField(field string, fieldValue interface{}) {
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

func (o Updateworkplanbid) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "BidWindowStartDate","BidWindowEndDate","EffectiveDate", }

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
	type Alias Updateworkplanbid
	
	BidWindowStartDate := new(string)
	if o.BidWindowStartDate != nil {
		*BidWindowStartDate = timeutil.Strftime(o.BidWindowStartDate, "%Y-%m-%d")
	} else {
		BidWindowStartDate = nil
	}
	
	BidWindowEndDate := new(string)
	if o.BidWindowEndDate != nil {
		*BidWindowEndDate = timeutil.Strftime(o.BidWindowEndDate, "%Y-%m-%d")
	} else {
		BidWindowEndDate = nil
	}
	
	EffectiveDate := new(string)
	if o.EffectiveDate != nil {
		*EffectiveDate = timeutil.Strftime(o.EffectiveDate, "%Y-%m-%d")
	} else {
		EffectiveDate = nil
	}
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Forecast *Bushorttermforecastweekreference `json:"forecast,omitempty"`
		
		BidWindowStartDate *string `json:"bidWindowStartDate,omitempty"`
		
		BidWindowEndDate *string `json:"bidWindowEndDate,omitempty"`
		
		EffectiveDate *string `json:"effectiveDate,omitempty"`
		
		AgentRankingType *string `json:"agentRankingType,omitempty"`
		
		RankingTiebreakerType *string `json:"rankingTiebreakerType,omitempty"`
		
		WorkPlanFieldsVisibleToAgents *Listwrapperagentworkplanfield `json:"workPlanFieldsVisibleToAgents,omitempty"`
		
		Status *string `json:"status,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		Forecast: o.Forecast,
		
		BidWindowStartDate: BidWindowStartDate,
		
		BidWindowEndDate: BidWindowEndDate,
		
		EffectiveDate: EffectiveDate,
		
		AgentRankingType: o.AgentRankingType,
		
		RankingTiebreakerType: o.RankingTiebreakerType,
		
		WorkPlanFieldsVisibleToAgents: o.WorkPlanFieldsVisibleToAgents,
		
		Status: o.Status,
		Alias:    (Alias)(o),
	})
}

func (o *Updateworkplanbid) UnmarshalJSON(b []byte) error {
	var UpdateworkplanbidMap map[string]interface{}
	err := json.Unmarshal(b, &UpdateworkplanbidMap)
	if err != nil {
		return err
	}
	
	if Name, ok := UpdateworkplanbidMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Forecast, ok := UpdateworkplanbidMap["forecast"].(map[string]interface{}); ok {
		ForecastString, _ := json.Marshal(Forecast)
		json.Unmarshal(ForecastString, &o.Forecast)
	}
	
	if bidWindowStartDateString, ok := UpdateworkplanbidMap["bidWindowStartDate"].(string); ok {
		BidWindowStartDate, _ := time.Parse("2006-01-02", bidWindowStartDateString)
		o.BidWindowStartDate = &BidWindowStartDate
	}
	
	if bidWindowEndDateString, ok := UpdateworkplanbidMap["bidWindowEndDate"].(string); ok {
		BidWindowEndDate, _ := time.Parse("2006-01-02", bidWindowEndDateString)
		o.BidWindowEndDate = &BidWindowEndDate
	}
	
	if effectiveDateString, ok := UpdateworkplanbidMap["effectiveDate"].(string); ok {
		EffectiveDate, _ := time.Parse("2006-01-02", effectiveDateString)
		o.EffectiveDate = &EffectiveDate
	}
	
	if AgentRankingType, ok := UpdateworkplanbidMap["agentRankingType"].(string); ok {
		o.AgentRankingType = &AgentRankingType
	}
    
	if RankingTiebreakerType, ok := UpdateworkplanbidMap["rankingTiebreakerType"].(string); ok {
		o.RankingTiebreakerType = &RankingTiebreakerType
	}
    
	if WorkPlanFieldsVisibleToAgents, ok := UpdateworkplanbidMap["workPlanFieldsVisibleToAgents"].(map[string]interface{}); ok {
		WorkPlanFieldsVisibleToAgentsString, _ := json.Marshal(WorkPlanFieldsVisibleToAgents)
		json.Unmarshal(WorkPlanFieldsVisibleToAgentsString, &o.WorkPlanFieldsVisibleToAgents)
	}
	
	if Status, ok := UpdateworkplanbidMap["status"].(string); ok {
		o.Status = &Status
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Updateworkplanbid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
