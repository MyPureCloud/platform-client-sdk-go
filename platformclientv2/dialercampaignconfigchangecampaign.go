package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercampaignconfigchangecampaign
type Dialercampaignconfigchangecampaign struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`


	// ContactList
	ContactList *Dialercampaignconfigchangeurireference `json:"contactList,omitempty"`


	// Queue
	Queue *Dialercampaignconfigchangeurireference `json:"queue,omitempty"`


	// DialingMode
	DialingMode *string `json:"dialingMode,omitempty"`


	// Script
	Script *Dialercampaignconfigchangeurireference `json:"script,omitempty"`


	// EdgeGroup
	EdgeGroup *Dialercampaignconfigchangeurireference `json:"edgeGroup,omitempty"`


	// Site
	Site *Dialercampaignconfigchangeurireference `json:"site,omitempty"`


	// CampaignStatus
	CampaignStatus *string `json:"campaignStatus,omitempty"`


	// PhoneColumns
	PhoneColumns *[]Dialercampaignconfigchangephonecolumn `json:"phoneColumns,omitempty"`


	// AbandonRate
	AbandonRate *float32 `json:"abandonRate,omitempty"`


	// DncLists
	DncLists *[]Dialercampaignconfigchangeurireference `json:"dncLists,omitempty"`


	// CallableTimeSet
	CallableTimeSet *Dialercampaignconfigchangeurireference `json:"callableTimeSet,omitempty"`


	// CallAnalysisResponseSet
	CallAnalysisResponseSet *Dialercampaignconfigchangeurireference `json:"callAnalysisResponseSet,omitempty"`


	// CallerName
	CallerName *string `json:"callerName,omitempty"`


	// CallerAddress
	CallerAddress *string `json:"callerAddress,omitempty"`


	// OutboundLineCount
	OutboundLineCount *int `json:"outboundLineCount,omitempty"`


	// Errors
	Errors *[]Dialercampaignconfigchangeresterrordetail `json:"errors,omitempty"`


	// RuleSets
	RuleSets *[]Dialercampaignconfigchangeurireference `json:"ruleSets,omitempty"`


	// SkipPreviewDisabled
	SkipPreviewDisabled *bool `json:"skipPreviewDisabled,omitempty"`


	// PreviewTimeOutSeconds
	PreviewTimeOutSeconds *int `json:"previewTimeOutSeconds,omitempty"`


	// SingleNumberPreview
	SingleNumberPreview *bool `json:"singleNumberPreview,omitempty"`


	// ContactSort
	ContactSort *Dialercampaignconfigchangecontactsort `json:"contactSort,omitempty"`


	// ContactSorts
	ContactSorts *[]Dialercampaignconfigchangecontactsort `json:"contactSorts,omitempty"`


	// NoAnswerTimeout
	NoAnswerTimeout *int `json:"noAnswerTimeout,omitempty"`


	// CallAnalysisLanguage
	CallAnalysisLanguage *string `json:"callAnalysisLanguage,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// ContactListFilters
	ContactListFilters *[]Dialercampaignconfigchangeurireference `json:"contactListFilters,omitempty"`


	// Division
	Division *Dialercampaignconfigchangeurireference `json:"division,omitempty"`


	// AgentOwnedColumn
	AgentOwnedColumn *string `json:"agentOwnedColumn,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercampaignconfigchangecampaign) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
