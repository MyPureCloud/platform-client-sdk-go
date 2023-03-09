Platform API version: 6711


# Major Changes (19 changes)

**/api/v2/recording/localkeys/settings/{settingsId}** (1 change)

* Path /api/v2/recording/localkeys/settings/{settingsId} was removed

**/api/v2/recording/localkeys/settings** (1 change)

* Path /api/v2/recording/localkeys/settings was removed

**GET /api/v2/knowledge/guest/sessions/{sessionId}/documents** (1 change)

* Parameter includeSubcategories was removed

**POST /api/v2/notifications/channels/{channelId}/subscriptions** (1 change)

* Parameter ignoreErrors was added

**PUT /api/v2/notifications/channels/{channelId}/subscriptions** (1 change)

* Parameter ignoreErrors was added

**GET /api/v2/quality/evaluations/query** (1 change)

* Parameter assigneeUserId was added

**GET /api/v2/routing/sms/phonenumbers/{addressId}** (1 change)

* Parameter expand was added

**GET /api/v2/telephony/providers/edges** (1 change)

* Parameter showCloudMedia was added

**LocalEncryptionConfigurationListing** (1 change)

* Model LocalEncryptionConfigurationListing was removed

**ChannelMetadata** (1 change)

* Model ChannelMetadata was removed

**Reason** (1 change)

* Model Reason was removed

**PresenceSetting** (1 change)

* Model PresenceSetting was removed

**OpenMessageContent** (1 change)

* Property attachment was changed from ContentAttachment to ConversationContentAttachment

**OpenMessagingChannel** (1 change)

* Property metadata was changed from ChannelMetadata to ConversationChannelMetadata

**OpenNormalizedMessage** (1 change)

* Property reasons was changed from Reason[] to ConversationReason[]

**EventSetting** (1 change)

* Property presence was removed

**ApprovalNamespace** (1 change)

* Enum value contacts was removed from property namespace

**LimitChangeRequestDetails** (1 change)

* Enum value contacts was removed from property namespace

**StatusChange** (1 change)

* Enum value contacts was removed from property namespace


# Minor Changes (121 changes)

**/api/v2/conversations/participants/attributes/search** (2 changes)

* Path was added
* Operation POST was added

**POST /api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages** (1 change)

* Response 409 was added

**POST /api/v2/conversations/messages/agentless** (1 change)

* Response 409 was added

**/api/v2/outbound/dnclists/{dncListId}/customexclusioncolumns** (3 changes)

* Path was added
* Operation DELETE was added
* Operation PATCH was added

**/api/v2/journey/outcomes/predictors/{predictorId}** (3 changes)

* Path was added
* Operation GET was added
* Operation DELETE was added

**/api/v2/journey/outcomes/predictors** (3 changes)

* Path was added
* Operation GET was added
* Operation POST was added

**POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents** (1 change)

* Response 409 was added

**/api/v2/speechandtextanalytics/programs/{programId}/transcriptionengines** (3 changes)

* Path was added
* Operation GET was added
* Operation PUT was added

**/api/v2/speechandtextanalytics/programs/transcriptionengines/dialects** (2 changes)

* Path was added
* Operation GET was added

**/api/v2/infrastructureascode/jobs/{jobId}** (2 changes)

* Path was added
* Operation GET was added

**/api/v2/infrastructureascode/jobs** (3 changes)

* Path was added
* Operation GET was added
* Operation POST was added

**/api/v2/users/{userId}/skillgroups** (2 changes)

* Path was added
* Operation GET was added

**/api/v2/learning/assignments/{assignmentId}/reschedule** (2 changes)

* Path was added
* Operation PATCH was added

**/api/v2/learning/scheduleslots/query** (2 changes)

* Path was added
* Operation POST was added

**/api/v2/workforcemanagement/agents/me/possibleworkshifts** (2 changes)

* Path was added
* Operation POST was added

**Station** (1 change)

* Optional property webRtcRequireMediaHelper was added

**ViewFilter** (3 changes)

* Optional property alertRuleIds was added
* Optional property evaluationFormContextIds was added
* Optional property evaluationStatuses was added

**AuditQueryExecutionStatusResponse** (1 change)

* Enum value GDPR was added to property serviceName

**AuditQueryRequest** (1 change)

* Enum value GDPR was added to property serviceName

**AuditLogMessage** (3 changes)

* Enum value GDPR was added to property serviceName
* Enum value RestoreAll was added to property action
* Enum value RestoreDeleted was added to property action

**InitiatingAction** (2 changes)

* Enum value RestoreAll was added to property actionContext
* Enum value RestoreDeleted was added to property actionContext

**AuditRealtimeQueryRequest** (1 change)

* Enum value GDPR was added to property serviceName

**AuditQueryEntity** (2 changes)

* Enum value RestoreAll was added to property actions
* Enum value RestoreDeleted was added to property actions

**AuditQueryService** (1 change)

* Enum value GDPR was added to property name

**InboundRoute** (1 change)

* Optional property allowMultipleActions was added

**InstagramId** (1 change)

* Model was added

**InstagramScopedId** (1 change)

