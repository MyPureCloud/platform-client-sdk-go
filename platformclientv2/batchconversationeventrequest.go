package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Batchconversationeventrequest - A maximum of 100 events are allowed per request
type Batchconversationeventrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EndTransferEvents - Voice - EndTransfer events for this batch
	EndTransferEvents *[]Endtransferevent `json:"endTransferEvents,omitempty"`

	// PhoneTransferEvents - Voice - PhoneTransfer events for this batch
	PhoneTransferEvents *[]Phonetransferevent `json:"phoneTransferEvents,omitempty"`

	// ProgressTransferEvents - Voice - ProgressTransfer events for this batch
	ProgressTransferEvents *[]Progresstransferevent `json:"progressTransferEvents,omitempty"`

	// RoutingTransferEvents - Voice - RoutingTransfer events for this batch
	RoutingTransferEvents *[]Routingtransferevent `json:"routingTransferEvents,omitempty"`

	// UserTransferEvents - Voice - UserTransfer events for this batch
	UserTransferEvents *[]Usertransferevent `json:"userTransferEvents,omitempty"`

	// CommunicationAnsweredEvents - Voice - CommunicationAnswered events for this batch
	CommunicationAnsweredEvents *[]Communicationansweredevent `json:"communicationAnsweredEvents,omitempty"`

	// CommunicationDispositionAppliedEvents - Voice - CommunicationDispositionApplied events for this batch
	CommunicationDispositionAppliedEvents *[]Communicationdispositionappliedevent `json:"communicationDispositionAppliedEvents,omitempty"`

	// HoldUpdatedEvents - Voice - HoldUpdated events for this batch
	HoldUpdatedEvents *[]Holdupdatedevent `json:"holdUpdatedEvents,omitempty"`

	// ExternalEstablishedEvents - Voice - ExternalEstablished events for this batch
	ExternalEstablishedEvents *[]Externalestablishedevent `json:"externalEstablishedEvents,omitempty"`

	// IvrEstablishedEvents - Voice - IvrEstablished events for this batch
	IvrEstablishedEvents *[]Ivrestablishedevent `json:"ivrEstablishedEvents,omitempty"`

	// PhoneEstablishedEvents - Voice - PhoneEstablished events for this batch
	PhoneEstablishedEvents *[]Phoneestablishedevent `json:"phoneEstablishedEvents,omitempty"`

	// RoutingEstablishedEvents - Voice - RoutingEstablished events for this batch
	RoutingEstablishedEvents *[]Routingestablishedevent `json:"routingEstablishedEvents,omitempty"`

	// UserEstablishedEvents - Voice - UserEstablished events for this batch
	UserEstablishedEvents *[]Userestablishedevent `json:"userEstablishedEvents,omitempty"`

	// AudioUpdatedEvents - Voice - AudioUpdated events for this batch
	AudioUpdatedEvents *[]Audioupdatedevent `json:"audioUpdatedEvents,omitempty"`

	// CommunicationEndedEvents - Voice - CommunicationEnded events for this batch
	CommunicationEndedEvents *[]Communicationendedevent `json:"communicationEndedEvents,omitempty"`

	// ConsultTransferEvents - Voice - ConsultTransfer events for this batch
	ConsultTransferEvents *[]Consulttransferevent `json:"consultTransferEvents,omitempty"`

	// ProgressConsultTransferEvents - Voice - ProgressConsultTransfer events for this batch
	ProgressConsultTransferEvents *[]Progressconsulttransferevent `json:"progressConsultTransferEvents,omitempty"`

	// EndConsultTransferEvents - Voice - EndConsultTransfer events for this batch
	EndConsultTransferEvents *[]Endconsulttransferevent `json:"endConsultTransferEvents,omitempty"`

	// EmailBeginTransmittingEvents - Email - EmailBeginTransmittingEvent events for this batch
	EmailBeginTransmittingEvents *[]Emailbegintransmittingevent `json:"emailBeginTransmittingEvents,omitempty"`

	// EmailCommunicationEndedEvents - Email - EmailCommunicationEndedEvent events for this batch
	EmailCommunicationEndedEvents *[]Emailcommunicationendedevent `json:"emailCommunicationEndedEvents,omitempty"`

	// EmailExternalEstablishedEvents - Email - EmailExternalEstablishedEvent events for this batch
	EmailExternalEstablishedEvents *[]Emailexternalestablishedevent `json:"emailExternalEstablishedEvents,omitempty"`

	// EmailFlowEstablishedEvents - Email - EmailFlowEstablishedEvent events for this batch
	EmailFlowEstablishedEvents *[]Emailflowestablishedevent `json:"emailFlowEstablishedEvents,omitempty"`

	// EmailRoutingEstablishedEvents - Email - EmailRoutingEstablishedEvent events for this batch
	EmailRoutingEstablishedEvents *[]Emailroutingestablishedevent `json:"emailRoutingEstablishedEvents,omitempty"`

	// EmailUserEstablishedEvents - Email - EmailUserEstablishedEvent events for this batch
	EmailUserEstablishedEvents *[]Emailuserestablishedevent `json:"emailUserEstablishedEvents,omitempty"`

	// EmailCommunicationAnsweredEvents - Email - EmailCommunicationAnsweredEvent events for this batch
	EmailCommunicationAnsweredEvents *[]Emailcommunicationansweredevent `json:"emailCommunicationAnsweredEvents,omitempty"`

	// EmailCommunicationDispositionAppliedEvents - Email - EmailCommunicationDispositionAppliedEvent events for this batch
	EmailCommunicationDispositionAppliedEvents *[]Emailcommunicationdispositionappliedevent `json:"emailCommunicationDispositionAppliedEvents,omitempty"`

	// EmailCommunicationSentMessageEvents - Email - EmailCommunicationSentMessageEvent events for this batch
	EmailCommunicationSentMessageEvents *[]Emailcommunicationsentmessageevent `json:"emailCommunicationSentMessageEvents,omitempty"`

	// EmailHoldUpdatedEvents - Email - EmailHoldUpdatedEvent events for this batch
	EmailHoldUpdatedEvents *[]Emailholdupdatedevent `json:"emailHoldUpdatedEvents,omitempty"`

	// EmailEndTransferEvents - Email - EmailEndTransferEvent events for this batch
	EmailEndTransferEvents *[]Emailendtransferevent `json:"emailEndTransferEvents,omitempty"`

	// EmailProgressTransferEvents - Email - EmailProgressTransferEvent events for this batch
	EmailProgressTransferEvents *[]Emailprogresstransferevent `json:"emailProgressTransferEvents,omitempty"`

	// EmailRoutingTransferEvents - Email - EmailRoutingTransferEvent events for this batch
	EmailRoutingTransferEvents *[]Emailroutingtransferevent `json:"emailRoutingTransferEvents,omitempty"`

	// EmailUserTransferEvents - Email - EmailUserTransferEvent events for this batch
	EmailUserTransferEvents *[]Emailusertransferevent `json:"emailUserTransferEvents,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Batchconversationeventrequest) SetField(field string, fieldValue interface{}) {
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

func (o Batchconversationeventrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Batchconversationeventrequest
	
	return json.Marshal(&struct { 
		EndTransferEvents *[]Endtransferevent `json:"endTransferEvents,omitempty"`
		
		PhoneTransferEvents *[]Phonetransferevent `json:"phoneTransferEvents,omitempty"`
		
		ProgressTransferEvents *[]Progresstransferevent `json:"progressTransferEvents,omitempty"`
		
		RoutingTransferEvents *[]Routingtransferevent `json:"routingTransferEvents,omitempty"`
		
		UserTransferEvents *[]Usertransferevent `json:"userTransferEvents,omitempty"`
		
		CommunicationAnsweredEvents *[]Communicationansweredevent `json:"communicationAnsweredEvents,omitempty"`
		
		CommunicationDispositionAppliedEvents *[]Communicationdispositionappliedevent `json:"communicationDispositionAppliedEvents,omitempty"`
		
		HoldUpdatedEvents *[]Holdupdatedevent `json:"holdUpdatedEvents,omitempty"`
		
		ExternalEstablishedEvents *[]Externalestablishedevent `json:"externalEstablishedEvents,omitempty"`
		
		IvrEstablishedEvents *[]Ivrestablishedevent `json:"ivrEstablishedEvents,omitempty"`
		
		PhoneEstablishedEvents *[]Phoneestablishedevent `json:"phoneEstablishedEvents,omitempty"`
		
		RoutingEstablishedEvents *[]Routingestablishedevent `json:"routingEstablishedEvents,omitempty"`
		
		UserEstablishedEvents *[]Userestablishedevent `json:"userEstablishedEvents,omitempty"`
		
		AudioUpdatedEvents *[]Audioupdatedevent `json:"audioUpdatedEvents,omitempty"`
		
		CommunicationEndedEvents *[]Communicationendedevent `json:"communicationEndedEvents,omitempty"`
		
		ConsultTransferEvents *[]Consulttransferevent `json:"consultTransferEvents,omitempty"`
		
		ProgressConsultTransferEvents *[]Progressconsulttransferevent `json:"progressConsultTransferEvents,omitempty"`
		
		EndConsultTransferEvents *[]Endconsulttransferevent `json:"endConsultTransferEvents,omitempty"`
		
		EmailBeginTransmittingEvents *[]Emailbegintransmittingevent `json:"emailBeginTransmittingEvents,omitempty"`
		
		EmailCommunicationEndedEvents *[]Emailcommunicationendedevent `json:"emailCommunicationEndedEvents,omitempty"`
		
		EmailExternalEstablishedEvents *[]Emailexternalestablishedevent `json:"emailExternalEstablishedEvents,omitempty"`
		
		EmailFlowEstablishedEvents *[]Emailflowestablishedevent `json:"emailFlowEstablishedEvents,omitempty"`
		
		EmailRoutingEstablishedEvents *[]Emailroutingestablishedevent `json:"emailRoutingEstablishedEvents,omitempty"`
		
		EmailUserEstablishedEvents *[]Emailuserestablishedevent `json:"emailUserEstablishedEvents,omitempty"`
		
		EmailCommunicationAnsweredEvents *[]Emailcommunicationansweredevent `json:"emailCommunicationAnsweredEvents,omitempty"`
		
		EmailCommunicationDispositionAppliedEvents *[]Emailcommunicationdispositionappliedevent `json:"emailCommunicationDispositionAppliedEvents,omitempty"`
		
		EmailCommunicationSentMessageEvents *[]Emailcommunicationsentmessageevent `json:"emailCommunicationSentMessageEvents,omitempty"`
		
		EmailHoldUpdatedEvents *[]Emailholdupdatedevent `json:"emailHoldUpdatedEvents,omitempty"`
		
		EmailEndTransferEvents *[]Emailendtransferevent `json:"emailEndTransferEvents,omitempty"`
		
		EmailProgressTransferEvents *[]Emailprogresstransferevent `json:"emailProgressTransferEvents,omitempty"`
		
		EmailRoutingTransferEvents *[]Emailroutingtransferevent `json:"emailRoutingTransferEvents,omitempty"`
		
		EmailUserTransferEvents *[]Emailusertransferevent `json:"emailUserTransferEvents,omitempty"`
		Alias
	}{ 
		EndTransferEvents: o.EndTransferEvents,
		
		PhoneTransferEvents: o.PhoneTransferEvents,
		
		ProgressTransferEvents: o.ProgressTransferEvents,
		
		RoutingTransferEvents: o.RoutingTransferEvents,
		
		UserTransferEvents: o.UserTransferEvents,
		
		CommunicationAnsweredEvents: o.CommunicationAnsweredEvents,
		
		CommunicationDispositionAppliedEvents: o.CommunicationDispositionAppliedEvents,
		
		HoldUpdatedEvents: o.HoldUpdatedEvents,
		
		ExternalEstablishedEvents: o.ExternalEstablishedEvents,
		
		IvrEstablishedEvents: o.IvrEstablishedEvents,
		
		PhoneEstablishedEvents: o.PhoneEstablishedEvents,
		
		RoutingEstablishedEvents: o.RoutingEstablishedEvents,
		
		UserEstablishedEvents: o.UserEstablishedEvents,
		
		AudioUpdatedEvents: o.AudioUpdatedEvents,
		
		CommunicationEndedEvents: o.CommunicationEndedEvents,
		
		ConsultTransferEvents: o.ConsultTransferEvents,
		
		ProgressConsultTransferEvents: o.ProgressConsultTransferEvents,
		
		EndConsultTransferEvents: o.EndConsultTransferEvents,
		
		EmailBeginTransmittingEvents: o.EmailBeginTransmittingEvents,
		
		EmailCommunicationEndedEvents: o.EmailCommunicationEndedEvents,
		
		EmailExternalEstablishedEvents: o.EmailExternalEstablishedEvents,
		
		EmailFlowEstablishedEvents: o.EmailFlowEstablishedEvents,
		
		EmailRoutingEstablishedEvents: o.EmailRoutingEstablishedEvents,
		
		EmailUserEstablishedEvents: o.EmailUserEstablishedEvents,
		
		EmailCommunicationAnsweredEvents: o.EmailCommunicationAnsweredEvents,
		
		EmailCommunicationDispositionAppliedEvents: o.EmailCommunicationDispositionAppliedEvents,
		
		EmailCommunicationSentMessageEvents: o.EmailCommunicationSentMessageEvents,
		
		EmailHoldUpdatedEvents: o.EmailHoldUpdatedEvents,
		
		EmailEndTransferEvents: o.EmailEndTransferEvents,
		
		EmailProgressTransferEvents: o.EmailProgressTransferEvents,
		
		EmailRoutingTransferEvents: o.EmailRoutingTransferEvents,
		
		EmailUserTransferEvents: o.EmailUserTransferEvents,
		Alias:    (Alias)(o),
	})
}

func (o *Batchconversationeventrequest) UnmarshalJSON(b []byte) error {
	var BatchconversationeventrequestMap map[string]interface{}
	err := json.Unmarshal(b, &BatchconversationeventrequestMap)
	if err != nil {
		return err
	}
	
	if EndTransferEvents, ok := BatchconversationeventrequestMap["endTransferEvents"].([]interface{}); ok {
		EndTransferEventsString, _ := json.Marshal(EndTransferEvents)
		json.Unmarshal(EndTransferEventsString, &o.EndTransferEvents)
	}
	
	if PhoneTransferEvents, ok := BatchconversationeventrequestMap["phoneTransferEvents"].([]interface{}); ok {
		PhoneTransferEventsString, _ := json.Marshal(PhoneTransferEvents)
		json.Unmarshal(PhoneTransferEventsString, &o.PhoneTransferEvents)
	}
	
	if ProgressTransferEvents, ok := BatchconversationeventrequestMap["progressTransferEvents"].([]interface{}); ok {
		ProgressTransferEventsString, _ := json.Marshal(ProgressTransferEvents)
		json.Unmarshal(ProgressTransferEventsString, &o.ProgressTransferEvents)
	}
	
	if RoutingTransferEvents, ok := BatchconversationeventrequestMap["routingTransferEvents"].([]interface{}); ok {
		RoutingTransferEventsString, _ := json.Marshal(RoutingTransferEvents)
		json.Unmarshal(RoutingTransferEventsString, &o.RoutingTransferEvents)
	}
	
	if UserTransferEvents, ok := BatchconversationeventrequestMap["userTransferEvents"].([]interface{}); ok {
		UserTransferEventsString, _ := json.Marshal(UserTransferEvents)
		json.Unmarshal(UserTransferEventsString, &o.UserTransferEvents)
	}
	
	if CommunicationAnsweredEvents, ok := BatchconversationeventrequestMap["communicationAnsweredEvents"].([]interface{}); ok {
		CommunicationAnsweredEventsString, _ := json.Marshal(CommunicationAnsweredEvents)
		json.Unmarshal(CommunicationAnsweredEventsString, &o.CommunicationAnsweredEvents)
	}
	
	if CommunicationDispositionAppliedEvents, ok := BatchconversationeventrequestMap["communicationDispositionAppliedEvents"].([]interface{}); ok {
		CommunicationDispositionAppliedEventsString, _ := json.Marshal(CommunicationDispositionAppliedEvents)
		json.Unmarshal(CommunicationDispositionAppliedEventsString, &o.CommunicationDispositionAppliedEvents)
	}
	
	if HoldUpdatedEvents, ok := BatchconversationeventrequestMap["holdUpdatedEvents"].([]interface{}); ok {
		HoldUpdatedEventsString, _ := json.Marshal(HoldUpdatedEvents)
		json.Unmarshal(HoldUpdatedEventsString, &o.HoldUpdatedEvents)
	}
	
	if ExternalEstablishedEvents, ok := BatchconversationeventrequestMap["externalEstablishedEvents"].([]interface{}); ok {
		ExternalEstablishedEventsString, _ := json.Marshal(ExternalEstablishedEvents)
		json.Unmarshal(ExternalEstablishedEventsString, &o.ExternalEstablishedEvents)
	}
	
	if IvrEstablishedEvents, ok := BatchconversationeventrequestMap["ivrEstablishedEvents"].([]interface{}); ok {
		IvrEstablishedEventsString, _ := json.Marshal(IvrEstablishedEvents)
		json.Unmarshal(IvrEstablishedEventsString, &o.IvrEstablishedEvents)
	}
	
	if PhoneEstablishedEvents, ok := BatchconversationeventrequestMap["phoneEstablishedEvents"].([]interface{}); ok {
		PhoneEstablishedEventsString, _ := json.Marshal(PhoneEstablishedEvents)
		json.Unmarshal(PhoneEstablishedEventsString, &o.PhoneEstablishedEvents)
	}
	
	if RoutingEstablishedEvents, ok := BatchconversationeventrequestMap["routingEstablishedEvents"].([]interface{}); ok {
		RoutingEstablishedEventsString, _ := json.Marshal(RoutingEstablishedEvents)
		json.Unmarshal(RoutingEstablishedEventsString, &o.RoutingEstablishedEvents)
	}
	
	if UserEstablishedEvents, ok := BatchconversationeventrequestMap["userEstablishedEvents"].([]interface{}); ok {
		UserEstablishedEventsString, _ := json.Marshal(UserEstablishedEvents)
		json.Unmarshal(UserEstablishedEventsString, &o.UserEstablishedEvents)
	}
	
	if AudioUpdatedEvents, ok := BatchconversationeventrequestMap["audioUpdatedEvents"].([]interface{}); ok {
		AudioUpdatedEventsString, _ := json.Marshal(AudioUpdatedEvents)
		json.Unmarshal(AudioUpdatedEventsString, &o.AudioUpdatedEvents)
	}
	
	if CommunicationEndedEvents, ok := BatchconversationeventrequestMap["communicationEndedEvents"].([]interface{}); ok {
		CommunicationEndedEventsString, _ := json.Marshal(CommunicationEndedEvents)
		json.Unmarshal(CommunicationEndedEventsString, &o.CommunicationEndedEvents)
	}
	
	if ConsultTransferEvents, ok := BatchconversationeventrequestMap["consultTransferEvents"].([]interface{}); ok {
		ConsultTransferEventsString, _ := json.Marshal(ConsultTransferEvents)
		json.Unmarshal(ConsultTransferEventsString, &o.ConsultTransferEvents)
	}
	
	if ProgressConsultTransferEvents, ok := BatchconversationeventrequestMap["progressConsultTransferEvents"].([]interface{}); ok {
		ProgressConsultTransferEventsString, _ := json.Marshal(ProgressConsultTransferEvents)
		json.Unmarshal(ProgressConsultTransferEventsString, &o.ProgressConsultTransferEvents)
	}
	
	if EndConsultTransferEvents, ok := BatchconversationeventrequestMap["endConsultTransferEvents"].([]interface{}); ok {
		EndConsultTransferEventsString, _ := json.Marshal(EndConsultTransferEvents)
		json.Unmarshal(EndConsultTransferEventsString, &o.EndConsultTransferEvents)
	}
	
	if EmailBeginTransmittingEvents, ok := BatchconversationeventrequestMap["emailBeginTransmittingEvents"].([]interface{}); ok {
		EmailBeginTransmittingEventsString, _ := json.Marshal(EmailBeginTransmittingEvents)
		json.Unmarshal(EmailBeginTransmittingEventsString, &o.EmailBeginTransmittingEvents)
	}
	
	if EmailCommunicationEndedEvents, ok := BatchconversationeventrequestMap["emailCommunicationEndedEvents"].([]interface{}); ok {
		EmailCommunicationEndedEventsString, _ := json.Marshal(EmailCommunicationEndedEvents)
		json.Unmarshal(EmailCommunicationEndedEventsString, &o.EmailCommunicationEndedEvents)
	}
	
	if EmailExternalEstablishedEvents, ok := BatchconversationeventrequestMap["emailExternalEstablishedEvents"].([]interface{}); ok {
		EmailExternalEstablishedEventsString, _ := json.Marshal(EmailExternalEstablishedEvents)
		json.Unmarshal(EmailExternalEstablishedEventsString, &o.EmailExternalEstablishedEvents)
	}
	
	if EmailFlowEstablishedEvents, ok := BatchconversationeventrequestMap["emailFlowEstablishedEvents"].([]interface{}); ok {
		EmailFlowEstablishedEventsString, _ := json.Marshal(EmailFlowEstablishedEvents)
		json.Unmarshal(EmailFlowEstablishedEventsString, &o.EmailFlowEstablishedEvents)
	}
	
	if EmailRoutingEstablishedEvents, ok := BatchconversationeventrequestMap["emailRoutingEstablishedEvents"].([]interface{}); ok {
		EmailRoutingEstablishedEventsString, _ := json.Marshal(EmailRoutingEstablishedEvents)
		json.Unmarshal(EmailRoutingEstablishedEventsString, &o.EmailRoutingEstablishedEvents)
	}
	
	if EmailUserEstablishedEvents, ok := BatchconversationeventrequestMap["emailUserEstablishedEvents"].([]interface{}); ok {
		EmailUserEstablishedEventsString, _ := json.Marshal(EmailUserEstablishedEvents)
		json.Unmarshal(EmailUserEstablishedEventsString, &o.EmailUserEstablishedEvents)
	}
	
	if EmailCommunicationAnsweredEvents, ok := BatchconversationeventrequestMap["emailCommunicationAnsweredEvents"].([]interface{}); ok {
		EmailCommunicationAnsweredEventsString, _ := json.Marshal(EmailCommunicationAnsweredEvents)
		json.Unmarshal(EmailCommunicationAnsweredEventsString, &o.EmailCommunicationAnsweredEvents)
	}
	
	if EmailCommunicationDispositionAppliedEvents, ok := BatchconversationeventrequestMap["emailCommunicationDispositionAppliedEvents"].([]interface{}); ok {
		EmailCommunicationDispositionAppliedEventsString, _ := json.Marshal(EmailCommunicationDispositionAppliedEvents)
		json.Unmarshal(EmailCommunicationDispositionAppliedEventsString, &o.EmailCommunicationDispositionAppliedEvents)
	}
	
	if EmailCommunicationSentMessageEvents, ok := BatchconversationeventrequestMap["emailCommunicationSentMessageEvents"].([]interface{}); ok {
		EmailCommunicationSentMessageEventsString, _ := json.Marshal(EmailCommunicationSentMessageEvents)
		json.Unmarshal(EmailCommunicationSentMessageEventsString, &o.EmailCommunicationSentMessageEvents)
	}
	
	if EmailHoldUpdatedEvents, ok := BatchconversationeventrequestMap["emailHoldUpdatedEvents"].([]interface{}); ok {
		EmailHoldUpdatedEventsString, _ := json.Marshal(EmailHoldUpdatedEvents)
		json.Unmarshal(EmailHoldUpdatedEventsString, &o.EmailHoldUpdatedEvents)
	}
	
	if EmailEndTransferEvents, ok := BatchconversationeventrequestMap["emailEndTransferEvents"].([]interface{}); ok {
		EmailEndTransferEventsString, _ := json.Marshal(EmailEndTransferEvents)
		json.Unmarshal(EmailEndTransferEventsString, &o.EmailEndTransferEvents)
	}
	
	if EmailProgressTransferEvents, ok := BatchconversationeventrequestMap["emailProgressTransferEvents"].([]interface{}); ok {
		EmailProgressTransferEventsString, _ := json.Marshal(EmailProgressTransferEvents)
		json.Unmarshal(EmailProgressTransferEventsString, &o.EmailProgressTransferEvents)
	}
	
	if EmailRoutingTransferEvents, ok := BatchconversationeventrequestMap["emailRoutingTransferEvents"].([]interface{}); ok {
		EmailRoutingTransferEventsString, _ := json.Marshal(EmailRoutingTransferEvents)
		json.Unmarshal(EmailRoutingTransferEventsString, &o.EmailRoutingTransferEvents)
	}
	
	if EmailUserTransferEvents, ok := BatchconversationeventrequestMap["emailUserTransferEvents"].([]interface{}); ok {
		EmailUserTransferEventsString, _ := json.Marshal(EmailUserTransferEvents)
		json.Unmarshal(EmailUserTransferEventsString, &o.EmailUserTransferEvents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Batchconversationeventrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
