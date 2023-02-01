package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createcallbackcommand
type Createcallbackcommand struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ScriptId - The identifier of the script to be used for the callback
	ScriptId *string `json:"scriptId,omitempty"`

	// QueueId - The identifier of the queue to be used for the callback. Either queueId or routingData is required.
	QueueId *string `json:"queueId,omitempty"`

	// RoutingData - The routing data to be used for the callback. Either queueId or routingData is required.
	RoutingData *Routingdata `json:"routingData,omitempty"`

	// CallbackUserName - The name of the party to be called back.
	CallbackUserName *string `json:"callbackUserName,omitempty"`

	// CallbackNumbers - A list of phone numbers for the callback.
	CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`

	// CallbackScheduledTime - The scheduled date-time for the callback as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	CallbackScheduledTime *time.Time `json:"callbackScheduledTime,omitempty"`

	// CountryCode - The country code to be associated with the callback numbers.
	CountryCode *string `json:"countryCode,omitempty"`

	// ValidateCallbackNumbers - Whether or not to validate the callback numbers for phone number format.
	ValidateCallbackNumbers *bool `json:"validateCallbackNumbers,omitempty"`

	// Data - A map of key-value pairs containing additional data that can be associated to the callback. These values will appear in the attributes property on the conversation participant. Example: { \"notes\": \"ready to close the deal!\", \"customerPreferredName\": \"Doc\" }
	Data *map[string]string `json:"data,omitempty"`

	// CallerId - The phone number displayed to recipients when a phone call is placed as part of the callback. Must conform to the E.164 format. May be overridden by other settings in the system such as external trunk settings. Telco support for \"callerId\" varies.
	CallerId *string `json:"callerId,omitempty"`

	// CallerIdName - The name displayed to recipients when a phone call is placed as part of the callback. May be overridden by other settings in the system such as external trunk settings. Telco support for \"callerIdName\" varies.
	CallerIdName *string `json:"callerIdName,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createcallbackcommand) SetField(field string, fieldValue interface{}) {
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

func (o Createcallbackcommand) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CallbackScheduledTime", }
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
	type Alias Createcallbackcommand
	
	CallbackScheduledTime := new(string)
	if o.CallbackScheduledTime != nil {
		
		*CallbackScheduledTime = timeutil.Strftime(o.CallbackScheduledTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CallbackScheduledTime = nil
	}
	
	return json.Marshal(&struct { 
		ScriptId *string `json:"scriptId,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		RoutingData *Routingdata `json:"routingData,omitempty"`
		
		CallbackUserName *string `json:"callbackUserName,omitempty"`
		
		CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`
		
		CallbackScheduledTime *string `json:"callbackScheduledTime,omitempty"`
		
		CountryCode *string `json:"countryCode,omitempty"`
		
		ValidateCallbackNumbers *bool `json:"validateCallbackNumbers,omitempty"`
		
		Data *map[string]string `json:"data,omitempty"`
		
		CallerId *string `json:"callerId,omitempty"`
		
		CallerIdName *string `json:"callerIdName,omitempty"`
		Alias
	}{ 
		ScriptId: o.ScriptId,
		
		QueueId: o.QueueId,
		
		RoutingData: o.RoutingData,
		
		CallbackUserName: o.CallbackUserName,
		
		CallbackNumbers: o.CallbackNumbers,
		
		CallbackScheduledTime: CallbackScheduledTime,
		
		CountryCode: o.CountryCode,
		
		ValidateCallbackNumbers: o.ValidateCallbackNumbers,
		
		Data: o.Data,
		
		CallerId: o.CallerId,
		
		CallerIdName: o.CallerIdName,
		Alias:    (Alias)(o),
	})
}

func (o *Createcallbackcommand) UnmarshalJSON(b []byte) error {
	var CreatecallbackcommandMap map[string]interface{}
	err := json.Unmarshal(b, &CreatecallbackcommandMap)
	if err != nil {
		return err
	}
	
	if ScriptId, ok := CreatecallbackcommandMap["scriptId"].(string); ok {
		o.ScriptId = &ScriptId
	}
    
	if QueueId, ok := CreatecallbackcommandMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if RoutingData, ok := CreatecallbackcommandMap["routingData"].(map[string]interface{}); ok {
		RoutingDataString, _ := json.Marshal(RoutingData)
		json.Unmarshal(RoutingDataString, &o.RoutingData)
	}
	
	if CallbackUserName, ok := CreatecallbackcommandMap["callbackUserName"].(string); ok {
		o.CallbackUserName = &CallbackUserName
	}
    
	if CallbackNumbers, ok := CreatecallbackcommandMap["callbackNumbers"].([]interface{}); ok {
		CallbackNumbersString, _ := json.Marshal(CallbackNumbers)
		json.Unmarshal(CallbackNumbersString, &o.CallbackNumbers)
	}
	
	if callbackScheduledTimeString, ok := CreatecallbackcommandMap["callbackScheduledTime"].(string); ok {
		CallbackScheduledTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", callbackScheduledTimeString)
		o.CallbackScheduledTime = &CallbackScheduledTime
	}
	
	if CountryCode, ok := CreatecallbackcommandMap["countryCode"].(string); ok {
		o.CountryCode = &CountryCode
	}
    
	if ValidateCallbackNumbers, ok := CreatecallbackcommandMap["validateCallbackNumbers"].(bool); ok {
		o.ValidateCallbackNumbers = &ValidateCallbackNumbers
	}
    
	if Data, ok := CreatecallbackcommandMap["data"].(map[string]interface{}); ok {
		DataString, _ := json.Marshal(Data)
		json.Unmarshal(DataString, &o.Data)
	}
	
	if CallerId, ok := CreatecallbackcommandMap["callerId"].(string); ok {
		o.CallerId = &CallerId
	}
    
	if CallerIdName, ok := CreatecallbackcommandMap["callerIdName"].(string); ok {
		o.CallerIdName = &CallerIdName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createcallbackcommand) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
