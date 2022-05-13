package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createcallbackonconversationcommand
type Createcallbackonconversationcommand struct { 
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

func (o *Createcallbackonconversationcommand) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createcallbackonconversationcommand
	
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
		*Alias
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
		Alias:    (*Alias)(o),
	})
}

func (o *Createcallbackonconversationcommand) UnmarshalJSON(b []byte) error {
	var CreatecallbackonconversationcommandMap map[string]interface{}
	err := json.Unmarshal(b, &CreatecallbackonconversationcommandMap)
	if err != nil {
		return err
	}
	
	if ScriptId, ok := CreatecallbackonconversationcommandMap["scriptId"].(string); ok {
		o.ScriptId = &ScriptId
	}
    
	if QueueId, ok := CreatecallbackonconversationcommandMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if RoutingData, ok := CreatecallbackonconversationcommandMap["routingData"].(map[string]interface{}); ok {
		RoutingDataString, _ := json.Marshal(RoutingData)
		json.Unmarshal(RoutingDataString, &o.RoutingData)
	}
	
	if CallbackUserName, ok := CreatecallbackonconversationcommandMap["callbackUserName"].(string); ok {
		o.CallbackUserName = &CallbackUserName
	}
    
	if CallbackNumbers, ok := CreatecallbackonconversationcommandMap["callbackNumbers"].([]interface{}); ok {
		CallbackNumbersString, _ := json.Marshal(CallbackNumbers)
		json.Unmarshal(CallbackNumbersString, &o.CallbackNumbers)
	}
	
	if callbackScheduledTimeString, ok := CreatecallbackonconversationcommandMap["callbackScheduledTime"].(string); ok {
		CallbackScheduledTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", callbackScheduledTimeString)
		o.CallbackScheduledTime = &CallbackScheduledTime
	}
	
	if CountryCode, ok := CreatecallbackonconversationcommandMap["countryCode"].(string); ok {
		o.CountryCode = &CountryCode
	}
    
	if ValidateCallbackNumbers, ok := CreatecallbackonconversationcommandMap["validateCallbackNumbers"].(bool); ok {
		o.ValidateCallbackNumbers = &ValidateCallbackNumbers
	}
    
	if Data, ok := CreatecallbackonconversationcommandMap["data"].(map[string]interface{}); ok {
		DataString, _ := json.Marshal(Data)
		json.Unmarshal(DataString, &o.Data)
	}
	
	if CallerId, ok := CreatecallbackonconversationcommandMap["callerId"].(string); ok {
		o.CallerId = &CallerId
	}
    
	if CallerIdName, ok := CreatecallbackonconversationcommandMap["callerIdName"].(string); ok {
		o.CallerIdName = &CallerIdName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createcallbackonconversationcommand) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
