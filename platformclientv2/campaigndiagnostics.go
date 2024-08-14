package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Campaigndiagnostics
type Campaigndiagnostics struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CallableContacts - Campaign properties that can impact which contacts are callable
	CallableContacts *Callablecontactsdiagnostic `json:"callableContacts,omitempty"`

	// QueueUtilizationDiagnostic - Information regarding the campaign's queue
	QueueUtilizationDiagnostic *Queueutilizationdiagnostic `json:"queueUtilizationDiagnostic,omitempty"`

	// RuleSetDiagnostics - Information regarding the campaign's rule sets
	RuleSetDiagnostics *[]Rulesetdiagnostic `json:"ruleSetDiagnostics,omitempty"`

	// OutstandingInteractionsCount - Current number of outstanding interactions on the campaign
	OutstandingInteractionsCount *int `json:"outstandingInteractionsCount,omitempty"`

	// ScheduledInteractionsCount - Current number of scheduled interactions on the campaign
	ScheduledInteractionsCount *int `json:"scheduledInteractionsCount,omitempty"`

	// TimeZoneRescheduledCallsCount - Current number of time zone rescheduled calls on the campaign
	TimeZoneRescheduledCallsCount *int `json:"timeZoneRescheduledCallsCount,omitempty"`

	// FilteredOutContactsCount - Number of contacts that don't match filter. This is currently supported only for Campaigns with dynamic filter on.
	FilteredOutContactsCount *int `json:"filteredOutContactsCount,omitempty"`

	// CampaignSkillStatistics - Information regarding the campaign's skills
	CampaignSkillStatistics *Campaignskillstatistics `json:"campaignSkillStatistics,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Campaigndiagnostics) SetField(field string, fieldValue interface{}) {
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

func (o Campaigndiagnostics) MarshalJSON() ([]byte, error) {
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
	type Alias Campaigndiagnostics
	
	return json.Marshal(&struct { 
		CallableContacts *Callablecontactsdiagnostic `json:"callableContacts,omitempty"`
		
		QueueUtilizationDiagnostic *Queueutilizationdiagnostic `json:"queueUtilizationDiagnostic,omitempty"`
		
		RuleSetDiagnostics *[]Rulesetdiagnostic `json:"ruleSetDiagnostics,omitempty"`
		
		OutstandingInteractionsCount *int `json:"outstandingInteractionsCount,omitempty"`
		
		ScheduledInteractionsCount *int `json:"scheduledInteractionsCount,omitempty"`
		
		TimeZoneRescheduledCallsCount *int `json:"timeZoneRescheduledCallsCount,omitempty"`
		
		FilteredOutContactsCount *int `json:"filteredOutContactsCount,omitempty"`
		
		CampaignSkillStatistics *Campaignskillstatistics `json:"campaignSkillStatistics,omitempty"`
		Alias
	}{ 
		CallableContacts: o.CallableContacts,
		
		QueueUtilizationDiagnostic: o.QueueUtilizationDiagnostic,
		
		RuleSetDiagnostics: o.RuleSetDiagnostics,
		
		OutstandingInteractionsCount: o.OutstandingInteractionsCount,
		
		ScheduledInteractionsCount: o.ScheduledInteractionsCount,
		
		TimeZoneRescheduledCallsCount: o.TimeZoneRescheduledCallsCount,
		
		FilteredOutContactsCount: o.FilteredOutContactsCount,
		
		CampaignSkillStatistics: o.CampaignSkillStatistics,
		Alias:    (Alias)(o),
	})
}

func (o *Campaigndiagnostics) UnmarshalJSON(b []byte) error {
	var CampaigndiagnosticsMap map[string]interface{}
	err := json.Unmarshal(b, &CampaigndiagnosticsMap)
	if err != nil {
		return err
	}
	
	if CallableContacts, ok := CampaigndiagnosticsMap["callableContacts"].(map[string]interface{}); ok {
		CallableContactsString, _ := json.Marshal(CallableContacts)
		json.Unmarshal(CallableContactsString, &o.CallableContacts)
	}
	
	if QueueUtilizationDiagnostic, ok := CampaigndiagnosticsMap["queueUtilizationDiagnostic"].(map[string]interface{}); ok {
		QueueUtilizationDiagnosticString, _ := json.Marshal(QueueUtilizationDiagnostic)
		json.Unmarshal(QueueUtilizationDiagnosticString, &o.QueueUtilizationDiagnostic)
	}
	
	if RuleSetDiagnostics, ok := CampaigndiagnosticsMap["ruleSetDiagnostics"].([]interface{}); ok {
		RuleSetDiagnosticsString, _ := json.Marshal(RuleSetDiagnostics)
		json.Unmarshal(RuleSetDiagnosticsString, &o.RuleSetDiagnostics)
	}
	
	if OutstandingInteractionsCount, ok := CampaigndiagnosticsMap["outstandingInteractionsCount"].(float64); ok {
		OutstandingInteractionsCountInt := int(OutstandingInteractionsCount)
		o.OutstandingInteractionsCount = &OutstandingInteractionsCountInt
	}
	
	if ScheduledInteractionsCount, ok := CampaigndiagnosticsMap["scheduledInteractionsCount"].(float64); ok {
		ScheduledInteractionsCountInt := int(ScheduledInteractionsCount)
		o.ScheduledInteractionsCount = &ScheduledInteractionsCountInt
	}
	
	if TimeZoneRescheduledCallsCount, ok := CampaigndiagnosticsMap["timeZoneRescheduledCallsCount"].(float64); ok {
		TimeZoneRescheduledCallsCountInt := int(TimeZoneRescheduledCallsCount)
		o.TimeZoneRescheduledCallsCount = &TimeZoneRescheduledCallsCountInt
	}
	
	if FilteredOutContactsCount, ok := CampaigndiagnosticsMap["filteredOutContactsCount"].(float64); ok {
		FilteredOutContactsCountInt := int(FilteredOutContactsCount)
		o.FilteredOutContactsCount = &FilteredOutContactsCountInt
	}
	
	if CampaignSkillStatistics, ok := CampaigndiagnosticsMap["campaignSkillStatistics"].(map[string]interface{}); ok {
		CampaignSkillStatisticsString, _ := json.Marshal(CampaignSkillStatistics)
		json.Unmarshal(CampaignSkillStatisticsString, &o.CampaignSkillStatistics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Campaigndiagnostics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
