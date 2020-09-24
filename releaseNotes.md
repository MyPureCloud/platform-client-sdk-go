Platform API version: 4219


# Major Changes (5 changes)

**GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}** (1 change)

* Parameter excludeCapabilities was added

**WfmTimeZone** (1 change)

* Model WfmTimeZone was removed

**WfmAgent** (1 change)

* Property timeZone was removed

**EntityListing** (1 change)

* Property entities was changed from object[] to DataTableImportJob[]

**WfmHistoricalAdherenceQueryForUsers** (1 change)

* Required property teamIds was added


# Minor Changes (93 changes)

**/api/v2/users/development/activities/aggregates/query** (2 changes)

* Path was added
* Operation POST was added

**/api/v2/users/development/activities/me** (2 changes)

* Path was added
* Operation GET was added

**/api/v2/users/development/activities** (2 changes)

* Path was added
* Operation GET was added

**/api/v2/users/development/activities/{activityId}** (2 changes)

* Path was added
* Operation GET was added

**/api/v2/conversations/{conversationId}/assign** (2 changes)

* Path was added
* Operation POST was added

**/api/v2/flows/milestones** (3 changes)

* Path was added
* Operation GET was added
* Operation POST was added

**/api/v2/flows/milestones/{milestoneId}** (4 changes)

* Path was added
* Operation GET was added
* Operation PUT was added
* Operation DELETE was added

**Call** (2 changes)

* Optional property afterCallWorkRequired was added
* Optional property agentAssistantId was added

**Callback** (1 change)

* Optional property afterCallWorkRequired was added

**Cobrowsesession** (1 change)

* Optional property afterCallWorkRequired was added

**ConversationChat** (1 change)

* Optional property afterCallWorkRequired was added

**Email** (1 change)

* Optional property afterCallWorkRequired was added

**Message** (1 change)

* Optional property afterCallWorkRequired was added

**Queue** (2 changes)

* Optional property enableTranscription was added
* Optional property enableManualAssignment was added

**Screenshare** (1 change)

* Optional property afterCallWorkRequired was added

**SocialExpression** (1 change)

* Optional property afterCallWorkRequired was added

**Video** (1 change)

* Optional property afterCallWorkRequired was added

**QueueRequest** (2 changes)

* Optional property enableTranscription was added
* Optional property enableManualAssignment was added

**AuditQueryExecutionStatusResponse** (1 change)

* Enum value ResponseManagement was added to property serviceName

**AuditQueryRequest** (1 change)

* Enum value ResponseManagement was added to property serviceName

**AuditLogMessage** (4 changes)

* Enum value ResponseManagement was added to property serviceName
* Enum value SessionType was added to property entityType
* Enum value EventType was added to property entityType
* Enum value Response was added to property entityType

**AuditRealtimeQueryRequest** (1 change)

* Enum value ResponseManagement was added to property serviceName

**AuditQueryEntity** (3 changes)

* Enum value SessionType was added to property name
* Enum value EventType was added to property name
* Enum value Response was added to property name

**AuditQueryService** (1 change)

* Enum value ResponseManagement was added to property name

**KnowledgeBase** (1 change)

* Enum value de-DE was added to property coreLanguage

**AvailableTopic** (1 change)

* Optional property publicApiTemplateUriPaths was added

**WfmAgent** (1 change)

* Optional property workPlanRotation was added

**WorkPlanRotationReference** (1 change)

* Model was added

**ViewFilter** (1 change)

* Optional property hasAgentAssistId was added

**WfmHistoricalAdherenceQuery** (1 change)

* Optional property teamIds was added

**FlowAggregateQueryPredicate** (1 change)

* Enum value flowMilestoneId was added to property dimension

**FlowAggregationQuery** (3 changes)

* Enum value flowMilestoneId was added to property groupBy
* Enum value nFlowMilestone was added to property metrics
* Enum value oFlowMilestone was added to property metrics

**FlowAggregationView** (2 changes)

* Enum value nFlowMilestone was added to property target
* Enum value oFlowMilestone was added to property target

**BuAgentScheduleRescheduleResponse** (1 change)

* Optional property workPlansPerWeek was added

**ConversationDetailQueryPredicate** (2 changes)

* Enum value nFlowMilestone was added to property metric
* Enum value oFlowMilestone was added to property metric

**Dependency** (1 change)

* Enum value FLOWMILESTONE was added to property type

**DependencyObject** (1 change)

* Enum value FLOWMILESTONE was added to property type

**KnowledgeCategory** (1 change)

* Enum value de-DE was added to property languageCode

**KnowledgeDocument** (1 change)

* Enum value de-DE was added to property languageCode

**KnowledgeSearchDocument** (1 change)

* Enum value de-DE was added to property languageCode

**EntityListing** (4 changes)

* Optional property pageSize was added
* Optional property pageNumber was added
* Optional property total was added
* Optional property pageCount was added

**BuAgentScheduleQueryResponse** (1 change)

* Optional property workPlansPerWeek was added

**UserQueue** (2 changes)

* Optional property enableTranscription was added
* Optional property enableManualAssignment was added

**DevelopmentActivityAggregateQueryResponseData** (1 change)

* Model was added

**DevelopmentActivityAggregateQueryResponseGroupedData** (1 change)

* Model was added

**DevelopmentActivityAggregateQueryResponseMetric** (1 change)

* Model was added

**DevelopmentActivityAggregateQueryResponseStatistics** (1 change)

* Model was added

**DevelopmentActivityAggregateResponse** (1 change)

* Model was added

**DevelopmentActivityAggregateParam** (1 change)

* Model was added

**DevelopmentActivityAggregateQueryRequestClause** (1 change)

* Model was added

**DevelopmentActivityAggregateQueryRequestFilter** (1 change)

* Model was added

**DevelopmentActivityAggregateQueryRequestPredicate** (1 change)

* Model was added

**ReportingExportJobResponse** (2 changes)

* Enum value FAILED_AS_EXPORT_FILE_SIZE_IS_GREATER_THAN_10MB was added to property exportErrorMessagesType
* Optional property emailErrorDescription was added

**DevelopmentActivity** (1 change)

* Model was added

**DevelopmentActivityListing** (1 change)

* Model was added

**ConversationUser** (1 change)

* Model was added

**CreateQueueRequest** (2 changes)

* Optional property enableTranscription was added
* Optional property enableManualAssignment was added

**FlowMilestone** (1 change)

* Model was added

**FlowMilestoneListing** (1 change)

* Model was added

**QueueObservationQuery** (2 changes)

* Enum value oAlerting was added to property metrics
* Enum value oAlerting was added to property detailMetrics

**CallBasic** (2 changes)

* Optional property afterCallWorkRequired was added
* Optional property agentAssistantId was added

**CallbackBasic** (1 change)

* Optional property afterCallWorkRequired was added

**KnowledgeExtendedCategory** (1 change)

* Enum value de-DE was added to property languageCode


# Point Changes (0 changes)
