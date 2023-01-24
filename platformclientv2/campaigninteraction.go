package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Campaigninteraction
type Campaigninteraction struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Campaign
	Campaign *Domainentityref `json:"campaign,omitempty"`

	// Agent
	Agent *Domainentityref `json:"agent,omitempty"`

	// Contact
	Contact *Domainentityref `json:"contact,omitempty"`

	// DestinationAddress
	DestinationAddress *string `json:"destinationAddress,omitempty"`

	// ActivePreviewCall - Boolean value if there is an active preview call on the interaction
	ActivePreviewCall *bool `json:"activePreviewCall,omitempty"`

	// LastActivePreviewWrapupTime - The time when the last preview of the interaction was wrapped up. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	LastActivePreviewWrapupTime *time.Time `json:"lastActivePreviewWrapupTime,omitempty"`

	// CreationTime - The time when dialer created the interaction. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreationTime *time.Time `json:"creationTime,omitempty"`

	// CallPlacedTime - The time when the agent or system places the call. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CallPlacedTime *time.Time `json:"callPlacedTime,omitempty"`

	// CallRoutedTime - The time when the agent was connected to the call. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CallRoutedTime *time.Time `json:"callRoutedTime,omitempty"`

	// PreviewConnectedTime - The time when the customer and routing participant are connected. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	PreviewConnectedTime *time.Time `json:"previewConnectedTime,omitempty"`

	// Queue
	Queue *Domainentityref `json:"queue,omitempty"`

	// Script
	Script *Domainentityref `json:"script,omitempty"`

	// Disposition - Describes what happened with call analysis for instance: disposition.classification.callable.person, disposition.classification.callable.noanswer
	Disposition *string `json:"disposition,omitempty"`

	// CallerName
	CallerName *string `json:"callerName,omitempty"`

	// CallerAddress
	CallerAddress *string `json:"callerAddress,omitempty"`

	// PreviewPopDeliveredTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	PreviewPopDeliveredTime *time.Time `json:"previewPopDeliveredTime,omitempty"`

	// Conversation
	Conversation *Conversationbasic `json:"conversation,omitempty"`

	// DialerSystemParticipantId - conversation participant id that is the dialer system participant to monitor the call from dialer perspective
	DialerSystemParticipantId *string `json:"dialerSystemParticipantId,omitempty"`

	// DialingMode
	DialingMode *string `json:"dialingMode,omitempty"`

	// Skills - Any skills that are attached to the call for routing
	Skills *[]Domainentityref `json:"skills,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Campaigninteraction) SetField(field string, fieldValue interface{}) {
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

func (o Campaigninteraction) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "LastActivePreviewWrapupTime","CreationTime","CallPlacedTime","CallRoutedTime","PreviewConnectedTime","PreviewPopDeliveredTime", }
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
	type Alias Campaigninteraction
	
	LastActivePreviewWrapupTime := new(string)
	if o.LastActivePreviewWrapupTime != nil {
		
		*LastActivePreviewWrapupTime = timeutil.Strftime(o.LastActivePreviewWrapupTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		LastActivePreviewWrapupTime = nil
	}
	
	CreationTime := new(string)
	if o.CreationTime != nil {
		
		*CreationTime = timeutil.Strftime(o.CreationTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreationTime = nil
	}
	
	CallPlacedTime := new(string)
	if o.CallPlacedTime != nil {
		
		*CallPlacedTime = timeutil.Strftime(o.CallPlacedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CallPlacedTime = nil
	}
	
	CallRoutedTime := new(string)
	if o.CallRoutedTime != nil {
		
		*CallRoutedTime = timeutil.Strftime(o.CallRoutedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CallRoutedTime = nil
	}
	
	PreviewConnectedTime := new(string)
	if o.PreviewConnectedTime != nil {
		
		*PreviewConnectedTime = timeutil.Strftime(o.PreviewConnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		PreviewConnectedTime = nil
	}
	
	PreviewPopDeliveredTime := new(string)
	if o.PreviewPopDeliveredTime != nil {
		
		*PreviewPopDeliveredTime = timeutil.Strftime(o.PreviewPopDeliveredTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		PreviewPopDeliveredTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Campaign *Domainentityref `json:"campaign,omitempty"`
		
		Agent *Domainentityref `json:"agent,omitempty"`
		
		Contact *Domainentityref `json:"contact,omitempty"`
		
		DestinationAddress *string `json:"destinationAddress,omitempty"`
		
		ActivePreviewCall *bool `json:"activePreviewCall,omitempty"`
		
		LastActivePreviewWrapupTime *string `json:"lastActivePreviewWrapupTime,omitempty"`
		
		CreationTime *string `json:"creationTime,omitempty"`
		
		CallPlacedTime *string `json:"callPlacedTime,omitempty"`
		
		CallRoutedTime *string `json:"callRoutedTime,omitempty"`
		
		PreviewConnectedTime *string `json:"previewConnectedTime,omitempty"`
		
		Queue *Domainentityref `json:"queue,omitempty"`
		
		Script *Domainentityref `json:"script,omitempty"`
		
		Disposition *string `json:"disposition,omitempty"`
		
		CallerName *string `json:"callerName,omitempty"`
		
		CallerAddress *string `json:"callerAddress,omitempty"`
		
		PreviewPopDeliveredTime *string `json:"previewPopDeliveredTime,omitempty"`
		
		Conversation *Conversationbasic `json:"conversation,omitempty"`
		
		DialerSystemParticipantId *string `json:"dialerSystemParticipantId,omitempty"`
		
		DialingMode *string `json:"dialingMode,omitempty"`
		
		Skills *[]Domainentityref `json:"skills,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Campaign: o.Campaign,
		
		Agent: o.Agent,
		
		Contact: o.Contact,
		
		DestinationAddress: o.DestinationAddress,
		
		ActivePreviewCall: o.ActivePreviewCall,
		
		LastActivePreviewWrapupTime: LastActivePreviewWrapupTime,
		
		CreationTime: CreationTime,
		
		CallPlacedTime: CallPlacedTime,
		
		CallRoutedTime: CallRoutedTime,
		
		PreviewConnectedTime: PreviewConnectedTime,
		
		Queue: o.Queue,
		
		Script: o.Script,
		
		Disposition: o.Disposition,
		
		CallerName: o.CallerName,
		
		CallerAddress: o.CallerAddress,
		
		PreviewPopDeliveredTime: PreviewPopDeliveredTime,
		
		Conversation: o.Conversation,
		
		DialerSystemParticipantId: o.DialerSystemParticipantId,
		
		DialingMode: o.DialingMode,
		
		Skills: o.Skills,
		Alias:    (Alias)(o),
	})
}

