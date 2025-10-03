package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Outboundvoicecampaignpostcontactoutboundvoicecampaignpostcontactevent
type Outboundvoicecampaignpostcontactoutboundvoicecampaignpostcontactevent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EventTime
	EventTime *int `json:"eventTime,omitempty"`

	// VoiceAttributes
	VoiceAttributes *Outboundvoicecampaignpostcontactvoiceattributes `json:"voiceAttributes,omitempty"`

	// WrapupCode
	WrapupCode *string `json:"wrapupCode,omitempty"`

	// OutboundCampaignId
	OutboundCampaignId *string `json:"outboundCampaignId,omitempty"`

	// DialingMode
	DialingMode *string `json:"dialingMode,omitempty"`

	// AgentScriptId
	AgentScriptId *string `json:"agentScriptId,omitempty"`

	// DivisionId
	DivisionId *string `json:"divisionId,omitempty"`

	// OutboundContactListId
	OutboundContactListId *string `json:"outboundContactListId,omitempty"`

	// OutboundContactListFilterId
	OutboundContactListFilterId *string `json:"outboundContactListFilterId,omitempty"`

	// OutboundQueueId
	OutboundQueueId *string `json:"outboundQueueId,omitempty"`

	// OutboundContactId
	OutboundContactId *string `json:"outboundContactId,omitempty"`

	// IsCampaignAlwaysRunning
	IsCampaignAlwaysRunning *bool `json:"isCampaignAlwaysRunning,omitempty"`

	// IsCampaignSkillBased
	IsCampaignSkillBased *bool `json:"isCampaignSkillBased,omitempty"`

	// IsCampaignDynamicSorting
	IsCampaignDynamicSorting *bool `json:"isCampaignDynamicSorting,omitempty"`

	// IsCampaignDynamicFiltering
	IsCampaignDynamicFiltering *bool `json:"isCampaignDynamicFiltering,omitempty"`

	// OutboundCampaignHealthMask
	OutboundCampaignHealthMask *int `json:"outboundCampaignHealthMask,omitempty"`

	// IsReCall
	IsReCall *bool `json:"isReCall,omitempty"`

	// ScheduledDateTime
	ScheduledDateTime *string `json:"scheduledDateTime,omitempty"`

	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`

	// IsRightPartyContact
	IsRightPartyContact *bool `json:"isRightPartyContact,omitempty"`

	// CallAnalysisDisposition
	CallAnalysisDisposition *string `json:"callAnalysisDisposition,omitempty"`

	// AgentId
	AgentId *string `json:"agentId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Outboundvoicecampaignpostcontactoutboundvoicecampaignpostcontactevent) SetField(field string, fieldValue interface{}) {
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

func (o Outboundvoicecampaignpostcontactoutboundvoicecampaignpostcontactevent) MarshalJSON() ([]byte, error) {
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
	type Alias Outboundvoicecampaignpostcontactoutboundvoicecampaignpostcontactevent
	
	return json.Marshal(&struct { 
		EventTime *int `json:"eventTime,omitempty"`
		
		VoiceAttributes *Outboundvoicecampaignpostcontactvoiceattributes `json:"voiceAttributes,omitempty"`
		
		WrapupCode *string `json:"wrapupCode,omitempty"`
		
		OutboundCampaignId *string `json:"outboundCampaignId,omitempty"`
		
		DialingMode *string `json:"dialingMode,omitempty"`
		
		AgentScriptId *string `json:"agentScriptId,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		OutboundContactListId *string `json:"outboundContactListId,omitempty"`
		
		OutboundContactListFilterId *string `json:"outboundContactListFilterId,omitempty"`
		
		OutboundQueueId *string `json:"outboundQueueId,omitempty"`
		
		OutboundContactId *string `json:"outboundContactId,omitempty"`
		
		IsCampaignAlwaysRunning *bool `json:"isCampaignAlwaysRunning,omitempty"`
		
		IsCampaignSkillBased *bool `json:"isCampaignSkillBased,omitempty"`
		
		IsCampaignDynamicSorting *bool `json:"isCampaignDynamicSorting,omitempty"`
		
		IsCampaignDynamicFiltering *bool `json:"isCampaignDynamicFiltering,omitempty"`
		
		OutboundCampaignHealthMask *int `json:"outboundCampaignHealthMask,omitempty"`
		
		IsReCall *bool `json:"isReCall,omitempty"`
		
		ScheduledDateTime *string `json:"scheduledDateTime,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		IsRightPartyContact *bool `json:"isRightPartyContact,omitempty"`
		
		CallAnalysisDisposition *string `json:"callAnalysisDisposition,omitempty"`
		
		AgentId *string `json:"agentId,omitempty"`
		Alias
	}{ 
		EventTime: o.EventTime,
		
		VoiceAttributes: o.VoiceAttributes,
		
		WrapupCode: o.WrapupCode,
		
		OutboundCampaignId: o.OutboundCampaignId,
		
		DialingMode: o.DialingMode,
		
		AgentScriptId: o.AgentScriptId,
		
		DivisionId: o.DivisionId,
		
		OutboundContactListId: o.OutboundContactListId,
		
		OutboundContactListFilterId: o.OutboundContactListFilterId,
		
		OutboundQueueId: o.OutboundQueueId,
		
		OutboundContactId: o.OutboundContactId,
		
		IsCampaignAlwaysRunning: o.IsCampaignAlwaysRunning,
		
		IsCampaignSkillBased: o.IsCampaignSkillBased,
		
		IsCampaignDynamicSorting: o.IsCampaignDynamicSorting,
		
		IsCampaignDynamicFiltering: o.IsCampaignDynamicFiltering,
		
		OutboundCampaignHealthMask: o.OutboundCampaignHealthMask,
		
		IsReCall: o.IsReCall,
		
		ScheduledDateTime: o.ScheduledDateTime,
		
		ConversationId: o.ConversationId,
		
		IsRightPartyContact: o.IsRightPartyContact,
		
		CallAnalysisDisposition: o.CallAnalysisDisposition,
		
		AgentId: o.AgentId,
		Alias:    (Alias)(o),
	})
}

