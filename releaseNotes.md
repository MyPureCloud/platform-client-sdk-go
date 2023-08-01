Platform API version: 7205


# Major Changes (3 changes)

**SupportedContentProfile** (1 change)

* Model SupportedContentProfile was removed

**WebDeployment** (1 change)

* Property supportedContentProfile was removed

**ExpandableWebDeployment** (1 change)

* Property supportedContentProfile was removed


# Minor Changes (21 changes)

**AnalyticsEvaluation** (1 change)

* Optional property assigneeApplicable was added

**EvaluationDetailQueryPredicate** (1 change)

* Enum value assigneeApplicable was added to property dimension

**ViewFilter** (2 changes)

* Optional property hasPciData was added
* Optional property hasPiiData was added

**Limit** (1 change)

* Enum value media.communications was added to property namespace

**Evaluation** (1 change)

* Optional property assigneeApplicable was added

**ApprovalNamespace** (1 change)

* Enum value media.communications was added to property namespace

**LimitChangeRequestDetails** (1 change)

* Enum value media.communications was added to property namespace

**StatusChange** (1 change)

* Enum value media.communications was added to property namespace

**EvaluationResponse** (1 change)

* Optional property assigneeApplicable was added

**RecordingMessagingMessage** (2 changes)

* Optional property cards was added
* Optional property contentType was added

**GeneralProgramJobRequest** (5 changes)

* Enum value de-CH was added to property dialect
* Enum value en-HK was added to property dialect
* Enum value en-IE was added to property dialect
* Enum value en-NZ was added to property dialect
* Enum value en-SG was added to property dialect

**LearningAssignmentStep** (1 change)

* Model was added

**LearningAssignmentStepScoStructure** (1 change)

* Model was added

**LearningAssignmentStepSignedCookie** (1 change)

* Model was added

**LearningShareableContentObject** (1 change)

* Model was added


# Point Changes (11 changes)

**GET /api/v2/routing/queues/{queueId}/members** (4 changes)

* Description was changed for parameter name
* Description was changed for parameter profileSkills
* Description was changed for parameter skills
* Description was changed for parameter languages

**GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/activitycodes** (1 change)

* Summary was changed

**GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits** (2 changes)

* Description was changed for parameter feature
* Description was changed for parameter divisionId

**GET /api/v2/workforcemanagement/businessunits** (2 changes)

* Description was changed for parameter feature
* Description was changed for parameter divisionId

**GET /api/v2/workforcemanagement/managementunits** (2 changes)

* Description was changed for parameter feature
* Description was changed for parameter divisionId
