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
	// EndTransferEvents - EndTransfer events for this batch
	EndTransferEvents *[]Endtransferevent `json:"endTransferEvents,omitempty"`

	// PhoneTransferEvents - PhoneTransfer events for this batch
	PhoneTransferEvents *[]Phonetransferevent `json:"phoneTransferEvents,omitempty"`

	// ProgressTransferEvents - ProgressTransfer events for this batch
	ProgressTransferEvents *[]Progresstransferevent `json:"progressTransferEvents,omitempty"`

	// RoutingTransferEvents - RoutingTransfer events for this batch
	RoutingTransferEvents *[]Routingtransferevent `json:"routingTransferEvents,omitempty"`

	// UserTransferEvents - UserTransfer events for this batch
	UserTransferEvents *[]Usertransferevent `json:"userTransferEvents,omitempty"`

	// CommunicationAnsweredEvents - CommunicationAnswered events for this batch
	CommunicationAnsweredEvents *[]Communicationansweredevent `json:"communicationAnsweredEvents,omitempty"`

	// CommunicationDispositionAppliedEvents - CommunicationDispositionApplied events for this batch
	CommunicationDispositionAppliedEvents *[]Communicationdispositionappliedevent `json:"communicationDispositionAppliedEvents,omitempty"`

	// HoldUpdatedEvents - HoldUpdated events for this batch
	HoldUpdatedEvents *[]Holdupdatedevent `json:"holdUpdatedEvents,omitempty"`

	// ExternalEstablishedEvents - ExternalEstablished events for this batch
	ExternalEstablishedEvents *[]Externalestablishedevent `json:"externalEstablishedEvents,omitempty"`

	// IvrEstablishedEvents - IvrEstablished events for this batch
	IvrEstablishedEvents *[]Ivrestablishedevent `json:"ivrEstablishedEvents,omitempty"`

	// PhoneEstablishedEvents - PhoneEstablished events for this batch
	PhoneEstablishedEvents *[]Phoneestablishedevent `json:"phoneEstablishedEvents,omitempty"`

	// RoutingEstablishedEvents - RoutingEstablished events for this batch
	RoutingEstablishedEvents *[]Routingestablishedevent `json:"routingEstablishedEvents,omitempty"`

	// UserEstablishedEvents - UserEstablished events for this batch
	UserEstablishedEvents *[]Userestablishedevent `json:"userEstablishedEvents,omitempty"`

	// AudioUpdatedEvents - AudioUpdated events for this batch
	AudioUpdatedEvents *[]Audioupdatedevent `json:"audioUpdatedEvents,omitempty"`

	// CommunicationEndedEvents - CommunicationEnded events for this batch
	CommunicationEndedEvents *[]Communicationendedevent `json:"communicationEndedEvents,omitempty"`

	// ConsultTransferEvents - ConsultTransfer events for this batch
	ConsultTransferEvents *[]Consulttransferevent `json:"consultTransferEvents,omitempty"`

	// ProgressConsultTransferEvents - ProgressConsultTransfer events for this batch
	ProgressConsultTransferEvents *[]Progressconsulttransferevent `json:"progressConsultTransferEvents,omitempty"`

	// EndConsultTransferEvents - EndConsultTransfer events for this batch
	EndConsultTransferEvents *[]Endconsulttransferevent `json:"endConsultTransferEvents,omitempty"`
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
			if contains(dateTimeFields, fieldName) {
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
	

	return nil
}

// String returns a JSON representation of the model
func (o *Batchconversationeventrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
