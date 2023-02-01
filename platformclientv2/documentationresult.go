package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentationresult
type Documentationresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *int `json:"id,omitempty"`

	// Categories - The category of the documentation entity. Will be returned in responses for certain entities.
	Categories *[]int `json:"categories,omitempty"`

	// Description - The description of the documentation entity. Will be returned in responses for certain entities.
	Description *string `json:"description,omitempty"`

	// Content - The text or html content for the documentation entity. Will be returned in responses for certain entities.
	Content *string `json:"content,omitempty"`

	// Excerpt - The excerpt of the documentation entity. Will be returned in responses for certain entities.
	Excerpt *string `json:"excerpt,omitempty"`

	// Link - URL link for the documentation entity. Will be returned in responses for certain entities.
	Link *string `json:"link,omitempty"`

	// Modified - The modified date for the documentation entity. Will be returned in responses for certain entities. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Modified *time.Time `json:"modified,omitempty"`

	// Name - The name of the documentation entity. Will be returned in responses for certain entities.
	Name *string `json:"name,omitempty"`

	// Service - The service of the documentation entity. Will be returned in responses for certain entities.
	Service *[]int `json:"service,omitempty"`

	// Slug - The slug of the documentation entity. Will be returned in responses for certain entities.
	Slug *string `json:"slug,omitempty"`

	// Title - The title of the documentation entity. Will be returned in responses for certain entities.
	Title *string `json:"title,omitempty"`

	// GetType - The search type. Will be returned in responses for certain entities.
	GetType *string `json:"get_type,omitempty"`

	// FacetFeature - The facet feature of the documentation entity. Will be returned in responses for certain entities.
	FacetFeature *[]int `json:"facet_feature,omitempty"`

	// FacetRole - The facet role of the documentation entity. Will be returned in responses for certain entities.
	FacetRole *[]int `json:"facet_role,omitempty"`

	// FacetService - The facet service of the documentation entity. Will be returned in responses for certain entities.
	FacetService *[]int `json:"facet_service,omitempty"`

	// FaqCategories - The faq categories of the documentation entity. Will be returned in responses for certain entities.
	FaqCategories *[]int `json:"faq_categories,omitempty"`

	// ReleasenoteCategory - The releasenote category of the documentation entity. Will be returned in responses for certain entities.
	ReleasenoteCategory *[]int `json:"releasenote_category,omitempty"`

	// ReleasenoteTag - The releasenote tag of the documentation entity. Will be returned in responses for certain entities.
	ReleasenoteTag *[]int `json:"releasenote_tag,omitempty"`

	// ServiceArea - The service area of the documentation entity. Will be returned in responses for certain entities.
	ServiceArea *[]int `json:"service-area,omitempty"`

	// VideoCategories - The video categories of the documentation entity. Will be returned in responses for certain entities.
	VideoCategories *[]int `json:"video_categories,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Documentationresult) SetField(field string, fieldValue interface{}) {
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

func (o Documentationresult) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "Modified", }
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
	type Alias Documentationresult
	
	Modified := new(string)
	if o.Modified != nil {
		
		*Modified = timeutil.Strftime(o.Modified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Modified = nil
	}
	
	return json.Marshal(&struct { 
		Id *int `json:"id,omitempty"`
		
		Categories *[]int `json:"categories,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Content *string `json:"content,omitempty"`
		
		Excerpt *string `json:"excerpt,omitempty"`
		
		Link *string `json:"link,omitempty"`
		
		Modified *string `json:"modified,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Service *[]int `json:"service,omitempty"`
		
		Slug *string `json:"slug,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		GetType *string `json:"get_type,omitempty"`
		
		FacetFeature *[]int `json:"facet_feature,omitempty"`
		
		FacetRole *[]int `json:"facet_role,omitempty"`
		
		FacetService *[]int `json:"facet_service,omitempty"`
		
		FaqCategories *[]int `json:"faq_categories,omitempty"`
		
		ReleasenoteCategory *[]int `json:"releasenote_category,omitempty"`
		
		ReleasenoteTag *[]int `json:"releasenote_tag,omitempty"`
		
		ServiceArea *[]int `json:"service-area,omitempty"`
		
		VideoCategories *[]int `json:"video_categories,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Categories: o.Categories,
		
		Description: o.Description,
		
		Content: o.Content,
		
		Excerpt: o.Excerpt,
		
		Link: o.Link,
		
		Modified: Modified,
		
		Name: o.Name,
		
		Service: o.Service,
		
		Slug: o.Slug,
		
		Title: o.Title,
		
		GetType: o.GetType,
		
		FacetFeature: o.FacetFeature,
		
		FacetRole: o.FacetRole,
		
		FacetService: o.FacetService,
		
		FaqCategories: o.FaqCategories,
		
		ReleasenoteCategory: o.ReleasenoteCategory,
		
		ReleasenoteTag: o.ReleasenoteTag,
		
		ServiceArea: o.ServiceArea,
		
		VideoCategories: o.VideoCategories,
		Alias:    (Alias)(o),
	})
}

func (o *Documentationresult) UnmarshalJSON(b []byte) error {
	var DocumentationresultMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentationresultMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DocumentationresultMap["id"].(float64); ok {
		IdInt := int(Id)
		o.Id = &IdInt
	}
	
	if Categories, ok := DocumentationresultMap["categories"].([]interface{}); ok {
		CategoriesString, _ := json.Marshal(Categories)
		json.Unmarshal(CategoriesString, &o.Categories)
	}
	
	if Description, ok := DocumentationresultMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Content, ok := DocumentationresultMap["content"].(string); ok {
		o.Content = &Content
	}
    
	if Excerpt, ok := DocumentationresultMap["excerpt"].(string); ok {
		o.Excerpt = &Excerpt
	}
    
	if Link, ok := DocumentationresultMap["link"].(string); ok {
		o.Link = &Link
	}
    
	if modifiedString, ok := DocumentationresultMap["modified"].(string); ok {
		Modified, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedString)
		o.Modified = &Modified
	}
	
	if Name, ok := DocumentationresultMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Service, ok := DocumentationresultMap["service"].([]interface{}); ok {
		ServiceString, _ := json.Marshal(Service)
		json.Unmarshal(ServiceString, &o.Service)
	}
	
	if Slug, ok := DocumentationresultMap["slug"].(string); ok {
		o.Slug = &Slug
	}
    
	if Title, ok := DocumentationresultMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if GetType, ok := DocumentationresultMap["get_type"].(string); ok {
		o.GetType = &GetType
	}
    
	if FacetFeature, ok := DocumentationresultMap["facet_feature"].([]interface{}); ok {
		FacetFeatureString, _ := json.Marshal(FacetFeature)
		json.Unmarshal(FacetFeatureString, &o.FacetFeature)
	}
	
	if FacetRole, ok := DocumentationresultMap["facet_role"].([]interface{}); ok {
		FacetRoleString, _ := json.Marshal(FacetRole)
		json.Unmarshal(FacetRoleString, &o.FacetRole)
	}
	
	if FacetService, ok := DocumentationresultMap["facet_service"].([]interface{}); ok {
		FacetServiceString, _ := json.Marshal(FacetService)
		json.Unmarshal(FacetServiceString, &o.FacetService)
	}
	
	if FaqCategories, ok := DocumentationresultMap["faq_categories"].([]interface{}); ok {
		FaqCategoriesString, _ := json.Marshal(FaqCategories)
		json.Unmarshal(FaqCategoriesString, &o.FaqCategories)
	}
	
	if ReleasenoteCategory, ok := DocumentationresultMap["releasenote_category"].([]interface{}); ok {
		ReleasenoteCategoryString, _ := json.Marshal(ReleasenoteCategory)
		json.Unmarshal(ReleasenoteCategoryString, &o.ReleasenoteCategory)
	}
	
	if ReleasenoteTag, ok := DocumentationresultMap["releasenote_tag"].([]interface{}); ok {
		ReleasenoteTagString, _ := json.Marshal(ReleasenoteTag)
		json.Unmarshal(ReleasenoteTagString, &o.ReleasenoteTag)
	}
	
	if ServiceArea, ok := DocumentationresultMap["service-area"].([]interface{}); ok {
		ServiceAreaString, _ := json.Marshal(ServiceArea)
		json.Unmarshal(ServiceAreaString, &o.ServiceArea)
	}
	
	if VideoCategories, ok := DocumentationresultMap["video_categories"].([]interface{}); ok {
		VideoCategoriesString, _ := json.Marshal(VideoCategories)
		json.Unmarshal(VideoCategoriesString, &o.VideoCategories)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentationresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
