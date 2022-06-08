package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeysessioneventsnotificationsessionevent
type Journeysessioneventsnotificationsessionevent struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// CreatedDate
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// EndedDate
	EndedDate *time.Time `json:"endedDate,omitempty"`


	// ExternalContact
	ExternalContact *Journeysessioneventsnotificationexternalcontact `json:"externalContact,omitempty"`


	// CustomerId
	CustomerId *string `json:"customerId,omitempty"`


	// CustomerIdType
	CustomerIdType *string `json:"customerIdType,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// ExternalId
	ExternalId *string `json:"externalId,omitempty"`


	// ExternalUrl
	ExternalUrl *string `json:"externalUrl,omitempty"`


	// OutcomeAchievements
	OutcomeAchievements *[]Journeysessioneventsnotificationoutcomeachievement `json:"outcomeAchievements,omitempty"`


	// SegmentAssignments
	SegmentAssignments *[]Journeysessioneventsnotificationsegmentassignment `json:"segmentAssignments,omitempty"`


	// Attributes
	Attributes *map[string]Journeysessioneventsnotificationcustomeventattribute `json:"attributes,omitempty"`


	// AttributeLists
	AttributeLists *map[string]Journeysessioneventsnotificationcustomeventattributelist `json:"attributeLists,omitempty"`


	// AwayDate
	AwayDate *time.Time `json:"awayDate,omitempty"`


	// Browser
	Browser *Journeysessioneventsnotificationbrowser `json:"browser,omitempty"`


	// Device
	Device *Journeysessioneventsnotificationdevice `json:"device,omitempty"`


	// Geolocation
	Geolocation *Journeysessioneventsnotificationgeolocation `json:"geolocation,omitempty"`


	// IdleDate
	IdleDate *time.Time `json:"idleDate,omitempty"`


	// IpAddress
	IpAddress *string `json:"ipAddress,omitempty"`


	// IpOrganization
	IpOrganization *string `json:"ipOrganization,omitempty"`


	// LastPage
	LastPage *Journeysessioneventsnotificationpage `json:"lastPage,omitempty"`


	// MktCampaign
	MktCampaign *Journeysessioneventsnotificationmktcampaign `json:"mktCampaign,omitempty"`


	// Referrer
	Referrer *Journeysessioneventsnotificationreferrer `json:"referrer,omitempty"`


	// SearchTerms
	SearchTerms *[]string `json:"searchTerms,omitempty"`


	// UserAgentString
	UserAgentString *string `json:"userAgentString,omitempty"`


	// DurationInSeconds
	DurationInSeconds *int `json:"durationInSeconds,omitempty"`


	// EventCount
	EventCount *int `json:"eventCount,omitempty"`


	// PageviewCount
	PageviewCount *int `json:"pageviewCount,omitempty"`


	// ScreenviewCount
	ScreenviewCount *int `json:"screenviewCount,omitempty"`


	// LastEvent
	LastEvent *Journeysessioneventsnotificationsessionlastevent `json:"lastEvent,omitempty"`


	// Conversation
	Conversation *Journeysessioneventsnotificationconversation `json:"conversation,omitempty"`


	// OriginatingDirection
	OriginatingDirection *string `json:"originatingDirection,omitempty"`


	// ConversationSubject
	ConversationSubject *string `json:"conversationSubject,omitempty"`


	// LastUserDisposition
	LastUserDisposition *Journeysessioneventsnotificationconversationuserdisposition `json:"lastUserDisposition,omitempty"`


	// LastConnectedUser
	LastConnectedUser *Journeysessioneventsnotificationuser `json:"lastConnectedUser,omitempty"`


	// LastConnectedQueue
	LastConnectedQueue *Journeysessioneventsnotificationconnectedqueue `json:"lastConnectedQueue,omitempty"`


	// ConversationChannels
	ConversationChannels *[]Journeysessioneventsnotificationconversationchannel `json:"conversationChannels,omitempty"`


	// LastUserDisconnectType
	LastUserDisconnectType *string `json:"lastUserDisconnectType,omitempty"`


	// LastAcdOutcome
	LastAcdOutcome *string `json:"lastAcdOutcome,omitempty"`


	// Authenticated
	Authenticated *bool `json:"authenticated,omitempty"`

}

func (o *Journeysessioneventsnotificationsessionevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeysessioneventsnotificationsessionevent
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	EndedDate := new(string)
	if o.EndedDate != nil {
		
		*EndedDate = timeutil.Strftime(o.EndedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndedDate = nil
	}
	
	AwayDate := new(string)
	if o.AwayDate != nil {
		
		*AwayDate = timeutil.Strftime(o.AwayDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		AwayDate = nil
	}
	
	IdleDate := new(string)
	if o.IdleDate != nil {
		
		*IdleDate = timeutil.Strftime(o.IdleDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		IdleDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		EndedDate *string `json:"endedDate,omitempty"`
		
		ExternalContact *Journeysessioneventsnotificationexternalcontact `json:"externalContact,omitempty"`
		
		CustomerId *string `json:"customerId,omitempty"`
		
		CustomerIdType *string `json:"customerIdType,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		ExternalId *string `json:"externalId,omitempty"`
		
		ExternalUrl *string `json:"externalUrl,omitempty"`
		
		OutcomeAchievements *[]Journeysessioneventsnotificationoutcomeachievement `json:"outcomeAchievements,omitempty"`
		
		SegmentAssignments *[]Journeysessioneventsnotificationsegmentassignment `json:"segmentAssignments,omitempty"`
		
		Attributes *map[string]Journeysessioneventsnotificationcustomeventattribute `json:"attributes,omitempty"`
		
		AttributeLists *map[string]Journeysessioneventsnotificationcustomeventattributelist `json:"attributeLists,omitempty"`
		
		AwayDate *string `json:"awayDate,omitempty"`
		
		Browser *Journeysessioneventsnotificationbrowser `json:"browser,omitempty"`
		
		Device *Journeysessioneventsnotificationdevice `json:"device,omitempty"`
		
		Geolocation *Journeysessioneventsnotificationgeolocation `json:"geolocation,omitempty"`
		
		IdleDate *string `json:"idleDate,omitempty"`
		
		IpAddress *string `json:"ipAddress,omitempty"`
		
		IpOrganization *string `json:"ipOrganization,omitempty"`
		
		LastPage *Journeysessioneventsnotificationpage `json:"lastPage,omitempty"`
		
		MktCampaign *Journeysessioneventsnotificationmktcampaign `json:"mktCampaign,omitempty"`
		
		Referrer *Journeysessioneventsnotificationreferrer `json:"referrer,omitempty"`
		
		SearchTerms *[]string `json:"searchTerms,omitempty"`
		
		UserAgentString *string `json:"userAgentString,omitempty"`
		
		DurationInSeconds *int `json:"durationInSeconds,omitempty"`
		
		EventCount *int `json:"eventCount,omitempty"`
		
		PageviewCount *int `json:"pageviewCount,omitempty"`
		
		ScreenviewCount *int `json:"screenviewCount,omitempty"`
		
		LastEvent *Journeysessioneventsnotificationsessionlastevent `json:"lastEvent,omitempty"`
		
		Conversation *Journeysessioneventsnotificationconversation `json:"conversation,omitempty"`
		
		OriginatingDirection *string `json:"originatingDirection,omitempty"`
		
		ConversationSubject *string `json:"conversationSubject,omitempty"`
		
		LastUserDisposition *Journeysessioneventsnotificationconversationuserdisposition `json:"lastUserDisposition,omitempty"`
		
		LastConnectedUser *Journeysessioneventsnotificationuser `json:"lastConnectedUser,omitempty"`
		
		LastConnectedQueue *Journeysessioneventsnotificationconnectedqueue `json:"lastConnectedQueue,omitempty"`
		
		ConversationChannels *[]Journeysessioneventsnotificationconversationchannel `json:"conversationChannels,omitempty"`
		
		LastUserDisconnectType *string `json:"lastUserDisconnectType,omitempty"`
		
		LastAcdOutcome *string `json:"lastAcdOutcome,omitempty"`
		
		Authenticated *bool `json:"authenticated,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		SelfUri: o.SelfUri,
		
		CreatedDate: CreatedDate,
		
		EndedDate: EndedDate,
		
		ExternalContact: o.ExternalContact,
		
		CustomerId: o.CustomerId,
		
		CustomerIdType: o.CustomerIdType,
		
		VarType: o.VarType,
		
		ExternalId: o.ExternalId,
		
		ExternalUrl: o.ExternalUrl,
		
		OutcomeAchievements: o.OutcomeAchievements,
		
		SegmentAssignments: o.SegmentAssignments,
		
		Attributes: o.Attributes,
		
		AttributeLists: o.AttributeLists,
		
		AwayDate: AwayDate,
		
		Browser: o.Browser,
		
		Device: o.Device,
		
		Geolocation: o.Geolocation,
		
		IdleDate: IdleDate,
		
		IpAddress: o.IpAddress,
		
		IpOrganization: o.IpOrganization,
		
		LastPage: o.LastPage,
		
		MktCampaign: o.MktCampaign,
		
		Referrer: o.Referrer,
		
		SearchTerms: o.SearchTerms,
		
		UserAgentString: o.UserAgentString,
		
		DurationInSeconds: o.DurationInSeconds,
		
		EventCount: o.EventCount,
		
		PageviewCount: o.PageviewCount,
		
		ScreenviewCount: o.ScreenviewCount,
		
		LastEvent: o.LastEvent,
		
		Conversation: o.Conversation,
		
		OriginatingDirection: o.OriginatingDirection,
		
		ConversationSubject: o.ConversationSubject,
		
		LastUserDisposition: o.LastUserDisposition,
		
		LastConnectedUser: o.LastConnectedUser,
		
		LastConnectedQueue: o.LastConnectedQueue,
		
		ConversationChannels: o.ConversationChannels,
		
		LastUserDisconnectType: o.LastUserDisconnectType,
		
		LastAcdOutcome: o.LastAcdOutcome,
		
		Authenticated: o.Authenticated,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeysessioneventsnotificationsessionevent) UnmarshalJSON(b []byte) error {
	var JourneysessioneventsnotificationsessioneventMap map[string]interface{}
	err := json.Unmarshal(b, &JourneysessioneventsnotificationsessioneventMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JourneysessioneventsnotificationsessioneventMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if SelfUri, ok := JourneysessioneventsnotificationsessioneventMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if createdDateString, ok := JourneysessioneventsnotificationsessioneventMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if endedDateString, ok := JourneysessioneventsnotificationsessioneventMap["endedDate"].(string); ok {
		EndedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endedDateString)
		o.EndedDate = &EndedDate
	}
	
	if ExternalContact, ok := JourneysessioneventsnotificationsessioneventMap["externalContact"].(map[string]interface{}); ok {
		ExternalContactString, _ := json.Marshal(ExternalContact)
		json.Unmarshal(ExternalContactString, &o.ExternalContact)
	}
	
	if CustomerId, ok := JourneysessioneventsnotificationsessioneventMap["customerId"].(string); ok {
		o.CustomerId = &CustomerId
	}
    
	if CustomerIdType, ok := JourneysessioneventsnotificationsessioneventMap["customerIdType"].(string); ok {
		o.CustomerIdType = &CustomerIdType
	}
    
	if VarType, ok := JourneysessioneventsnotificationsessioneventMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if ExternalId, ok := JourneysessioneventsnotificationsessioneventMap["externalId"].(string); ok {
		o.ExternalId = &ExternalId
	}
    
	if ExternalUrl, ok := JourneysessioneventsnotificationsessioneventMap["externalUrl"].(string); ok {
		o.ExternalUrl = &ExternalUrl
	}
    
	if OutcomeAchievements, ok := JourneysessioneventsnotificationsessioneventMap["outcomeAchievements"].([]interface{}); ok {
		OutcomeAchievementsString, _ := json.Marshal(OutcomeAchievements)
		json.Unmarshal(OutcomeAchievementsString, &o.OutcomeAchievements)
	}
	
	if SegmentAssignments, ok := JourneysessioneventsnotificationsessioneventMap["segmentAssignments"].([]interface{}); ok {
		SegmentAssignmentsString, _ := json.Marshal(SegmentAssignments)
		json.Unmarshal(SegmentAssignmentsString, &o.SegmentAssignments)
	}
	
	if Attributes, ok := JourneysessioneventsnotificationsessioneventMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if AttributeLists, ok := JourneysessioneventsnotificationsessioneventMap["attributeLists"].(map[string]interface{}); ok {
		AttributeListsString, _ := json.Marshal(AttributeLists)
		json.Unmarshal(AttributeListsString, &o.AttributeLists)
	}
	
	if awayDateString, ok := JourneysessioneventsnotificationsessioneventMap["awayDate"].(string); ok {
		AwayDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", awayDateString)
		o.AwayDate = &AwayDate
	}
	
	if Browser, ok := JourneysessioneventsnotificationsessioneventMap["browser"].(map[string]interface{}); ok {
		BrowserString, _ := json.Marshal(Browser)
		json.Unmarshal(BrowserString, &o.Browser)
	}
	
	if Device, ok := JourneysessioneventsnotificationsessioneventMap["device"].(map[string]interface{}); ok {
		DeviceString, _ := json.Marshal(Device)
		json.Unmarshal(DeviceString, &o.Device)
	}
	
	if Geolocation, ok := JourneysessioneventsnotificationsessioneventMap["geolocation"].(map[string]interface{}); ok {
		GeolocationString, _ := json.Marshal(Geolocation)
		json.Unmarshal(GeolocationString, &o.Geolocation)
	}
	
	if idleDateString, ok := JourneysessioneventsnotificationsessioneventMap["idleDate"].(string); ok {
		IdleDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", idleDateString)
		o.IdleDate = &IdleDate
	}
	
	if IpAddress, ok := JourneysessioneventsnotificationsessioneventMap["ipAddress"].(string); ok {
		o.IpAddress = &IpAddress
	}
    
	if IpOrganization, ok := JourneysessioneventsnotificationsessioneventMap["ipOrganization"].(string); ok {
		o.IpOrganization = &IpOrganization
	}
    
	if LastPage, ok := JourneysessioneventsnotificationsessioneventMap["lastPage"].(map[string]interface{}); ok {
		LastPageString, _ := json.Marshal(LastPage)
		json.Unmarshal(LastPageString, &o.LastPage)
	}
	
	if MktCampaign, ok := JourneysessioneventsnotificationsessioneventMap["mktCampaign"].(map[string]interface{}); ok {
		MktCampaignString, _ := json.Marshal(MktCampaign)
		json.Unmarshal(MktCampaignString, &o.MktCampaign)
	}
	
	if Referrer, ok := JourneysessioneventsnotificationsessioneventMap["referrer"].(map[string]interface{}); ok {
		ReferrerString, _ := json.Marshal(Referrer)
		json.Unmarshal(ReferrerString, &o.Referrer)
	}
	
	if SearchTerms, ok := JourneysessioneventsnotificationsessioneventMap["searchTerms"].([]interface{}); ok {
		SearchTermsString, _ := json.Marshal(SearchTerms)
		json.Unmarshal(SearchTermsString, &o.SearchTerms)
	}
	
	if UserAgentString, ok := JourneysessioneventsnotificationsessioneventMap["userAgentString"].(string); ok {
		o.UserAgentString = &UserAgentString
	}
    
	if DurationInSeconds, ok := JourneysessioneventsnotificationsessioneventMap["durationInSeconds"].(float64); ok {
		DurationInSecondsInt := int(DurationInSeconds)
		o.DurationInSeconds = &DurationInSecondsInt
	}
	
	if EventCount, ok := JourneysessioneventsnotificationsessioneventMap["eventCount"].(float64); ok {
		EventCountInt := int(EventCount)
		o.EventCount = &EventCountInt
	}
	
	if PageviewCount, ok := JourneysessioneventsnotificationsessioneventMap["pageviewCount"].(float64); ok {
		PageviewCountInt := int(PageviewCount)
		o.PageviewCount = &PageviewCountInt
	}
	
	if ScreenviewCount, ok := JourneysessioneventsnotificationsessioneventMap["screenviewCount"].(float64); ok {
		ScreenviewCountInt := int(ScreenviewCount)
		o.ScreenviewCount = &ScreenviewCountInt
	}
	
	if LastEvent, ok := JourneysessioneventsnotificationsessioneventMap["lastEvent"].(map[string]interface{}); ok {
		LastEventString, _ := json.Marshal(LastEvent)
		json.Unmarshal(LastEventString, &o.LastEvent)
	}
	
	if Conversation, ok := JourneysessioneventsnotificationsessioneventMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if OriginatingDirection, ok := JourneysessioneventsnotificationsessioneventMap["originatingDirection"].(string); ok {
		o.OriginatingDirection = &OriginatingDirection
	}
    
	if ConversationSubject, ok := JourneysessioneventsnotificationsessioneventMap["conversationSubject"].(string); ok {
		o.ConversationSubject = &ConversationSubject
	}
    
	if LastUserDisposition, ok := JourneysessioneventsnotificationsessioneventMap["lastUserDisposition"].(map[string]interface{}); ok {
		LastUserDispositionString, _ := json.Marshal(LastUserDisposition)
		json.Unmarshal(LastUserDispositionString, &o.LastUserDisposition)
	}
	
	if LastConnectedUser, ok := JourneysessioneventsnotificationsessioneventMap["lastConnectedUser"].(map[string]interface{}); ok {
		LastConnectedUserString, _ := json.Marshal(LastConnectedUser)
		json.Unmarshal(LastConnectedUserString, &o.LastConnectedUser)
	}
	
	if LastConnectedQueue, ok := JourneysessioneventsnotificationsessioneventMap["lastConnectedQueue"].(map[string]interface{}); ok {
		LastConnectedQueueString, _ := json.Marshal(LastConnectedQueue)
		json.Unmarshal(LastConnectedQueueString, &o.LastConnectedQueue)
	}
	
	if ConversationChannels, ok := JourneysessioneventsnotificationsessioneventMap["conversationChannels"].([]interface{}); ok {
		ConversationChannelsString, _ := json.Marshal(ConversationChannels)
		json.Unmarshal(ConversationChannelsString, &o.ConversationChannels)
	}
	
	if LastUserDisconnectType, ok := JourneysessioneventsnotificationsessioneventMap["lastUserDisconnectType"].(string); ok {
		o.LastUserDisconnectType = &LastUserDisconnectType
	}
    
	if LastAcdOutcome, ok := JourneysessioneventsnotificationsessioneventMap["lastAcdOutcome"].(string); ok {
		o.LastAcdOutcome = &LastAcdOutcome
	}
    
	if Authenticated, ok := JourneysessioneventsnotificationsessioneventMap["authenticated"].(bool); ok {
		o.Authenticated = &Authenticated
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeysessioneventsnotificationsessionevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
