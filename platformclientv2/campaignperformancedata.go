package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Campaignperformancedata
type Campaignperformancedata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Campaign - Identifier of the campaign
	Campaign *Domainentityref `json:"campaign,omitempty"`

	// Division - The division the campaign belongs to
	Division *Addressableentityref `json:"division,omitempty"`

	// ContactRate - Information regarding the campaign's connect rate
	ContactRate *Campaignperformancedatacontactrate `json:"contactRate,omitempty"`

	// IdleAgents - Number of available agents not currently being utilized
	IdleAgents *int `json:"idleAgents,omitempty"`

	// EffectiveIdleAgents - Number of effective available agents not currently being utilized
	EffectiveIdleAgents *float32 `json:"effectiveIdleAgents,omitempty"`

	// AdjustedCallsPerAgent - Calls per agent adjusted by pace
	AdjustedCallsPerAgent *float32 `json:"adjustedCallsPerAgent,omitempty"`

	// OutstandingCalls - Number of campaign calls currently ongoing
	OutstandingCalls *int `json:"outstandingCalls,omitempty"`

	// ScheduledCalls - Number of campaign calls currently scheduled
	ScheduledCalls *int `json:"scheduledCalls,omitempty"`

	// RightPartyContactsCount - Information on the campaign's number of Right Party Contacts
	RightPartyContactsCount *int `json:"rightPartyContactsCount,omitempty"`

	// CampaignStatus - The status of the campaign
	CampaignStatus *string `json:"campaignStatus,omitempty"`

	// DialingMode - The strategy this Campaign will use for dialing
	DialingMode *string `json:"dialingMode,omitempty"`

	// Progress - Information on the campaign's progress
	Progress *Campaignperformancedataprogress `json:"progress,omitempty"`

	// LinesUtilization - Information on the campaign's lines utilization
	LinesUtilization *Campaignlinesutilization `json:"linesUtilization,omitempty"`

	// BusinessCategoryMetrics - Information on the campaign's business category metrics
	BusinessCategoryMetrics *Campaignbusinesscategorymetrics `json:"businessCategoryMetrics,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Campaignperformancedata) SetField(field string, fieldValue interface{}) {
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

func (o Campaignperformancedata) MarshalJSON() ([]byte, error) {
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
	type Alias Campaignperformancedata
	
	return json.Marshal(&struct { 
		Campaign *Domainentityref `json:"campaign,omitempty"`
		
		Division *Addressableentityref `json:"division,omitempty"`
		
		ContactRate *Campaignperformancedatacontactrate `json:"contactRate,omitempty"`
		
		IdleAgents *int `json:"idleAgents,omitempty"`
		
		EffectiveIdleAgents *float32 `json:"effectiveIdleAgents,omitempty"`
		
		AdjustedCallsPerAgent *float32 `json:"adjustedCallsPerAgent,omitempty"`
		
		OutstandingCalls *int `json:"outstandingCalls,omitempty"`
		
		ScheduledCalls *int `json:"scheduledCalls,omitempty"`
		
		RightPartyContactsCount *int `json:"rightPartyContactsCount,omitempty"`
		
		CampaignStatus *string `json:"campaignStatus,omitempty"`
		
		DialingMode *string `json:"dialingMode,omitempty"`
		
		Progress *Campaignperformancedataprogress `json:"progress,omitempty"`
		
		LinesUtilization *Campaignlinesutilization `json:"linesUtilization,omitempty"`
		
		BusinessCategoryMetrics *Campaignbusinesscategorymetrics `json:"businessCategoryMetrics,omitempty"`
		Alias
	}{ 
		Campaign: o.Campaign,
		
		Division: o.Division,
		
		ContactRate: o.ContactRate,
		
		IdleAgents: o.IdleAgents,
		
		EffectiveIdleAgents: o.EffectiveIdleAgents,
		
		AdjustedCallsPerAgent: o.AdjustedCallsPerAgent,
		
		OutstandingCalls: o.OutstandingCalls,
		
		ScheduledCalls: o.ScheduledCalls,
		
		RightPartyContactsCount: o.RightPartyContactsCount,
		
		CampaignStatus: o.CampaignStatus,
		
		DialingMode: o.DialingMode,
		
		Progress: o.Progress,
		
		LinesUtilization: o.LinesUtilization,
		
		BusinessCategoryMetrics: o.BusinessCategoryMetrics,
		Alias:    (Alias)(o),
	})
}

func (o *Campaignperformancedata) UnmarshalJSON(b []byte) error {
	var CampaignperformancedataMap map[string]interface{}
	err := json.Unmarshal(b, &CampaignperformancedataMap)
	if err != nil {
		return err
	}
	
	if Campaign, ok := CampaignperformancedataMap["campaign"].(map[string]interface{}); ok {
		CampaignString, _ := json.Marshal(Campaign)
		json.Unmarshal(CampaignString, &o.Campaign)
	}
	
	if Division, ok := CampaignperformancedataMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if ContactRate, ok := CampaignperformancedataMap["contactRate"].(map[string]interface{}); ok {
		ContactRateString, _ := json.Marshal(ContactRate)
		json.Unmarshal(ContactRateString, &o.ContactRate)
	}
	
	if IdleAgents, ok := CampaignperformancedataMap["idleAgents"].(float64); ok {
		IdleAgentsInt := int(IdleAgents)
		o.IdleAgents = &IdleAgentsInt
	}
	
	if EffectiveIdleAgents, ok := CampaignperformancedataMap["effectiveIdleAgents"].(float64); ok {
		EffectiveIdleAgentsFloat32 := float32(EffectiveIdleAgents)
		o.EffectiveIdleAgents = &EffectiveIdleAgentsFloat32
	}
    
	if AdjustedCallsPerAgent, ok := CampaignperformancedataMap["adjustedCallsPerAgent"].(float64); ok {
		AdjustedCallsPerAgentFloat32 := float32(AdjustedCallsPerAgent)
		o.AdjustedCallsPerAgent = &AdjustedCallsPerAgentFloat32
	}
    
	if OutstandingCalls, ok := CampaignperformancedataMap["outstandingCalls"].(float64); ok {
		OutstandingCallsInt := int(OutstandingCalls)
		o.OutstandingCalls = &OutstandingCallsInt
	}
	
	if ScheduledCalls, ok := CampaignperformancedataMap["scheduledCalls"].(float64); ok {
		ScheduledCallsInt := int(ScheduledCalls)
		o.ScheduledCalls = &ScheduledCallsInt
	}
	
	if RightPartyContactsCount, ok := CampaignperformancedataMap["rightPartyContactsCount"].(float64); ok {
		RightPartyContactsCountInt := int(RightPartyContactsCount)
		o.RightPartyContactsCount = &RightPartyContactsCountInt
	}
	
	if CampaignStatus, ok := CampaignperformancedataMap["campaignStatus"].(string); ok {
		o.CampaignStatus = &CampaignStatus
	}
    
	if DialingMode, ok := CampaignperformancedataMap["dialingMode"].(string); ok {
		o.DialingMode = &DialingMode
	}
    
	if Progress, ok := CampaignperformancedataMap["progress"].(map[string]interface{}); ok {
		ProgressString, _ := json.Marshal(Progress)
		json.Unmarshal(ProgressString, &o.Progress)
	}
	
	if LinesUtilization, ok := CampaignperformancedataMap["linesUtilization"].(map[string]interface{}); ok {
		LinesUtilizationString, _ := json.Marshal(LinesUtilization)
		json.Unmarshal(LinesUtilizationString, &o.LinesUtilization)
	}
	
	if BusinessCategoryMetrics, ok := CampaignperformancedataMap["businessCategoryMetrics"].(map[string]interface{}); ok {
		BusinessCategoryMetricsString, _ := json.Marshal(BusinessCategoryMetrics)
		json.Unmarshal(BusinessCategoryMetricsString, &o.BusinessCategoryMetrics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Campaignperformancedata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
