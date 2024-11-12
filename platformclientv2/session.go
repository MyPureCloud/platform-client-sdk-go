package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Session
type Session struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The ID of the session.
	Id *string `json:"id,omitempty"`

	// CustomerId - Primary identifier of the customer in the source where the events for the session originate from.
	CustomerId *string `json:"customerId,omitempty"`

	// CustomerIdType - Type of source customer identifier (e.g. cookie, email, phone).
	CustomerIdType *string `json:"customerIdType,omitempty"`

	// VarType - Session types indicate the type or category of sessions (e.g. web, app).
	VarType *string `json:"type,omitempty"`

	// ExternalId - Unique identifier in the external system where the events for the session originate from.
	ExternalId *string `json:"externalId,omitempty"`

	// ExternalUrl - A URL that identifies an external system-of-record resource that may have more detailed information on the session.
	ExternalUrl *string `json:"externalUrl,omitempty"`

	// ShortId - Shortened numeric identifier of 4-6 digits.
	ShortId *string `json:"shortId,omitempty"`

	// OutcomeAchievements - List of the outcome achievements by the customer in this session.
	OutcomeAchievements *[]Outcomeachievement `json:"outcomeAchievements,omitempty"`

	// SegmentAssignments - List of the segment assignments to the customer in this session.
	SegmentAssignments *[]Sessionsegmentassignment `json:"segmentAssignments,omitempty"`

	// Attributes - Attributes projected from the session's event stream.
	Attributes *map[string]Customeventattribute `json:"attributes,omitempty"`

	// AttributeLists - List-type attributes projected from the session's event stream.
	AttributeLists *map[string]Customeventattributelist `json:"attributeLists,omitempty"`

	// Browser - Customer's browser.
	Browser *Browser `json:"browser,omitempty"`

	// Device - Customer's device.
	Device *Device `json:"device,omitempty"`

	// Geolocation - Customer's geolocation.
	Geolocation *Journeygeolocation `json:"geolocation,omitempty"`

	// IpAddress - Customer's IP address.
	IpAddress *string `json:"ipAddress,omitempty"`

	// IpOrganization - Customer's IP-based organization or ISP name.
	IpOrganization *string `json:"ipOrganization,omitempty"`

	// LastPage - The webpage where the customer's last web interaction occurred.
	LastPage *Journeypage `json:"lastPage,omitempty"`

	// MktCampaign - Marketing / traffic source information.
	MktCampaign *Journeycampaign `json:"mktCampaign,omitempty"`

	// Referrer - Identifies the page URL that originally generated the request for the current page being viewed.
	Referrer *Referrer `json:"referrer,omitempty"`

	// App - Application that the customer is interacting with (for app sessions).
	App *Journeyapp `json:"app,omitempty"`

	// SdkLibrary - SDK library used to generate the events for the session (for app and web sessions).
	SdkLibrary *Sdklibrary `json:"sdkLibrary,omitempty"`

	// NetworkConnectivity - Information relating to the device's network connectivity (for app sessions).
	NetworkConnectivity *Networkconnectivity `json:"networkConnectivity,omitempty"`

	// SearchTerms - Search terms associated with the session.
	SearchTerms *[]string `json:"searchTerms,omitempty"`

	// UserAgentString - String identifying the user agent.
	UserAgentString *string `json:"userAgentString,omitempty"`

	// DurationInSeconds - Indicates how long the session has been active (valid for an individual device).
	DurationInSeconds *int `json:"durationInSeconds,omitempty"`

	// EventCount - The count of all events performed during the session.
	EventCount *int `json:"eventCount,omitempty"`

	// PageviewCount - The count of all pageviews performed during the session.
	PageviewCount *int `json:"pageviewCount,omitempty"`

	// ScreenviewCount - The count of all screenviews performed during the session.
	ScreenviewCount *int `json:"screenviewCount,omitempty"`

	// LastEvent - Information about the most recent event in this session.
	LastEvent *Sessionlastevent `json:"lastEvent,omitempty"`

	// LastConnectedQueue - The last queue connected to this session.
	LastConnectedQueue *Connectedqueue `json:"lastConnectedQueue,omitempty"`

	// LastConnectedUser - The last user connected to this session.
	LastConnectedUser *Connecteduser `json:"lastConnectedUser,omitempty"`

	// LastUserDisposition - The last user disposition connected to this session.
	LastUserDisposition *Conversationuserdisposition `json:"lastUserDisposition,omitempty"`

	// ConversationChannels - Represents the channels used for this conversation.
	ConversationChannels *[]Conversationchannel `json:"conversationChannels,omitempty"`

	// OriginatingDirection - The original direction of the conversation.
	OriginatingDirection *string `json:"originatingDirection,omitempty"`

	// ConversationSubject - The subject for the conversation, for example an email subject.
	ConversationSubject *string `json:"conversationSubject,omitempty"`

	// LastUserDisconnectType - Disconnect reason for the last user connected to the conversation.
	LastUserDisconnectType *string `json:"lastUserDisconnectType,omitempty"`

	// LastAcdOutcome - Last ACD outcome for the conversation.
	LastAcdOutcome *string `json:"lastAcdOutcome,omitempty"`

	// Authenticated - Indicates whether or not the session is authenticated.
	Authenticated *bool `json:"authenticated,omitempty"`

	// LastScreen - The app screen name where the customer's last app interaction occurred.
	LastScreen *string `json:"lastScreen,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

	// CreatedDate - Timestamp indicating when the session was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`

	// EndedDate - Timestamp indicating when the session was ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndedDate *time.Time `json:"endedDate,omitempty"`

	// ExternalContact - The external contact associated with this session.
	ExternalContact *Addressableentityref `json:"externalContact,omitempty"`

	// AwayDate - Timestamp indicating when the visitor should be considered as away. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	AwayDate *time.Time `json:"awayDate,omitempty"`

	// IdleDate - Timestamp indicating when the visitor should be considered as idle. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	IdleDate *time.Time `json:"idleDate,omitempty"`

	// Conversation - The conversation for this session.
	Conversation *Addressableentityref `json:"conversation,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Session) SetField(field string, fieldValue interface{}) {
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

func (o Session) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CreatedDate","EndedDate","AwayDate","IdleDate", }
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
	type Alias Session
	
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
		
		CustomerId *string `json:"customerId,omitempty"`
		
		CustomerIdType *string `json:"customerIdType,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		ExternalId *string `json:"externalId,omitempty"`
		
		ExternalUrl *string `json:"externalUrl,omitempty"`
		
		ShortId *string `json:"shortId,omitempty"`
		
		OutcomeAchievements *[]Outcomeachievement `json:"outcomeAchievements,omitempty"`
		
		SegmentAssignments *[]Sessionsegmentassignment `json:"segmentAssignments,omitempty"`
		
		Attributes *map[string]Customeventattribute `json:"attributes,omitempty"`
		
		AttributeLists *map[string]Customeventattributelist `json:"attributeLists,omitempty"`
		
		Browser *Browser `json:"browser,omitempty"`
		
		Device *Device `json:"device,omitempty"`
		
		Geolocation *Journeygeolocation `json:"geolocation,omitempty"`
		
		IpAddress *string `json:"ipAddress,omitempty"`
		
		IpOrganization *string `json:"ipOrganization,omitempty"`
		
		LastPage *Journeypage `json:"lastPage,omitempty"`
		
		MktCampaign *Journeycampaign `json:"mktCampaign,omitempty"`
		
		Referrer *Referrer `json:"referrer,omitempty"`
		
		App *Journeyapp `json:"app,omitempty"`
		
		SdkLibrary *Sdklibrary `json:"sdkLibrary,omitempty"`
		
		NetworkConnectivity *Networkconnectivity `json:"networkConnectivity,omitempty"`
		
		SearchTerms *[]string `json:"searchTerms,omitempty"`
		
		UserAgentString *string `json:"userAgentString,omitempty"`
		
		DurationInSeconds *int `json:"durationInSeconds,omitempty"`
		
		EventCount *int `json:"eventCount,omitempty"`
		
		PageviewCount *int `json:"pageviewCount,omitempty"`
		
		ScreenviewCount *int `json:"screenviewCount,omitempty"`
		
		LastEvent *Sessionlastevent `json:"lastEvent,omitempty"`
		
		LastConnectedQueue *Connectedqueue `json:"lastConnectedQueue,omitempty"`
		
		LastConnectedUser *Connecteduser `json:"lastConnectedUser,omitempty"`
		
		LastUserDisposition *Conversationuserdisposition `json:"lastUserDisposition,omitempty"`
		
		ConversationChannels *[]Conversationchannel `json:"conversationChannels,omitempty"`
		
		OriginatingDirection *string `json:"originatingDirection,omitempty"`
		
		ConversationSubject *string `json:"conversationSubject,omitempty"`
		
		LastUserDisconnectType *string `json:"lastUserDisconnectType,omitempty"`
		
		LastAcdOutcome *string `json:"lastAcdOutcome,omitempty"`
		
		Authenticated *bool `json:"authenticated,omitempty"`
		
		LastScreen *string `json:"lastScreen,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		EndedDate *string `json:"endedDate,omitempty"`
		
		ExternalContact *Addressableentityref `json:"externalContact,omitempty"`
		
		AwayDate *string `json:"awayDate,omitempty"`
		
		IdleDate *string `json:"idleDate,omitempty"`
		
		Conversation *Addressableentityref `json:"conversation,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		CustomerId: o.CustomerId,
		
		CustomerIdType: o.CustomerIdType,
		
		VarType: o.VarType,
		
		ExternalId: o.ExternalId,
		
		ExternalUrl: o.ExternalUrl,
		
		ShortId: o.ShortId,
		
		OutcomeAchievements: o.OutcomeAchievements,
		
		SegmentAssignments: o.SegmentAssignments,
		
		Attributes: o.Attributes,
		
		AttributeLists: o.AttributeLists,
		
		Browser: o.Browser,
		
		Device: o.Device,
		
		Geolocation: o.Geolocation,
		
		IpAddress: o.IpAddress,
		
		IpOrganization: o.IpOrganization,
		
		LastPage: o.LastPage,
		
		MktCampaign: o.MktCampaign,
		
		Referrer: o.Referrer,
		
		App: o.App,
		
		SdkLibrary: o.SdkLibrary,
		
		NetworkConnectivity: o.NetworkConnectivity,
		
		SearchTerms: o.SearchTerms,
		
		UserAgentString: o.UserAgentString,
		
		DurationInSeconds: o.DurationInSeconds,
		
		EventCount: o.EventCount,
		
		PageviewCount: o.PageviewCount,
		
		ScreenviewCount: o.ScreenviewCount,
		
		LastEvent: o.LastEvent,
		
		LastConnectedQueue: o.LastConnectedQueue,
		
		LastConnectedUser: o.LastConnectedUser,
		
		LastUserDisposition: o.LastUserDisposition,
		
		ConversationChannels: o.ConversationChannels,
		
		OriginatingDirection: o.OriginatingDirection,
		
		ConversationSubject: o.ConversationSubject,
		
		LastUserDisconnectType: o.LastUserDisconnectType,
		
		LastAcdOutcome: o.LastAcdOutcome,
		
		Authenticated: o.Authenticated,
		
		LastScreen: o.LastScreen,
		
		SelfUri: o.SelfUri,
		
		CreatedDate: CreatedDate,
		
		EndedDate: EndedDate,
		
		ExternalContact: o.ExternalContact,
		
		AwayDate: AwayDate,
		
		IdleDate: IdleDate,
		
		Conversation: o.Conversation,
		Alias:    (Alias)(o),
	})
}

