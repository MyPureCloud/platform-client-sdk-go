package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queuerequest
type Queuerequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The queue name
	Name *string `json:"name,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Writabledivision `json:"division,omitempty"`


	// Description - The queue description.
	Description *string `json:"description,omitempty"`


	// DateCreated - The date the queue was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - The date of the last modification to the queue. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// ModifiedBy - The ID of the user that last modified the queue.
	ModifiedBy *string `json:"modifiedBy,omitempty"`


	// CreatedBy - The ID of the user that created the queue.
	CreatedBy *string `json:"createdBy,omitempty"`


	// MemberCount - The total number of members (joined or unjoined) in the queue.
	MemberCount *int `json:"memberCount,omitempty"`


	// JoinedMemberCount - The number of joined members in the queue.
	JoinedMemberCount *int `json:"joinedMemberCount,omitempty"`


	// MediaSettings - The media settings for the queue. Valid key values: CALL, CALLBACK, CHAT, EMAIL, MESSAGE, SOCIAL_EXPRESSION, VIDEO_COMM
	MediaSettings *map[string]Mediasetting `json:"mediaSettings,omitempty"`


	// RoutingRules - The routing rules for the queue, used for routing to known or preferred agents.
	RoutingRules *[]Routingrule `json:"routingRules,omitempty"`


	// Bullseye - The bulls-eye settings for the queue.
	Bullseye *Bullseye `json:"bullseye,omitempty"`


	// AcwSettings - The ACW settings for the queue.
	AcwSettings *Acwsettings `json:"acwSettings,omitempty"`


	// SkillEvaluationMethod - The skill evaluation method to use when routing conversations.
	SkillEvaluationMethod *string `json:"skillEvaluationMethod,omitempty"`


	// QueueFlow - The in-queue flow to use for conversations waiting in queue.
	QueueFlow *Domainentityref `json:"queueFlow,omitempty"`


	// WhisperPrompt - The prompt used for whisper on the queue, if configured.
	WhisperPrompt *Domainentityref `json:"whisperPrompt,omitempty"`


	// AutoAnswerOnly - Specifies whether the configured whisper should play for all ACD calls, or only for those which are auto-answered.
	AutoAnswerOnly *bool `json:"autoAnswerOnly,omitempty"`


	// EnableTranscription - Indicates whether voice transcription is enabled for this queue.
	EnableTranscription *bool `json:"enableTranscription,omitempty"`


	// EnableManualAssignment - Indicates whether manual assignment is enabled for this queue.
	EnableManualAssignment *bool `json:"enableManualAssignment,omitempty"`


	// CallingPartyName - The name to use for caller identification for outbound calls from this queue.
	CallingPartyName *string `json:"callingPartyName,omitempty"`


	// CallingPartyNumber - The phone number to use for caller identification for outbound calls from this queue.
	CallingPartyNumber *string `json:"callingPartyNumber,omitempty"`


	// DefaultScripts - The default script Ids for the communication types.
	DefaultScripts *map[string]Script `json:"defaultScripts,omitempty"`


	// OutboundMessagingAddresses - The messaging addresses for the queue.
	OutboundMessagingAddresses *Queuemessagingaddresses `json:"outboundMessagingAddresses,omitempty"`


	// OutboundEmailAddress
	OutboundEmailAddress *Queueemailaddress `json:"outboundEmailAddress,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Queuerequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queuerequest
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Writabledivision `json:"division,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ModifiedBy *string `json:"modifiedBy,omitempty"`
		
		CreatedBy *string `json:"createdBy,omitempty"`
		
		MemberCount *int `json:"memberCount,omitempty"`
		
		JoinedMemberCount *int `json:"joinedMemberCount,omitempty"`
		
		MediaSettings *map[string]Mediasetting `json:"mediaSettings,omitempty"`
		
		RoutingRules *[]Routingrule `json:"routingRules,omitempty"`
		
		Bullseye *Bullseye `json:"bullseye,omitempty"`
		
		AcwSettings *Acwsettings `json:"acwSettings,omitempty"`
		
		SkillEvaluationMethod *string `json:"skillEvaluationMethod,omitempty"`
		
		QueueFlow *Domainentityref `json:"queueFlow,omitempty"`
		
		WhisperPrompt *Domainentityref `json:"whisperPrompt,omitempty"`
		
		AutoAnswerOnly *bool `json:"autoAnswerOnly,omitempty"`
		
		EnableTranscription *bool `json:"enableTranscription,omitempty"`
		
		EnableManualAssignment *bool `json:"enableManualAssignment,omitempty"`
		
		CallingPartyName *string `json:"callingPartyName,omitempty"`
		
		CallingPartyNumber *string `json:"callingPartyNumber,omitempty"`
		
		DefaultScripts *map[string]Script `json:"defaultScripts,omitempty"`
		
		OutboundMessagingAddresses *Queuemessagingaddresses `json:"outboundMessagingAddresses,omitempty"`
		
		OutboundEmailAddress *Queueemailaddress `json:"outboundEmailAddress,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		Description: o.Description,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ModifiedBy: o.ModifiedBy,
		
		CreatedBy: o.CreatedBy,
		
		MemberCount: o.MemberCount,
		
		JoinedMemberCount: o.JoinedMemberCount,
		
		MediaSettings: o.MediaSettings,
		
		RoutingRules: o.RoutingRules,
		
		Bullseye: o.Bullseye,
		
		AcwSettings: o.AcwSettings,
		
		SkillEvaluationMethod: o.SkillEvaluationMethod,
		
		QueueFlow: o.QueueFlow,
		
		WhisperPrompt: o.WhisperPrompt,
		
		AutoAnswerOnly: o.AutoAnswerOnly,
		
		EnableTranscription: o.EnableTranscription,
		
		EnableManualAssignment: o.EnableManualAssignment,
		
		CallingPartyName: o.CallingPartyName,
		
		CallingPartyNumber: o.CallingPartyNumber,
		
		DefaultScripts: o.DefaultScripts,
		
		OutboundMessagingAddresses: o.OutboundMessagingAddresses,
		
		OutboundEmailAddress: o.OutboundEmailAddress,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Queuerequest) UnmarshalJSON(b []byte) error {
	var QueuerequestMap map[string]interface{}
	err := json.Unmarshal(b, &QueuerequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueuerequestMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := QueuerequestMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Division, ok := QueuerequestMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := QueuerequestMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if dateCreatedString, ok := QueuerequestMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := QueuerequestMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := QueuerequestMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
	
	if CreatedBy, ok := QueuerequestMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
	
	if MemberCount, ok := QueuerequestMap["memberCount"].(float64); ok {
		MemberCountInt := int(MemberCount)
		o.MemberCount = &MemberCountInt
	}
	
	if JoinedMemberCount, ok := QueuerequestMap["joinedMemberCount"].(float64); ok {
		JoinedMemberCountInt := int(JoinedMemberCount)
		o.JoinedMemberCount = &JoinedMemberCountInt
	}
	
	if MediaSettings, ok := QueuerequestMap["mediaSettings"].(map[string]interface{}); ok {
		MediaSettingsString, _ := json.Marshal(MediaSettings)
		json.Unmarshal(MediaSettingsString, &o.MediaSettings)
	}
	
	if RoutingRules, ok := QueuerequestMap["routingRules"].([]interface{}); ok {
		RoutingRulesString, _ := json.Marshal(RoutingRules)
		json.Unmarshal(RoutingRulesString, &o.RoutingRules)
	}
	
	if Bullseye, ok := QueuerequestMap["bullseye"].(map[string]interface{}); ok {
		BullseyeString, _ := json.Marshal(Bullseye)
		json.Unmarshal(BullseyeString, &o.Bullseye)
	}
	
	if AcwSettings, ok := QueuerequestMap["acwSettings"].(map[string]interface{}); ok {
		AcwSettingsString, _ := json.Marshal(AcwSettings)
		json.Unmarshal(AcwSettingsString, &o.AcwSettings)
	}
	
	if SkillEvaluationMethod, ok := QueuerequestMap["skillEvaluationMethod"].(string); ok {
		o.SkillEvaluationMethod = &SkillEvaluationMethod
	}
	
	if QueueFlow, ok := QueuerequestMap["queueFlow"].(map[string]interface{}); ok {
		QueueFlowString, _ := json.Marshal(QueueFlow)
		json.Unmarshal(QueueFlowString, &o.QueueFlow)
	}
	
	if WhisperPrompt, ok := QueuerequestMap["whisperPrompt"].(map[string]interface{}); ok {
		WhisperPromptString, _ := json.Marshal(WhisperPrompt)
		json.Unmarshal(WhisperPromptString, &o.WhisperPrompt)
	}
	
	if AutoAnswerOnly, ok := QueuerequestMap["autoAnswerOnly"].(bool); ok {
		o.AutoAnswerOnly = &AutoAnswerOnly
	}
	
	if EnableTranscription, ok := QueuerequestMap["enableTranscription"].(bool); ok {
		o.EnableTranscription = &EnableTranscription
	}
	
	if EnableManualAssignment, ok := QueuerequestMap["enableManualAssignment"].(bool); ok {
		o.EnableManualAssignment = &EnableManualAssignment
	}
	
	if CallingPartyName, ok := QueuerequestMap["callingPartyName"].(string); ok {
		o.CallingPartyName = &CallingPartyName
	}
	
	if CallingPartyNumber, ok := QueuerequestMap["callingPartyNumber"].(string); ok {
		o.CallingPartyNumber = &CallingPartyNumber
	}
	
	if DefaultScripts, ok := QueuerequestMap["defaultScripts"].(map[string]interface{}); ok {
		DefaultScriptsString, _ := json.Marshal(DefaultScripts)
		json.Unmarshal(DefaultScriptsString, &o.DefaultScripts)
	}
	
	if OutboundMessagingAddresses, ok := QueuerequestMap["outboundMessagingAddresses"].(map[string]interface{}); ok {
		OutboundMessagingAddressesString, _ := json.Marshal(OutboundMessagingAddresses)
		json.Unmarshal(OutboundMessagingAddressesString, &o.OutboundMessagingAddresses)
	}
	
	if OutboundEmailAddress, ok := QueuerequestMap["outboundEmailAddress"].(map[string]interface{}); ok {
		OutboundEmailAddressString, _ := json.Marshal(OutboundEmailAddress)
		json.Unmarshal(OutboundEmailAddressString, &o.OutboundEmailAddress)
	}
	
	if SelfUri, ok := QueuerequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queuerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
