package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceexceptioninfo
type Wfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceexceptioninfo struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// StartOffsetSeconds
	StartOffsetSeconds *int `json:"startOffsetSeconds,omitempty"`

	// EndOffsetSeconds
	EndOffsetSeconds *int `json:"endOffsetSeconds,omitempty"`

	// ScheduledActivityCodeId
	ScheduledActivityCodeId *string `json:"scheduledActivityCodeId,omitempty"`

	// ScheduledActivityCategory
	ScheduledActivityCategory *string `json:"scheduledActivityCategory,omitempty"`

	// ScheduledSecondaryPresenceLookupIds
	ScheduledSecondaryPresenceLookupIds *[]string `json:"scheduledSecondaryPresenceLookupIds,omitempty"`

	// ActualActivityCodeId
	ActualActivityCodeId *string `json:"actualActivityCodeId,omitempty"`

	// ActualActivityCategory
	ActualActivityCategory *string `json:"actualActivityCategory,omitempty"`

	// SystemPresence
	SystemPresence *string `json:"systemPresence,omitempty"`

	// RoutingStatus
	RoutingStatus *string `json:"routingStatus,omitempty"`

	// Impact
	Impact *string `json:"impact,omitempty"`

	// SecondaryPresenceLookupId
	SecondaryPresenceLookupId *string `json:"secondaryPresenceLookupId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceexceptioninfo) SetField(field string, fieldValue interface{}) {
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

func (o Wfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceexceptioninfo) MarshalJSON() ([]byte, error) {
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
	type Alias Wfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceexceptioninfo
	
	return json.Marshal(&struct { 
		StartOffsetSeconds *int `json:"startOffsetSeconds,omitempty"`
		
		EndOffsetSeconds *int `json:"endOffsetSeconds,omitempty"`
		
		ScheduledActivityCodeId *string `json:"scheduledActivityCodeId,omitempty"`
		
		ScheduledActivityCategory *string `json:"scheduledActivityCategory,omitempty"`
		
		ScheduledSecondaryPresenceLookupIds *[]string `json:"scheduledSecondaryPresenceLookupIds,omitempty"`
		
		ActualActivityCodeId *string `json:"actualActivityCodeId,omitempty"`
		
		ActualActivityCategory *string `json:"actualActivityCategory,omitempty"`
		
		SystemPresence *string `json:"systemPresence,omitempty"`
		
		RoutingStatus *string `json:"routingStatus,omitempty"`
		
		Impact *string `json:"impact,omitempty"`
		
		SecondaryPresenceLookupId *string `json:"secondaryPresenceLookupId,omitempty"`
		Alias
	}{ 
		StartOffsetSeconds: o.StartOffsetSeconds,
		
		EndOffsetSeconds: o.EndOffsetSeconds,
		
		ScheduledActivityCodeId: o.ScheduledActivityCodeId,
		
		ScheduledActivityCategory: o.ScheduledActivityCategory,
		
		ScheduledSecondaryPresenceLookupIds: o.ScheduledSecondaryPresenceLookupIds,
		
		ActualActivityCodeId: o.ActualActivityCodeId,
		
		ActualActivityCategory: o.ActualActivityCategory,
		
		SystemPresence: o.SystemPresence,
		
		RoutingStatus: o.RoutingStatus,
		
		Impact: o.Impact,
		
		SecondaryPresenceLookupId: o.SecondaryPresenceLookupId,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceexceptioninfo) UnmarshalJSON(b []byte) error {
	var WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceexceptioninfoMap map[string]interface{}
	err := json.Unmarshal(b, &WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceexceptioninfoMap)
	if err != nil {
		return err
	}
	
	if StartOffsetSeconds, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceexceptioninfoMap["startOffsetSeconds"].(float64); ok {
		StartOffsetSecondsInt := int(StartOffsetSeconds)
		o.StartOffsetSeconds = &StartOffsetSecondsInt
	}
	
	if EndOffsetSeconds, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceexceptioninfoMap["endOffsetSeconds"].(float64); ok {
		EndOffsetSecondsInt := int(EndOffsetSeconds)
		o.EndOffsetSeconds = &EndOffsetSecondsInt
	}
	
	if ScheduledActivityCodeId, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceexceptioninfoMap["scheduledActivityCodeId"].(string); ok {
		o.ScheduledActivityCodeId = &ScheduledActivityCodeId
	}
    
	if ScheduledActivityCategory, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceexceptioninfoMap["scheduledActivityCategory"].(string); ok {
		o.ScheduledActivityCategory = &ScheduledActivityCategory
	}
    
	if ScheduledSecondaryPresenceLookupIds, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceexceptioninfoMap["scheduledSecondaryPresenceLookupIds"].([]interface{}); ok {
		ScheduledSecondaryPresenceLookupIdsString, _ := json.Marshal(ScheduledSecondaryPresenceLookupIds)
		json.Unmarshal(ScheduledSecondaryPresenceLookupIdsString, &o.ScheduledSecondaryPresenceLookupIds)
	}
	
	if ActualActivityCodeId, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceexceptioninfoMap["actualActivityCodeId"].(string); ok {
		o.ActualActivityCodeId = &ActualActivityCodeId
	}
    
	if ActualActivityCategory, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceexceptioninfoMap["actualActivityCategory"].(string); ok {
		o.ActualActivityCategory = &ActualActivityCategory
	}
    
	if SystemPresence, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceexceptioninfoMap["systemPresence"].(string); ok {
		o.SystemPresence = &SystemPresence
	}
    
	if RoutingStatus, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceexceptioninfoMap["routingStatus"].(string); ok {
		o.RoutingStatus = &RoutingStatus
	}
    
	if Impact, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceexceptioninfoMap["impact"].(string); ok {
		o.Impact = &Impact
	}
    
	if SecondaryPresenceLookupId, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceexceptioninfoMap["secondaryPresenceLookupId"].(string); ok {
		o.SecondaryPresenceLookupId = &SecondaryPresenceLookupId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceexceptioninfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
