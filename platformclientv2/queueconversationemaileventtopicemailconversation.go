package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationemaileventtopicemailconversation
type Queueconversationemaileventtopicemailconversation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Participants
	Participants *[]Queueconversationemaileventtopicemailmediaparticipant `json:"participants,omitempty"`

	// OtherMediaUris
	OtherMediaUris *[]string `json:"otherMediaUris,omitempty"`

	// Address
	Address *string `json:"address,omitempty"`

	// UtilizationLabelId
	UtilizationLabelId *string `json:"utilizationLabelId,omitempty"`

	// Divisions
	Divisions *[]Queueconversationemaileventtopicconversationdivisionmembership `json:"divisions,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Queueconversationemaileventtopicemailconversation) SetField(field string, fieldValue interface{}) {
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

func (o Queueconversationemaileventtopicemailconversation) MarshalJSON() ([]byte, error) {
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
	type Alias Queueconversationemaileventtopicemailconversation
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Participants *[]Queueconversationemaileventtopicemailmediaparticipant `json:"participants,omitempty"`
		
		OtherMediaUris *[]string `json:"otherMediaUris,omitempty"`
		
		Address *string `json:"address,omitempty"`
		
		UtilizationLabelId *string `json:"utilizationLabelId,omitempty"`
		
		Divisions *[]Queueconversationemaileventtopicconversationdivisionmembership `json:"divisions,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Participants: o.Participants,
		
		OtherMediaUris: o.OtherMediaUris,
		
		Address: o.Address,
		
		UtilizationLabelId: o.UtilizationLabelId,
		
		Divisions: o.Divisions,
		Alias:    (Alias)(o),
	})
}

func (o *Queueconversationemaileventtopicemailconversation) UnmarshalJSON(b []byte) error {
	var QueueconversationemaileventtopicemailconversationMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationemaileventtopicemailconversationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationemaileventtopicemailconversationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := QueueconversationemaileventtopicemailconversationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Participants, ok := QueueconversationemaileventtopicemailconversationMap["participants"].([]interface{}); ok {
		ParticipantsString, _ := json.Marshal(Participants)
		json.Unmarshal(ParticipantsString, &o.Participants)
	}
	
	if OtherMediaUris, ok := QueueconversationemaileventtopicemailconversationMap["otherMediaUris"].([]interface{}); ok {
		OtherMediaUrisString, _ := json.Marshal(OtherMediaUris)
		json.Unmarshal(OtherMediaUrisString, &o.OtherMediaUris)
	}
	
	if Address, ok := QueueconversationemaileventtopicemailconversationMap["address"].(string); ok {
		o.Address = &Address
	}
    
	if UtilizationLabelId, ok := QueueconversationemaileventtopicemailconversationMap["utilizationLabelId"].(string); ok {
		o.UtilizationLabelId = &UtilizationLabelId
	}
    
	if Divisions, ok := QueueconversationemaileventtopicemailconversationMap["divisions"].([]interface{}); ok {
		DivisionsString, _ := json.Marshal(Divisions)
		json.Unmarshal(DivisionsString, &o.Divisions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationemaileventtopicemailconversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
