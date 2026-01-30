package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanbiddingadminnotificationtopicworkplanbiddingnotification
type Workplanbiddingadminnotificationtopicworkplanbiddingnotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// BuId
	BuId *string `json:"buId,omitempty"`

	// Status
	Status *string `json:"status,omitempty"`

	// BidType
	BidType *string `json:"bidType,omitempty"`

	// BidWindowStartDate
	BidWindowStartDate *string `json:"bidWindowStartDate,omitempty"`

	// BidWindowEndDate
	BidWindowEndDate *string `json:"bidWindowEndDate,omitempty"`

	// EffectiveDate
	EffectiveDate *string `json:"effectiveDate,omitempty"`

	// AgentRankingType
	AgentRankingType *string `json:"agentRankingType,omitempty"`

	// RankingTiebreakerType
	RankingTiebreakerType *string `json:"rankingTiebreakerType,omitempty"`

	// WorkPlanFieldsVisibleToAgents
	WorkPlanFieldsVisibleToAgents *[]string `json:"workPlanFieldsVisibleToAgents,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workplanbiddingadminnotificationtopicworkplanbiddingnotification) SetField(field string, fieldValue interface{}) {
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

func (o Workplanbiddingadminnotificationtopicworkplanbiddingnotification) MarshalJSON() ([]byte, error) {
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
	type Alias Workplanbiddingadminnotificationtopicworkplanbiddingnotification
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		BuId *string `json:"buId,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		BidType *string `json:"bidType,omitempty"`
		
		BidWindowStartDate *string `json:"bidWindowStartDate,omitempty"`
		
		BidWindowEndDate *string `json:"bidWindowEndDate,omitempty"`
		
		EffectiveDate *string `json:"effectiveDate,omitempty"`
		
		AgentRankingType *string `json:"agentRankingType,omitempty"`
		
		RankingTiebreakerType *string `json:"rankingTiebreakerType,omitempty"`
		
		WorkPlanFieldsVisibleToAgents *[]string `json:"workPlanFieldsVisibleToAgents,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		BuId: o.BuId,
		
		Status: o.Status,
		
		BidType: o.BidType,
		
		BidWindowStartDate: o.BidWindowStartDate,
		
		BidWindowEndDate: o.BidWindowEndDate,
		
		EffectiveDate: o.EffectiveDate,
		
		AgentRankingType: o.AgentRankingType,
		
		RankingTiebreakerType: o.RankingTiebreakerType,
		
		WorkPlanFieldsVisibleToAgents: o.WorkPlanFieldsVisibleToAgents,
		Alias:    (Alias)(o),
	})
}

func (o *Workplanbiddingadminnotificationtopicworkplanbiddingnotification) UnmarshalJSON(b []byte) error {
	var WorkplanbiddingadminnotificationtopicworkplanbiddingnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WorkplanbiddingadminnotificationtopicworkplanbiddingnotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WorkplanbiddingadminnotificationtopicworkplanbiddingnotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if BuId, ok := WorkplanbiddingadminnotificationtopicworkplanbiddingnotificationMap["buId"].(string); ok {
		o.BuId = &BuId
	}
    
	if Status, ok := WorkplanbiddingadminnotificationtopicworkplanbiddingnotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if BidType, ok := WorkplanbiddingadminnotificationtopicworkplanbiddingnotificationMap["bidType"].(string); ok {
		o.BidType = &BidType
	}
    
	if BidWindowStartDate, ok := WorkplanbiddingadminnotificationtopicworkplanbiddingnotificationMap["bidWindowStartDate"].(string); ok {
		o.BidWindowStartDate = &BidWindowStartDate
	}
    
	if BidWindowEndDate, ok := WorkplanbiddingadminnotificationtopicworkplanbiddingnotificationMap["bidWindowEndDate"].(string); ok {
		o.BidWindowEndDate = &BidWindowEndDate
	}
    
	if EffectiveDate, ok := WorkplanbiddingadminnotificationtopicworkplanbiddingnotificationMap["effectiveDate"].(string); ok {
		o.EffectiveDate = &EffectiveDate
	}
    
	if AgentRankingType, ok := WorkplanbiddingadminnotificationtopicworkplanbiddingnotificationMap["agentRankingType"].(string); ok {
		o.AgentRankingType = &AgentRankingType
	}
    
	if RankingTiebreakerType, ok := WorkplanbiddingadminnotificationtopicworkplanbiddingnotificationMap["rankingTiebreakerType"].(string); ok {
		o.RankingTiebreakerType = &RankingTiebreakerType
	}
    
	if WorkPlanFieldsVisibleToAgents, ok := WorkplanbiddingadminnotificationtopicworkplanbiddingnotificationMap["workPlanFieldsVisibleToAgents"].([]interface{}); ok {
		WorkPlanFieldsVisibleToAgentsString, _ := json.Marshal(WorkPlanFieldsVisibleToAgents)
		json.Unmarshal(WorkPlanFieldsVisibleToAgentsString, &o.WorkPlanFieldsVisibleToAgents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workplanbiddingadminnotificationtopicworkplanbiddingnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
