package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Outboundmessagingcampaignpostcontactsmseventtopicoutboundmessagingcampaignpostcontactsmsevent
type Outboundmessagingcampaignpostcontactsmseventtopicoutboundmessagingcampaignpostcontactsmsevent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EventTime
	EventTime *int `json:"eventTime,omitempty"`

	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`

	// OutboundCampaignType
	OutboundCampaignType *string `json:"outboundCampaignType,omitempty"`

	// SmsAttributes
	SmsAttributes *Outboundmessagingcampaignpostcontactsmseventtopicsmsattributes `json:"smsAttributes,omitempty"`

	// WrapupCode
	WrapupCode *string `json:"wrapupCode,omitempty"`

	// OutboundCampaignId
	OutboundCampaignId *string `json:"outboundCampaignId,omitempty"`

	// DivisionId
	DivisionId *string `json:"divisionId,omitempty"`

	// ContentTemplateId
	ContentTemplateId *string `json:"contentTemplateId,omitempty"`

	// OutboundContactListId
	OutboundContactListId *string `json:"outboundContactListId,omitempty"`

	// OutboundContactId
	OutboundContactId *string `json:"outboundContactId,omitempty"`

	// IsCampaignAlwaysRunning
	IsCampaignAlwaysRunning *bool `json:"isCampaignAlwaysRunning,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Outboundmessagingcampaignpostcontactsmseventtopicoutboundmessagingcampaignpostcontactsmsevent) SetField(field string, fieldValue interface{}) {
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

func (o Outboundmessagingcampaignpostcontactsmseventtopicoutboundmessagingcampaignpostcontactsmsevent) MarshalJSON() ([]byte, error) {
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
	type Alias Outboundmessagingcampaignpostcontactsmseventtopicoutboundmessagingcampaignpostcontactsmsevent
	
	return json.Marshal(&struct { 
		EventTime *int `json:"eventTime,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		OutboundCampaignType *string `json:"outboundCampaignType,omitempty"`
		
		SmsAttributes *Outboundmessagingcampaignpostcontactsmseventtopicsmsattributes `json:"smsAttributes,omitempty"`
		
		WrapupCode *string `json:"wrapupCode,omitempty"`
		
		OutboundCampaignId *string `json:"outboundCampaignId,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		ContentTemplateId *string `json:"contentTemplateId,omitempty"`
		
		OutboundContactListId *string `json:"outboundContactListId,omitempty"`
		
		OutboundContactId *string `json:"outboundContactId,omitempty"`
		
		IsCampaignAlwaysRunning *bool `json:"isCampaignAlwaysRunning,omitempty"`
		Alias
	}{ 
		EventTime: o.EventTime,
		
		ConversationId: o.ConversationId,
		
		OutboundCampaignType: o.OutboundCampaignType,
		
		SmsAttributes: o.SmsAttributes,
		
		WrapupCode: o.WrapupCode,
		
		OutboundCampaignId: o.OutboundCampaignId,
		
		DivisionId: o.DivisionId,
		
		ContentTemplateId: o.ContentTemplateId,
		
		OutboundContactListId: o.OutboundContactListId,
		
		OutboundContactId: o.OutboundContactId,
		
		IsCampaignAlwaysRunning: o.IsCampaignAlwaysRunning,
		Alias:    (Alias)(o),
	})
}

func (o *Outboundmessagingcampaignpostcontactsmseventtopicoutboundmessagingcampaignpostcontactsmsevent) UnmarshalJSON(b []byte) error {
	var OutboundmessagingcampaignpostcontactsmseventtopicoutboundmessagingcampaignpostcontactsmseventMap map[string]interface{}
	err := json.Unmarshal(b, &OutboundmessagingcampaignpostcontactsmseventtopicoutboundmessagingcampaignpostcontactsmseventMap)
	if err != nil {
		return err
	}
	
	if EventTime, ok := OutboundmessagingcampaignpostcontactsmseventtopicoutboundmessagingcampaignpostcontactsmseventMap["eventTime"].(float64); ok {
		EventTimeInt := int(EventTime)
		o.EventTime = &EventTimeInt
	}
	
	if ConversationId, ok := OutboundmessagingcampaignpostcontactsmseventtopicoutboundmessagingcampaignpostcontactsmseventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if OutboundCampaignType, ok := OutboundmessagingcampaignpostcontactsmseventtopicoutboundmessagingcampaignpostcontactsmseventMap["outboundCampaignType"].(string); ok {
		o.OutboundCampaignType = &OutboundCampaignType
	}
    
	if SmsAttributes, ok := OutboundmessagingcampaignpostcontactsmseventtopicoutboundmessagingcampaignpostcontactsmseventMap["smsAttributes"].(map[string]interface{}); ok {
		SmsAttributesString, _ := json.Marshal(SmsAttributes)
		json.Unmarshal(SmsAttributesString, &o.SmsAttributes)
	}
	
	if WrapupCode, ok := OutboundmessagingcampaignpostcontactsmseventtopicoutboundmessagingcampaignpostcontactsmseventMap["wrapupCode"].(string); ok {
		o.WrapupCode = &WrapupCode
	}
    
	if OutboundCampaignId, ok := OutboundmessagingcampaignpostcontactsmseventtopicoutboundmessagingcampaignpostcontactsmseventMap["outboundCampaignId"].(string); ok {
		o.OutboundCampaignId = &OutboundCampaignId
	}
    
	if DivisionId, ok := OutboundmessagingcampaignpostcontactsmseventtopicoutboundmessagingcampaignpostcontactsmseventMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if ContentTemplateId, ok := OutboundmessagingcampaignpostcontactsmseventtopicoutboundmessagingcampaignpostcontactsmseventMap["contentTemplateId"].(string); ok {
		o.ContentTemplateId = &ContentTemplateId
	}
    
	if OutboundContactListId, ok := OutboundmessagingcampaignpostcontactsmseventtopicoutboundmessagingcampaignpostcontactsmseventMap["outboundContactListId"].(string); ok {
		o.OutboundContactListId = &OutboundContactListId
	}
    
	if OutboundContactId, ok := OutboundmessagingcampaignpostcontactsmseventtopicoutboundmessagingcampaignpostcontactsmseventMap["outboundContactId"].(string); ok {
		o.OutboundContactId = &OutboundContactId
	}
    
	if IsCampaignAlwaysRunning, ok := OutboundmessagingcampaignpostcontactsmseventtopicoutboundmessagingcampaignpostcontactsmseventMap["isCampaignAlwaysRunning"].(bool); ok {
		o.IsCampaignAlwaysRunning = &IsCampaignAlwaysRunning
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Outboundmessagingcampaignpostcontactsmseventtopicoutboundmessagingcampaignpostcontactsmsevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
