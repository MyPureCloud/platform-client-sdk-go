package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmusernotificationtopicwfmusernotification
type Wfmusernotificationtopicwfmusernotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// MutableGroupId
	MutableGroupId *string `json:"mutableGroupId,omitempty"`

	// Timestamp
	Timestamp *time.Time `json:"timestamp,omitempty"`

	// VarType
	VarType *string `json:"type,omitempty"`

	// ShiftTrade
	ShiftTrade *Wfmusernotificationtopicshifttradenotification `json:"shiftTrade,omitempty"`

	// TimeOffRequest
	TimeOffRequest *Wfmusernotificationtopictimeoffrequestnotification `json:"timeOffRequest,omitempty"`

	// AdherenceExplanation
	AdherenceExplanation *Wfmusernotificationtopicadherenceexplanationnotification `json:"adherenceExplanation,omitempty"`

	// AgentNotification
	AgentNotification *bool `json:"agentNotification,omitempty"`

	// OtherNotificationIdsInGroup
	OtherNotificationIdsInGroup *[]string `json:"otherNotificationIdsInGroup,omitempty"`

	// MarkedAsRead
	MarkedAsRead *bool `json:"markedAsRead,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmusernotificationtopicwfmusernotification) SetField(field string, fieldValue interface{}) {
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

func (o Wfmusernotificationtopicwfmusernotification) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "Timestamp", }
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
	type Alias Wfmusernotificationtopicwfmusernotification
	
	Timestamp := new(string)
	if o.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(o.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Timestamp = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		MutableGroupId *string `json:"mutableGroupId,omitempty"`
		
		Timestamp *string `json:"timestamp,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		ShiftTrade *Wfmusernotificationtopicshifttradenotification `json:"shiftTrade,omitempty"`
		
		TimeOffRequest *Wfmusernotificationtopictimeoffrequestnotification `json:"timeOffRequest,omitempty"`
		
		AdherenceExplanation *Wfmusernotificationtopicadherenceexplanationnotification `json:"adherenceExplanation,omitempty"`
		
		AgentNotification *bool `json:"agentNotification,omitempty"`
		
		OtherNotificationIdsInGroup *[]string `json:"otherNotificationIdsInGroup,omitempty"`
		
		MarkedAsRead *bool `json:"markedAsRead,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		MutableGroupId: o.MutableGroupId,
		
		Timestamp: Timestamp,
		
		VarType: o.VarType,
		
		ShiftTrade: o.ShiftTrade,
		
		TimeOffRequest: o.TimeOffRequest,
		
		AdherenceExplanation: o.AdherenceExplanation,
		
		AgentNotification: o.AgentNotification,
		
		OtherNotificationIdsInGroup: o.OtherNotificationIdsInGroup,
		
		MarkedAsRead: o.MarkedAsRead,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmusernotificationtopicwfmusernotification) UnmarshalJSON(b []byte) error {
	var WfmusernotificationtopicwfmusernotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmusernotificationtopicwfmusernotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmusernotificationtopicwfmusernotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if MutableGroupId, ok := WfmusernotificationtopicwfmusernotificationMap["mutableGroupId"].(string); ok {
		o.MutableGroupId = &MutableGroupId
	}
    
	if timestampString, ok := WfmusernotificationtopicwfmusernotificationMap["timestamp"].(string); ok {
		Timestamp, _ := time.Parse("2006-01-02T15:04:05.999999Z", timestampString)
		o.Timestamp = &Timestamp
	}
	
	if VarType, ok := WfmusernotificationtopicwfmusernotificationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if ShiftTrade, ok := WfmusernotificationtopicwfmusernotificationMap["shiftTrade"].(map[string]interface{}); ok {
		ShiftTradeString, _ := json.Marshal(ShiftTrade)
		json.Unmarshal(ShiftTradeString, &o.ShiftTrade)
	}
	
	if TimeOffRequest, ok := WfmusernotificationtopicwfmusernotificationMap["timeOffRequest"].(map[string]interface{}); ok {
		TimeOffRequestString, _ := json.Marshal(TimeOffRequest)
		json.Unmarshal(TimeOffRequestString, &o.TimeOffRequest)
	}
	
	if AdherenceExplanation, ok := WfmusernotificationtopicwfmusernotificationMap["adherenceExplanation"].(map[string]interface{}); ok {
		AdherenceExplanationString, _ := json.Marshal(AdherenceExplanation)
		json.Unmarshal(AdherenceExplanationString, &o.AdherenceExplanation)
	}
	
	if AgentNotification, ok := WfmusernotificationtopicwfmusernotificationMap["agentNotification"].(bool); ok {
		o.AgentNotification = &AgentNotification
	}
    
	if OtherNotificationIdsInGroup, ok := WfmusernotificationtopicwfmusernotificationMap["otherNotificationIdsInGroup"].([]interface{}); ok {
		OtherNotificationIdsInGroupString, _ := json.Marshal(OtherNotificationIdsInGroup)
		json.Unmarshal(OtherNotificationIdsInGroupString, &o.OtherNotificationIdsInGroup)
	}
	
	if MarkedAsRead, ok := WfmusernotificationtopicwfmusernotificationMap["markedAsRead"].(bool); ok {
		o.MarkedAsRead = &MarkedAsRead
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmusernotificationtopicwfmusernotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
