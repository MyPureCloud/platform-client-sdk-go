Platform API version: 5632


# Major Changes (6 changes)

**/api/v2/gamification/profiles/{performanceProfileId}** (1 change)

* Path /api/v2/gamification/profiles/{performanceProfileId} was removed

**/api/v2/gamification/profiles/{performanceProfileId}/activate** (1 change)

* Path /api/v2/gamification/profiles/{performanceProfileId}/activate was removed

**/api/v2/gamification/profiles/{performanceProfileId}/deactivate** (1 change)

* Path /api/v2/gamification/profiles/{performanceProfileId}/deactivate was removed

**Calibration** (1 change)

* Property conversation was changed from Conversation to ConversationReference

**Evaluation** (1 change)

* Property conversation was changed from Conversation to ConversationReference

**CalibrationCreate** (1 change)

* Property conversation was changed from Conversation to ConversationReference


# Minor Changes (14 changes)

**POST /api/v2/conversations/messaging/integrations/whatsapp** (1 change)

* Response 202 was added

**/api/v2/gamification/profiles/{profileId}** (3 changes)

* Path was added
* Operation GET was added
* Operation PUT was added

**/api/v2/gamification/profiles/{profileId}/activate** (2 changes)

* Path was added
* Operation POST was added

**/api/v2/gamification/profiles/{profileId}/deactivate** (2 changes)

* Path was added
* Operation POST was added

**MemberGroup** (1 change)

* Enum value SKILLGROUP was added to property type

**WhatsAppAvailablePhoneNumberDetails** (1 change)

* Model was added

**WhatsAppAvailablePhoneNumberDetailsListing** (1 change)

* Model was added

**WhatsAppIntegration** (1 change)

* Optional property availablePhoneNumbers was added

**WhatsAppIntegrationUpdateRequest** (2 changes)

* name is no longer readonly
* Optional property phoneNumber was added


# Point Changes (4 changes)

**PATCH /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}** (2 changes)

* Description was changed
* Summary was changed

**POST /api/v2/recording/jobs** (2 changes)

* Description was changed
* Summary was changed
