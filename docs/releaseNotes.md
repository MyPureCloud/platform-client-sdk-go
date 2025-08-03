Platform API version: 9484




# Major Changes (7 changes)

**GET /api/v2/taskmanagement/workitems/{workitemId}/users/{userId}/wrapups** (1 change)

* Response 200 type was changed from WorkitemWrapup to WorkitemWrapupEntityListing

**GET /api/v2/businessrules/decisiontables/{tableId}/versions** (1 change)

* Parameter divisionIds was removed

**GET /api/v2/businessrules/decisiontables** (1 change)

* Parameter name was added

**/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows/{rowId}** (1 change)

* Operation PATCH was removed

**UpdateDecisionTableRowRequest** (1 change)

* Model UpdateDecisionTableRowRequest was removed

**NluDomain** (2 changes)

* Property draftVersion was changed from NluDomainVersion to NluDomainVersionReference
* Property lastPublishedVersion was changed from NluDomainVersion to NluDomainVersionReference


# Minor Changes (49 changes)

**/api/v2/outbound/schedules/whatsappcampaigns/{whatsAppCampaignId}** (4 changes)

* Path was added
* Operation GET was added
* Operation PUT was added
* Operation DELETE was added

**/api/v2/outbound/schedules/whatsappcampaigns** (2 changes)

* Path was added
* Operation GET was added

**/api/v2/workforcemanagement/businessunits/{businessUnitId}/staffinggroups/planninggroups/query** (2 changes)

* Path was added
* Operation POST was added

**/api/v2/analytics/actions/aggregates/jobs/{jobId}** (1 change)

* Operation delete was added. Summary: Delete/cancel an async request for action aggregates

**/api/v2/analytics/agentcopilots/aggregates/jobs/{jobId}** (1 change)

* Operation delete was added. Summary: Delete/cancel an async request for agent copilot aggregates

**/api/v2/analytics/bots/aggregates/jobs/{jobId}** (1 change)

* Operation delete was added. Summary: Delete/cancel an async request for bot aggregates

**/api/v2/analytics/conversations/aggregates/jobs/{jobId}** (1 change)

* Operation delete was added. Summary: Delete/cancel an async request for conversation aggregates

**/api/v2/analytics/evaluations/aggregates/jobs/{jobId}** (1 change)

* Operation delete was added. Summary: Delete/cancel an async request for evaluation aggregates

**/api/v2/analytics/flowexecutions/aggregates/jobs/{jobId}** (1 change)

* Operation delete was added. Summary: Delete/cancel an async request for flow execution aggregates

**/api/v2/analytics/flows/aggregates/jobs/{jobId}** (1 change)

* Operation delete was added. Summary: Delete/cancel an async request for flow aggregates

**/api/v2/analytics/journeys/aggregates/jobs/{jobId}** (1 change)

* Operation delete was added. Summary: Delete/cancel an async request for journey aggregates

**/api/v2/analytics/knowledge/aggregates/jobs/{jobId}** (1 change)

* Operation delete was added. Summary: Delete/cancel an async request for knowledge aggregates

**/api/v2/analytics/resolutions/aggregates/jobs/{jobId}** (1 change)

* Operation delete was added. Summary: Delete/cancel an async request for resolution aggregates

**/api/v2/analytics/summaries/aggregates/jobs/{jobId}** (1 change)

* Operation delete was added. Summary: Delete/cancel an async request for summary aggregates

**/api/v2/analytics/surveys/aggregates/jobs/{jobId}** (1 change)

* Operation delete was added. Summary: Delete/cancel an async request for survey aggregates

**/api/v2/analytics/taskmanagement/aggregates/jobs/{jobId}** (1 change)

* Operation delete was added. Summary: Delete/cancel an async request for task management aggregates

**/api/v2/analytics/transcripts/aggregates/jobs/{jobId}** (1 change)

* Operation delete was added. Summary: Delete/cancel an async request for transcript aggregates

**/api/v2/analytics/users/aggregates/jobs/{jobId}** (1 change)

* Operation delete was added. Summary: Delete/cancel an async request for user aggregates

**Queue** (1 change)

* Optional property conditionalGroupActivation was added

**Limit** (1 change)

* Enum value workforce.management.agent.availability was added to property namespace

**WhatsAppCampaignSchedule** (1 change)

* Model was added

**WhatsAppCampaignScheduleEntityListing** (1 change)

* Model was added

**IntentReference** (1 change)

* Model was added

**NluDomainVersionReference** (1 change)

* Model was added

**ApprovalNamespace** (1 change)

* Enum value workforce.management.agent.availability was added to property namespace

**LimitChangeRequestDetails** (1 change)

* Enum value workforce.management.agent.availability was added to property namespace

**StatusChange** (1 change)

* Enum value workforce.management.agent.availability was added to property namespace

**QueueRequest** (1 change)

* Optional property conditionalGroupActivation was added

**UserQueue** (1 change)

* Optional property conditionalGroupActivation was added

**CreateQueueRequest** (1 change)

* Optional property conditionalGroupActivation was added

**Flow** (1 change)

