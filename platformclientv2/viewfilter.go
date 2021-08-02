package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Viewfilter
type Viewfilter struct { 
	// MediaTypes - The media types are used to filter the view
	MediaTypes *[]string `json:"mediaTypes,omitempty"`


	// QueueIds - The queue ids are used to filter the view
	QueueIds *[]string `json:"queueIds,omitempty"`


	// SkillIds - The skill ids are used to filter the view
	SkillIds *[]string `json:"skillIds,omitempty"`


	// SkillGroups - The skill groups used to filter the view
	SkillGroups *[]string `json:"skillGroups,omitempty"`


	// LanguageIds - The language ids are used to filter the view
	LanguageIds *[]string `json:"languageIds,omitempty"`


	// LanguageGroups - The language groups used to filter the view
	LanguageGroups *[]string `json:"languageGroups,omitempty"`


	// Directions - The directions are used to filter the view
	Directions *[]string `json:"directions,omitempty"`


	// OriginatingDirections - The list of orginating directions used to filter the view
	OriginatingDirections *[]string `json:"originatingDirections,omitempty"`


	// WrapUpCodes - The wrap up codes are used to filter the view
	WrapUpCodes *[]string `json:"wrapUpCodes,omitempty"`


	// DnisList - The dnis list is used to filter the view
	DnisList *[]string `json:"dnisList,omitempty"`


	// SessionDnisList - The list of session dnis used to filter the view
	SessionDnisList *[]string `json:"sessionDnisList,omitempty"`


	// FilterQueuesByUserIds - The user ids are used to fetch associated queues for the view
	FilterQueuesByUserIds *[]string `json:"filterQueuesByUserIds,omitempty"`


	// FilterUsersByQueueIds - The queue ids are used to fetch associated users for the view
	FilterUsersByQueueIds *[]string `json:"filterUsersByQueueIds,omitempty"`


	// UserIds - The user ids are used to filter the view
	UserIds *[]string `json:"userIds,omitempty"`


	// AddressTos - The address To values are used to filter the view
	AddressTos *[]string `json:"addressTos,omitempty"`


	// AddressFroms - The address from values are used to filter the view
	AddressFroms *[]string `json:"addressFroms,omitempty"`


	// OutboundCampaignIds - The outbound campaign ids are used to filter the view
	OutboundCampaignIds *[]string `json:"outboundCampaignIds,omitempty"`


	// OutboundContactListIds - The outbound contact list ids are used to filter the view
	OutboundContactListIds *[]string `json:"outboundContactListIds,omitempty"`


	// ContactIds - The contact ids are used to filter the view
	ContactIds *[]string `json:"contactIds,omitempty"`


	// ExternalContactIds - The external contact ids are used to filter the view
	ExternalContactIds *[]string `json:"externalContactIds,omitempty"`


	// ExternalOrgIds - The external org ids are used to filter the view
	ExternalOrgIds *[]string `json:"externalOrgIds,omitempty"`


	// AniList - The ani list ids are used to filter the view
	AniList *[]string `json:"aniList,omitempty"`


	// DurationsMilliseconds - The durations in milliseconds used to filter the view
	DurationsMilliseconds *[]Numericrange `json:"durationsMilliseconds,omitempty"`


	// AcdDurationsMilliseconds - The acd durations in milliseconds used to filter the view
	AcdDurationsMilliseconds *[]Numericrange `json:"acdDurationsMilliseconds,omitempty"`


	// TalkDurationsMilliseconds - The talk durations in milliseconds used to filter the view
	TalkDurationsMilliseconds *[]Numericrange `json:"talkDurationsMilliseconds,omitempty"`


	// AcwDurationsMilliseconds - The acw durations in milliseconds used to filter the view
	AcwDurationsMilliseconds *[]Numericrange `json:"acwDurationsMilliseconds,omitempty"`


	// HandleDurationsMilliseconds - The handle durations in milliseconds used to filter the view
	HandleDurationsMilliseconds *[]Numericrange `json:"handleDurationsMilliseconds,omitempty"`


	// HoldDurationsMilliseconds - The hold durations in milliseconds used to filter the view
	HoldDurationsMilliseconds *[]Numericrange `json:"holdDurationsMilliseconds,omitempty"`


	// AbandonDurationsMilliseconds - The abandon durations in milliseconds used to filter the view
	AbandonDurationsMilliseconds *[]Numericrange `json:"abandonDurationsMilliseconds,omitempty"`


	// EvaluationScore - The evaluationScore is used to filter the view
	EvaluationScore *Numericrange `json:"evaluationScore,omitempty"`


	// EvaluationCriticalScore - The evaluationCriticalScore is used to filter the view
	EvaluationCriticalScore *Numericrange `json:"evaluationCriticalScore,omitempty"`


	// EvaluationFormIds - The evaluation form ids are used to filter the view
	EvaluationFormIds *[]string `json:"evaluationFormIds,omitempty"`


	// EvaluatedAgentIds - The evaluated agent ids are used to filter the view
	EvaluatedAgentIds *[]string `json:"evaluatedAgentIds,omitempty"`


	// EvaluatorIds - The evaluator ids are used to filter the view
	EvaluatorIds *[]string `json:"evaluatorIds,omitempty"`


	// Transferred - Indicates filtering for transfers
	Transferred *bool `json:"transferred,omitempty"`


	// Abandoned - Indicates filtering for abandons
	Abandoned *bool `json:"abandoned,omitempty"`


	// Answered - Indicates filtering for answered interactions
	Answered *bool `json:"answered,omitempty"`


	// MessageTypes - The message media types used to filter the view
	MessageTypes *[]string `json:"messageTypes,omitempty"`


	// DivisionIds - The divison Ids used to filter the view
	DivisionIds *[]string `json:"divisionIds,omitempty"`


	// SurveyFormIds - The survey form ids used to filter the view
	SurveyFormIds *[]string `json:"surveyFormIds,omitempty"`


	// SurveyTotalScore - The survey total score used to filter the view
	SurveyTotalScore *Numericrange `json:"surveyTotalScore,omitempty"`


	// SurveyNpsScore - The survey NPS score used to filter the view
	SurveyNpsScore *Numericrange `json:"surveyNpsScore,omitempty"`


	// Mos - The desired range for mos values
	Mos *Numericrange `json:"mos,omitempty"`


	// SurveyQuestionGroupScore - The survey question group score used to filter the view
	SurveyQuestionGroupScore *Numericrange `json:"surveyQuestionGroupScore,omitempty"`


	// SurveyPromoterScore - The survey promoter score used to filter the view
	SurveyPromoterScore *Numericrange `json:"surveyPromoterScore,omitempty"`


	// SurveyFormContextIds - The list of survey form context ids used to filter the view
	SurveyFormContextIds *[]string `json:"surveyFormContextIds,omitempty"`


	// ConversationIds - The list of conversation ids used to filter the view
	ConversationIds *[]string `json:"conversationIds,omitempty"`


	// SipCallIds - The list of SIP call ids used to filter the view
	SipCallIds *[]string `json:"sipCallIds,omitempty"`


	// IsEnded - Indicates filtering for ended
	IsEnded *bool `json:"isEnded,omitempty"`


	// IsSurveyed - Indicates filtering for survey
	IsSurveyed *bool `json:"isSurveyed,omitempty"`


	// SurveyScores - The list of survey score ranges used to filter the view
	SurveyScores *[]Numericrange `json:"surveyScores,omitempty"`


	// PromoterScores - The list of promoter score ranges used to filter the view
	PromoterScores *[]Numericrange `json:"promoterScores,omitempty"`


	// IsCampaign - Indicates filtering for campaign
	IsCampaign *bool `json:"isCampaign,omitempty"`


	// SurveyStatuses - The list of survey statuses used to filter the view
	SurveyStatuses *[]string `json:"surveyStatuses,omitempty"`


	// ConversationProperties - A grouping of conversation level filters
	ConversationProperties *Conversationproperties `json:"conversationProperties,omitempty"`


	// IsBlindTransferred - Indicates filtering for blind transferred
	IsBlindTransferred *bool `json:"isBlindTransferred,omitempty"`


	// IsConsulted - Indicates filtering for consulted
	IsConsulted *bool `json:"isConsulted,omitempty"`


	// IsConsultTransferred - Indicates filtering for consult transferred
	IsConsultTransferred *bool `json:"isConsultTransferred,omitempty"`


	// RemoteParticipants - The list of remote participants used to filter the view
	RemoteParticipants *[]string `json:"remoteParticipants,omitempty"`


	// FlowIds - The list of flow Ids
	FlowIds *[]string `json:"flowIds,omitempty"`


	// FlowOutcomeIds - A list of outcome ids of the flow
	FlowOutcomeIds *[]string `json:"flowOutcomeIds,omitempty"`


	// FlowOutcomeValues - A list of outcome values of the flow
	FlowOutcomeValues *[]string `json:"flowOutcomeValues,omitempty"`


	// FlowDestinationTypes - The list of destination types of the flow
	FlowDestinationTypes *[]string `json:"flowDestinationTypes,omitempty"`


	// FlowDisconnectReasons - The list of reasons for the flow to disconnect
	FlowDisconnectReasons *[]string `json:"flowDisconnectReasons,omitempty"`


	// FlowTypes - A list of types of the flow
	FlowTypes *[]string `json:"flowTypes,omitempty"`


	// FlowEntryTypes - A list of types of the flow entry
	FlowEntryTypes *[]string `json:"flowEntryTypes,omitempty"`


	// FlowEntryReasons - A list of reasons of flow entry
	FlowEntryReasons *[]string `json:"flowEntryReasons,omitempty"`


	// FlowVersions - A list of versions of a flow
	FlowVersions *[]string `json:"flowVersions,omitempty"`


	// GroupIds - A list of directory group ids
	GroupIds *[]string `json:"groupIds,omitempty"`


	// HasJourneyCustomerId - Indicates filtering for journey customer id
	HasJourneyCustomerId *bool `json:"hasJourneyCustomerId,omitempty"`


	// HasJourneyActionMapId - Indicates filtering for Journey action map id
	HasJourneyActionMapId *bool `json:"hasJourneyActionMapId,omitempty"`


	// HasJourneyVisitId - Indicates filtering for Journey visit id
	HasJourneyVisitId *bool `json:"hasJourneyVisitId,omitempty"`


	// HasMedia - Indicates filtering for presence of MMS media
	HasMedia *bool `json:"hasMedia,omitempty"`


	// RoleIds - The role Ids used to filter the view
	RoleIds *[]string `json:"roleIds,omitempty"`


	// ReportsTos - The report to user IDs used to filter the view
	ReportsTos *[]string `json:"reportsTos,omitempty"`


	// LocationIds - The location Ids used to filter the view
	LocationIds *[]string `json:"locationIds,omitempty"`


	// FlowOutTypes - A list of flow out types
	FlowOutTypes *[]string `json:"flowOutTypes,omitempty"`


	// ProviderList - A list of providers
	ProviderList *[]string `json:"providerList,omitempty"`


	// CallbackNumberList - A list of callback numbers or substrings of numbers (ex: [\"317\", \"13172222222\"])
	CallbackNumberList *[]string `json:"callbackNumberList,omitempty"`


	// CallbackInterval - An interval of time to filter for scheduled callbacks. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	CallbackInterval *string `json:"callbackInterval,omitempty"`


	// UsedRoutingTypes - A list of routing types used
	UsedRoutingTypes *[]string `json:"usedRoutingTypes,omitempty"`


	// RequestedRoutingTypes - A list of routing types requested
	RequestedRoutingTypes *[]string `json:"requestedRoutingTypes,omitempty"`


	// HasAgentAssistId - Indicates filtering for agent assist id
	HasAgentAssistId *bool `json:"hasAgentAssistId,omitempty"`


	// Transcripts - A list of transcript contents requested
	Transcripts *[]Transcripts `json:"transcripts,omitempty"`


	// TranscriptLanguages - A list of transcript languages requested
	TranscriptLanguages *[]string `json:"transcriptLanguages,omitempty"`


	// ParticipantPurposes - A list of participant purpose requested
	ParticipantPurposes *[]string `json:"participantPurposes,omitempty"`


	// ShowFirstQueue - Indicates filtering for first queue data
	ShowFirstQueue *bool `json:"showFirstQueue,omitempty"`


	// TeamIds - The team ids used to filter the view data
	TeamIds *[]string `json:"teamIds,omitempty"`


	// FilterUsersByTeamIds - The team ids are used to fetch associated users for the view
	FilterUsersByTeamIds *[]string `json:"filterUsersByTeamIds,omitempty"`


	// JourneyActionMapIds - The journey action map ids are used to fetch action maps for the associated view
	JourneyActionMapIds *[]string `json:"journeyActionMapIds,omitempty"`


	// JourneyOutcomeIds - The journey outcome ids are used to fetch outcomes for the associated view
	JourneyOutcomeIds *[]string `json:"journeyOutcomeIds,omitempty"`


	// JourneySegmentIds - The journey segment ids are used to fetch segments for the associated view
	JourneySegmentIds *[]string `json:"journeySegmentIds,omitempty"`


	// JourneyActionMapTypes - The journey action map types are used to filter action map data for the associated view
	JourneyActionMapTypes *[]string `json:"journeyActionMapTypes,omitempty"`


	// DevelopmentRoleList - The list of development roles used to filter agent development view
	DevelopmentRoleList *[]string `json:"developmentRoleList,omitempty"`


	// DevelopmentTypeList - The list of development types used to filter agent development view
	DevelopmentTypeList *[]string `json:"developmentTypeList,omitempty"`


	// DevelopmentStatusList - The list of development status used to filter agent development view
	DevelopmentStatusList *[]string `json:"developmentStatusList,omitempty"`


	// DevelopmentModuleIds - The list of development moduleIds used to filter agent development view
	DevelopmentModuleIds *[]string `json:"developmentModuleIds,omitempty"`


	// DevelopmentActivityOverdue - Indicates filtering for development activities
	DevelopmentActivityOverdue *bool `json:"developmentActivityOverdue,omitempty"`


	// CustomerSentimentScore - The customer sentiment score used to filter the view
	CustomerSentimentScore *Numericrange `json:"customerSentimentScore,omitempty"`


	// CustomerSentimentTrend - The customer sentiment trend used to filter the view
	CustomerSentimentTrend *Numericrange `json:"customerSentimentTrend,omitempty"`


	// FlowTransferTargets - The list of transfer targets used to filter flow data
	FlowTransferTargets *[]string `json:"flowTransferTargets,omitempty"`


	// DevelopmentName - Filter for development name
	DevelopmentName *string `json:"developmentName,omitempty"`


	// TopicIds - Represents the topics detected in the transcript
	TopicIds *[]string `json:"topicIds,omitempty"`


	// ExternalTags - The list of external Tags used to filter conversation data
	ExternalTags *[]string `json:"externalTags,omitempty"`


	// IsNotResponding - Indicates filtering for not responding users
	IsNotResponding *bool `json:"isNotResponding,omitempty"`


	// IsAuthenticated - Indicates filtering for the authenticated chat
	IsAuthenticated *bool `json:"isAuthenticated,omitempty"`


	// BotIds - The list of bot IDs used to filter bot views
	BotIds *[]string `json:"botIds,omitempty"`


	// BotVersions - The list of bot versions used to filter bot views
	BotVersions *[]string `json:"botVersions,omitempty"`


	// BotMessageTypes - The list of bot message types used to filter bot views
	BotMessageTypes *[]string `json:"botMessageTypes,omitempty"`


	// BotProviderList - The list of bot providers used to filter bot views
	BotProviderList *[]string `json:"botProviderList,omitempty"`


	// BotProductList - The list of bot products used to filter bot views
	BotProductList *[]string `json:"botProductList,omitempty"`


	// BotRecognitionFailureReasonList - The list of bot recognition failure reasons used to filter bot views
	BotRecognitionFailureReasonList *[]string `json:"botRecognitionFailureReasonList,omitempty"`


	// BotIntentList - The list of bot intents used to filter bot views
	BotIntentList *[]string `json:"botIntentList,omitempty"`


	// BotFinalIntentList - The list of bot final intents used to filter bot views
	BotFinalIntentList *[]string `json:"botFinalIntentList,omitempty"`


	// BotSlotList - The list of bot slots used to filter bot views
	BotSlotList *[]string `json:"botSlotList,omitempty"`


	// BotResultList - The list of bot results used to filter bot views
	BotResultList *[]string `json:"botResultList,omitempty"`


	// BlockedReasons - The list of blocked reason used to filter action map constraints views
	BlockedReasons *[]string `json:"blockedReasons,omitempty"`

}

// String returns a JSON representation of the model
func (o *Viewfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
