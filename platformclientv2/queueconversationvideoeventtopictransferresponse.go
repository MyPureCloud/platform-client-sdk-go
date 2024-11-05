package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationvideoeventtopictransferresponse
type Queueconversationvideoeventtopictransferresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The id of the command.
	Id *string `json:"id,omitempty"`

	// State
	State *string `json:"state,omitempty"`

	// DateIssued - The date/time that this command was issued.
	DateIssued *time.Time `json:"dateIssued,omitempty"`

	// Initiator
	Initiator *Queueconversationvideoeventtopictransferinitiator `json:"initiator,omitempty"`

	// ModifiedBy
	ModifiedBy *Queueconversationvideoeventtopictransfermodifedby `json:"modifiedBy,omitempty"`

	// Destination
	Destination *Queueconversationvideoeventtopictransferdestination `json:"destination,omitempty"`

	// TransferType - The type of transfer to perform.
	TransferType *string `json:"transferType,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Queueconversationvideoeventtopictransferresponse) SetField(field string, fieldValue interface{}) {
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

func (o Queueconversationvideoeventtopictransferresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateIssued", }
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
	type Alias Queueconversationvideoeventtopictransferresponse
	
	DateIssued := new(string)
	if o.DateIssued != nil {
		
		*DateIssued = timeutil.Strftime(o.DateIssued, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateIssued = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		DateIssued *string `json:"dateIssued,omitempty"`
		
		Initiator *Queueconversationvideoeventtopictransferinitiator `json:"initiator,omitempty"`
		
		ModifiedBy *Queueconversationvideoeventtopictransfermodifedby `json:"modifiedBy,omitempty"`
		
		Destination *Queueconversationvideoeventtopictransferdestination `json:"destination,omitempty"`
		
		TransferType *string `json:"transferType,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		State: o.State,
		
		DateIssued: DateIssued,
		
		Initiator: o.Initiator,
		
		ModifiedBy: o.ModifiedBy,
		
		Destination: o.Destination,
		
		TransferType: o.TransferType,
		Alias:    (Alias)(o),
	})
}

func (o *Queueconversationvideoeventtopictransferresponse) UnmarshalJSON(b []byte) error {
	var QueueconversationvideoeventtopictransferresponseMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationvideoeventtopictransferresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationvideoeventtopictransferresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if State, ok := QueueconversationvideoeventtopictransferresponseMap["state"].(string); ok {
		o.State = &State
	}
    
	if dateIssuedString, ok := QueueconversationvideoeventtopictransferresponseMap["dateIssued"].(string); ok {
		DateIssued, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateIssuedString)
		o.DateIssued = &DateIssued
	}
	
	if Initiator, ok := QueueconversationvideoeventtopictransferresponseMap["initiator"].(map[string]interface{}); ok {
		InitiatorString, _ := json.Marshal(Initiator)
		json.Unmarshal(InitiatorString, &o.Initiator)
	}
	
	if ModifiedBy, ok := QueueconversationvideoeventtopictransferresponseMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if Destination, ok := QueueconversationvideoeventtopictransferresponseMap["destination"].(map[string]interface{}); ok {
		DestinationString, _ := json.Marshal(Destination)
		json.Unmarshal(DestinationString, &o.Destination)
	}
	
	if TransferType, ok := QueueconversationvideoeventtopictransferresponseMap["transferType"].(string); ok {
		o.TransferType = &TransferType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopictransferresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
