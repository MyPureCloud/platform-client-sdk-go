package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Opensocialmediachannel - Channel-specific information that describes the message and the message channel/provider.
type Opensocialmediachannel struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The topic ID.
	Id *string `json:"id,omitempty"`

	// Platform - The provider type.
	Platform *string `json:"platform,omitempty"`

	// VarType - Specifies if this message is part of a private or public conversation.
	VarType *string `json:"type,omitempty"`

	// MessageId - Unique provider ID of the message such as a Facebook message ID.
	MessageId *string `json:"messageId,omitempty"`

	// To - Information about the recipient the message is sent to.
	To *Opensocialmediarecipient `json:"to,omitempty"`

	// From - Information about the recipient the message is received from.
	From *Opensocialmediarecipient `json:"from,omitempty"`

	// Time - Original time of the event. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Time *time.Time `json:"time,omitempty"`

	// Metadata - Information about the channel.
	Metadata *interface{} `json:"metadata,omitempty"`

	// PublicMetadata - Meta data of this public post. For example, used to define where in the thread the post exists.
	PublicMetadata *Opensocialmediapublicmetadata `json:"publicMetadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Opensocialmediachannel) SetField(field string, fieldValue interface{}) {
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

func (o Opensocialmediachannel) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "Time", }
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
	type Alias Opensocialmediachannel
	
	Time := new(string)
	if o.Time != nil {
		
		*Time = timeutil.Strftime(o.Time, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Time = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Platform *string `json:"platform,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		MessageId *string `json:"messageId,omitempty"`
		
		To *Opensocialmediarecipient `json:"to,omitempty"`
		
		From *Opensocialmediarecipient `json:"from,omitempty"`
		
		Time *string `json:"time,omitempty"`
		
		Metadata *interface{} `json:"metadata,omitempty"`
		
		PublicMetadata *Opensocialmediapublicmetadata `json:"publicMetadata,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Platform: o.Platform,
		
		VarType: o.VarType,
		
		MessageId: o.MessageId,
		
		To: o.To,
		
		From: o.From,
		
		Time: Time,
		
		Metadata: o.Metadata,
		
		PublicMetadata: o.PublicMetadata,
		Alias:    (Alias)(o),
	})
}

func (o *Opensocialmediachannel) UnmarshalJSON(b []byte) error {
	var OpensocialmediachannelMap map[string]interface{}
	err := json.Unmarshal(b, &OpensocialmediachannelMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OpensocialmediachannelMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Platform, ok := OpensocialmediachannelMap["platform"].(string); ok {
		o.Platform = &Platform
	}
    
	if VarType, ok := OpensocialmediachannelMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if MessageId, ok := OpensocialmediachannelMap["messageId"].(string); ok {
		o.MessageId = &MessageId
	}
    
	if To, ok := OpensocialmediachannelMap["to"].(map[string]interface{}); ok {
		ToString, _ := json.Marshal(To)
		json.Unmarshal(ToString, &o.To)
	}
	
	if From, ok := OpensocialmediachannelMap["from"].(map[string]interface{}); ok {
		FromString, _ := json.Marshal(From)
		json.Unmarshal(FromString, &o.From)
	}
	
	if timeString, ok := OpensocialmediachannelMap["time"].(string); ok {
		Time, _ := time.Parse("2006-01-02T15:04:05.999999Z", timeString)
		o.Time = &Time
	}
	
	if Metadata, ok := OpensocialmediachannelMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if PublicMetadata, ok := OpensocialmediachannelMap["publicMetadata"].(map[string]interface{}); ok {
		PublicMetadataString, _ := json.Marshal(PublicMetadata)
		json.Unmarshal(PublicMetadataString, &o.PublicMetadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Opensocialmediachannel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
