package platformclientv2
import (
	"time"
	"encoding/json"
)

// Documentationresult
type Documentationresult struct { 
	// Id - The globally unique identifier for the object.
	Id *int32 `json:"id,omitempty"`


	// Categories - The category of the documentation entity. Will be returned in responses for certain entities.
	Categories *[]int32 `json:"categories,omitempty"`


	// Description - The description of the documentation entity. Will be returned in responses for certain entities.
	Description *string `json:"description,omitempty"`


	// Content - The text or html content for the documentation entity. Will be returned in responses for certain entities.
	Content *string `json:"content,omitempty"`


	// Excerpt - The excerpt of the documentation entity. Will be returned in responses for certain entities.
	Excerpt *string `json:"excerpt,omitempty"`


	// Link - URL link for the documentation entity. Will be returned in responses for certain entities.
	Link *string `json:"link,omitempty"`


	// Modified - The modified date for the documentation entity. Will be returned in responses for certain entities. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	Modified *time.Time `json:"modified,omitempty"`


	// Name - The name of the documentation entity. Will be returned in responses for certain entities.
	Name *string `json:"name,omitempty"`


	// Service - The service of the documentation entity. Will be returned in responses for certain entities.
	Service *[]int32 `json:"service,omitempty"`


	// Slug - The slug of the documentation entity. Will be returned in responses for certain entities.
	Slug *string `json:"slug,omitempty"`


	// Title - The title of the documentation entity. Will be returned in responses for certain entities.
	Title *string `json:"title,omitempty"`


	// GetType - The search type. Will be returned in responses for certain entities.
	GetType *string `json:"get_type,omitempty"`


	// FacetFeature - The facet feature of the documentation entity. Will be returned in responses for certain entities.
	FacetFeature *[]int32 `json:"facet_feature,omitempty"`


	// FacetRole - The facet role of the documentation entity. Will be returned in responses for certain entities.
	FacetRole *[]int32 `json:"facet_role,omitempty"`


	// FacetService - The facet service of the documentation entity. Will be returned in responses for certain entities.
	FacetService *[]int32 `json:"facet_service,omitempty"`


	// FaqCategories - The faq categories of the documentation entity. Will be returned in responses for certain entities.
	FaqCategories *[]int32 `json:"faq_categories,omitempty"`


	// ReleasenoteCategory - The releasenote category of the documentation entity. Will be returned in responses for certain entities.
	ReleasenoteCategory *[]int32 `json:"releasenote_category,omitempty"`


	// ReleasenoteTag - The releasenote tag of the documentation entity. Will be returned in responses for certain entities.
	ReleasenoteTag *[]int32 `json:"releasenote_tag,omitempty"`


	// ServiceArea - The service area of the documentation entity. Will be returned in responses for certain entities.
	ServiceArea *[]int32 `json:"service-area,omitempty"`


	// VideoCategories - The video categories of the documentation entity. Will be returned in responses for certain entities.
	VideoCategories *[]int32 `json:"video_categories,omitempty"`

}

// String returns a JSON representation of the model
func (o *Documentationresult) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
