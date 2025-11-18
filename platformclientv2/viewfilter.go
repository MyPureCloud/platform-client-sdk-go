package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Viewfilter
type Viewfilter struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// MediaTypes - The media types are used to filter the view
	MediaTypes *[]string `json:"mediaTypes,omitempty"`

	// QueueIds - The queue ids are used to filter the view
	QueueIds *[]string `json:"queueIds,omitempty"`

	// SkillIds - The skill ids are used to filter the view
	SkillIds *[]string `json:"skillIds,omitempty"`

	// AssignedSkillIds - The assigned user skill ids are used to filter the view
	AssignedSkillIds *[]string `json:"assignedSkillIds,omitempty"`

	// SkillGroups - The skill groups used to filter the view
	SkillGroups *[]string `json:"skillGroups,omitempty"`

	// LanguageIds - The language ids are used to filter the view
	LanguageIds *[]string `json:"languageIds,omitempty"`

	// AssignedLanguageIds - The assigned user language ids are used to filter the view
	AssignedLanguageIds *[]string `json:"assignedLanguageIds,omitempty"`

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

	// ManagementUnitIds - The management unit ids are used to filter the view
	ManagementUnitIds *[]string `json:"managementUnitIds,omitempty"`

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

	// IsRecorded - Indicates filtering for recorded
	IsRecorded *bool `json:"isRecorded,omitempty"`

	// HasEvaluation - Indicates filtering for evaluation
	HasEvaluation *bool `json:"hasEvaluation,omitempty"`

	// HasScoredEvaluation - Indicates filtering for scored evaluation
	HasScoredEvaluation *bool `json:"hasScoredEvaluation,omitempty"`

	// EmailDeliveryStatusList - The list of email delivery statuses used to filter views
	EmailDeliveryStatusList *[]string `json:"emailDeliveryStatusList,omitempty"`

	// IsAgentOwnedCallback - Indicates filtering for agent owned callback interactions
	IsAgentOwnedCallback *bool `json:"isAgentOwnedCallback,omitempty"`

	// AgentCallbackOwnerIds - The list of callback owners used to filter interactions
	AgentCallbackOwnerIds *[]string `json:"agentCallbackOwnerIds,omitempty"`

	// TranscriptTopics - The list of transcript topics requested in filter
	TranscriptTopics *[]Transcripttopics `json:"transcriptTopics,omitempty"`

	// JourneyFrequencyCapReasons - The list of frequency cap reasons to filter offer constraints
	JourneyFrequencyCapReasons *[]string `json:"journeyFrequencyCapReasons,omitempty"`

	// JourneyBlockingActionMapIds - The list of blocking action maps to filter offer constraints
	JourneyBlockingActionMapIds *[]string `json:"journeyBlockingActionMapIds,omitempty"`

	// JourneyActionTargetIds - The list of action targets to filter offer constraints
	JourneyActionTargetIds *[]string `json:"journeyActionTargetIds,omitempty"`

	// JourneyBlockingScheduleGroupIds - The list of blocking schedule groups to filter offer constraints
	JourneyBlockingScheduleGroupIds *[]string `json:"journeyBlockingScheduleGroupIds,omitempty"`

	// JourneyBlockingEmergencyScheduleGroupIds - The list of emergency schedule groups to filter offer constraints
	JourneyBlockingEmergencyScheduleGroupIds *[]string `json:"journeyBlockingEmergencyScheduleGroupIds,omitempty"`

	// JourneyUrlEqualConditions - The list of url equal conditions to filter offer constraints
	JourneyUrlEqualConditions *[]string `json:"journeyUrlEqualConditions,omitempty"`

	// JourneyUrlNotEqualConditions - The list of url not equal conditions to filter offer constraints
	JourneyUrlNotEqualConditions *[]string `json:"journeyUrlNotEqualConditions,omitempty"`

	// JourneyUrlStartsWithConditions - The list of url starts with conditions to filter offer constraints
	JourneyUrlStartsWithConditions *[]string `json:"journeyUrlStartsWithConditions,omitempty"`

	// JourneyUrlEndsWithConditions - The list of url ends with conditions to filter offer constraints
	JourneyUrlEndsWithConditions *[]string `json:"journeyUrlEndsWithConditions,omitempty"`

	// JourneyUrlContainsAnyConditions - The list of url contains any conditions to filter offer constraints
	JourneyUrlContainsAnyConditions *[]string `json:"journeyUrlContainsAnyConditions,omitempty"`

	// JourneyUrlNotContainsAnyConditions - The list of url not contains any conditions to filter offer constraints
	JourneyUrlNotContainsAnyConditions *[]string `json:"journeyUrlNotContainsAnyConditions,omitempty"`

	// JourneyUrlContainsAllConditions - The list of url contains all conditions to filter offer constraints
	JourneyUrlContainsAllConditions *[]string `json:"journeyUrlContainsAllConditions,omitempty"`

	// JourneyUrlNotContainsAllConditions - The list of url not contains all conditions to filter offer constraints
	JourneyUrlNotContainsAllConditions *[]string `json:"journeyUrlNotContainsAllConditions,omitempty"`

	// FlowMilestoneIds - The list of flow milestones to filter exports
	FlowMilestoneIds *[]string `json:"flowMilestoneIds,omitempty"`

	// IsAssessmentPassed - Filter to indicate if Agent passed assessment or not
	IsAssessmentPassed *bool `json:"isAssessmentPassed,omitempty"`

	// ConversationInitiators - The list to filter based on Brands (Bot/User/Agent) or End User who initiated the first message in the conversation
	ConversationInitiators *[]string `json:"conversationInitiators,omitempty"`

	// HasCustomerParticipated - Indicates if the customer has participated in an initiated conversation
	HasCustomerParticipated *bool `json:"hasCustomerParticipated,omitempty"`

	// IsAcdInteraction - Filter to indicate if interaction was ACD or non-ACD
	IsAcdInteraction *bool `json:"isAcdInteraction,omitempty"`

	// HasFax - Filters to indicate if interaction has FAX
	HasFax *bool `json:"hasFax,omitempty"`

	// DataActionIds - The list of Data Action IDs 
	DataActionIds *[]string `json:"dataActionIds,omitempty"`

	// ActionCategoryName - Deprecated - Please use integrationIds instead
	ActionCategoryName *string `json:"actionCategoryName,omitempty"`

	// IntegrationIds - The list of integration IDs for Data Action
	IntegrationIds *[]string `json:"integrationIds,omitempty"`

	// ResponseStatuses - The list of Response codes for Data Action
	ResponseStatuses *[]string `json:"responseStatuses,omitempty"`

	// AvailableDashboard - Filter to indicate the availability of the dashboard is public or private.
	AvailableDashboard *string `json:"availableDashboard,omitempty"`

	// FavouriteDashboard - Filter to indicate whether the dashboard is favorite or unfavorite.
	FavouriteDashboard *bool `json:"favouriteDashboard,omitempty"`

	// MyDashboard - Filter to indicate the dashboard owned by the user.
	MyDashboard *bool `json:"myDashboard,omitempty"`

	// StationErrors - The list of agent errors that are related to station
	StationErrors *[]string `json:"stationErrors,omitempty"`

	// CanonicalContactIds - The canonical contact ids are used to filter the view
	CanonicalContactIds *[]string `json:"canonicalContactIds,omitempty"`

	// AlertRuleIds - The list of Alert Rule IDs
	AlertRuleIds *[]string `json:"alertRuleIds,omitempty"`

	// EvaluationFormContextIds - The list of Evaluation Form Context IDs
	EvaluationFormContextIds *[]string `json:"evaluationFormContextIds,omitempty"`

	// EvaluationStatuses - The evaluation statuses that are used to filter the view
	EvaluationStatuses *[]string `json:"evaluationStatuses,omitempty"`

	// WorkbinIds - The list of Workbin IDs
	WorkbinIds *[]string `json:"workbinIds,omitempty"`

	// WorktypeIds - The list of Worktype IDs
	WorktypeIds *[]string `json:"worktypeIds,omitempty"`

	// WorkitemIds - The list of Workitem IDs
	WorkitemIds *[]string `json:"workitemIds,omitempty"`

	// WorkitemAssigneeIds - The list of Workitem Assignee IDs
	WorkitemAssigneeIds *[]string `json:"workitemAssigneeIds,omitempty"`

	// WorkitemStatuses - The list of Workitem Statuses IDs
	WorkitemStatuses *[]string `json:"workitemStatuses,omitempty"`

	// IsAnalyzedForSensitiveData - Deprecated - Use hasPciData or hasPiiData instead.
	IsAnalyzedForSensitiveData *bool `json:"isAnalyzedForSensitiveData,omitempty"`

	// HasSensitiveData - Deprecated. Use hasPciData or hasPiiData instead.
	HasSensitiveData *bool `json:"hasSensitiveData,omitempty"`

	// HasPciData - Filter to indicate the transcript contains Pci data.
	HasPciData *bool `json:"hasPciData,omitempty"`

	// HasPiiData - Filter to indicate the transcript contains Pii data.
	HasPiiData *bool `json:"hasPiiData,omitempty"`

	// SubPath - Filter for Sub Path
	SubPath *string `json:"subPath,omitempty"`

	// UserState - The user supplied state value in the view
	UserState *string `json:"userState,omitempty"`

	// IsClearedByCustomer - Filter to indicate if the customer cleared the conversation.
	IsClearedByCustomer *bool `json:"isClearedByCustomer,omitempty"`

	// EvaluationAssigneeIds - The evaluation assignee ids that are used to filter the view.
	EvaluationAssigneeIds *[]string `json:"evaluationAssigneeIds,omitempty"`

	// EvaluationAssigned - Filter to indicate that the user has no assigned evaluation.
	EvaluationAssigned *bool `json:"evaluationAssigned,omitempty"`

	// AssistantIds - The assistant ids that are used to filter the view.
	AssistantIds *[]string `json:"assistantIds,omitempty"`

	// KnowledgeBaseIds - The knowledge base ids that are used to filter the view.
	KnowledgeBaseIds *[]string `json:"knowledgeBaseIds,omitempty"`

	// IsParked - Filter to indicate if the interactions are parked.
	IsParked *bool `json:"isParked,omitempty"`

	// AgentEmpathyScore - The agentEmpathyScore is used to filter the view
	AgentEmpathyScore *Numericrange `json:"agentEmpathyScore,omitempty"`

	// SurveyTypes - The surveyTypes is used to filter the view
	SurveyTypes *[]string `json:"surveyTypes,omitempty"`

	// SurveyResponseStatuses - The list of Survey Response Status
	SurveyResponseStatuses *[]string `json:"surveyResponseStatuses,omitempty"`

	// BotFlowTypes - The botFlowTypes is used to filter the view
	BotFlowTypes *[]string `json:"botFlowTypes,omitempty"`

	// AgentTalkDurationMilliseconds - The agent talk durations in milliseconds used to filter the view
	AgentTalkDurationMilliseconds *[]Numericrange `json:"agentTalkDurationMilliseconds,omitempty"`

	// CustomerTalkDurationMilliseconds - The customer talk durations in milliseconds used to filter the view
	CustomerTalkDurationMilliseconds *[]Numericrange `json:"customerTalkDurationMilliseconds,omitempty"`

	// OvertalkDurationMilliseconds - The overtalk durations in milliseconds used to filter the view
	OvertalkDurationMilliseconds *[]Numericrange `json:"overtalkDurationMilliseconds,omitempty"`

	// SilenceDurationMilliseconds - The silence durations in milliseconds used to filter the view
	SilenceDurationMilliseconds *[]Numericrange `json:"silenceDurationMilliseconds,omitempty"`

	// AcdDurationMilliseconds - The acd durations in milliseconds used to filter the view
	AcdDurationMilliseconds *[]Numericrange `json:"acdDurationMilliseconds,omitempty"`

	// IvrDurationMilliseconds - The ivr durations in milliseconds used to filter the view
	IvrDurationMilliseconds *[]Numericrange `json:"ivrDurationMilliseconds,omitempty"`

	// OtherDurationMilliseconds - The other (hold/music) durations in milliseconds used to filter the view
	OtherDurationMilliseconds *[]Numericrange `json:"otherDurationMilliseconds,omitempty"`

	// AgentTalkPercentage - The agent talk percentage used to filter the view
	AgentTalkPercentage *Numericrange `json:"agentTalkPercentage,omitempty"`

	// CustomerTalkPercentage - The customer talk percentage used to filter the view
	CustomerTalkPercentage *Numericrange `json:"customerTalkPercentage,omitempty"`

	// OvertalkPercentage - The overtalk percentage used to filter the view
	OvertalkPercentage *Numericrange `json:"overtalkPercentage,omitempty"`

	// SilencePercentage - The silence percentage used to filter the view
	SilencePercentage *Numericrange `json:"silencePercentage,omitempty"`

	// AcdPercentage - The acd percentage used to filter the view
	AcdPercentage *Numericrange `json:"acdPercentage,omitempty"`

	// IvrPercentage - The ivr percentage used to filter the view
	IvrPercentage *Numericrange `json:"ivrPercentage,omitempty"`

	// OtherPercentage - The other (hold/music percentage used to filter the view
	OtherPercentage *Numericrange `json:"otherPercentage,omitempty"`

	// OvertalkInstances - The overtalk instance range used to filter the view
	OvertalkInstances *Numericrange `json:"overtalkInstances,omitempty"`

	// IsScreenRecorded - Filter to indicate if the screen is recorded
	IsScreenRecorded *bool `json:"isScreenRecorded,omitempty"`

	// ScreenMonitorUserIds - The list of Screen Monitor User Ids
	ScreenMonitorUserIds *[]string `json:"screenMonitorUserIds,omitempty"`

	// DashboardState - The state of dashboard being filtered
	DashboardState *string `json:"dashboardState,omitempty"`

	// DashboardType - The type of dashboard being filtered
	DashboardType *string `json:"dashboardType,omitempty"`

	// DashboardAccessFilter - The type of dashboard access being filtered
	DashboardAccessFilter *string `json:"dashboardAccessFilter,omitempty"`

	// TranscriptDurationMilliseconds - The transcript durations in milliseconds used to filter the view
	TranscriptDurationMilliseconds *[]Numericrange `json:"transcriptDurationMilliseconds,omitempty"`

	// WorkitemsStatuses - The list of workitem status with worktype
	WorkitemsStatuses *[]Workitemstatusfilter `json:"workitemsStatuses,omitempty"`

	// SocialCountries - List of countries for social filtering
	SocialCountries *[]string `json:"socialCountries,omitempty"`

	// SocialLanguages - List of languages for social filtering
	SocialLanguages *[]string `json:"socialLanguages,omitempty"`

	// SocialChannels - List of channels for social filtering
	SocialChannels *[]string `json:"socialChannels,omitempty"`

	// SocialSentimentCategory - The sentiment of the social post
	SocialSentimentCategory *[]string `json:"socialSentimentCategory,omitempty"`

	// SocialTopicIds - The list of topicIds for social filtering
	SocialTopicIds *[]string `json:"socialTopicIds,omitempty"`

	// SocialIngestionRuleIds - The list of ingestion ruleIds for social filtering
	SocialIngestionRuleIds *[]string `json:"socialIngestionRuleIds,omitempty"`

	// SocialConversationCreated - Filter to indicate if the post has created a conversation
	SocialConversationCreated *bool `json:"socialConversationCreated,omitempty"`

	// SocialContentType - The list of content Type for social filtering
	SocialContentType *[]string `json:"socialContentType,omitempty"`

	// SocialKeywords - The list of keywords for social filtering
	SocialKeywords *[]Socialkeyword `json:"socialKeywords,omitempty"`

	// SocialPostEscalated - Filter to indicate if the post is escalated
	SocialPostEscalated *bool `json:"socialPostEscalated,omitempty"`

	// SocialClassifications - Indicates if a social message was public or private
	SocialClassifications *[]string `json:"socialClassifications,omitempty"`

	// FilterUsersByManagerIds - The manager ids used to fetch associated users for the view
	FilterUsersByManagerIds *[]string `json:"filterUsersByManagerIds,omitempty"`

	// SlideshowIds - List of Dashboard slideshowIds to be filtered
	SlideshowIds *[]string `json:"slideshowIds,omitempty"`

	// Conferenced - Filter to indicate if the conversation has conference
	Conferenced *bool `json:"conferenced,omitempty"`

	// Video - Filter to indicate if the conversation has video
	Video *bool `json:"video,omitempty"`

	// LinkedInteraction - Filter to indicate if the conversation has linked interaction
	LinkedInteraction *bool `json:"linkedInteraction,omitempty"`

	// RecommendationSources - List of recommendation sources for filtering recommendation details pane
	RecommendationSources *[]string `json:"recommendationSources,omitempty"`

	// EvaluationRole - Sets the role when viewing agent evaluations
	EvaluationRole *string `json:"evaluationRole,omitempty"`

	// ComparisonQueueIds - The queue ids are used to for comparison to the primary queue filter in reporting
	ComparisonQueueIds *[]string `json:"comparisonQueueIds,omitempty"`

	// ViewMetrics - A list of metrics selected for the view
	ViewMetrics *[]string `json:"viewMetrics,omitempty"`

	// TimelineCategories - A list of timeline categories
	TimelineCategories *[]string `json:"timelineCategories,omitempty"`

	// Acw - Filter to indicate for acw state
	Acw *bool `json:"acw,omitempty"`

	// SegmentTypes - A list of filtered segment types
	SegmentTypes *[]string `json:"segmentTypes,omitempty"`

	// ProgramIds - A list of program ids for filtering
	ProgramIds *[]string `json:"programIds,omitempty"`

	// CategoryIds - A list of category ids for filtering
	CategoryIds *[]string `json:"categoryIds,omitempty"`

	// DeliveryPushed - Filter to indicate if push notification is sent
	DeliveryPushed *bool `json:"deliveryPushed,omitempty"`

	// SocialRatings - A set of ratings for Google Business Profile
	SocialRatings *[]float32 `json:"socialRatings,omitempty"`

	// VirtualAgentIds - A list of virtual agent ids for filtering.
	VirtualAgentIds *[]string `json:"virtualAgentIds,omitempty"`

	// EmpathyScoreCategories - A set of Empathy Score Categories for filtering
	EmpathyScoreCategories *[]string `json:"empathyScoreCategories,omitempty"`

	// SentimentScoreCategories - A set of Sentiment Score Categories  for filtering
	SentimentScoreCategories *[]string `json:"sentimentScoreCategories,omitempty"`

	// SentimentTrendCategories - A set of Sentiment Trend Categories for filtering
	SentimentTrendCategories *[]string `json:"sentimentTrendCategories,omitempty"`

	// ContentModerationFlags - A set of Content Moderation Flags for filtering
	ContentModerationFlags *[]string `json:"contentModerationFlags,omitempty"`

	// SessionExpired - Filter to indicate for if session is expired
	SessionExpired *bool `json:"sessionExpired,omitempty"`

	// EngagementSource - The engagement sources used to filter the view
	EngagementSource *[]string `json:"engagementSource,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Viewfilter) SetField(field string, fieldValue interface{}) {
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

func (o Viewfilter) MarshalJSON() ([]byte, error) {
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
	type Alias Viewfilter
	
	return json.Marshal(&struct { 
		MediaTypes *[]string `json:"mediaTypes,omitempty"`
		
		QueueIds *[]string `json:"queueIds,omitempty"`
		
		SkillIds *[]string `json:"skillIds,omitempty"`
		
		AssignedSkillIds *[]string `json:"assignedSkillIds,omitempty"`
		
		SkillGroups *[]string `json:"skillGroups,omitempty"`
		
		LanguageIds *[]string `json:"languageIds,omitempty"`
		
		AssignedLanguageIds *[]string `json:"assignedLanguageIds,omitempty"`
		
		LanguageGroups *[]string `json:"languageGroups,omitempty"`
		
		Directions *[]string `json:"directions,omitempty"`
		
		OriginatingDirections *[]string `json:"originatingDirections,omitempty"`
		
		WrapUpCodes *[]string `json:"wrapUpCodes,omitempty"`
		
		DnisList *[]string `json:"dnisList,omitempty"`
		
		SessionDnisList *[]string `json:"sessionDnisList,omitempty"`
		
		FilterQueuesByUserIds *[]string `json:"filterQueuesByUserIds,omitempty"`
		
		FilterUsersByQueueIds *[]string `json:"filterUsersByQueueIds,omitempty"`
		
		UserIds *[]string `json:"userIds,omitempty"`
		
		ManagementUnitIds *[]string `json:"managementUnitIds,omitempty"`
		
		AddressTos *[]string `json:"addressTos,omitempty"`
		
		AddressFroms *[]string `json:"addressFroms,omitempty"`
		
		OutboundCampaignIds *[]string `json:"outboundCampaignIds,omitempty"`
		
		OutboundContactListIds *[]string `json:"outboundContactListIds,omitempty"`
		
		ContactIds *[]string `json:"contactIds,omitempty"`
		
		ExternalContactIds *[]string `json:"externalContactIds,omitempty"`
		
		ExternalOrgIds *[]string `json:"externalOrgIds,omitempty"`
		
		AniList *[]string `json:"aniList,omitempty"`
		
		DurationsMilliseconds *[]Numericrange `json:"durationsMilliseconds,omitempty"`
		
		AcdDurationsMilliseconds *[]Numericrange `json:"acdDurationsMilliseconds,omitempty"`
		
		TalkDurationsMilliseconds *[]Numericrange `json:"talkDurationsMilliseconds,omitempty"`
		
		AcwDurationsMilliseconds *[]Numericrange `json:"acwDurationsMilliseconds,omitempty"`
		
		HandleDurationsMilliseconds *[]Numericrange `json:"handleDurationsMilliseconds,omitempty"`
		
		HoldDurationsMilliseconds *[]Numericrange `json:"holdDurationsMilliseconds,omitempty"`
		
		AbandonDurationsMilliseconds *[]Numericrange `json:"abandonDurationsMilliseconds,omitempty"`
		
		EvaluationScore *Numericrange `json:"evaluationScore,omitempty"`
		
		EvaluationCriticalScore *Numericrange `json:"evaluationCriticalScore,omitempty"`
		
		EvaluationFormIds *[]string `json:"evaluationFormIds,omitempty"`
		
		EvaluatedAgentIds *[]string `json:"evaluatedAgentIds,omitempty"`
		
		EvaluatorIds *[]string `json:"evaluatorIds,omitempty"`
		
		Transferred *bool `json:"transferred,omitempty"`
		
		Abandoned *bool `json:"abandoned,omitempty"`
		
		Answered *bool `json:"answered,omitempty"`
		
		MessageTypes *[]string `json:"messageTypes,omitempty"`
		
		DivisionIds *[]string `json:"divisionIds,omitempty"`
		
		SurveyFormIds *[]string `json:"surveyFormIds,omitempty"`
		
		SurveyTotalScore *Numericrange `json:"surveyTotalScore,omitempty"`
		
		SurveyNpsScore *Numericrange `json:"surveyNpsScore,omitempty"`
		
		Mos *Numericrange `json:"mos,omitempty"`
		
		SurveyQuestionGroupScore *Numericrange `json:"surveyQuestionGroupScore,omitempty"`
		
		SurveyPromoterScore *Numericrange `json:"surveyPromoterScore,omitempty"`
		
		SurveyFormContextIds *[]string `json:"surveyFormContextIds,omitempty"`
		
		ConversationIds *[]string `json:"conversationIds,omitempty"`
		
		SipCallIds *[]string `json:"sipCallIds,omitempty"`
		
		IsEnded *bool `json:"isEnded,omitempty"`
		
		IsSurveyed *bool `json:"isSurveyed,omitempty"`
		
		SurveyScores *[]Numericrange `json:"surveyScores,omitempty"`
		
		PromoterScores *[]Numericrange `json:"promoterScores,omitempty"`
		
		IsCampaign *bool `json:"isCampaign,omitempty"`
		
		SurveyStatuses *[]string `json:"surveyStatuses,omitempty"`
		
		ConversationProperties *Conversationproperties `json:"conversationProperties,omitempty"`
		
		IsBlindTransferred *bool `json:"isBlindTransferred,omitempty"`
		
		IsConsulted *bool `json:"isConsulted,omitempty"`
		
		IsConsultTransferred *bool `json:"isConsultTransferred,omitempty"`
		
		RemoteParticipants *[]string `json:"remoteParticipants,omitempty"`
		
		FlowIds *[]string `json:"flowIds,omitempty"`
		
		FlowOutcomeIds *[]string `json:"flowOutcomeIds,omitempty"`
		
		FlowOutcomeValues *[]string `json:"flowOutcomeValues,omitempty"`
		
		FlowDestinationTypes *[]string `json:"flowDestinationTypes,omitempty"`
		
		FlowDisconnectReasons *[]string `json:"flowDisconnectReasons,omitempty"`
		
		FlowTypes *[]string `json:"flowTypes,omitempty"`
		
		FlowEntryTypes *[]string `json:"flowEntryTypes,omitempty"`
		
		FlowEntryReasons *[]string `json:"flowEntryReasons,omitempty"`
		
		FlowVersions *[]string `json:"flowVersions,omitempty"`
		
		GroupIds *[]string `json:"groupIds,omitempty"`
		
		HasJourneyCustomerId *bool `json:"hasJourneyCustomerId,omitempty"`
		
		HasJourneyActionMapId *bool `json:"hasJourneyActionMapId,omitempty"`
		
		HasJourneyVisitId *bool `json:"hasJourneyVisitId,omitempty"`
		
		HasMedia *bool `json:"hasMedia,omitempty"`
		
		RoleIds *[]string `json:"roleIds,omitempty"`
		
		ReportsTos *[]string `json:"reportsTos,omitempty"`
		
		LocationIds *[]string `json:"locationIds,omitempty"`
		
		FlowOutTypes *[]string `json:"flowOutTypes,omitempty"`
		
		ProviderList *[]string `json:"providerList,omitempty"`
		
		CallbackNumberList *[]string `json:"callbackNumberList,omitempty"`
		
		CallbackInterval *string `json:"callbackInterval,omitempty"`
		
		UsedRoutingTypes *[]string `json:"usedRoutingTypes,omitempty"`
		
		RequestedRoutingTypes *[]string `json:"requestedRoutingTypes,omitempty"`
		
		HasAgentAssistId *bool `json:"hasAgentAssistId,omitempty"`
		
		Transcripts *[]Transcripts `json:"transcripts,omitempty"`
		
		TranscriptLanguages *[]string `json:"transcriptLanguages,omitempty"`
		
		ParticipantPurposes *[]string `json:"participantPurposes,omitempty"`
		
		ShowFirstQueue *bool `json:"showFirstQueue,omitempty"`
		
		TeamIds *[]string `json:"teamIds,omitempty"`
		
		FilterUsersByTeamIds *[]string `json:"filterUsersByTeamIds,omitempty"`
		
		JourneyActionMapIds *[]string `json:"journeyActionMapIds,omitempty"`
		
		JourneyOutcomeIds *[]string `json:"journeyOutcomeIds,omitempty"`
		
		JourneySegmentIds *[]string `json:"journeySegmentIds,omitempty"`
		
		JourneyActionMapTypes *[]string `json:"journeyActionMapTypes,omitempty"`
		
		DevelopmentRoleList *[]string `json:"developmentRoleList,omitempty"`
		
		DevelopmentTypeList *[]string `json:"developmentTypeList,omitempty"`
		
		DevelopmentStatusList *[]string `json:"developmentStatusList,omitempty"`
		
		DevelopmentModuleIds *[]string `json:"developmentModuleIds,omitempty"`
		
		DevelopmentActivityOverdue *bool `json:"developmentActivityOverdue,omitempty"`
		
		CustomerSentimentScore *Numericrange `json:"customerSentimentScore,omitempty"`
		
		CustomerSentimentTrend *Numericrange `json:"customerSentimentTrend,omitempty"`
		
		FlowTransferTargets *[]string `json:"flowTransferTargets,omitempty"`
		
		DevelopmentName *string `json:"developmentName,omitempty"`
		
		TopicIds *[]string `json:"topicIds,omitempty"`
		
		ExternalTags *[]string `json:"externalTags,omitempty"`
		
		IsNotResponding *bool `json:"isNotResponding,omitempty"`
		
		IsAuthenticated *bool `json:"isAuthenticated,omitempty"`
		
		BotIds *[]string `json:"botIds,omitempty"`
		
		BotVersions *[]string `json:"botVersions,omitempty"`
		
		BotMessageTypes *[]string `json:"botMessageTypes,omitempty"`
		
		BotProviderList *[]string `json:"botProviderList,omitempty"`
		
		BotProductList *[]string `json:"botProductList,omitempty"`
		
		BotRecognitionFailureReasonList *[]string `json:"botRecognitionFailureReasonList,omitempty"`
		
		BotIntentList *[]string `json:"botIntentList,omitempty"`
		
		BotFinalIntentList *[]string `json:"botFinalIntentList,omitempty"`
		
		BotSlotList *[]string `json:"botSlotList,omitempty"`
		
		BotResultList *[]string `json:"botResultList,omitempty"`
		
		BlockedReasons *[]string `json:"blockedReasons,omitempty"`
		
		IsRecorded *bool `json:"isRecorded,omitempty"`
		
		HasEvaluation *bool `json:"hasEvaluation,omitempty"`
		
		HasScoredEvaluation *bool `json:"hasScoredEvaluation,omitempty"`
		
		EmailDeliveryStatusList *[]string `json:"emailDeliveryStatusList,omitempty"`
		
		IsAgentOwnedCallback *bool `json:"isAgentOwnedCallback,omitempty"`
		
		AgentCallbackOwnerIds *[]string `json:"agentCallbackOwnerIds,omitempty"`
		
		TranscriptTopics *[]Transcripttopics `json:"transcriptTopics,omitempty"`
		
		JourneyFrequencyCapReasons *[]string `json:"journeyFrequencyCapReasons,omitempty"`
		
		JourneyBlockingActionMapIds *[]string `json:"journeyBlockingActionMapIds,omitempty"`
		
		JourneyActionTargetIds *[]string `json:"journeyActionTargetIds,omitempty"`
		
		JourneyBlockingScheduleGroupIds *[]string `json:"journeyBlockingScheduleGroupIds,omitempty"`
		
		JourneyBlockingEmergencyScheduleGroupIds *[]string `json:"journeyBlockingEmergencyScheduleGroupIds,omitempty"`
		
		JourneyUrlEqualConditions *[]string `json:"journeyUrlEqualConditions,omitempty"`
		
		JourneyUrlNotEqualConditions *[]string `json:"journeyUrlNotEqualConditions,omitempty"`
		
		JourneyUrlStartsWithConditions *[]string `json:"journeyUrlStartsWithConditions,omitempty"`
		
		JourneyUrlEndsWithConditions *[]string `json:"journeyUrlEndsWithConditions,omitempty"`
		
		JourneyUrlContainsAnyConditions *[]string `json:"journeyUrlContainsAnyConditions,omitempty"`
		
		JourneyUrlNotContainsAnyConditions *[]string `json:"journeyUrlNotContainsAnyConditions,omitempty"`
		
		JourneyUrlContainsAllConditions *[]string `json:"journeyUrlContainsAllConditions,omitempty"`
		
		JourneyUrlNotContainsAllConditions *[]string `json:"journeyUrlNotContainsAllConditions,omitempty"`
		
		FlowMilestoneIds *[]string `json:"flowMilestoneIds,omitempty"`
		
		IsAssessmentPassed *bool `json:"isAssessmentPassed,omitempty"`
		
		ConversationInitiators *[]string `json:"conversationInitiators,omitempty"`
		
		HasCustomerParticipated *bool `json:"hasCustomerParticipated,omitempty"`
		
		IsAcdInteraction *bool `json:"isAcdInteraction,omitempty"`
		
		HasFax *bool `json:"hasFax,omitempty"`
		
		DataActionIds *[]string `json:"dataActionIds,omitempty"`
		
		ActionCategoryName *string `json:"actionCategoryName,omitempty"`
		
		IntegrationIds *[]string `json:"integrationIds,omitempty"`
		
		ResponseStatuses *[]string `json:"responseStatuses,omitempty"`
		
		AvailableDashboard *string `json:"availableDashboard,omitempty"`
		
		FavouriteDashboard *bool `json:"favouriteDashboard,omitempty"`
		
		MyDashboard *bool `json:"myDashboard,omitempty"`
		
		StationErrors *[]string `json:"stationErrors,omitempty"`
		
		CanonicalContactIds *[]string `json:"canonicalContactIds,omitempty"`
		
		AlertRuleIds *[]string `json:"alertRuleIds,omitempty"`
		
		EvaluationFormContextIds *[]string `json:"evaluationFormContextIds,omitempty"`
		
		EvaluationStatuses *[]string `json:"evaluationStatuses,omitempty"`
		
		WorkbinIds *[]string `json:"workbinIds,omitempty"`
		
		WorktypeIds *[]string `json:"worktypeIds,omitempty"`
		
		WorkitemIds *[]string `json:"workitemIds,omitempty"`
		
		WorkitemAssigneeIds *[]string `json:"workitemAssigneeIds,omitempty"`
		
		WorkitemStatuses *[]string `json:"workitemStatuses,omitempty"`
		
		IsAnalyzedForSensitiveData *bool `json:"isAnalyzedForSensitiveData,omitempty"`
		
		HasSensitiveData *bool `json:"hasSensitiveData,omitempty"`
		
		HasPciData *bool `json:"hasPciData,omitempty"`
		
		HasPiiData *bool `json:"hasPiiData,omitempty"`
		
		SubPath *string `json:"subPath,omitempty"`
		
		UserState *string `json:"userState,omitempty"`
		
		IsClearedByCustomer *bool `json:"isClearedByCustomer,omitempty"`
		
		EvaluationAssigneeIds *[]string `json:"evaluationAssigneeIds,omitempty"`
		
		EvaluationAssigned *bool `json:"evaluationAssigned,omitempty"`
		
		AssistantIds *[]string `json:"assistantIds,omitempty"`
		
		KnowledgeBaseIds *[]string `json:"knowledgeBaseIds,omitempty"`
		
		IsParked *bool `json:"isParked,omitempty"`
		
		AgentEmpathyScore *Numericrange `json:"agentEmpathyScore,omitempty"`
		
		SurveyTypes *[]string `json:"surveyTypes,omitempty"`
		
		SurveyResponseStatuses *[]string `json:"surveyResponseStatuses,omitempty"`
		
		BotFlowTypes *[]string `json:"botFlowTypes,omitempty"`
		
		AgentTalkDurationMilliseconds *[]Numericrange `json:"agentTalkDurationMilliseconds,omitempty"`
		
		CustomerTalkDurationMilliseconds *[]Numericrange `json:"customerTalkDurationMilliseconds,omitempty"`
		
		OvertalkDurationMilliseconds *[]Numericrange `json:"overtalkDurationMilliseconds,omitempty"`
		
		SilenceDurationMilliseconds *[]Numericrange `json:"silenceDurationMilliseconds,omitempty"`
		
		AcdDurationMilliseconds *[]Numericrange `json:"acdDurationMilliseconds,omitempty"`
		
		IvrDurationMilliseconds *[]Numericrange `json:"ivrDurationMilliseconds,omitempty"`
		
		OtherDurationMilliseconds *[]Numericrange `json:"otherDurationMilliseconds,omitempty"`
		
		AgentTalkPercentage *Numericrange `json:"agentTalkPercentage,omitempty"`
		
		CustomerTalkPercentage *Numericrange `json:"customerTalkPercentage,omitempty"`
		
		OvertalkPercentage *Numericrange `json:"overtalkPercentage,omitempty"`
		
		SilencePercentage *Numericrange `json:"silencePercentage,omitempty"`
		
		AcdPercentage *Numericrange `json:"acdPercentage,omitempty"`
		
		IvrPercentage *Numericrange `json:"ivrPercentage,omitempty"`
		
		OtherPercentage *Numericrange `json:"otherPercentage,omitempty"`
		
		OvertalkInstances *Numericrange `json:"overtalkInstances,omitempty"`
		
		IsScreenRecorded *bool `json:"isScreenRecorded,omitempty"`
		
		ScreenMonitorUserIds *[]string `json:"screenMonitorUserIds,omitempty"`
		
		DashboardState *string `json:"dashboardState,omitempty"`
		
		DashboardType *string `json:"dashboardType,omitempty"`
		
		DashboardAccessFilter *string `json:"dashboardAccessFilter,omitempty"`
		
		TranscriptDurationMilliseconds *[]Numericrange `json:"transcriptDurationMilliseconds,omitempty"`
		
		WorkitemsStatuses *[]Workitemstatusfilter `json:"workitemsStatuses,omitempty"`
		
		SocialCountries *[]string `json:"socialCountries,omitempty"`
		
		SocialLanguages *[]string `json:"socialLanguages,omitempty"`
		
		SocialChannels *[]string `json:"socialChannels,omitempty"`
		
		SocialSentimentCategory *[]string `json:"socialSentimentCategory,omitempty"`
		
		SocialTopicIds *[]string `json:"socialTopicIds,omitempty"`
		
		SocialIngestionRuleIds *[]string `json:"socialIngestionRuleIds,omitempty"`
		
		SocialConversationCreated *bool `json:"socialConversationCreated,omitempty"`
		
		SocialContentType *[]string `json:"socialContentType,omitempty"`
		
		SocialKeywords *[]Socialkeyword `json:"socialKeywords,omitempty"`
		
		SocialPostEscalated *bool `json:"socialPostEscalated,omitempty"`
		
		SocialClassifications *[]string `json:"socialClassifications,omitempty"`
		
		FilterUsersByManagerIds *[]string `json:"filterUsersByManagerIds,omitempty"`
		
		SlideshowIds *[]string `json:"slideshowIds,omitempty"`
		
		Conferenced *bool `json:"conferenced,omitempty"`
		
		Video *bool `json:"video,omitempty"`
		
		LinkedInteraction *bool `json:"linkedInteraction,omitempty"`
		
		RecommendationSources *[]string `json:"recommendationSources,omitempty"`
		
		EvaluationRole *string `json:"evaluationRole,omitempty"`
		
		ComparisonQueueIds *[]string `json:"comparisonQueueIds,omitempty"`
		
		ViewMetrics *[]string `json:"viewMetrics,omitempty"`
		
		TimelineCategories *[]string `json:"timelineCategories,omitempty"`
		
		Acw *bool `json:"acw,omitempty"`
		
		SegmentTypes *[]string `json:"segmentTypes,omitempty"`
		
		ProgramIds *[]string `json:"programIds,omitempty"`
		
		CategoryIds *[]string `json:"categoryIds,omitempty"`
		
		DeliveryPushed *bool `json:"deliveryPushed,omitempty"`
		
		SocialRatings *[]float32 `json:"socialRatings,omitempty"`
		
		VirtualAgentIds *[]string `json:"virtualAgentIds,omitempty"`
		
		EmpathyScoreCategories *[]string `json:"empathyScoreCategories,omitempty"`
		
		SentimentScoreCategories *[]string `json:"sentimentScoreCategories,omitempty"`
		
		SentimentTrendCategories *[]string `json:"sentimentTrendCategories,omitempty"`
		
		ContentModerationFlags *[]string `json:"contentModerationFlags,omitempty"`
		
		SessionExpired *bool `json:"sessionExpired,omitempty"`
		
		EngagementSource *[]string `json:"engagementSource,omitempty"`
		Alias
	}{ 
		MediaTypes: o.MediaTypes,
		
		QueueIds: o.QueueIds,
		
		SkillIds: o.SkillIds,
		
		AssignedSkillIds: o.AssignedSkillIds,
		
		SkillGroups: o.SkillGroups,
		
		LanguageIds: o.LanguageIds,
		
		AssignedLanguageIds: o.AssignedLanguageIds,
		
		LanguageGroups: o.LanguageGroups,
		
		Directions: o.Directions,
		
		OriginatingDirections: o.OriginatingDirections,
		
		WrapUpCodes: o.WrapUpCodes,
		
		DnisList: o.DnisList,
		
		SessionDnisList: o.SessionDnisList,
		
		FilterQueuesByUserIds: o.FilterQueuesByUserIds,
		
		FilterUsersByQueueIds: o.FilterUsersByQueueIds,
		
		UserIds: o.UserIds,
		
		ManagementUnitIds: o.ManagementUnitIds,
		
		AddressTos: o.AddressTos,
		
		AddressFroms: o.AddressFroms,
		
		OutboundCampaignIds: o.OutboundCampaignIds,
		
		OutboundContactListIds: o.OutboundContactListIds,
		
		ContactIds: o.ContactIds,
		
		ExternalContactIds: o.ExternalContactIds,
		
		ExternalOrgIds: o.ExternalOrgIds,
		
		AniList: o.AniList,
		
		DurationsMilliseconds: o.DurationsMilliseconds,
		
		AcdDurationsMilliseconds: o.AcdDurationsMilliseconds,
		
		TalkDurationsMilliseconds: o.TalkDurationsMilliseconds,
		
		AcwDurationsMilliseconds: o.AcwDurationsMilliseconds,
		
		HandleDurationsMilliseconds: o.HandleDurationsMilliseconds,
		
		HoldDurationsMilliseconds: o.HoldDurationsMilliseconds,
		
		AbandonDurationsMilliseconds: o.AbandonDurationsMilliseconds,
		
		EvaluationScore: o.EvaluationScore,
		
		EvaluationCriticalScore: o.EvaluationCriticalScore,
		
		EvaluationFormIds: o.EvaluationFormIds,
		
		EvaluatedAgentIds: o.EvaluatedAgentIds,
		
		EvaluatorIds: o.EvaluatorIds,
		
		Transferred: o.Transferred,
		
		Abandoned: o.Abandoned,
		
		Answered: o.Answered,
		
		MessageTypes: o.MessageTypes,
		
		DivisionIds: o.DivisionIds,
		
		SurveyFormIds: o.SurveyFormIds,
		
		SurveyTotalScore: o.SurveyTotalScore,
		
		SurveyNpsScore: o.SurveyNpsScore,
		
		Mos: o.Mos,
		
		SurveyQuestionGroupScore: o.SurveyQuestionGroupScore,
		
		SurveyPromoterScore: o.SurveyPromoterScore,
		
		SurveyFormContextIds: o.SurveyFormContextIds,
		
		ConversationIds: o.ConversationIds,
		
		SipCallIds: o.SipCallIds,
		
		IsEnded: o.IsEnded,
		
		IsSurveyed: o.IsSurveyed,
		
		SurveyScores: o.SurveyScores,
		
		PromoterScores: o.PromoterScores,
		
		IsCampaign: o.IsCampaign,
		
		SurveyStatuses: o.SurveyStatuses,
		
		ConversationProperties: o.ConversationProperties,
		
		IsBlindTransferred: o.IsBlindTransferred,
		
		IsConsulted: o.IsConsulted,
		
		IsConsultTransferred: o.IsConsultTransferred,
		
		RemoteParticipants: o.RemoteParticipants,
		
		FlowIds: o.FlowIds,
		
		FlowOutcomeIds: o.FlowOutcomeIds,
		
		FlowOutcomeValues: o.FlowOutcomeValues,
		
		FlowDestinationTypes: o.FlowDestinationTypes,
		
		FlowDisconnectReasons: o.FlowDisconnectReasons,
		
		FlowTypes: o.FlowTypes,
		
		FlowEntryTypes: o.FlowEntryTypes,
		
		FlowEntryReasons: o.FlowEntryReasons,
		
		FlowVersions: o.FlowVersions,
		
		GroupIds: o.GroupIds,
		
		HasJourneyCustomerId: o.HasJourneyCustomerId,
		
		HasJourneyActionMapId: o.HasJourneyActionMapId,
		
		HasJourneyVisitId: o.HasJourneyVisitId,
		
		HasMedia: o.HasMedia,
		
		RoleIds: o.RoleIds,
		
		ReportsTos: o.ReportsTos,
		
		LocationIds: o.LocationIds,
		
		FlowOutTypes: o.FlowOutTypes,
		
		ProviderList: o.ProviderList,
		
		CallbackNumberList: o.CallbackNumberList,
		
		CallbackInterval: o.CallbackInterval,
		
		UsedRoutingTypes: o.UsedRoutingTypes,
		
		RequestedRoutingTypes: o.RequestedRoutingTypes,
		
		HasAgentAssistId: o.HasAgentAssistId,
		
		Transcripts: o.Transcripts,
		
		TranscriptLanguages: o.TranscriptLanguages,
		
		ParticipantPurposes: o.ParticipantPurposes,
		
		ShowFirstQueue: o.ShowFirstQueue,
		
		TeamIds: o.TeamIds,
		
		FilterUsersByTeamIds: o.FilterUsersByTeamIds,
		
		JourneyActionMapIds: o.JourneyActionMapIds,
		
		JourneyOutcomeIds: o.JourneyOutcomeIds,
		
		JourneySegmentIds: o.JourneySegmentIds,
		
		JourneyActionMapTypes: o.JourneyActionMapTypes,
		
		DevelopmentRoleList: o.DevelopmentRoleList,
		
		DevelopmentTypeList: o.DevelopmentTypeList,
		
		DevelopmentStatusList: o.DevelopmentStatusList,
		
		DevelopmentModuleIds: o.DevelopmentModuleIds,
		
		DevelopmentActivityOverdue: o.DevelopmentActivityOverdue,
		
		CustomerSentimentScore: o.CustomerSentimentScore,
		
		CustomerSentimentTrend: o.CustomerSentimentTrend,
		
		FlowTransferTargets: o.FlowTransferTargets,
		
		DevelopmentName: o.DevelopmentName,
		
		TopicIds: o.TopicIds,
		
		ExternalTags: o.ExternalTags,
		
		IsNotResponding: o.IsNotResponding,
		
		IsAuthenticated: o.IsAuthenticated,
		
		BotIds: o.BotIds,
		
		BotVersions: o.BotVersions,
		
		BotMessageTypes: o.BotMessageTypes,
		
		BotProviderList: o.BotProviderList,
		
		BotProductList: o.BotProductList,
		
		BotRecognitionFailureReasonList: o.BotRecognitionFailureReasonList,
		
		BotIntentList: o.BotIntentList,
		
		BotFinalIntentList: o.BotFinalIntentList,
		
		BotSlotList: o.BotSlotList,
		
		BotResultList: o.BotResultList,
		
		BlockedReasons: o.BlockedReasons,
		
		IsRecorded: o.IsRecorded,
		
		HasEvaluation: o.HasEvaluation,
		
		HasScoredEvaluation: o.HasScoredEvaluation,
		
		EmailDeliveryStatusList: o.EmailDeliveryStatusList,
		
		IsAgentOwnedCallback: o.IsAgentOwnedCallback,
		
		AgentCallbackOwnerIds: o.AgentCallbackOwnerIds,
		
		TranscriptTopics: o.TranscriptTopics,
		
		JourneyFrequencyCapReasons: o.JourneyFrequencyCapReasons,
		
		JourneyBlockingActionMapIds: o.JourneyBlockingActionMapIds,
		
		JourneyActionTargetIds: o.JourneyActionTargetIds,
		
		JourneyBlockingScheduleGroupIds: o.JourneyBlockingScheduleGroupIds,
		
		JourneyBlockingEmergencyScheduleGroupIds: o.JourneyBlockingEmergencyScheduleGroupIds,
		
		JourneyUrlEqualConditions: o.JourneyUrlEqualConditions,
		
		JourneyUrlNotEqualConditions: o.JourneyUrlNotEqualConditions,
		
		JourneyUrlStartsWithConditions: o.JourneyUrlStartsWithConditions,
		
		JourneyUrlEndsWithConditions: o.JourneyUrlEndsWithConditions,
		
		JourneyUrlContainsAnyConditions: o.JourneyUrlContainsAnyConditions,
		
		JourneyUrlNotContainsAnyConditions: o.JourneyUrlNotContainsAnyConditions,
		
		JourneyUrlContainsAllConditions: o.JourneyUrlContainsAllConditions,
		
		JourneyUrlNotContainsAllConditions: o.JourneyUrlNotContainsAllConditions,
		
		FlowMilestoneIds: o.FlowMilestoneIds,
		
		IsAssessmentPassed: o.IsAssessmentPassed,
		
		ConversationInitiators: o.ConversationInitiators,
		
		HasCustomerParticipated: o.HasCustomerParticipated,
		
		IsAcdInteraction: o.IsAcdInteraction,
		
		HasFax: o.HasFax,
		
		DataActionIds: o.DataActionIds,
		
		ActionCategoryName: o.ActionCategoryName,
		
		IntegrationIds: o.IntegrationIds,
		
		ResponseStatuses: o.ResponseStatuses,
		
		AvailableDashboard: o.AvailableDashboard,
		
		FavouriteDashboard: o.FavouriteDashboard,
		
		MyDashboard: o.MyDashboard,
		
		StationErrors: o.StationErrors,
		
		CanonicalContactIds: o.CanonicalContactIds,
		
		AlertRuleIds: o.AlertRuleIds,
		
		EvaluationFormContextIds: o.EvaluationFormContextIds,
		
		EvaluationStatuses: o.EvaluationStatuses,
		
		WorkbinIds: o.WorkbinIds,
		
		WorktypeIds: o.WorktypeIds,
		
		WorkitemIds: o.WorkitemIds,
		
		WorkitemAssigneeIds: o.WorkitemAssigneeIds,
		
		WorkitemStatuses: o.WorkitemStatuses,
		
		IsAnalyzedForSensitiveData: o.IsAnalyzedForSensitiveData,
		
		HasSensitiveData: o.HasSensitiveData,
		
		HasPciData: o.HasPciData,
		
		HasPiiData: o.HasPiiData,
		
		SubPath: o.SubPath,
		
		UserState: o.UserState,
		
		IsClearedByCustomer: o.IsClearedByCustomer,
		
		EvaluationAssigneeIds: o.EvaluationAssigneeIds,
		
		EvaluationAssigned: o.EvaluationAssigned,
		
		AssistantIds: o.AssistantIds,
		
		KnowledgeBaseIds: o.KnowledgeBaseIds,
		
		IsParked: o.IsParked,
		
		AgentEmpathyScore: o.AgentEmpathyScore,
		
		SurveyTypes: o.SurveyTypes,
		
		SurveyResponseStatuses: o.SurveyResponseStatuses,
		
		BotFlowTypes: o.BotFlowTypes,
		
		AgentTalkDurationMilliseconds: o.AgentTalkDurationMilliseconds,
		
		CustomerTalkDurationMilliseconds: o.CustomerTalkDurationMilliseconds,
		
		OvertalkDurationMilliseconds: o.OvertalkDurationMilliseconds,
		
		SilenceDurationMilliseconds: o.SilenceDurationMilliseconds,
		
		AcdDurationMilliseconds: o.AcdDurationMilliseconds,
		
		IvrDurationMilliseconds: o.IvrDurationMilliseconds,
		
		OtherDurationMilliseconds: o.OtherDurationMilliseconds,
		
		AgentTalkPercentage: o.AgentTalkPercentage,
		
		CustomerTalkPercentage: o.CustomerTalkPercentage,
		
		OvertalkPercentage: o.OvertalkPercentage,
		
		SilencePercentage: o.SilencePercentage,
		
		AcdPercentage: o.AcdPercentage,
		
		IvrPercentage: o.IvrPercentage,
		
		OtherPercentage: o.OtherPercentage,
		
		OvertalkInstances: o.OvertalkInstances,
		
		IsScreenRecorded: o.IsScreenRecorded,
		
		ScreenMonitorUserIds: o.ScreenMonitorUserIds,
		
		DashboardState: o.DashboardState,
		
		DashboardType: o.DashboardType,
		
		DashboardAccessFilter: o.DashboardAccessFilter,
		
		TranscriptDurationMilliseconds: o.TranscriptDurationMilliseconds,
		
		WorkitemsStatuses: o.WorkitemsStatuses,
		
		SocialCountries: o.SocialCountries,
		
		SocialLanguages: o.SocialLanguages,
		
		SocialChannels: o.SocialChannels,
		
		SocialSentimentCategory: o.SocialSentimentCategory,
		
		SocialTopicIds: o.SocialTopicIds,
		
		SocialIngestionRuleIds: o.SocialIngestionRuleIds,
		
		SocialConversationCreated: o.SocialConversationCreated,
		
		SocialContentType: o.SocialContentType,
		
		SocialKeywords: o.SocialKeywords,
		
		SocialPostEscalated: o.SocialPostEscalated,
		
		SocialClassifications: o.SocialClassifications,
		
		FilterUsersByManagerIds: o.FilterUsersByManagerIds,
		
		SlideshowIds: o.SlideshowIds,
		
		Conferenced: o.Conferenced,
		
		Video: o.Video,
		
		LinkedInteraction: o.LinkedInteraction,
		
		RecommendationSources: o.RecommendationSources,
		
		EvaluationRole: o.EvaluationRole,
		
		ComparisonQueueIds: o.ComparisonQueueIds,
		
		ViewMetrics: o.ViewMetrics,
		
		TimelineCategories: o.TimelineCategories,
		
		Acw: o.Acw,
		
		SegmentTypes: o.SegmentTypes,
		
		ProgramIds: o.ProgramIds,
		
		CategoryIds: o.CategoryIds,
		
		DeliveryPushed: o.DeliveryPushed,
		
		SocialRatings: o.SocialRatings,
		
		VirtualAgentIds: o.VirtualAgentIds,
		
		EmpathyScoreCategories: o.EmpathyScoreCategories,
		
		SentimentScoreCategories: o.SentimentScoreCategories,
		
		SentimentTrendCategories: o.SentimentTrendCategories,
		
		ContentModerationFlags: o.ContentModerationFlags,
		
		SessionExpired: o.SessionExpired,
		
		EngagementSource: o.EngagementSource,
		Alias:    (Alias)(o),
	})
}

func (o *Viewfilter) UnmarshalJSON(b []byte) error {
	var ViewfilterMap map[string]interface{}
	err := json.Unmarshal(b, &ViewfilterMap)
	if err != nil {
		return err
	}
	
	if MediaTypes, ok := ViewfilterMap["mediaTypes"].([]interface{}); ok {
		MediaTypesString, _ := json.Marshal(MediaTypes)
		json.Unmarshal(MediaTypesString, &o.MediaTypes)
	}
	
	if QueueIds, ok := ViewfilterMap["queueIds"].([]interface{}); ok {
		QueueIdsString, _ := json.Marshal(QueueIds)
		json.Unmarshal(QueueIdsString, &o.QueueIds)
	}
	
	if SkillIds, ok := ViewfilterMap["skillIds"].([]interface{}); ok {
		SkillIdsString, _ := json.Marshal(SkillIds)
		json.Unmarshal(SkillIdsString, &o.SkillIds)
	}
	
	if AssignedSkillIds, ok := ViewfilterMap["assignedSkillIds"].([]interface{}); ok {
		AssignedSkillIdsString, _ := json.Marshal(AssignedSkillIds)
		json.Unmarshal(AssignedSkillIdsString, &o.AssignedSkillIds)
	}
	
	if SkillGroups, ok := ViewfilterMap["skillGroups"].([]interface{}); ok {
		SkillGroupsString, _ := json.Marshal(SkillGroups)
		json.Unmarshal(SkillGroupsString, &o.SkillGroups)
	}
	
	if LanguageIds, ok := ViewfilterMap["languageIds"].([]interface{}); ok {
		LanguageIdsString, _ := json.Marshal(LanguageIds)
		json.Unmarshal(LanguageIdsString, &o.LanguageIds)
	}
	
	if AssignedLanguageIds, ok := ViewfilterMap["assignedLanguageIds"].([]interface{}); ok {
		AssignedLanguageIdsString, _ := json.Marshal(AssignedLanguageIds)
		json.Unmarshal(AssignedLanguageIdsString, &o.AssignedLanguageIds)
	}
	
	if LanguageGroups, ok := ViewfilterMap["languageGroups"].([]interface{}); ok {
		LanguageGroupsString, _ := json.Marshal(LanguageGroups)
		json.Unmarshal(LanguageGroupsString, &o.LanguageGroups)
	}
	
	if Directions, ok := ViewfilterMap["directions"].([]interface{}); ok {
		DirectionsString, _ := json.Marshal(Directions)
		json.Unmarshal(DirectionsString, &o.Directions)
	}
	
	if OriginatingDirections, ok := ViewfilterMap["originatingDirections"].([]interface{}); ok {
		OriginatingDirectionsString, _ := json.Marshal(OriginatingDirections)
		json.Unmarshal(OriginatingDirectionsString, &o.OriginatingDirections)
	}
	
	if WrapUpCodes, ok := ViewfilterMap["wrapUpCodes"].([]interface{}); ok {
		WrapUpCodesString, _ := json.Marshal(WrapUpCodes)
		json.Unmarshal(WrapUpCodesString, &o.WrapUpCodes)
	}
	
	if DnisList, ok := ViewfilterMap["dnisList"].([]interface{}); ok {
		DnisListString, _ := json.Marshal(DnisList)
		json.Unmarshal(DnisListString, &o.DnisList)
	}
	
	if SessionDnisList, ok := ViewfilterMap["sessionDnisList"].([]interface{}); ok {
		SessionDnisListString, _ := json.Marshal(SessionDnisList)
		json.Unmarshal(SessionDnisListString, &o.SessionDnisList)
	}
	
	if FilterQueuesByUserIds, ok := ViewfilterMap["filterQueuesByUserIds"].([]interface{}); ok {
		FilterQueuesByUserIdsString, _ := json.Marshal(FilterQueuesByUserIds)
		json.Unmarshal(FilterQueuesByUserIdsString, &o.FilterQueuesByUserIds)
	}
	
	if FilterUsersByQueueIds, ok := ViewfilterMap["filterUsersByQueueIds"].([]interface{}); ok {
		FilterUsersByQueueIdsString, _ := json.Marshal(FilterUsersByQueueIds)
		json.Unmarshal(FilterUsersByQueueIdsString, &o.FilterUsersByQueueIds)
	}
	
	if UserIds, ok := ViewfilterMap["userIds"].([]interface{}); ok {
		UserIdsString, _ := json.Marshal(UserIds)
		json.Unmarshal(UserIdsString, &o.UserIds)
	}
	
	if ManagementUnitIds, ok := ViewfilterMap["managementUnitIds"].([]interface{}); ok {
		ManagementUnitIdsString, _ := json.Marshal(ManagementUnitIds)
		json.Unmarshal(ManagementUnitIdsString, &o.ManagementUnitIds)
	}
	
	if AddressTos, ok := ViewfilterMap["addressTos"].([]interface{}); ok {
		AddressTosString, _ := json.Marshal(AddressTos)
		json.Unmarshal(AddressTosString, &o.AddressTos)
	}
	
	if AddressFroms, ok := ViewfilterMap["addressFroms"].([]interface{}); ok {
		AddressFromsString, _ := json.Marshal(AddressFroms)
		json.Unmarshal(AddressFromsString, &o.AddressFroms)
	}
	
	if OutboundCampaignIds, ok := ViewfilterMap["outboundCampaignIds"].([]interface{}); ok {
		OutboundCampaignIdsString, _ := json.Marshal(OutboundCampaignIds)
		json.Unmarshal(OutboundCampaignIdsString, &o.OutboundCampaignIds)
	}
	
	if OutboundContactListIds, ok := ViewfilterMap["outboundContactListIds"].([]interface{}); ok {
		OutboundContactListIdsString, _ := json.Marshal(OutboundContactListIds)
		json.Unmarshal(OutboundContactListIdsString, &o.OutboundContactListIds)
	}
	
	if ContactIds, ok := ViewfilterMap["contactIds"].([]interface{}); ok {
		ContactIdsString, _ := json.Marshal(ContactIds)
		json.Unmarshal(ContactIdsString, &o.ContactIds)
	}
	
	if ExternalContactIds, ok := ViewfilterMap["externalContactIds"].([]interface{}); ok {
		ExternalContactIdsString, _ := json.Marshal(ExternalContactIds)
		json.Unmarshal(ExternalContactIdsString, &o.ExternalContactIds)
	}
	
	if ExternalOrgIds, ok := ViewfilterMap["externalOrgIds"].([]interface{}); ok {
		ExternalOrgIdsString, _ := json.Marshal(ExternalOrgIds)
		json.Unmarshal(ExternalOrgIdsString, &o.ExternalOrgIds)
	}
	
	if AniList, ok := ViewfilterMap["aniList"].([]interface{}); ok {
		AniListString, _ := json.Marshal(AniList)
		json.Unmarshal(AniListString, &o.AniList)
	}
	
	if DurationsMilliseconds, ok := ViewfilterMap["durationsMilliseconds"].([]interface{}); ok {
		DurationsMillisecondsString, _ := json.Marshal(DurationsMilliseconds)
		json.Unmarshal(DurationsMillisecondsString, &o.DurationsMilliseconds)
	}
	
	if AcdDurationsMilliseconds, ok := ViewfilterMap["acdDurationsMilliseconds"].([]interface{}); ok {
		AcdDurationsMillisecondsString, _ := json.Marshal(AcdDurationsMilliseconds)
		json.Unmarshal(AcdDurationsMillisecondsString, &o.AcdDurationsMilliseconds)
	}
	
	if TalkDurationsMilliseconds, ok := ViewfilterMap["talkDurationsMilliseconds"].([]interface{}); ok {
		TalkDurationsMillisecondsString, _ := json.Marshal(TalkDurationsMilliseconds)
		json.Unmarshal(TalkDurationsMillisecondsString, &o.TalkDurationsMilliseconds)
	}
	
	if AcwDurationsMilliseconds, ok := ViewfilterMap["acwDurationsMilliseconds"].([]interface{}); ok {
		AcwDurationsMillisecondsString, _ := json.Marshal(AcwDurationsMilliseconds)
		json.Unmarshal(AcwDurationsMillisecondsString, &o.AcwDurationsMilliseconds)
	}
	
	if HandleDurationsMilliseconds, ok := ViewfilterMap["handleDurationsMilliseconds"].([]interface{}); ok {
		HandleDurationsMillisecondsString, _ := json.Marshal(HandleDurationsMilliseconds)
		json.Unmarshal(HandleDurationsMillisecondsString, &o.HandleDurationsMilliseconds)
	}
	
	if HoldDurationsMilliseconds, ok := ViewfilterMap["holdDurationsMilliseconds"].([]interface{}); ok {
		HoldDurationsMillisecondsString, _ := json.Marshal(HoldDurationsMilliseconds)
		json.Unmarshal(HoldDurationsMillisecondsString, &o.HoldDurationsMilliseconds)
	}
	
	if AbandonDurationsMilliseconds, ok := ViewfilterMap["abandonDurationsMilliseconds"].([]interface{}); ok {
		AbandonDurationsMillisecondsString, _ := json.Marshal(AbandonDurationsMilliseconds)
		json.Unmarshal(AbandonDurationsMillisecondsString, &o.AbandonDurationsMilliseconds)
	}
	
	if EvaluationScore, ok := ViewfilterMap["evaluationScore"].(map[string]interface{}); ok {
		EvaluationScoreString, _ := json.Marshal(EvaluationScore)
		json.Unmarshal(EvaluationScoreString, &o.EvaluationScore)
	}
	
	if EvaluationCriticalScore, ok := ViewfilterMap["evaluationCriticalScore"].(map[string]interface{}); ok {
		EvaluationCriticalScoreString, _ := json.Marshal(EvaluationCriticalScore)
		json.Unmarshal(EvaluationCriticalScoreString, &o.EvaluationCriticalScore)
	}
	
	if EvaluationFormIds, ok := ViewfilterMap["evaluationFormIds"].([]interface{}); ok {
		EvaluationFormIdsString, _ := json.Marshal(EvaluationFormIds)
		json.Unmarshal(EvaluationFormIdsString, &o.EvaluationFormIds)
	}
	
	if EvaluatedAgentIds, ok := ViewfilterMap["evaluatedAgentIds"].([]interface{}); ok {
		EvaluatedAgentIdsString, _ := json.Marshal(EvaluatedAgentIds)
		json.Unmarshal(EvaluatedAgentIdsString, &o.EvaluatedAgentIds)
	}
	
	if EvaluatorIds, ok := ViewfilterMap["evaluatorIds"].([]interface{}); ok {
		EvaluatorIdsString, _ := json.Marshal(EvaluatorIds)
		json.Unmarshal(EvaluatorIdsString, &o.EvaluatorIds)
	}
	
	if Transferred, ok := ViewfilterMap["transferred"].(bool); ok {
		o.Transferred = &Transferred
	}
    
	if Abandoned, ok := ViewfilterMap["abandoned"].(bool); ok {
		o.Abandoned = &Abandoned
	}
    
	if Answered, ok := ViewfilterMap["answered"].(bool); ok {
		o.Answered = &Answered
	}
    
	if MessageTypes, ok := ViewfilterMap["messageTypes"].([]interface{}); ok {
		MessageTypesString, _ := json.Marshal(MessageTypes)
		json.Unmarshal(MessageTypesString, &o.MessageTypes)
	}
	
	if DivisionIds, ok := ViewfilterMap["divisionIds"].([]interface{}); ok {
		DivisionIdsString, _ := json.Marshal(DivisionIds)
		json.Unmarshal(DivisionIdsString, &o.DivisionIds)
	}
	
	if SurveyFormIds, ok := ViewfilterMap["surveyFormIds"].([]interface{}); ok {
		SurveyFormIdsString, _ := json.Marshal(SurveyFormIds)
		json.Unmarshal(SurveyFormIdsString, &o.SurveyFormIds)
	}
	
	if SurveyTotalScore, ok := ViewfilterMap["surveyTotalScore"].(map[string]interface{}); ok {
		SurveyTotalScoreString, _ := json.Marshal(SurveyTotalScore)
		json.Unmarshal(SurveyTotalScoreString, &o.SurveyTotalScore)
	}
	
	if SurveyNpsScore, ok := ViewfilterMap["surveyNpsScore"].(map[string]interface{}); ok {
		SurveyNpsScoreString, _ := json.Marshal(SurveyNpsScore)
		json.Unmarshal(SurveyNpsScoreString, &o.SurveyNpsScore)
	}
	
	if Mos, ok := ViewfilterMap["mos"].(map[string]interface{}); ok {
		MosString, _ := json.Marshal(Mos)
		json.Unmarshal(MosString, &o.Mos)
	}
	
	if SurveyQuestionGroupScore, ok := ViewfilterMap["surveyQuestionGroupScore"].(map[string]interface{}); ok {
		SurveyQuestionGroupScoreString, _ := json.Marshal(SurveyQuestionGroupScore)
		json.Unmarshal(SurveyQuestionGroupScoreString, &o.SurveyQuestionGroupScore)
	}
	
	if SurveyPromoterScore, ok := ViewfilterMap["surveyPromoterScore"].(map[string]interface{}); ok {
		SurveyPromoterScoreString, _ := json.Marshal(SurveyPromoterScore)
		json.Unmarshal(SurveyPromoterScoreString, &o.SurveyPromoterScore)
	}
	
	if SurveyFormContextIds, ok := ViewfilterMap["surveyFormContextIds"].([]interface{}); ok {
		SurveyFormContextIdsString, _ := json.Marshal(SurveyFormContextIds)
		json.Unmarshal(SurveyFormContextIdsString, &o.SurveyFormContextIds)
	}
	
	if ConversationIds, ok := ViewfilterMap["conversationIds"].([]interface{}); ok {
		ConversationIdsString, _ := json.Marshal(ConversationIds)
		json.Unmarshal(ConversationIdsString, &o.ConversationIds)
	}
	
	if SipCallIds, ok := ViewfilterMap["sipCallIds"].([]interface{}); ok {
		SipCallIdsString, _ := json.Marshal(SipCallIds)
		json.Unmarshal(SipCallIdsString, &o.SipCallIds)
	}
	
	if IsEnded, ok := ViewfilterMap["isEnded"].(bool); ok {
		o.IsEnded = &IsEnded
	}
    
	if IsSurveyed, ok := ViewfilterMap["isSurveyed"].(bool); ok {
		o.IsSurveyed = &IsSurveyed
	}
    
	if SurveyScores, ok := ViewfilterMap["surveyScores"].([]interface{}); ok {
		SurveyScoresString, _ := json.Marshal(SurveyScores)
		json.Unmarshal(SurveyScoresString, &o.SurveyScores)
	}
	
	if PromoterScores, ok := ViewfilterMap["promoterScores"].([]interface{}); ok {
		PromoterScoresString, _ := json.Marshal(PromoterScores)
		json.Unmarshal(PromoterScoresString, &o.PromoterScores)
	}
	
	if IsCampaign, ok := ViewfilterMap["isCampaign"].(bool); ok {
		o.IsCampaign = &IsCampaign
	}
    
	if SurveyStatuses, ok := ViewfilterMap["surveyStatuses"].([]interface{}); ok {
		SurveyStatusesString, _ := json.Marshal(SurveyStatuses)
		json.Unmarshal(SurveyStatusesString, &o.SurveyStatuses)
	}
	
	if ConversationProperties, ok := ViewfilterMap["conversationProperties"].(map[string]interface{}); ok {
		ConversationPropertiesString, _ := json.Marshal(ConversationProperties)
		json.Unmarshal(ConversationPropertiesString, &o.ConversationProperties)
	}
	
	if IsBlindTransferred, ok := ViewfilterMap["isBlindTransferred"].(bool); ok {
		o.IsBlindTransferred = &IsBlindTransferred
	}
    
	if IsConsulted, ok := ViewfilterMap["isConsulted"].(bool); ok {
		o.IsConsulted = &IsConsulted
	}
    
	if IsConsultTransferred, ok := ViewfilterMap["isConsultTransferred"].(bool); ok {
		o.IsConsultTransferred = &IsConsultTransferred
	}
    
	if RemoteParticipants, ok := ViewfilterMap["remoteParticipants"].([]interface{}); ok {
		RemoteParticipantsString, _ := json.Marshal(RemoteParticipants)
		json.Unmarshal(RemoteParticipantsString, &o.RemoteParticipants)
	}
	
	if FlowIds, ok := ViewfilterMap["flowIds"].([]interface{}); ok {
		FlowIdsString, _ := json.Marshal(FlowIds)
		json.Unmarshal(FlowIdsString, &o.FlowIds)
	}
	
	if FlowOutcomeIds, ok := ViewfilterMap["flowOutcomeIds"].([]interface{}); ok {
		FlowOutcomeIdsString, _ := json.Marshal(FlowOutcomeIds)
		json.Unmarshal(FlowOutcomeIdsString, &o.FlowOutcomeIds)
	}
	
	if FlowOutcomeValues, ok := ViewfilterMap["flowOutcomeValues"].([]interface{}); ok {
		FlowOutcomeValuesString, _ := json.Marshal(FlowOutcomeValues)
		json.Unmarshal(FlowOutcomeValuesString, &o.FlowOutcomeValues)
	}
	
	if FlowDestinationTypes, ok := ViewfilterMap["flowDestinationTypes"].([]interface{}); ok {
		FlowDestinationTypesString, _ := json.Marshal(FlowDestinationTypes)
		json.Unmarshal(FlowDestinationTypesString, &o.FlowDestinationTypes)
	}
	
	if FlowDisconnectReasons, ok := ViewfilterMap["flowDisconnectReasons"].([]interface{}); ok {
		FlowDisconnectReasonsString, _ := json.Marshal(FlowDisconnectReasons)
		json.Unmarshal(FlowDisconnectReasonsString, &o.FlowDisconnectReasons)
	}
	
	if FlowTypes, ok := ViewfilterMap["flowTypes"].([]interface{}); ok {
		FlowTypesString, _ := json.Marshal(FlowTypes)
		json.Unmarshal(FlowTypesString, &o.FlowTypes)
	}
	
	if FlowEntryTypes, ok := ViewfilterMap["flowEntryTypes"].([]interface{}); ok {
		FlowEntryTypesString, _ := json.Marshal(FlowEntryTypes)
		json.Unmarshal(FlowEntryTypesString, &o.FlowEntryTypes)
	}
	
	if FlowEntryReasons, ok := ViewfilterMap["flowEntryReasons"].([]interface{}); ok {
		FlowEntryReasonsString, _ := json.Marshal(FlowEntryReasons)
		json.Unmarshal(FlowEntryReasonsString, &o.FlowEntryReasons)
	}
	
	if FlowVersions, ok := ViewfilterMap["flowVersions"].([]interface{}); ok {
		FlowVersionsString, _ := json.Marshal(FlowVersions)
		json.Unmarshal(FlowVersionsString, &o.FlowVersions)
	}
	
	if GroupIds, ok := ViewfilterMap["groupIds"].([]interface{}); ok {
		GroupIdsString, _ := json.Marshal(GroupIds)
		json.Unmarshal(GroupIdsString, &o.GroupIds)
	}
	
	if HasJourneyCustomerId, ok := ViewfilterMap["hasJourneyCustomerId"].(bool); ok {
		o.HasJourneyCustomerId = &HasJourneyCustomerId
	}
    
	if HasJourneyActionMapId, ok := ViewfilterMap["hasJourneyActionMapId"].(bool); ok {
		o.HasJourneyActionMapId = &HasJourneyActionMapId
	}
    
	if HasJourneyVisitId, ok := ViewfilterMap["hasJourneyVisitId"].(bool); ok {
		o.HasJourneyVisitId = &HasJourneyVisitId
	}
    
	if HasMedia, ok := ViewfilterMap["hasMedia"].(bool); ok {
		o.HasMedia = &HasMedia
	}
    
	if RoleIds, ok := ViewfilterMap["roleIds"].([]interface{}); ok {
		RoleIdsString, _ := json.Marshal(RoleIds)
		json.Unmarshal(RoleIdsString, &o.RoleIds)
	}
	
	if ReportsTos, ok := ViewfilterMap["reportsTos"].([]interface{}); ok {
		ReportsTosString, _ := json.Marshal(ReportsTos)
		json.Unmarshal(ReportsTosString, &o.ReportsTos)
	}
	
	if LocationIds, ok := ViewfilterMap["locationIds"].([]interface{}); ok {
		LocationIdsString, _ := json.Marshal(LocationIds)
		json.Unmarshal(LocationIdsString, &o.LocationIds)
	}
	
	if FlowOutTypes, ok := ViewfilterMap["flowOutTypes"].([]interface{}); ok {
		FlowOutTypesString, _ := json.Marshal(FlowOutTypes)
		json.Unmarshal(FlowOutTypesString, &o.FlowOutTypes)
	}
	
	if ProviderList, ok := ViewfilterMap["providerList"].([]interface{}); ok {
		ProviderListString, _ := json.Marshal(ProviderList)
		json.Unmarshal(ProviderListString, &o.ProviderList)
	}
	
	if CallbackNumberList, ok := ViewfilterMap["callbackNumberList"].([]interface{}); ok {
		CallbackNumberListString, _ := json.Marshal(CallbackNumberList)
		json.Unmarshal(CallbackNumberListString, &o.CallbackNumberList)
	}
	
	if CallbackInterval, ok := ViewfilterMap["callbackInterval"].(string); ok {
		o.CallbackInterval = &CallbackInterval
	}
    
	if UsedRoutingTypes, ok := ViewfilterMap["usedRoutingTypes"].([]interface{}); ok {
		UsedRoutingTypesString, _ := json.Marshal(UsedRoutingTypes)
		json.Unmarshal(UsedRoutingTypesString, &o.UsedRoutingTypes)
	}
	
	if RequestedRoutingTypes, ok := ViewfilterMap["requestedRoutingTypes"].([]interface{}); ok {
		RequestedRoutingTypesString, _ := json.Marshal(RequestedRoutingTypes)
		json.Unmarshal(RequestedRoutingTypesString, &o.RequestedRoutingTypes)
	}
	
	if HasAgentAssistId, ok := ViewfilterMap["hasAgentAssistId"].(bool); ok {
		o.HasAgentAssistId = &HasAgentAssistId
	}
    
	if Transcripts, ok := ViewfilterMap["transcripts"].([]interface{}); ok {
		TranscriptsString, _ := json.Marshal(Transcripts)
		json.Unmarshal(TranscriptsString, &o.Transcripts)
	}
	
	if TranscriptLanguages, ok := ViewfilterMap["transcriptLanguages"].([]interface{}); ok {
		TranscriptLanguagesString, _ := json.Marshal(TranscriptLanguages)
		json.Unmarshal(TranscriptLanguagesString, &o.TranscriptLanguages)
	}
	
	if ParticipantPurposes, ok := ViewfilterMap["participantPurposes"].([]interface{}); ok {
		ParticipantPurposesString, _ := json.Marshal(ParticipantPurposes)
		json.Unmarshal(ParticipantPurposesString, &o.ParticipantPurposes)
	}
	
	if ShowFirstQueue, ok := ViewfilterMap["showFirstQueue"].(bool); ok {
		o.ShowFirstQueue = &ShowFirstQueue
	}
    
	if TeamIds, ok := ViewfilterMap["teamIds"].([]interface{}); ok {
		TeamIdsString, _ := json.Marshal(TeamIds)
		json.Unmarshal(TeamIdsString, &o.TeamIds)
	}
	
	if FilterUsersByTeamIds, ok := ViewfilterMap["filterUsersByTeamIds"].([]interface{}); ok {
		FilterUsersByTeamIdsString, _ := json.Marshal(FilterUsersByTeamIds)
		json.Unmarshal(FilterUsersByTeamIdsString, &o.FilterUsersByTeamIds)
	}
	
	if JourneyActionMapIds, ok := ViewfilterMap["journeyActionMapIds"].([]interface{}); ok {
		JourneyActionMapIdsString, _ := json.Marshal(JourneyActionMapIds)
		json.Unmarshal(JourneyActionMapIdsString, &o.JourneyActionMapIds)
	}
	
	if JourneyOutcomeIds, ok := ViewfilterMap["journeyOutcomeIds"].([]interface{}); ok {
		JourneyOutcomeIdsString, _ := json.Marshal(JourneyOutcomeIds)
		json.Unmarshal(JourneyOutcomeIdsString, &o.JourneyOutcomeIds)
	}
	
	if JourneySegmentIds, ok := ViewfilterMap["journeySegmentIds"].([]interface{}); ok {
		JourneySegmentIdsString, _ := json.Marshal(JourneySegmentIds)
		json.Unmarshal(JourneySegmentIdsString, &o.JourneySegmentIds)
	}
	
	if JourneyActionMapTypes, ok := ViewfilterMap["journeyActionMapTypes"].([]interface{}); ok {
		JourneyActionMapTypesString, _ := json.Marshal(JourneyActionMapTypes)
		json.Unmarshal(JourneyActionMapTypesString, &o.JourneyActionMapTypes)
	}
	
	if DevelopmentRoleList, ok := ViewfilterMap["developmentRoleList"].([]interface{}); ok {
		DevelopmentRoleListString, _ := json.Marshal(DevelopmentRoleList)
		json.Unmarshal(DevelopmentRoleListString, &o.DevelopmentRoleList)
	}
	
	if DevelopmentTypeList, ok := ViewfilterMap["developmentTypeList"].([]interface{}); ok {
		DevelopmentTypeListString, _ := json.Marshal(DevelopmentTypeList)
		json.Unmarshal(DevelopmentTypeListString, &o.DevelopmentTypeList)
	}
	
	if DevelopmentStatusList, ok := ViewfilterMap["developmentStatusList"].([]interface{}); ok {
		DevelopmentStatusListString, _ := json.Marshal(DevelopmentStatusList)
		json.Unmarshal(DevelopmentStatusListString, &o.DevelopmentStatusList)
	}
	
	if DevelopmentModuleIds, ok := ViewfilterMap["developmentModuleIds"].([]interface{}); ok {
		DevelopmentModuleIdsString, _ := json.Marshal(DevelopmentModuleIds)
		json.Unmarshal(DevelopmentModuleIdsString, &o.DevelopmentModuleIds)
	}
	
	if DevelopmentActivityOverdue, ok := ViewfilterMap["developmentActivityOverdue"].(bool); ok {
		o.DevelopmentActivityOverdue = &DevelopmentActivityOverdue
	}
    
	if CustomerSentimentScore, ok := ViewfilterMap["customerSentimentScore"].(map[string]interface{}); ok {
		CustomerSentimentScoreString, _ := json.Marshal(CustomerSentimentScore)
		json.Unmarshal(CustomerSentimentScoreString, &o.CustomerSentimentScore)
	}
	
	if CustomerSentimentTrend, ok := ViewfilterMap["customerSentimentTrend"].(map[string]interface{}); ok {
		CustomerSentimentTrendString, _ := json.Marshal(CustomerSentimentTrend)
		json.Unmarshal(CustomerSentimentTrendString, &o.CustomerSentimentTrend)
	}
	
	if FlowTransferTargets, ok := ViewfilterMap["flowTransferTargets"].([]interface{}); ok {
		FlowTransferTargetsString, _ := json.Marshal(FlowTransferTargets)
		json.Unmarshal(FlowTransferTargetsString, &o.FlowTransferTargets)
	}
	
	if DevelopmentName, ok := ViewfilterMap["developmentName"].(string); ok {
		o.DevelopmentName = &DevelopmentName
	}
    
	if TopicIds, ok := ViewfilterMap["topicIds"].([]interface{}); ok {
		TopicIdsString, _ := json.Marshal(TopicIds)
		json.Unmarshal(TopicIdsString, &o.TopicIds)
	}
	
	if ExternalTags, ok := ViewfilterMap["externalTags"].([]interface{}); ok {
		ExternalTagsString, _ := json.Marshal(ExternalTags)
		json.Unmarshal(ExternalTagsString, &o.ExternalTags)
	}
	
	if IsNotResponding, ok := ViewfilterMap["isNotResponding"].(bool); ok {
		o.IsNotResponding = &IsNotResponding
	}
    
	if IsAuthenticated, ok := ViewfilterMap["isAuthenticated"].(bool); ok {
		o.IsAuthenticated = &IsAuthenticated
	}
    
	if BotIds, ok := ViewfilterMap["botIds"].([]interface{}); ok {
		BotIdsString, _ := json.Marshal(BotIds)
		json.Unmarshal(BotIdsString, &o.BotIds)
	}
	
	if BotVersions, ok := ViewfilterMap["botVersions"].([]interface{}); ok {
		BotVersionsString, _ := json.Marshal(BotVersions)
		json.Unmarshal(BotVersionsString, &o.BotVersions)
	}
	
	if BotMessageTypes, ok := ViewfilterMap["botMessageTypes"].([]interface{}); ok {
		BotMessageTypesString, _ := json.Marshal(BotMessageTypes)
		json.Unmarshal(BotMessageTypesString, &o.BotMessageTypes)
	}
	
	if BotProviderList, ok := ViewfilterMap["botProviderList"].([]interface{}); ok {
		BotProviderListString, _ := json.Marshal(BotProviderList)
		json.Unmarshal(BotProviderListString, &o.BotProviderList)
	}
	
	if BotProductList, ok := ViewfilterMap["botProductList"].([]interface{}); ok {
		BotProductListString, _ := json.Marshal(BotProductList)
		json.Unmarshal(BotProductListString, &o.BotProductList)
	}
	
	if BotRecognitionFailureReasonList, ok := ViewfilterMap["botRecognitionFailureReasonList"].([]interface{}); ok {
		BotRecognitionFailureReasonListString, _ := json.Marshal(BotRecognitionFailureReasonList)
		json.Unmarshal(BotRecognitionFailureReasonListString, &o.BotRecognitionFailureReasonList)
	}
	
	if BotIntentList, ok := ViewfilterMap["botIntentList"].([]interface{}); ok {
		BotIntentListString, _ := json.Marshal(BotIntentList)
		json.Unmarshal(BotIntentListString, &o.BotIntentList)
	}
	
	if BotFinalIntentList, ok := ViewfilterMap["botFinalIntentList"].([]interface{}); ok {
		BotFinalIntentListString, _ := json.Marshal(BotFinalIntentList)
		json.Unmarshal(BotFinalIntentListString, &o.BotFinalIntentList)
	}
	
	if BotSlotList, ok := ViewfilterMap["botSlotList"].([]interface{}); ok {
		BotSlotListString, _ := json.Marshal(BotSlotList)
		json.Unmarshal(BotSlotListString, &o.BotSlotList)
	}
	
	if BotResultList, ok := ViewfilterMap["botResultList"].([]interface{}); ok {
		BotResultListString, _ := json.Marshal(BotResultList)
		json.Unmarshal(BotResultListString, &o.BotResultList)
	}
	
	if BlockedReasons, ok := ViewfilterMap["blockedReasons"].([]interface{}); ok {
		BlockedReasonsString, _ := json.Marshal(BlockedReasons)
		json.Unmarshal(BlockedReasonsString, &o.BlockedReasons)
	}
	
	if IsRecorded, ok := ViewfilterMap["isRecorded"].(bool); ok {
		o.IsRecorded = &IsRecorded
	}
    
	if HasEvaluation, ok := ViewfilterMap["hasEvaluation"].(bool); ok {
		o.HasEvaluation = &HasEvaluation
	}
    
	if HasScoredEvaluation, ok := ViewfilterMap["hasScoredEvaluation"].(bool); ok {
		o.HasScoredEvaluation = &HasScoredEvaluation
	}
    
	if EmailDeliveryStatusList, ok := ViewfilterMap["emailDeliveryStatusList"].([]interface{}); ok {
		EmailDeliveryStatusListString, _ := json.Marshal(EmailDeliveryStatusList)
		json.Unmarshal(EmailDeliveryStatusListString, &o.EmailDeliveryStatusList)
	}
	
	if IsAgentOwnedCallback, ok := ViewfilterMap["isAgentOwnedCallback"].(bool); ok {
		o.IsAgentOwnedCallback = &IsAgentOwnedCallback
	}
    
	if AgentCallbackOwnerIds, ok := ViewfilterMap["agentCallbackOwnerIds"].([]interface{}); ok {
		AgentCallbackOwnerIdsString, _ := json.Marshal(AgentCallbackOwnerIds)
		json.Unmarshal(AgentCallbackOwnerIdsString, &o.AgentCallbackOwnerIds)
	}
	
	if TranscriptTopics, ok := ViewfilterMap["transcriptTopics"].([]interface{}); ok {
		TranscriptTopicsString, _ := json.Marshal(TranscriptTopics)
		json.Unmarshal(TranscriptTopicsString, &o.TranscriptTopics)
	}
	
	if JourneyFrequencyCapReasons, ok := ViewfilterMap["journeyFrequencyCapReasons"].([]interface{}); ok {
		JourneyFrequencyCapReasonsString, _ := json.Marshal(JourneyFrequencyCapReasons)
		json.Unmarshal(JourneyFrequencyCapReasonsString, &o.JourneyFrequencyCapReasons)
	}
	
	if JourneyBlockingActionMapIds, ok := ViewfilterMap["journeyBlockingActionMapIds"].([]interface{}); ok {
		JourneyBlockingActionMapIdsString, _ := json.Marshal(JourneyBlockingActionMapIds)
		json.Unmarshal(JourneyBlockingActionMapIdsString, &o.JourneyBlockingActionMapIds)
	}
	
	if JourneyActionTargetIds, ok := ViewfilterMap["journeyActionTargetIds"].([]interface{}); ok {
		JourneyActionTargetIdsString, _ := json.Marshal(JourneyActionTargetIds)
		json.Unmarshal(JourneyActionTargetIdsString, &o.JourneyActionTargetIds)
	}
	
	if JourneyBlockingScheduleGroupIds, ok := ViewfilterMap["journeyBlockingScheduleGroupIds"].([]interface{}); ok {
		JourneyBlockingScheduleGroupIdsString, _ := json.Marshal(JourneyBlockingScheduleGroupIds)
		json.Unmarshal(JourneyBlockingScheduleGroupIdsString, &o.JourneyBlockingScheduleGroupIds)
	}
	
	if JourneyBlockingEmergencyScheduleGroupIds, ok := ViewfilterMap["journeyBlockingEmergencyScheduleGroupIds"].([]interface{}); ok {
		JourneyBlockingEmergencyScheduleGroupIdsString, _ := json.Marshal(JourneyBlockingEmergencyScheduleGroupIds)
		json.Unmarshal(JourneyBlockingEmergencyScheduleGroupIdsString, &o.JourneyBlockingEmergencyScheduleGroupIds)
	}
	
	if JourneyUrlEqualConditions, ok := ViewfilterMap["journeyUrlEqualConditions"].([]interface{}); ok {
		JourneyUrlEqualConditionsString, _ := json.Marshal(JourneyUrlEqualConditions)
		json.Unmarshal(JourneyUrlEqualConditionsString, &o.JourneyUrlEqualConditions)
	}
	
	if JourneyUrlNotEqualConditions, ok := ViewfilterMap["journeyUrlNotEqualConditions"].([]interface{}); ok {
		JourneyUrlNotEqualConditionsString, _ := json.Marshal(JourneyUrlNotEqualConditions)
		json.Unmarshal(JourneyUrlNotEqualConditionsString, &o.JourneyUrlNotEqualConditions)
	}
	
	if JourneyUrlStartsWithConditions, ok := ViewfilterMap["journeyUrlStartsWithConditions"].([]interface{}); ok {
		JourneyUrlStartsWithConditionsString, _ := json.Marshal(JourneyUrlStartsWithConditions)
		json.Unmarshal(JourneyUrlStartsWithConditionsString, &o.JourneyUrlStartsWithConditions)
	}
	
	if JourneyUrlEndsWithConditions, ok := ViewfilterMap["journeyUrlEndsWithConditions"].([]interface{}); ok {
		JourneyUrlEndsWithConditionsString, _ := json.Marshal(JourneyUrlEndsWithConditions)
		json.Unmarshal(JourneyUrlEndsWithConditionsString, &o.JourneyUrlEndsWithConditions)
	}
	
	if JourneyUrlContainsAnyConditions, ok := ViewfilterMap["journeyUrlContainsAnyConditions"].([]interface{}); ok {
		JourneyUrlContainsAnyConditionsString, _ := json.Marshal(JourneyUrlContainsAnyConditions)
		json.Unmarshal(JourneyUrlContainsAnyConditionsString, &o.JourneyUrlContainsAnyConditions)
	}
	
	if JourneyUrlNotContainsAnyConditions, ok := ViewfilterMap["journeyUrlNotContainsAnyConditions"].([]interface{}); ok {
		JourneyUrlNotContainsAnyConditionsString, _ := json.Marshal(JourneyUrlNotContainsAnyConditions)
		json.Unmarshal(JourneyUrlNotContainsAnyConditionsString, &o.JourneyUrlNotContainsAnyConditions)
	}
	
	if JourneyUrlContainsAllConditions, ok := ViewfilterMap["journeyUrlContainsAllConditions"].([]interface{}); ok {
		JourneyUrlContainsAllConditionsString, _ := json.Marshal(JourneyUrlContainsAllConditions)
		json.Unmarshal(JourneyUrlContainsAllConditionsString, &o.JourneyUrlContainsAllConditions)
	}
	
	if JourneyUrlNotContainsAllConditions, ok := ViewfilterMap["journeyUrlNotContainsAllConditions"].([]interface{}); ok {
		JourneyUrlNotContainsAllConditionsString, _ := json.Marshal(JourneyUrlNotContainsAllConditions)
		json.Unmarshal(JourneyUrlNotContainsAllConditionsString, &o.JourneyUrlNotContainsAllConditions)
	}
	
	if FlowMilestoneIds, ok := ViewfilterMap["flowMilestoneIds"].([]interface{}); ok {
		FlowMilestoneIdsString, _ := json.Marshal(FlowMilestoneIds)
		json.Unmarshal(FlowMilestoneIdsString, &o.FlowMilestoneIds)
	}
	
	if IsAssessmentPassed, ok := ViewfilterMap["isAssessmentPassed"].(bool); ok {
		o.IsAssessmentPassed = &IsAssessmentPassed
	}
    
	if ConversationInitiators, ok := ViewfilterMap["conversationInitiators"].([]interface{}); ok {
		ConversationInitiatorsString, _ := json.Marshal(ConversationInitiators)
		json.Unmarshal(ConversationInitiatorsString, &o.ConversationInitiators)
	}
	
	if HasCustomerParticipated, ok := ViewfilterMap["hasCustomerParticipated"].(bool); ok {
		o.HasCustomerParticipated = &HasCustomerParticipated
	}
    
	if IsAcdInteraction, ok := ViewfilterMap["isAcdInteraction"].(bool); ok {
		o.IsAcdInteraction = &IsAcdInteraction
	}
    
	if HasFax, ok := ViewfilterMap["hasFax"].(bool); ok {
		o.HasFax = &HasFax
	}
    
	if DataActionIds, ok := ViewfilterMap["dataActionIds"].([]interface{}); ok {
		DataActionIdsString, _ := json.Marshal(DataActionIds)
		json.Unmarshal(DataActionIdsString, &o.DataActionIds)
	}
	
	if ActionCategoryName, ok := ViewfilterMap["actionCategoryName"].(string); ok {
		o.ActionCategoryName = &ActionCategoryName
	}
    
	if IntegrationIds, ok := ViewfilterMap["integrationIds"].([]interface{}); ok {
		IntegrationIdsString, _ := json.Marshal(IntegrationIds)
		json.Unmarshal(IntegrationIdsString, &o.IntegrationIds)
	}
	
	if ResponseStatuses, ok := ViewfilterMap["responseStatuses"].([]interface{}); ok {
		ResponseStatusesString, _ := json.Marshal(ResponseStatuses)
		json.Unmarshal(ResponseStatusesString, &o.ResponseStatuses)
	}
	
	if AvailableDashboard, ok := ViewfilterMap["availableDashboard"].(string); ok {
		o.AvailableDashboard = &AvailableDashboard
	}
    
	if FavouriteDashboard, ok := ViewfilterMap["favouriteDashboard"].(bool); ok {
		o.FavouriteDashboard = &FavouriteDashboard
	}
    
	if MyDashboard, ok := ViewfilterMap["myDashboard"].(bool); ok {
		o.MyDashboard = &MyDashboard
	}
    
	if StationErrors, ok := ViewfilterMap["stationErrors"].([]interface{}); ok {
		StationErrorsString, _ := json.Marshal(StationErrors)
		json.Unmarshal(StationErrorsString, &o.StationErrors)
	}
	
	if CanonicalContactIds, ok := ViewfilterMap["canonicalContactIds"].([]interface{}); ok {
		CanonicalContactIdsString, _ := json.Marshal(CanonicalContactIds)
		json.Unmarshal(CanonicalContactIdsString, &o.CanonicalContactIds)
	}
	
	if AlertRuleIds, ok := ViewfilterMap["alertRuleIds"].([]interface{}); ok {
		AlertRuleIdsString, _ := json.Marshal(AlertRuleIds)
		json.Unmarshal(AlertRuleIdsString, &o.AlertRuleIds)
	}
	
	if EvaluationFormContextIds, ok := ViewfilterMap["evaluationFormContextIds"].([]interface{}); ok {
		EvaluationFormContextIdsString, _ := json.Marshal(EvaluationFormContextIds)
		json.Unmarshal(EvaluationFormContextIdsString, &o.EvaluationFormContextIds)
	}
	
	if EvaluationStatuses, ok := ViewfilterMap["evaluationStatuses"].([]interface{}); ok {
		EvaluationStatusesString, _ := json.Marshal(EvaluationStatuses)
		json.Unmarshal(EvaluationStatusesString, &o.EvaluationStatuses)
	}
	
	if WorkbinIds, ok := ViewfilterMap["workbinIds"].([]interface{}); ok {
		WorkbinIdsString, _ := json.Marshal(WorkbinIds)
		json.Unmarshal(WorkbinIdsString, &o.WorkbinIds)
	}
	
	if WorktypeIds, ok := ViewfilterMap["worktypeIds"].([]interface{}); ok {
		WorktypeIdsString, _ := json.Marshal(WorktypeIds)
		json.Unmarshal(WorktypeIdsString, &o.WorktypeIds)
	}
	
	if WorkitemIds, ok := ViewfilterMap["workitemIds"].([]interface{}); ok {
		WorkitemIdsString, _ := json.Marshal(WorkitemIds)
		json.Unmarshal(WorkitemIdsString, &o.WorkitemIds)
	}
	
	if WorkitemAssigneeIds, ok := ViewfilterMap["workitemAssigneeIds"].([]interface{}); ok {
		WorkitemAssigneeIdsString, _ := json.Marshal(WorkitemAssigneeIds)
		json.Unmarshal(WorkitemAssigneeIdsString, &o.WorkitemAssigneeIds)
	}
	
	if WorkitemStatuses, ok := ViewfilterMap["workitemStatuses"].([]interface{}); ok {
		WorkitemStatusesString, _ := json.Marshal(WorkitemStatuses)
		json.Unmarshal(WorkitemStatusesString, &o.WorkitemStatuses)
	}
	
	if IsAnalyzedForSensitiveData, ok := ViewfilterMap["isAnalyzedForSensitiveData"].(bool); ok {
		o.IsAnalyzedForSensitiveData = &IsAnalyzedForSensitiveData
	}
    
	if HasSensitiveData, ok := ViewfilterMap["hasSensitiveData"].(bool); ok {
		o.HasSensitiveData = &HasSensitiveData
	}
    
	if HasPciData, ok := ViewfilterMap["hasPciData"].(bool); ok {
		o.HasPciData = &HasPciData
	}
    
	if HasPiiData, ok := ViewfilterMap["hasPiiData"].(bool); ok {
		o.HasPiiData = &HasPiiData
	}
    
	if SubPath, ok := ViewfilterMap["subPath"].(string); ok {
		o.SubPath = &SubPath
	}
    
	if UserState, ok := ViewfilterMap["userState"].(string); ok {
		o.UserState = &UserState
	}
    
	if IsClearedByCustomer, ok := ViewfilterMap["isClearedByCustomer"].(bool); ok {
		o.IsClearedByCustomer = &IsClearedByCustomer
	}
    
	if EvaluationAssigneeIds, ok := ViewfilterMap["evaluationAssigneeIds"].([]interface{}); ok {
		EvaluationAssigneeIdsString, _ := json.Marshal(EvaluationAssigneeIds)
		json.Unmarshal(EvaluationAssigneeIdsString, &o.EvaluationAssigneeIds)
	}
	
	if EvaluationAssigned, ok := ViewfilterMap["evaluationAssigned"].(bool); ok {
		o.EvaluationAssigned = &EvaluationAssigned
	}
    
	if AssistantIds, ok := ViewfilterMap["assistantIds"].([]interface{}); ok {
		AssistantIdsString, _ := json.Marshal(AssistantIds)
		json.Unmarshal(AssistantIdsString, &o.AssistantIds)
	}
	
	if KnowledgeBaseIds, ok := ViewfilterMap["knowledgeBaseIds"].([]interface{}); ok {
		KnowledgeBaseIdsString, _ := json.Marshal(KnowledgeBaseIds)
		json.Unmarshal(KnowledgeBaseIdsString, &o.KnowledgeBaseIds)
	}
	
	if IsParked, ok := ViewfilterMap["isParked"].(bool); ok {
		o.IsParked = &IsParked
	}
    
	if AgentEmpathyScore, ok := ViewfilterMap["agentEmpathyScore"].(map[string]interface{}); ok {
		AgentEmpathyScoreString, _ := json.Marshal(AgentEmpathyScore)
		json.Unmarshal(AgentEmpathyScoreString, &o.AgentEmpathyScore)
	}
	
	if SurveyTypes, ok := ViewfilterMap["surveyTypes"].([]interface{}); ok {
		SurveyTypesString, _ := json.Marshal(SurveyTypes)
		json.Unmarshal(SurveyTypesString, &o.SurveyTypes)
	}
	
	if SurveyResponseStatuses, ok := ViewfilterMap["surveyResponseStatuses"].([]interface{}); ok {
		SurveyResponseStatusesString, _ := json.Marshal(SurveyResponseStatuses)
		json.Unmarshal(SurveyResponseStatusesString, &o.SurveyResponseStatuses)
	}
	
	if BotFlowTypes, ok := ViewfilterMap["botFlowTypes"].([]interface{}); ok {
		BotFlowTypesString, _ := json.Marshal(BotFlowTypes)
		json.Unmarshal(BotFlowTypesString, &o.BotFlowTypes)
	}
	
	if AgentTalkDurationMilliseconds, ok := ViewfilterMap["agentTalkDurationMilliseconds"].([]interface{}); ok {
		AgentTalkDurationMillisecondsString, _ := json.Marshal(AgentTalkDurationMilliseconds)
		json.Unmarshal(AgentTalkDurationMillisecondsString, &o.AgentTalkDurationMilliseconds)
	}
	
	if CustomerTalkDurationMilliseconds, ok := ViewfilterMap["customerTalkDurationMilliseconds"].([]interface{}); ok {
		CustomerTalkDurationMillisecondsString, _ := json.Marshal(CustomerTalkDurationMilliseconds)
		json.Unmarshal(CustomerTalkDurationMillisecondsString, &o.CustomerTalkDurationMilliseconds)
	}
	
	if OvertalkDurationMilliseconds, ok := ViewfilterMap["overtalkDurationMilliseconds"].([]interface{}); ok {
		OvertalkDurationMillisecondsString, _ := json.Marshal(OvertalkDurationMilliseconds)
		json.Unmarshal(OvertalkDurationMillisecondsString, &o.OvertalkDurationMilliseconds)
	}
	
	if SilenceDurationMilliseconds, ok := ViewfilterMap["silenceDurationMilliseconds"].([]interface{}); ok {
		SilenceDurationMillisecondsString, _ := json.Marshal(SilenceDurationMilliseconds)
		json.Unmarshal(SilenceDurationMillisecondsString, &o.SilenceDurationMilliseconds)
	}
	
	if AcdDurationMilliseconds, ok := ViewfilterMap["acdDurationMilliseconds"].([]interface{}); ok {
		AcdDurationMillisecondsString, _ := json.Marshal(AcdDurationMilliseconds)
		json.Unmarshal(AcdDurationMillisecondsString, &o.AcdDurationMilliseconds)
	}
	
	if IvrDurationMilliseconds, ok := ViewfilterMap["ivrDurationMilliseconds"].([]interface{}); ok {
		IvrDurationMillisecondsString, _ := json.Marshal(IvrDurationMilliseconds)
		json.Unmarshal(IvrDurationMillisecondsString, &o.IvrDurationMilliseconds)
	}
	
	if OtherDurationMilliseconds, ok := ViewfilterMap["otherDurationMilliseconds"].([]interface{}); ok {
		OtherDurationMillisecondsString, _ := json.Marshal(OtherDurationMilliseconds)
		json.Unmarshal(OtherDurationMillisecondsString, &o.OtherDurationMilliseconds)
	}
	
	if AgentTalkPercentage, ok := ViewfilterMap["agentTalkPercentage"].(map[string]interface{}); ok {
		AgentTalkPercentageString, _ := json.Marshal(AgentTalkPercentage)
		json.Unmarshal(AgentTalkPercentageString, &o.AgentTalkPercentage)
	}
	
	if CustomerTalkPercentage, ok := ViewfilterMap["customerTalkPercentage"].(map[string]interface{}); ok {
		CustomerTalkPercentageString, _ := json.Marshal(CustomerTalkPercentage)
		json.Unmarshal(CustomerTalkPercentageString, &o.CustomerTalkPercentage)
	}
	
	if OvertalkPercentage, ok := ViewfilterMap["overtalkPercentage"].(map[string]interface{}); ok {
		OvertalkPercentageString, _ := json.Marshal(OvertalkPercentage)
		json.Unmarshal(OvertalkPercentageString, &o.OvertalkPercentage)
	}
	
	if SilencePercentage, ok := ViewfilterMap["silencePercentage"].(map[string]interface{}); ok {
		SilencePercentageString, _ := json.Marshal(SilencePercentage)
		json.Unmarshal(SilencePercentageString, &o.SilencePercentage)
	}
	
	if AcdPercentage, ok := ViewfilterMap["acdPercentage"].(map[string]interface{}); ok {
		AcdPercentageString, _ := json.Marshal(AcdPercentage)
		json.Unmarshal(AcdPercentageString, &o.AcdPercentage)
	}
	
	if IvrPercentage, ok := ViewfilterMap["ivrPercentage"].(map[string]interface{}); ok {
		IvrPercentageString, _ := json.Marshal(IvrPercentage)
		json.Unmarshal(IvrPercentageString, &o.IvrPercentage)
	}
	
	if OtherPercentage, ok := ViewfilterMap["otherPercentage"].(map[string]interface{}); ok {
		OtherPercentageString, _ := json.Marshal(OtherPercentage)
		json.Unmarshal(OtherPercentageString, &o.OtherPercentage)
	}
	
	if OvertalkInstances, ok := ViewfilterMap["overtalkInstances"].(map[string]interface{}); ok {
		OvertalkInstancesString, _ := json.Marshal(OvertalkInstances)
		json.Unmarshal(OvertalkInstancesString, &o.OvertalkInstances)
	}
	
	if IsScreenRecorded, ok := ViewfilterMap["isScreenRecorded"].(bool); ok {
		o.IsScreenRecorded = &IsScreenRecorded
	}
    
	if ScreenMonitorUserIds, ok := ViewfilterMap["screenMonitorUserIds"].([]interface{}); ok {
		ScreenMonitorUserIdsString, _ := json.Marshal(ScreenMonitorUserIds)
		json.Unmarshal(ScreenMonitorUserIdsString, &o.ScreenMonitorUserIds)
	}
	
	if DashboardState, ok := ViewfilterMap["dashboardState"].(string); ok {
		o.DashboardState = &DashboardState
	}
    
	if DashboardType, ok := ViewfilterMap["dashboardType"].(string); ok {
		o.DashboardType = &DashboardType
	}
    
	if DashboardAccessFilter, ok := ViewfilterMap["dashboardAccessFilter"].(string); ok {
		o.DashboardAccessFilter = &DashboardAccessFilter
	}
    
	if TranscriptDurationMilliseconds, ok := ViewfilterMap["transcriptDurationMilliseconds"].([]interface{}); ok {
		TranscriptDurationMillisecondsString, _ := json.Marshal(TranscriptDurationMilliseconds)
		json.Unmarshal(TranscriptDurationMillisecondsString, &o.TranscriptDurationMilliseconds)
	}
	
	if WorkitemsStatuses, ok := ViewfilterMap["workitemsStatuses"].([]interface{}); ok {
		WorkitemsStatusesString, _ := json.Marshal(WorkitemsStatuses)
		json.Unmarshal(WorkitemsStatusesString, &o.WorkitemsStatuses)
	}
	
	if SocialCountries, ok := ViewfilterMap["socialCountries"].([]interface{}); ok {
		SocialCountriesString, _ := json.Marshal(SocialCountries)
		json.Unmarshal(SocialCountriesString, &o.SocialCountries)
	}
	
	if SocialLanguages, ok := ViewfilterMap["socialLanguages"].([]interface{}); ok {
		SocialLanguagesString, _ := json.Marshal(SocialLanguages)
		json.Unmarshal(SocialLanguagesString, &o.SocialLanguages)
	}
	
	if SocialChannels, ok := ViewfilterMap["socialChannels"].([]interface{}); ok {
		SocialChannelsString, _ := json.Marshal(SocialChannels)
		json.Unmarshal(SocialChannelsString, &o.SocialChannels)
	}
	
	if SocialSentimentCategory, ok := ViewfilterMap["socialSentimentCategory"].([]interface{}); ok {
		SocialSentimentCategoryString, _ := json.Marshal(SocialSentimentCategory)
		json.Unmarshal(SocialSentimentCategoryString, &o.SocialSentimentCategory)
	}
	
	if SocialTopicIds, ok := ViewfilterMap["socialTopicIds"].([]interface{}); ok {
		SocialTopicIdsString, _ := json.Marshal(SocialTopicIds)
		json.Unmarshal(SocialTopicIdsString, &o.SocialTopicIds)
	}
	
	if SocialIngestionRuleIds, ok := ViewfilterMap["socialIngestionRuleIds"].([]interface{}); ok {
		SocialIngestionRuleIdsString, _ := json.Marshal(SocialIngestionRuleIds)
		json.Unmarshal(SocialIngestionRuleIdsString, &o.SocialIngestionRuleIds)
	}
	
	if SocialConversationCreated, ok := ViewfilterMap["socialConversationCreated"].(bool); ok {
		o.SocialConversationCreated = &SocialConversationCreated
	}
    
	if SocialContentType, ok := ViewfilterMap["socialContentType"].([]interface{}); ok {
		SocialContentTypeString, _ := json.Marshal(SocialContentType)
		json.Unmarshal(SocialContentTypeString, &o.SocialContentType)
	}
	
	if SocialKeywords, ok := ViewfilterMap["socialKeywords"].([]interface{}); ok {
		SocialKeywordsString, _ := json.Marshal(SocialKeywords)
		json.Unmarshal(SocialKeywordsString, &o.SocialKeywords)
	}
	
	if SocialPostEscalated, ok := ViewfilterMap["socialPostEscalated"].(bool); ok {
		o.SocialPostEscalated = &SocialPostEscalated
	}
    
	if SocialClassifications, ok := ViewfilterMap["socialClassifications"].([]interface{}); ok {
		SocialClassificationsString, _ := json.Marshal(SocialClassifications)
		json.Unmarshal(SocialClassificationsString, &o.SocialClassifications)
	}
	
	if FilterUsersByManagerIds, ok := ViewfilterMap["filterUsersByManagerIds"].([]interface{}); ok {
		FilterUsersByManagerIdsString, _ := json.Marshal(FilterUsersByManagerIds)
		json.Unmarshal(FilterUsersByManagerIdsString, &o.FilterUsersByManagerIds)
	}
	
	if SlideshowIds, ok := ViewfilterMap["slideshowIds"].([]interface{}); ok {
		SlideshowIdsString, _ := json.Marshal(SlideshowIds)
		json.Unmarshal(SlideshowIdsString, &o.SlideshowIds)
	}
	
	if Conferenced, ok := ViewfilterMap["conferenced"].(bool); ok {
		o.Conferenced = &Conferenced
	}
    
	if Video, ok := ViewfilterMap["video"].(bool); ok {
		o.Video = &Video
	}
    
	if LinkedInteraction, ok := ViewfilterMap["linkedInteraction"].(bool); ok {
		o.LinkedInteraction = &LinkedInteraction
	}
    
	if RecommendationSources, ok := ViewfilterMap["recommendationSources"].([]interface{}); ok {
		RecommendationSourcesString, _ := json.Marshal(RecommendationSources)
		json.Unmarshal(RecommendationSourcesString, &o.RecommendationSources)
	}
	
	if EvaluationRole, ok := ViewfilterMap["evaluationRole"].(string); ok {
		o.EvaluationRole = &EvaluationRole
	}
    
	if ComparisonQueueIds, ok := ViewfilterMap["comparisonQueueIds"].([]interface{}); ok {
		ComparisonQueueIdsString, _ := json.Marshal(ComparisonQueueIds)
		json.Unmarshal(ComparisonQueueIdsString, &o.ComparisonQueueIds)
	}
	
	if ViewMetrics, ok := ViewfilterMap["viewMetrics"].([]interface{}); ok {
		ViewMetricsString, _ := json.Marshal(ViewMetrics)
		json.Unmarshal(ViewMetricsString, &o.ViewMetrics)
	}
	
	if TimelineCategories, ok := ViewfilterMap["timelineCategories"].([]interface{}); ok {
		TimelineCategoriesString, _ := json.Marshal(TimelineCategories)
		json.Unmarshal(TimelineCategoriesString, &o.TimelineCategories)
	}
	
	if Acw, ok := ViewfilterMap["acw"].(bool); ok {
		o.Acw = &Acw
	}
    
	if SegmentTypes, ok := ViewfilterMap["segmentTypes"].([]interface{}); ok {
		SegmentTypesString, _ := json.Marshal(SegmentTypes)
		json.Unmarshal(SegmentTypesString, &o.SegmentTypes)
	}
	
	if ProgramIds, ok := ViewfilterMap["programIds"].([]interface{}); ok {
		ProgramIdsString, _ := json.Marshal(ProgramIds)
		json.Unmarshal(ProgramIdsString, &o.ProgramIds)
	}
	
	if CategoryIds, ok := ViewfilterMap["categoryIds"].([]interface{}); ok {
		CategoryIdsString, _ := json.Marshal(CategoryIds)
		json.Unmarshal(CategoryIdsString, &o.CategoryIds)
	}
	
	if DeliveryPushed, ok := ViewfilterMap["deliveryPushed"].(bool); ok {
		o.DeliveryPushed = &DeliveryPushed
	}
    
	if SocialRatings, ok := ViewfilterMap["socialRatings"].([]interface{}); ok {
		SocialRatingsString, _ := json.Marshal(SocialRatings)
		json.Unmarshal(SocialRatingsString, &o.SocialRatings)
	}
	
	if VirtualAgentIds, ok := ViewfilterMap["virtualAgentIds"].([]interface{}); ok {
		VirtualAgentIdsString, _ := json.Marshal(VirtualAgentIds)
		json.Unmarshal(VirtualAgentIdsString, &o.VirtualAgentIds)
	}
	
	if EmpathyScoreCategories, ok := ViewfilterMap["empathyScoreCategories"].([]interface{}); ok {
		EmpathyScoreCategoriesString, _ := json.Marshal(EmpathyScoreCategories)
		json.Unmarshal(EmpathyScoreCategoriesString, &o.EmpathyScoreCategories)
	}
	
	if SentimentScoreCategories, ok := ViewfilterMap["sentimentScoreCategories"].([]interface{}); ok {
		SentimentScoreCategoriesString, _ := json.Marshal(SentimentScoreCategories)
		json.Unmarshal(SentimentScoreCategoriesString, &o.SentimentScoreCategories)
	}
	
	if SentimentTrendCategories, ok := ViewfilterMap["sentimentTrendCategories"].([]interface{}); ok {
		SentimentTrendCategoriesString, _ := json.Marshal(SentimentTrendCategories)
		json.Unmarshal(SentimentTrendCategoriesString, &o.SentimentTrendCategories)
	}
	
	if ContentModerationFlags, ok := ViewfilterMap["contentModerationFlags"].([]interface{}); ok {
		ContentModerationFlagsString, _ := json.Marshal(ContentModerationFlags)
		json.Unmarshal(ContentModerationFlagsString, &o.ContentModerationFlags)
	}
	
	if SessionExpired, ok := ViewfilterMap["sessionExpired"].(bool); ok {
		o.SessionExpired = &SessionExpired
	}
    
	if EngagementSource, ok := ViewfilterMap["engagementSource"].([]interface{}); ok {
		EngagementSourceString, _ := json.Marshal(EngagementSource)
		json.Unmarshal(EngagementSourceString, &o.EngagementSource)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Viewfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
