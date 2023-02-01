package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Historicaladherenceexceptioninfo
type Historicaladherenceexceptioninfo struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// StartOffsetSeconds - Exception start offset in seconds relative to query start time
	StartOffsetSeconds *int `json:"startOffsetSeconds,omitempty"`

	// EndOffsetSeconds - Exception end offset in seconds relative to query start time
	EndOffsetSeconds *int `json:"endOffsetSeconds,omitempty"`

	// ScheduledActivityCodeId - The ID of the scheduled activity code for this user
	ScheduledActivityCodeId *string `json:"scheduledActivityCodeId,omitempty"`

	// ScheduledActivityCategory - Activity for which the user is scheduled
	ScheduledActivityCategory *string `json:"scheduledActivityCategory,omitempty"`

	// ScheduledSecondaryPresenceLookupIds - The lookup IDs used to retrieve the scheduled secondary statuses from map of lookup ID to corresponding secondary presence ID
	ScheduledSecondaryPresenceLookupIds *[]string `json:"scheduledSecondaryPresenceLookupIds,omitempty"`

	// ActualActivityCodeId - The ID of the actual activity code for this user
	ActualActivityCodeId *string `json:"actualActivityCodeId,omitempty"`

	// ActualActivityCategory - Activity for which the user is actually engaged
	ActualActivityCategory *string `json:"actualActivityCategory,omitempty"`

	// SystemPresence - Actual underlying system presence value
	SystemPresence *string `json:"systemPresence,omitempty"`

	// RoutingStatus - Actual underlying routing status, used to determine whether a user is actually in adherence when OnQueue
	RoutingStatus *string `json:"routingStatus,omitempty"`

	// Impact - The impact of the current adherence state for this user
	Impact *string `json:"impact,omitempty"`

	// SecondaryPresenceLookupId - The lookup ID used to retrieve the actual secondary status from map of lookup ID to corresponding secondary presence ID
	SecondaryPresenceLookupId *string `json:"secondaryPresenceLookupId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Historicaladherenceexceptioninfo) SetField(field string, fieldValue interface{}) {
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

func (o Historicaladherenceexceptioninfo) MarshalJSON() ([]byte, error) {
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
	type Alias Historicaladherenceexceptioninfo
	
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

func (o *Historicaladherenceexceptioninfo) UnmarshalJSON(b []byte) error {
	var HistoricaladherenceexceptioninfoMap map[string]interface{}
	err := json.Unmarshal(b, &HistoricaladherenceexceptioninfoMap)
	if err != nil {
		return err
	}
	
	if StartOffsetSeconds, ok := HistoricaladherenceexceptioninfoMap["startOffsetSeconds"].(float64); ok {
		StartOffsetSecondsInt := int(StartOffsetSeconds)
		o.StartOffsetSeconds = &StartOffsetSecondsInt
	}
	
	if EndOffsetSeconds, ok := HistoricaladherenceexceptioninfoMap["endOffsetSeconds"].(float64); ok {
		EndOffsetSecondsInt := int(EndOffsetSeconds)
		o.EndOffsetSeconds = &EndOffsetSecondsInt
	}
	
	if ScheduledActivityCodeId, ok := HistoricaladherenceexceptioninfoMap["scheduledActivityCodeId"].(string); ok {
		o.ScheduledActivityCodeId = &ScheduledActivityCodeId
	}
    
	if ScheduledActivityCategory, ok := HistoricaladherenceexceptioninfoMap["scheduledActivityCategory"].(string); ok {
		o.ScheduledActivityCategory = &ScheduledActivityCategory
	}
    
	if ScheduledSecondaryPresenceLookupIds, ok := HistoricaladherenceexceptioninfoMap["scheduledSecondaryPresenceLookupIds"].([]interface{}); ok {
		ScheduledSecondaryPresenceLookupIdsString, _ := json.Marshal(ScheduledSecondaryPresenceLookupIds)
		json.Unmarshal(ScheduledSecondaryPresenceLookupIdsString, &o.ScheduledSecondaryPresenceLookupIds)
	}
	
	if ActualActivityCodeId, ok := HistoricaladherenceexceptioninfoMap["actualActivityCodeId"].(string); ok {
		o.ActualActivityCodeId = &ActualActivityCodeId
	}
    
	if ActualActivityCategory, ok := HistoricaladherenceexceptioninfoMap["actualActivityCategory"].(string); ok {
		o.ActualActivityCategory = &ActualActivityCategory
	}
    
	if SystemPresence, ok := HistoricaladherenceexceptioninfoMap["systemPresence"].(string); ok {
		o.SystemPresence = &SystemPresence
	}
    
	if RoutingStatus, ok := HistoricaladherenceexceptioninfoMap["routingStatus"].(string); ok {
		o.RoutingStatus = &RoutingStatus
	}
    
	if Impact, ok := HistoricaladherenceexceptioninfoMap["impact"].(string); ok {
		o.Impact = &Impact
	}
    
	if SecondaryPresenceLookupId, ok := HistoricaladherenceexceptioninfoMap["secondaryPresenceLookupId"].(string); ok {
		o.SecondaryPresenceLookupId = &SecondaryPresenceLookupId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Historicaladherenceexceptioninfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