func (o *Session) UnmarshalJSON(b []byte) error {
	var SessionMap map[string]interface{}
	err := json.Unmarshal(b, &SessionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SessionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if CustomerId, ok := SessionMap["customerId"].(string); ok {
		o.CustomerId = &CustomerId
	}
    
	if CustomerIdType, ok := SessionMap["customerIdType"].(string); ok {
		o.CustomerIdType = &CustomerIdType
	}
    
	if VarType, ok := SessionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if ExternalId, ok := SessionMap["externalId"].(string); ok {
		o.ExternalId = &ExternalId
	}
    
	if ExternalUrl, ok := SessionMap["externalUrl"].(string); ok {
		o.ExternalUrl = &ExternalUrl
	}
    
	if ShortId, ok := SessionMap["shortId"].(string); ok {
		o.ShortId = &ShortId
	}
    
	if OutcomeAchievements, ok := SessionMap["outcomeAchievements"].([]interface{}); ok {
		OutcomeAchievementsString, _ := json.Marshal(OutcomeAchievements)
		json.Unmarshal(OutcomeAchievementsString, &o.OutcomeAchievements)
	}
	
	if SegmentAssignments, ok := SessionMap["segmentAssignments"].([]interface{}); ok {
		SegmentAssignmentsString, _ := json.Marshal(SegmentAssignments)
		json.Unmarshal(SegmentAssignmentsString, &o.SegmentAssignments)
	}
	
	if Attributes, ok := SessionMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if AttributeLists, ok := SessionMap["attributeLists"].(map[string]interface{}); ok {
		AttributeListsString, _ := json.Marshal(AttributeLists)
		json.Unmarshal(AttributeListsString, &o.AttributeLists)
	}
	
	if Browser, ok := SessionMap["browser"].(map[string]interface{}); ok {
		BrowserString, _ := json.Marshal(Browser)
		json.Unmarshal(BrowserString, &o.Browser)
	}
	
	if Device, ok := SessionMap["device"].(map[string]interface{}); ok {
		DeviceString, _ := json.Marshal(Device)
		json.Unmarshal(DeviceString, &o.Device)
	}
	
	if Geolocation, ok := SessionMap["geolocation"].(map[string]interface{}); ok {
		GeolocationString, _ := json.Marshal(Geolocation)
		json.Unmarshal(GeolocationString, &o.Geolocation)
	}
	
	if IpAddress, ok := SessionMap["ipAddress"].(string); ok {
		o.IpAddress = &IpAddress
	}
    
	if IpOrganization, ok := SessionMap["ipOrganization"].(string); ok {
		o.IpOrganization = &IpOrganization
	}
    
	if LastPage, ok := SessionMap["lastPage"].(map[string]interface{}); ok {
		LastPageString, _ := json.Marshal(LastPage)
		json.Unmarshal(LastPageString, &o.LastPage)
	}
	
	if MktCampaign, ok := SessionMap["mktCampaign"].(map[string]interface{}); ok {
		MktCampaignString, _ := json.Marshal(MktCampaign)
		json.Unmarshal(MktCampaignString, &o.MktCampaign)
	}
	
	if Referrer, ok := SessionMap["referrer"].(map[string]interface{}); ok {
		ReferrerString, _ := json.Marshal(Referrer)
		json.Unmarshal(ReferrerString, &o.Referrer)
	}
	
	if App, ok := SessionMap["app"].(map[string]interface{}); ok {
		AppString, _ := json.Marshal(App)
		json.Unmarshal(AppString, &o.App)
	}
	
	if SdkLibrary, ok := SessionMap["sdkLibrary"].(map[string]interface{}); ok {
		SdkLibraryString, _ := json.Marshal(SdkLibrary)
		json.Unmarshal(SdkLibraryString, &o.SdkLibrary)
	}
	
	if NetworkConnectivity, ok := SessionMap["networkConnectivity"].(map[string]interface{}); ok {
		NetworkConnectivityString, _ := json.Marshal(NetworkConnectivity)
		json.Unmarshal(NetworkConnectivityString, &o.NetworkConnectivity)
	}
	
	if SearchTerms, ok := SessionMap["searchTerms"].([]interface{}); ok {
		SearchTermsString, _ := json.Marshal(SearchTerms)
		json.Unmarshal(SearchTermsString, &o.SearchTerms)
	}
	
	if UserAgentString, ok := SessionMap["userAgentString"].(string); ok {
		o.UserAgentString = &UserAgentString
	}
    
	if DurationInSeconds, ok := SessionMap["durationInSeconds"].(float64); ok {
		DurationInSecondsInt := int(DurationInSeconds)
		o.DurationInSeconds = &DurationInSecondsInt
	}
	
	if EventCount, ok := SessionMap["eventCount"].(float64); ok {
		EventCountInt := int(EventCount)
		o.EventCount = &EventCountInt
	}
	
	if PageviewCount, ok := SessionMap["pageviewCount"].(float64); ok {
		PageviewCountInt := int(PageviewCount)
		o.PageviewCount = &PageviewCountInt
	}
	
	if ScreenviewCount, ok := SessionMap["screenviewCount"].(float64); ok {
		ScreenviewCountInt := int(ScreenviewCount)
		o.ScreenviewCount = &ScreenviewCountInt
	}
	
	if LastEvent, ok := SessionMap["lastEvent"].(map[string]interface{}); ok {
		LastEventString, _ := json.Marshal(LastEvent)
		json.Unmarshal(LastEventString, &o.LastEvent)
	}
	
	if LastConnectedQueue, ok := SessionMap["lastConnectedQueue"].(map[string]interface{}); ok {
		LastConnectedQueueString, _ := json.Marshal(LastConnectedQueue)
		json.Unmarshal(LastConnectedQueueString, &o.LastConnectedQueue)
	}
	
	if LastConnectedUser, ok := SessionMap["lastConnectedUser"].(map[string]interface{}); ok {
		LastConnectedUserString, _ := json.Marshal(LastConnectedUser)
		json.Unmarshal(LastConnectedUserString, &o.LastConnectedUser)
	}
	
	if LastUserDisposition, ok := SessionMap["lastUserDisposition"].(map[string]interface{}); ok {
		LastUserDispositionString, _ := json.Marshal(LastUserDisposition)
		json.Unmarshal(LastUserDispositionString, &o.LastUserDisposition)
	}
	
	if ConversationChannels, ok := SessionMap["conversationChannels"].([]interface{}); ok {
		ConversationChannelsString, _ := json.Marshal(ConversationChannels)
		json.Unmarshal(ConversationChannelsString, &o.ConversationChannels)
	}
	
	if OriginatingDirection, ok := SessionMap["originatingDirection"].(string); ok {
		o.OriginatingDirection = &OriginatingDirection
	}
    
	if ConversationSubject, ok := SessionMap["conversationSubject"].(string); ok {
		o.ConversationSubject = &ConversationSubject
	}
    
	if LastUserDisconnectType, ok := SessionMap["lastUserDisconnectType"].(string); ok {
		o.LastUserDisconnectType = &LastUserDisconnectType
	}
    
	if LastAcdOutcome, ok := SessionMap["lastAcdOutcome"].(string); ok {
		o.LastAcdOutcome = &LastAcdOutcome
	}
    
	if Authenticated, ok := SessionMap["authenticated"].(bool); ok {
		o.Authenticated = &Authenticated
	}
    
	if LastScreen, ok := SessionMap["lastScreen"].(string); ok {
		o.LastScreen = &LastScreen
	}
    
	if SelfUri, ok := SessionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if createdDateString, ok := SessionMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if endedDateString, ok := SessionMap["endedDate"].(string); ok {
		EndedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endedDateString)
		o.EndedDate = &EndedDate
	}
	
	if ExternalContact, ok := SessionMap["externalContact"].(map[string]interface{}); ok {
		ExternalContactString, _ := json.Marshal(ExternalContact)
		json.Unmarshal(ExternalContactString, &o.ExternalContact)
	}
	
	if awayDateString, ok := SessionMap["awayDate"].(string); ok {
		AwayDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", awayDateString)
		o.AwayDate = &AwayDate
	}
	
	if idleDateString, ok := SessionMap["idleDate"].(string); ok {
		IdleDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", idleDateString)
		o.IdleDate = &IdleDate
	}
	
	if Conversation, ok := SessionMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Session) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
