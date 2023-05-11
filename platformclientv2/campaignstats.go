package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Campaignstats
type Campaignstats struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ContactRate - Information regarding the campaign's connect rate
	ContactRate *Connectrate `json:"contactRate,omitempty"`

	// IdleAgents - Number of available agents not currently being utilized
	IdleAgents *int `json:"idleAgents,omitempty"`

	// EffectiveIdleAgents - Number of effective available agents not currently being utilized
	EffectiveIdleAgents *float64 `json:"effectiveIdleAgents,omitempty"`

	// AdjustedCallsPerAgent - Calls per agent adjusted by pace
	AdjustedCallsPerAgent *float64 `json:"adjustedCallsPerAgent,omitempty"`

	// OutstandingCalls - Number of campaign calls currently ongoing
	OutstandingCalls *int `json:"outstandingCalls,omitempty"`

	// ScheduledCalls - Number of campaign calls currently scheduled
	ScheduledCalls *int `json:"scheduledCalls,omitempty"`

	// TimeZoneRescheduledCalls - Number of campaign calls currently timezone rescheduled
	TimeZoneRescheduledCalls *int `json:"timeZoneRescheduledCalls,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Campaignstats) SetField(field string, fieldValue interface{}) {
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

func (o Campaignstats) MarshalJSON() ([]byte, error) {
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
	type Alias Campaignstats
	
	return json.Marshal(&struct { 
		ContactRate *Connectrate `json:"contactRate,omitempty"`
		
		IdleAgents *int `json:"idleAgents,omitempty"`
		
		EffectiveIdleAgents *float64 `json:"effectiveIdleAgents,omitempty"`
		
		AdjustedCallsPerAgent *float64 `json:"adjustedCallsPerAgent,omitempty"`
		
		OutstandingCalls *int `json:"outstandingCalls,omitempty"`
		
		ScheduledCalls *int `json:"scheduledCalls,omitempty"`
		
		TimeZoneRescheduledCalls *int `json:"timeZoneRescheduledCalls,omitempty"`
		Alias
	}{ 
		ContactRate: o.ContactRate,
		
		IdleAgents: o.IdleAgents,
		
		EffectiveIdleAgents: o.EffectiveIdleAgents,
		
		AdjustedCallsPerAgent: o.AdjustedCallsPerAgent,
		
		OutstandingCalls: o.OutstandingCalls,
		
		ScheduledCalls: o.ScheduledCalls,
		
		TimeZoneRescheduledCalls: o.TimeZoneRescheduledCalls,
		Alias:    (Alias)(o),
	})
}

func (o *Campaignstats) UnmarshalJSON(b []byte) error {
	var CampaignstatsMap map[string]interface{}
	err := json.Unmarshal(b, &CampaignstatsMap)
	if err != nil {
		return err
	}
	
	if ContactRate, ok := CampaignstatsMap["contactRate"].(map[string]interface{}); ok {
		ContactRateString, _ := json.Marshal(ContactRate)
		json.Unmarshal(ContactRateString, &o.ContactRate)
	}
	
	if IdleAgents, ok := CampaignstatsMap["idleAgents"].(float64); ok {
		IdleAgentsInt := int(IdleAgents)
		o.IdleAgents = &IdleAgentsInt
	}
	
	if EffectiveIdleAgents, ok := CampaignstatsMap["effectiveIdleAgents"].(float64); ok {
		o.EffectiveIdleAgents = &EffectiveIdleAgents
	}
    
	if AdjustedCallsPerAgent, ok := CampaignstatsMap["adjustedCallsPerAgent"].(float64); ok {
		o.AdjustedCallsPerAgent = &AdjustedCallsPerAgent
	}
    
	if OutstandingCalls, ok := CampaignstatsMap["outstandingCalls"].(float64); ok {
		OutstandingCallsInt := int(OutstandingCalls)
		o.OutstandingCalls = &OutstandingCallsInt
	}
	
	if ScheduledCalls, ok := CampaignstatsMap["scheduledCalls"].(float64); ok {
		ScheduledCallsInt := int(ScheduledCalls)
		o.ScheduledCalls = &ScheduledCallsInt
	}
	
	if TimeZoneRescheduledCalls, ok := CampaignstatsMap["timeZoneRescheduledCalls"].(float64); ok {
		TimeZoneRescheduledCallsInt := int(TimeZoneRescheduledCalls)
		o.TimeZoneRescheduledCalls = &TimeZoneRescheduledCallsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Campaignstats) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
