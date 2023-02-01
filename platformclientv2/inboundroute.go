package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Inboundroute
type Inboundroute struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Pattern - The search pattern that the mailbox name should match.
	Pattern *string `json:"pattern,omitempty"`

	// Queue - The queue to route the emails to.
	Queue *Domainentityref `json:"queue,omitempty"`

	// Priority - The priority to use for routing.
	Priority *int `json:"priority,omitempty"`

	// Skills - The skills to use for routing.
	Skills *[]Domainentityref `json:"skills,omitempty"`

	// Language - The language to use for routing.
	Language *Domainentityref `json:"language,omitempty"`

	// FromName - The sender name to use for outgoing replies.
	FromName *string `json:"fromName,omitempty"`

	// FromEmail - The sender email to use for outgoing replies.
	FromEmail *string `json:"fromEmail,omitempty"`

	// Flow - The flow to use for processing the email.
	Flow *Domainentityref `json:"flow,omitempty"`

	// ReplyEmailAddress - The route to use for email replies.
	ReplyEmailAddress *Queueemailaddress `json:"replyEmailAddress,omitempty"`

	// AutoBcc - The recipients that should be automatically blind copied on outbound emails associated with this InboundRoute.
	AutoBcc *[]Emailaddress `json:"autoBcc,omitempty"`

	// SpamFlow - The flow to use for processing inbound emails that have been marked as spam.
	SpamFlow *Domainentityref `json:"spamFlow,omitempty"`

	// Signature - The configuration for the canned response signature that will be appended to outbound emails sent via this route
	Signature *Signature `json:"signature,omitempty"`

	// HistoryInclusion - The configuration to indicate how the history of a conversation has to be included in a draft
	HistoryInclusion *string `json:"historyInclusion,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Inboundroute) SetField(field string, fieldValue interface{}) {
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

func (o Inboundroute) MarshalJSON() ([]byte, error) {
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
	type Alias Inboundroute
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Pattern *string `json:"pattern,omitempty"`
		
		Queue *Domainentityref `json:"queue,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Domainentityref `json:"skills,omitempty"`
		
		Language *Domainentityref `json:"language,omitempty"`
		
		FromName *string `json:"fromName,omitempty"`
		
		FromEmail *string `json:"fromEmail,omitempty"`
		
		Flow *Domainentityref `json:"flow,omitempty"`
		
		ReplyEmailAddress *Queueemailaddress `json:"replyEmailAddress,omitempty"`
		
		AutoBcc *[]Emailaddress `json:"autoBcc,omitempty"`
		
		SpamFlow *Domainentityref `json:"spamFlow,omitempty"`
		
		Signature *Signature `json:"signature,omitempty"`
		
		HistoryInclusion *string `json:"historyInclusion,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Pattern: o.Pattern,
		
		Queue: o.Queue,
		
		Priority: o.Priority,
		
		Skills: o.Skills,
		
		Language: o.Language,
		
		FromName: o.FromName,
		
		FromEmail: o.FromEmail,
		
		Flow: o.Flow,
		
		ReplyEmailAddress: o.ReplyEmailAddress,
		
		AutoBcc: o.AutoBcc,
		
		SpamFlow: o.SpamFlow,
		
		Signature: o.Signature,
		
		HistoryInclusion: o.HistoryInclusion,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Inboundroute) UnmarshalJSON(b []byte) error {
	var InboundrouteMap map[string]interface{}
	err := json.Unmarshal(b, &InboundrouteMap)
	if err != nil {
		return err
	}
	
	if Id, ok := InboundrouteMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := InboundrouteMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Pattern, ok := InboundrouteMap["pattern"].(string); ok {
		o.Pattern = &Pattern
	}
    
	if Queue, ok := InboundrouteMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Priority, ok := InboundrouteMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if Skills, ok := InboundrouteMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if Language, ok := InboundrouteMap["language"].(map[string]interface{}); ok {
		LanguageString, _ := json.Marshal(Language)
		json.Unmarshal(LanguageString, &o.Language)
	}
	
	if FromName, ok := InboundrouteMap["fromName"].(string); ok {
		o.FromName = &FromName
	}
    
	if FromEmail, ok := InboundrouteMap["fromEmail"].(string); ok {
		o.FromEmail = &FromEmail
	}
    
	if Flow, ok := InboundrouteMap["flow"].(map[string]interface{}); ok {
		FlowString, _ := json.Marshal(Flow)
		json.Unmarshal(FlowString, &o.Flow)
	}
	
	if ReplyEmailAddress, ok := InboundrouteMap["replyEmailAddress"].(map[string]interface{}); ok {
		ReplyEmailAddressString, _ := json.Marshal(ReplyEmailAddress)
		json.Unmarshal(ReplyEmailAddressString, &o.ReplyEmailAddress)
	}
	
	if AutoBcc, ok := InboundrouteMap["autoBcc"].([]interface{}); ok {
		AutoBccString, _ := json.Marshal(AutoBcc)
		json.Unmarshal(AutoBccString, &o.AutoBcc)
	}
	
	if SpamFlow, ok := InboundrouteMap["spamFlow"].(map[string]interface{}); ok {
		SpamFlowString, _ := json.Marshal(SpamFlow)
		json.Unmarshal(SpamFlowString, &o.SpamFlow)
	}
	
	if Signature, ok := InboundrouteMap["signature"].(map[string]interface{}); ok {
		SignatureString, _ := json.Marshal(Signature)
		json.Unmarshal(SignatureString, &o.Signature)
	}
	
	if HistoryInclusion, ok := InboundrouteMap["historyInclusion"].(string); ok {
		o.HistoryInclusion = &HistoryInclusion
	}
    
	if SelfUri, ok := InboundrouteMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Inboundroute) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