* Model was added

**Conversation** (1 change)

* Optional property recentTransfers was added

**Evaluation** (2 changes)

* Optional property assignee was added
* Optional property evaluationSource was added

**EvaluationScoringSet** (1 change)

* Optional property privateComments was added

**EvaluationSource** (1 change)

* Model was added

**TransferDestination** (1 change)

* Model was added

**TransferInitiator** (1 change)

* Model was added

**TransferResponse** (1 change)

* Model was added

**JsonCursorSearchResponse** (1 change)

* Model was added

**ConversationParticipantSearchCriteria** (1 change)

* Model was added

**ConversationParticipantSearchRequest** (1 change)

* Model was added

**CallConversation** (1 change)

* Optional property recentTransfers was added

**TransferRequest** (1 change)

* Optional property transferType was added

**CallbackConversation** (1 change)

* Optional property recentTransfers was added

**ChatConversation** (1 change)

* Optional property recentTransfers was added

**CobrowseConversation** (1 change)

* Optional property recentTransfers was added

**EmailMessage** (2 changes)

* Enum value Edited was added to property state
* Optional property draftType was added

**EmailConversation** (1 change)

* Optional property recentTransfers was added

**MessageConversation** (1 change)

* Optional property recentTransfers was added

**DncPatchCustomExclusionColumnsRequest** (1 change)

* Model was added

**DncList** (1 change)

* Optional property customExclusionColumn was added

**DncListCreate** (1 change)

* Optional property customExclusionColumn was added

**DialerAction** (1 change)

* Enum value APPEND_CUSTOM_ENTRY_TO_DNC_LIST was added to property actionTypeName

**AppendToDncActionSettings** (1 change)

* Optional property listType was added

**ActionConfig** (1 change)

* Optional property timeoutSeconds was added

**OutcomePredictor** (1 change)

* Model was added

**OutcomeRef** (1 change)

* Model was added

**OutcomePredictorRequest** (1 change)

* Model was added

**OutcomeRefRequest** (1 change)

* Model was added

**OutcomePredictorListing** (1 change)

* Model was added

**KnowledgeExportJobDocumentsFilter** (1 change)

* Optional property entities was added

**ApprovalNamespace** (2 changes)

* Enum value callback was added to property namespace
* Enum value presence was added to property namespace

**LimitChangeRequestDetails** (2 changes)

* Enum value callback was added to property namespace
* Enum value presence was added to property namespace

**StatusChange** (2 changes)

* Enum value callback was added to property namespace
* Enum value presence was added to property namespace

**ChannelTopic** (2 changes)

* Optional property state was added
* Optional property rejectionReason was added

**EvaluationResponse** (2 changes)

* Optional property assignee was added
* Optional property evaluationSource was added

**OrphanRecording** (3 changes)

* Enum value ap-east-1 was added to property region
* Enum value ap-southeast-3 was added to property region
* Enum value eu-west-3 was added to property region

**Check** (1 change)

* Enum value SalesAmountValue was added to property type

**ProgramTranscriptionEngines** (1 change)

* Model was added

**TranscriptionEngines** (1 change)

* Model was added

**TranscriptionEnginesRequest** (1 change)

* Model was added

**SupportedDialectsEntityListing** (1 change)

* Model was added

**InfrastructureascodeJob** (1 change)

* Model was added

**AcceleratorInput** (1 change)

* Model was added

**AcceleratorParameter** (1 change)

* Model was added

**UserSkillGroupEntityListing** (1 change)

* Model was added

**DevelopmentActivity** (1 change)

* Enum value External was added to property type

**LearningModule** (1 change)

* Enum value External was added to property type

**LearningAssignmentReschedule** (1 change)

* Model was added

**LearningModuleRequest** (1 change)

* Enum value External was added to property type

**AssignedLearningModule** (1 change)

* Enum value External was added to property type

**LearningScheduleSlotsQueryResponse** (1 change)

* Model was added

**LearningSlot** (1 change)

* Model was added

**LearningSlotFullDayTimeOffMarker** (1 change)

* Model was added

**LearningSlotScheduleActivity** (1 change)

* Model was added

**LearningSlotWfmScheduleActivity** (1 change)

* Model was added

**LearningScheduleSlotsQueryRequest** (1 change)

* Model was added

**WfmServiceGoalImpact** (1 change)

* Model was added

**WfmServiceGoalImpactSettings** (1 change)

* Model was added

**AgentPossibleWorkShiftsResponse** (1 change)

* Model was added

**DailyPossibleShift** (1 change)

* Model was added

**PossibleWorkShiftsForWeek** (1 change)

* Model was added

**AgentPossibleWorkShiftsRequest** (1 change)

* Model was added


# Point Changes (4 changes)

**POST /api/v2/conversations/{conversationId}/cobrowse** (1 change)

* Summary was changed

**GET /api/v2/quality/conversations/{conversationId}/evaluations/{evaluationId}** (1 change)

* Description was changed for parameter expand

**GET /api/v2/quality/evaluations/query** (1 change)

* Description was changed

**POST /api/v2/recording/jobs** (1 change)

* Description was changed
