package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeypage
type Journeypage struct { 
	// Url - The page URL.
	Url *string `json:"url,omitempty"`


	// Title - Title of the page.
	Title *string `json:"title,omitempty"`


	// Domain - Domain of the page's URL.
	Domain *string `json:"domain,omitempty"`


	// Fragment - Fragment or hash of the page's URL.
	Fragment *string `json:"fragment,omitempty"`


	// Hostname - Hostname of the page's URL.
	Hostname *string `json:"hostname,omitempty"`


	// Keywords - Keywords from the HTML {@code <meta>} tag of the page.
	Keywords *string `json:"keywords,omitempty"`


	// Lang - ISO 639-1 language code for the page as defined in the {@code <html>} tag.
	Lang *string `json:"lang,omitempty"`


	// Pathname - Path name of the page for the event.
	Pathname *string `json:"pathname,omitempty"`


	// QueryString - Query string that is passed to the page in the current event.
	QueryString *string `json:"queryString,omitempty"`


	// Breadcrumb - Hierarchy of the current page in relation to the website's structure.
	Breadcrumb *[]string `json:"breadcrumb,omitempty"`

}

func (o *Journeypage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeypage
	
	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Domain *string `json:"domain,omitempty"`
		
		Fragment *string `json:"fragment,omitempty"`
		
		Hostname *string `json:"hostname,omitempty"`
		
		Keywords *string `json:"keywords,omitempty"`
		
		Lang *string `json:"lang,omitempty"`
		
		Pathname *string `json:"pathname,omitempty"`
		
		QueryString *string `json:"queryString,omitempty"`
		
		Breadcrumb *[]string `json:"breadcrumb,omitempty"`
		*Alias
	}{ 
		Url: o.Url,
		
		Title: o.Title,
		
		Domain: o.Domain,
		
		Fragment: o.Fragment,
		
		Hostname: o.Hostname,
		
		Keywords: o.Keywords,
		
		Lang: o.Lang,
		
		Pathname: o.Pathname,
		
		QueryString: o.QueryString,
		
		Breadcrumb: o.Breadcrumb,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeypage) UnmarshalJSON(b []byte) error {
	var JourneypageMap map[string]interface{}
	err := json.Unmarshal(b, &JourneypageMap)
	if err != nil {
		return err
	}
	
	if Url, ok := JourneypageMap["url"].(string); ok {
		o.Url = &Url
	}
	
	if Title, ok := JourneypageMap["title"].(string); ok {
		o.Title = &Title
	}
	
	if Domain, ok := JourneypageMap["domain"].(string); ok {
		o.Domain = &Domain
	}
	
	if Fragment, ok := JourneypageMap["fragment"].(string); ok {
		o.Fragment = &Fragment
	}
	
	if Hostname, ok := JourneypageMap["hostname"].(string); ok {
		o.Hostname = &Hostname
	}
	
	if Keywords, ok := JourneypageMap["keywords"].(string); ok {
		o.Keywords = &Keywords
	}
	
	if Lang, ok := JourneypageMap["lang"].(string); ok {
		o.Lang = &Lang
	}
	
	if Pathname, ok := JourneypageMap["pathname"].(string); ok {
		o.Pathname = &Pathname
	}
	
	if QueryString, ok := JourneypageMap["queryString"].(string); ok {
		o.QueryString = &QueryString
	}
	
	if Breadcrumb, ok := JourneypageMap["breadcrumb"].([]interface{}); ok {
		BreadcrumbString, _ := json.Marshal(Breadcrumb)
		json.Unmarshal(BreadcrumbString, &o.Breadcrumb)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeypage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
