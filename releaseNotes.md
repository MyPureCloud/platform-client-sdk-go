Platform API version: 3829


# Major Changes (24 changes)

**GET /api/v2/analytics/conversations/details/jobs/{jobId}/results** (1 change)

* Parameter pageSize was added

**GET /api/v2/analytics/users/details/jobs/{jobId}/results** (1 change)

* Parameter pageSize was added

**GET /api/v2/authorization/divisionspermitted/me** (1 change)

* Has been deprecated

**GET /api/v2/quality/surveys/scorable** (1 change)

* Parameter customerSurveyUrl was made required

**GET /api/v2/audits/query/{transactionId}/results** (1 change)

* Parameter expand was added

**POST /api/v2/workforcemanagement/managementunits/{muId}/agentschedules/search** (1 change)

* Response 200 type was changed from UserScheduleContainer to BuAsyncAgentSchedulesSearchResponse

**Photo** (1 change)

* Model Photo was removed

**AnalyticsConversationQueryResponse** (1 change)

* Property conversations was changed from AnalyticsConversation[] to AnalyticsConversationWithoutAttributes[]

**ScimV2User** (1 change)

* Property photos was removed

**ScimV2CreateUser** (1 change)

* Property photos was removed

**ViewFilter** (10 changes)

* Property showSecondaryStatus was removed
* Property agentDurationSortOrder was removed
* Property waitingDurationSortOrder was removed
* Property interactingDurationSortOrder was removed
* Property agentName was removed
* Property skillsList was removed
* Property languageList was removed
* Property statusList was removed
* Property oauthClientIds was removed
* Property apiOperations was removed

**AuditQueryExecutionStatusResponse** (1 change)

* Enum value TopicsDefinitionsService was removed from property serviceName

**AuditQueryRequest** (1 change)

* Enum value TopicsDefinitionsService was removed from property serviceName

**AuditLogMessage** (2 changes)

* Property user was changed from AddressableEntityRef to DomainEntityRef
* Enum value TopicsDefinitionsService was removed from property serviceName


# Minor Changes (57 changes)

**PATCH /api/v2/users/{userId}/presences/{sourceId}** (1 change)

* Response 409 was added

**/api/v2/routing/users/{userId}/utilization** (4 changes)

* Path was added
* Operation GET was added
* Operation PUT was added
* Operation DELETE was added

**/api/v2/license/infer** (2 changes)

* Path was added
* Operation POST was added

**/api/v2/authorization/divisionspermitted/paged/me** (2 changes)

* Path was added
* Operation GET was added

**/api/v2/authorization/divisionspermitted/paged/{subjectId}** (2 changes)

* Path was added
* Operation GET was added

**/api/v2/flows/datatables/{datatableId}/import/jobs/{importJobId}** (2 changes)

* Path was added
* Operation GET was added

**/api/v2/flows/datatables/{datatableId}/import/jobs** (3 changes)

* Path was added
* Operation GET was added
* Operation POST was added

**/api/v2/flows/datatables/{datatableId}/export/jobs/{exportJobId}** (2 changes)

* Path was added
* Operation GET was added

**/api/v2/flows/datatables/{datatableId}/export/jobs** (2 changes)

* Path was added
* Operation POST was added

**Organization** (1 change)

* Optional property productPlatform was added

**Evaluation** (1 change)

* Optional property conversationEndDate was added

**EvaluationQuestionGroupScore** (4 changes)

* Optional property totalNonCriticalScore was added
* Optional property maxTotalNonCriticalScore was added
* Optional property totalNonCriticalScoreUnweighted was added
* Optional property maxTotalNonCriticalScoreUnweighted was added

**EvaluationScoringSet** (1 change)

* Optional property totalNonCriticalScore was added

**ConversationAggregateQueryPredicate** (1 change)

* Enum value teamId was added to property dimension

**ConversationAggregationQuery** (1 change)

* Enum value teamId was added to property groupBy

**AnalyticsParticipantWithoutAttributes** (1 change)

* Optional property teamId was added

**SegmentDetailQueryPredicate** (1 change)

* Enum value teamId was added to property dimension

**AnalyticsParticipant** (1 change)

* Optional property teamId was added

**CampaignRuleParameters** (1 change)

* Enum value external was added to property dialingMode

**EntityListing** (1 change)

* Model was added

**FlowAggregateQueryPredicate** (1 change)

* Enum value teamId was added to property dimension

**FlowAggregationQuery** (1 change)

* Enum value teamId was added to property groupBy

**ObservationValue** (1 change)

* Optional property teamId was added

**BuScheduleReference** (1 change)

* Model was added

**BuAgentScheduleActivity** (1 change)

* Model was added

**BuAgentScheduleShift** (1 change)

* Model was added

**BuFullDayTimeOffMarker** (1 change)

* Model was added

**CampaignProgress** (1 change)

* Optional property numberOfContactsMessaged was added

**Campaign** (1 change)

* Enum value external was added to property dialingMode

**DivsPermittedEntityListing** (1 change)

* Model was added

**BuAgentSchedulePublishedScheduleReference** (1 change)

* Model was added

**BuAgentScheduleSearchResponse** (1 change)

* Model was added

**BuAgentSchedulesSearchResponse** (1 change)

* Model was added

**BuAsyncAgentSchedulesSearchResponse** (1 change)

* Model was added

**DataTableImportJob** (1 change)

* Model was added

**ReportingExportJobResponse** (1 change)

* Enum value TOO_MANY_REQUESTS_FROM_AN_ORGANIZATION was added to property exportErrorMessagesType

**AvailableTopic** (2 changes)

* Optional property requiresCurrentUser was added
* Optional property requiresCurrentUserOrPermission was added

**AuditQueryExecutionStatusResponse** (1 change)

* Enum value TopicsDefinitions was added to property serviceName

**AuditQueryRequest** (1 change)

* Enum value TopicsDefinitions was added to property serviceName

**AuditLogMessage** (2 changes)

* Enum value TopicsDefinitions was added to property serviceName
* Enum value OAuthClientAuthorization was added to property entityType

**DataTableExportJob** (1 change)

* Model was added


# Point Changes (4 changes)

**GET /api/v2/authorization/divisionspermitted/me** (2 changes)

* Description was changed
* Summary was changed

**GET /api/v2/authorization/divisionspermitted/{subjectId}** (2 changes)

* Description was changed
* Summary was changed