* Optional property agenticVirtualAgentEnabled was added

**FlowVersion** (1 change)

* Optional property agenticVirtualAgentEnabled was added

**Dependency** (1 change)

* Enum value CONVERSATIONCUSTOMATTRIBUTESCHEMA was added to property type

**DependencyObject** (1 change)

* Enum value CONVERSATIONCUSTOMATTRIBUTESCHEMA was added to property type

**HistoryEntry** (1 change)

* Optional property agenticVirtualAgentEnabled was added

**WebMessagingButtonResponse** (1 change)

* Enum value ListPicker was added to property type

**WebMessagingContent** (2 changes)

* Enum value ListPicker was added to property contentType
* Optional property listPicker was added

**StaffingGroupResponse** (1 change)

* Optional property planningGroups was added

**CreateStaffingGroupRequest** (1 change)

* Optional property planningGroupIds was added

**UpdateStaffingGroupRequest** (1 change)

* Optional property planningGroupIds was added

**PlanningGroupToStaffingGroupsListing** (1 change)

* Model was added

**PlanningGroupToStaffingGroupsResponse** (1 change)

* Model was added

**QueryPlanningGroupToStaffingGroupsRequest** (1 change)

* Model was added


# Point Changes (46 changes)

**GET /api/v2/conversations/chats/{conversationId}/messages/{messageId}** (1 change)

* Description was changed

**GET /api/v2/conversations/chats/{conversationId}/messages** (1 change)

* Description was changed

**POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/messages** (1 change)

* Description was changed

**POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/typing** (1 change)

* Description was changed

**GET /api/v2/conversations/chats** (1 change)

* Description was changed

**POST /api/v2/conversations/chats** (1 change)

* Description was changed

**POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace** (1 change)

* Description was changed

**GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes** (1 change)

* Description was changed

**GET /api/v2/conversations/chats/{conversationId}** (1 change)

* Description was changed

**PATCH /api/v2/conversations/chats/{conversationId}** (1 change)

* Description was changed

**PUT /api/v2/conversations/chats/{conversationId}/recordingstate** (1 change)

* Description was changed

**PATCH /api/v2/conversations/chats/{conversationId}/participants/{participantId}/attributes** (1 change)

* Description was changed

**GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup** (1 change)

* Description was changed

**PATCH /api/v2/conversations/chats/{conversationId}/participants/{participantId}** (1 change)

* Description was changed

**PATCH /api/v2/conversations/chats/{conversationId}/participants/{participantId}/communications/{communicationId}** (1 change)

* Description was changed

**GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/communications/{communicationId}/wrapup** (1 change)

* Description was changed

**POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/communications/{communicationId}/wrapup** (1 change)

* Description was changed

**POST /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/replace** (1 change)

* Description was changed

**GET /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/wrapupcodes** (1 change)

* Description was changed

**PATCH /api/v2/conversations/cobrowsesessions/{conversationId}** (1 change)

* Description was changed

**PUT /api/v2/conversations/cobrowsesessions/{conversationId}/recordingstate** (1 change)

* Description was changed

**PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/attributes** (1 change)

* Description was changed

**GET /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/wrapup** (1 change)

* Description was changed

**PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/communications/{communicationId}** (1 change)

* Description was changed

**GET /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/communications/{communicationId}/wrapup** (1 change)

* Description was changed

**POST /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/communications/{communicationId}/wrapup** (1 change)

* Description was changed

**PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}** (1 change)

* Description was changed

**PUT /api/v2/conversations/screenshares/{conversationId}/recordingstate** (1 change)

* Description was changed

**GET /api/v2/conversations/screenshares/{conversationId}/participants/{participantId}/communications/{communicationId}/wrapup** (1 change)

* Description was changed

**POST /api/v2/conversations/screenshares/{conversationId}/participants/{participantId}/communications/{communicationId}/wrapup** (1 change)

* Description was changed

**POST /api/v2/webchat/guest/conversations** (1 change)

* Description was changed

**GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}** (1 change)

* Description was changed

**PATCH /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}** (1 change)

* Description was changed

**GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests** (1 change)

* Description was changed

**POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/messages** (1 change)

* Description was changed

**POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/typing** (1 change)

* Description was changed

**GET /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}** (1 change)

* Description was changed

**DELETE /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}** (1 change)

* Description was changed

**GET /api/v2/webchat/guest/conversations/{conversationId}/members** (1 change)

* Description was changed

**GET /api/v2/webchat/guest/conversations/{conversationId}/messages/{messageId}** (1 change)

* Description was changed

**GET /api/v2/webchat/guest/conversations/{conversationId}/messages** (1 change)

* Description was changed

**GET /api/v2/widgets/deployments/{deploymentId}** (1 change)

* Description was changed

**PUT /api/v2/widgets/deployments/{deploymentId}** (1 change)

* Description was changed

**DELETE /api/v2/widgets/deployments/{deploymentId}** (1 change)

* Description was changed

**GET /api/v2/widgets/deployments** (1 change)

* Description was changed

**POST /api/v2/widgets/deployments** (1 change)

* Description was changed