func (o *Outboundvoicecampaignpostcontactoutboundvoicecampaignpostcontactevent) UnmarshalJSON(b []byte) error {
	var OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap map[string]interface{}
	err := json.Unmarshal(b, &OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap)
	if err != nil {
		return err
	}
	
	if EventTime, ok := OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap["eventTime"].(float64); ok {
		EventTimeInt := int(EventTime)
		o.EventTime = &EventTimeInt
	}
	
	if VoiceAttributes, ok := OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap["voiceAttributes"].(map[string]interface{}); ok {
		VoiceAttributesString, _ := json.Marshal(VoiceAttributes)
		json.Unmarshal(VoiceAttributesString, &o.VoiceAttributes)
	}
	
	if WrapupCode, ok := OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap["wrapupCode"].(string); ok {
		o.WrapupCode = &WrapupCode
	}
    
	if OutboundCampaignId, ok := OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap["outboundCampaignId"].(string); ok {
		o.OutboundCampaignId = &OutboundCampaignId
	}
    
	if DialingMode, ok := OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap["dialingMode"].(string); ok {
		o.DialingMode = &DialingMode
	}
    
	if AgentScriptId, ok := OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap["agentScriptId"].(string); ok {
		o.AgentScriptId = &AgentScriptId
	}
    
	if DivisionId, ok := OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if OutboundContactListId, ok := OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap["outboundContactListId"].(string); ok {
		o.OutboundContactListId = &OutboundContactListId
	}
    
	if OutboundContactListFilterId, ok := OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap["outboundContactListFilterId"].(string); ok {
		o.OutboundContactListFilterId = &OutboundContactListFilterId
	}
    
	if OutboundQueueId, ok := OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap["outboundQueueId"].(string); ok {
		o.OutboundQueueId = &OutboundQueueId
	}
    
	if OutboundContactId, ok := OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap["outboundContactId"].(string); ok {
		o.OutboundContactId = &OutboundContactId
	}
    
	if IsCampaignAlwaysRunning, ok := OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap["isCampaignAlwaysRunning"].(bool); ok {
		o.IsCampaignAlwaysRunning = &IsCampaignAlwaysRunning
	}
    
	if IsCampaignSkillBased, ok := OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap["isCampaignSkillBased"].(bool); ok {
		o.IsCampaignSkillBased = &IsCampaignSkillBased
	}
    
	if IsCampaignDynamicSorting, ok := OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap["isCampaignDynamicSorting"].(bool); ok {
		o.IsCampaignDynamicSorting = &IsCampaignDynamicSorting
	}
    
	if IsCampaignDynamicFiltering, ok := OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap["isCampaignDynamicFiltering"].(bool); ok {
		o.IsCampaignDynamicFiltering = &IsCampaignDynamicFiltering
	}
    
	if OutboundCampaignHealthMask, ok := OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap["outboundCampaignHealthMask"].(float64); ok {
		OutboundCampaignHealthMaskInt := int(OutboundCampaignHealthMask)
		o.OutboundCampaignHealthMask = &OutboundCampaignHealthMaskInt
	}
	
	if IsReCall, ok := OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap["isReCall"].(bool); ok {
		o.IsReCall = &IsReCall
	}
    
	if ScheduledDateTime, ok := OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap["scheduledDateTime"].(string); ok {
		o.ScheduledDateTime = &ScheduledDateTime
	}
    
	if ConversationId, ok := OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if IsRightPartyContact, ok := OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap["isRightPartyContact"].(bool); ok {
		o.IsRightPartyContact = &IsRightPartyContact
	}
    
	if CallAnalysisDisposition, ok := OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap["callAnalysisDisposition"].(string); ok {
		o.CallAnalysisDisposition = &CallAnalysisDisposition
	}
    
	if AgentId, ok := OutboundvoicecampaignpostcontactoutboundvoicecampaignpostcontacteventMap["agentId"].(string); ok {
		o.AgentId = &AgentId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Outboundvoicecampaignpostcontactoutboundvoicecampaignpostcontactevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
