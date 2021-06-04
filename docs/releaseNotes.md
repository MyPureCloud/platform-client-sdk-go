Platform API version: 4753


# Major Changes (10 changes)

**POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers** (1 change)

* Parameter expirationDateTime was removed

**POST /api/v2/flows** (1 change)

* Parameter language was added

**GET /api/v2/gamification/metrics/{metricId}** (1 change)

* Parameter performance profile id was added

**PUT /api/v2/gamification/metrics/{metricId}** (1 change)

* Parameter performance profile id was added

**POST /api/v2/learning/assignments/bulkremove** (1 change)

* Response 204 was removed

**PerformanceProfile** (3 changes)

* Required property reportingIntervals was added
* Required property active was added
* Required property maxLeaderboardRankSize was added

**ButtonResponse** (1 change)

* Property id was removed

**QuickReply** (1 change)

* Property id was removed


# Minor Changes (55 changes)

**/api/v2/conversations/messages/inbound/open** (2 changes)

* Path was added
* Operation POST was added

**/api/v2/workforcemanagement/managementunits/{managementUnitId}** (4 changes)

* Path was added
* Operation GET was added
* Operation DELETE was added
* Operation PATCH was added

**/api/v2/conversations/messaging/integrations/open** (3 changes)

* Path was added
* Operation GET was added
* Operation POST was added

**/api/v2/conversations/messaging/integrations/open/{integrationId}** (4 changes)

* Path was added
* Operation GET was added
* Operation DELETE was added
* Operation PATCH was added

**POST /api/v2/learning/assignments/bulkremove** (1 change)

* Response 200 was added

**ViewFilter** (2 changes)

* Optional property isNotResponding was added
* Optional property isAuthenticated was added

**PerformanceProfile** (2 changes)

* Optional property division was added
* Optional property dateCreated was added

**ReportingInterval** (1 change)

* Model was added

**WorkPlanActivity** (1 change)

* Optional property validationId was added

**WorkPlanListItemResponse** (1 change)

* Optional property valid was added

**WorkPlanShift** (1 change)

* Optional property validationId was added

**WorkPlan** (1 change)

* Optional property valid was added

**WorkPlanValidationRequest** (1 change)

* Optional property valid was added

**OpenMessageContent** (1 change)

* Model was added

**OpenMessagingChannel** (1 change)

* Model was added

**OpenMessagingFromRecipient** (1 change)

* Model was added

**OpenMessagingToRecipient** (1 change)

* Model was added

**OpenNormalizedMessage** (1 change)

* Model was added

**AuditQueryExecutionStatusResponse** (1 change)

* Enum value Messaging was added to property serviceName

**AuditQueryRequest** (1 change)

* Enum value Messaging was added to property serviceName

**AuditLogMessage** (6 changes)

* Enum value Messaging was added to property serviceName
* Enum value ConversationPhoneNumber was added to property entityType
* Enum value ConversationRecipient was added to property entityType
* Enum value ConversationAccount was added to property entityType
* Enum value ConversationDefaultSupportedContent was added to property entityType
* Enum value ConversationThreadingWindow was added to property entityType

**AuditRealtimeQueryRequest** (1 change)

* Enum value Messaging was added to property serviceName

**AuditQueryEntity** (5 changes)

* Enum value ConversationPhoneNumber was added to property name
* Enum value ConversationRecipient was added to property name
* Enum value ConversationAccount was added to property name
* Enum value ConversationDefaultSupportedContent was added to property name
* Enum value ConversationThreadingWindow was added to property name

**AuditQueryService** (1 change)

* Enum value Messaging was added to property name

**Leaderboard** (1 change)

* Optional property userRank was added

**ManagementUnitSettingsRequest** (1 change)

* Model was added

**UpdateManagementUnitRequest** (1 change)

* Model was added

**OpenIntegration** (1 change)

* Model was added

**OpenIntegrationEntityListing** (1 change)

* Model was added

**OpenIntegrationRequest** (1 change)

* Model was added

**OpenIntegrationUpdateRequest** (1 change)

* Model was added

**DisallowedEntityLearningAssignmentReference** (1 change)

* Model was added

**LearningAssignmentBulkRemoveResponse** (1 change)

* Model was added

**LearningAssignmentEntity** (1 change)

* Model was added

**LearningAssignmentReference** (1 change)

* Model was added


# Point Changes (20 changes)

**GET /api/v2/gamification/scorecards** (1 change)

* Description was changed for parameter workday

**GET /api/v2/gamification/scorecards/attendance** (2 changes)

* Description was changed for parameter startWorkday
* Description was changed for parameter endWorkday

**GET /api/v2/gamification/scorecards/points/alltime** (1 change)

* Description was changed for parameter endWorkday

**GET /api/v2/gamification/scorecards/points/trends** (2 changes)

* Description was changed for parameter startWorkday
* Description was changed for parameter endWorkday

**GET /api/v2/gamification/scorecards/values/trends** (2 changes)

* Description was changed for parameter startWorkday
* Description was changed for parameter endWorkday

**GET /api/v2/gamification/leaderboard** (1 change)

* Summary was changed

**GET /api/v2/gamification/leaderboard/all** (1 change)

* Summary was changed

**GET /api/v2/gamification/scorecards/users/{userId}/attendance** (2 changes)

* Description was changed for parameter startWorkday
* Description was changed for parameter endWorkday

**GET /api/v2/gamification/scorecards/users/{userId}/points/trends** (2 changes)

* Description was changed for parameter startWorkday
* Description was changed for parameter endWorkday

**GET /api/v2/gamification/scorecards/users/{userId}/points/alltime** (1 change)

* Description was changed for parameter endWorkday

**GET /api/v2/gamification/scorecards/users/{userId}** (1 change)

* Description was changed for parameter workday

**GET /api/v2/gamification/scorecards/users/values/trends** (2 changes)

* Description was changed for parameter startWorkday
* Description was changed for parameter endWorkday

**GET /api/v2/gamification/scorecards/users/{userId}/values/trends** (2 changes)

* Description was changed for parameter startWorkday
* Description was changed for parameter endWorkday