func (o *Campaigninteraction) UnmarshalJSON(b []byte) error {
	var CampaigninteractionMap map[string]interface{}
	err := json.Unmarshal(b, &CampaigninteractionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CampaigninteractionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Campaign, ok := CampaigninteractionMap["campaign"].(map[string]interface{}); ok {
		CampaignString, _ := json.Marshal(Campaign)
		json.Unmarshal(CampaignString, &o.Campaign)
	}
	
	if Agent, ok := CampaigninteractionMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if Contact, ok := CampaigninteractionMap["contact"].(map[string]interface{}); ok {
		ContactString, _ := json.Marshal(Contact)
		json.Unmarshal(ContactString, &o.Contact)
	}
	
	if DestinationAddress, ok := CampaigninteractionMap["destinationAddress"].(string); ok {
		o.DestinationAddress = &DestinationAddress
	}
    
	if ActivePreviewCall, ok := CampaigninteractionMap["activePreviewCall"].(bool); ok {
		o.ActivePreviewCall = &ActivePreviewCall
	}
    
	if lastActivePreviewWrapupTimeString, ok := CampaigninteractionMap["lastActivePreviewWrapupTime"].(string); ok {
		LastActivePreviewWrapupTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", lastActivePreviewWrapupTimeString)
		o.LastActivePreviewWrapupTime = &LastActivePreviewWrapupTime
	}
	
	if creationTimeString, ok := CampaigninteractionMap["creationTime"].(string); ok {
		CreationTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", creationTimeString)
		o.CreationTime = &CreationTime
	}
	
	if callPlacedTimeString, ok := CampaigninteractionMap["callPlacedTime"].(string); ok {
		CallPlacedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", callPlacedTimeString)
		o.CallPlacedTime = &CallPlacedTime
	}
	
	if callRoutedTimeString, ok := CampaigninteractionMap["callRoutedTime"].(string); ok {
		CallRoutedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", callRoutedTimeString)
		o.CallRoutedTime = &CallRoutedTime
	}
	
	if previewConnectedTimeString, ok := CampaigninteractionMap["previewConnectedTime"].(string); ok {
		PreviewConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", previewConnectedTimeString)
		o.PreviewConnectedTime = &PreviewConnectedTime
	}
	
	if Queue, ok := CampaigninteractionMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Script, ok := CampaigninteractionMap["script"].(map[string]interface{}); ok {
		ScriptString, _ := json.Marshal(Script)
		json.Unmarshal(ScriptString, &o.Script)
	}
	
	if Disposition, ok := CampaigninteractionMap["disposition"].(string); ok {
		o.Disposition = &Disposition
	}
    
	if CallerName, ok := CampaigninteractionMap["callerName"].(string); ok {
		o.CallerName = &CallerName
	}
    
	if CallerAddress, ok := CampaigninteractionMap["callerAddress"].(string); ok {
		o.CallerAddress = &CallerAddress
	}
    
	if previewPopDeliveredTimeString, ok := CampaigninteractionMap["previewPopDeliveredTime"].(string); ok {
		PreviewPopDeliveredTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", previewPopDeliveredTimeString)
		o.PreviewPopDeliveredTime = &PreviewPopDeliveredTime
	}
	
	if Conversation, ok := CampaigninteractionMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if DialerSystemParticipantId, ok := CampaigninteractionMap["dialerSystemParticipantId"].(string); ok {
		o.DialerSystemParticipantId = &DialerSystemParticipantId
	}
    
	if DialingMode, ok := CampaigninteractionMap["dialingMode"].(string); ok {
		o.DialingMode = &DialingMode
	}
    
	if Skills, ok := CampaigninteractionMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Campaigninteraction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
