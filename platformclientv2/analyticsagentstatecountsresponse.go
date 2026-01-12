package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsagentstatecountsresponse
type Analyticsagentstatecountsresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SegmentCounts - List of count by segment types
	SegmentCounts *[]Agentstatesegmenttypecount `json:"segmentCounts,omitempty"`

	// PresenceCounts - List of count by presences
	PresenceCounts *[]Agentstatepresencecount `json:"presenceCounts,omitempty"`

	// RoutingStatusCounts - List of count by routing statuses
	RoutingStatusCounts *[]Agentstateroutingstatuscount `json:"routingStatusCounts,omitempty"`

	// IsOutOfOfficeCounts - List of count by out of office states
	IsOutOfOfficeCounts *[]Agentstateisoutofofficecount `json:"isOutOfOfficeCounts,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Analyticsagentstatecountsresponse) SetField(field string, fieldValue interface{}) {
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

func (o Analyticsagentstatecountsresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Analyticsagentstatecountsresponse
	
	return json.Marshal(&struct { 
		SegmentCounts *[]Agentstatesegmenttypecount `json:"segmentCounts,omitempty"`
		
		PresenceCounts *[]Agentstatepresencecount `json:"presenceCounts,omitempty"`
		
		RoutingStatusCounts *[]Agentstateroutingstatuscount `json:"routingStatusCounts,omitempty"`
		
		IsOutOfOfficeCounts *[]Agentstateisoutofofficecount `json:"isOutOfOfficeCounts,omitempty"`
		Alias
	}{ 
		SegmentCounts: o.SegmentCounts,
		
		PresenceCounts: o.PresenceCounts,
		
		RoutingStatusCounts: o.RoutingStatusCounts,
		
		IsOutOfOfficeCounts: o.IsOutOfOfficeCounts,
		Alias:    (Alias)(o),
	})
}

func (o *Analyticsagentstatecountsresponse) UnmarshalJSON(b []byte) error {
	var AnalyticsagentstatecountsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsagentstatecountsresponseMap)
	if err != nil {
		return err
	}
	
	if SegmentCounts, ok := AnalyticsagentstatecountsresponseMap["segmentCounts"].([]interface{}); ok {
		SegmentCountsString, _ := json.Marshal(SegmentCounts)
		json.Unmarshal(SegmentCountsString, &o.SegmentCounts)
	}
	
	if PresenceCounts, ok := AnalyticsagentstatecountsresponseMap["presenceCounts"].([]interface{}); ok {
		PresenceCountsString, _ := json.Marshal(PresenceCounts)
		json.Unmarshal(PresenceCountsString, &o.PresenceCounts)
	}
	
	if RoutingStatusCounts, ok := AnalyticsagentstatecountsresponseMap["routingStatusCounts"].([]interface{}); ok {
		RoutingStatusCountsString, _ := json.Marshal(RoutingStatusCounts)
		json.Unmarshal(RoutingStatusCountsString, &o.RoutingStatusCounts)
	}
	
	if IsOutOfOfficeCounts, ok := AnalyticsagentstatecountsresponseMap["isOutOfOfficeCounts"].([]interface{}); ok {
		IsOutOfOfficeCountsString, _ := json.Marshal(IsOutOfOfficeCounts)
		json.Unmarshal(IsOutOfOfficeCountsString, &o.IsOutOfOfficeCounts)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsagentstatecountsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
