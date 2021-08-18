package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchcontentoffer
type Patchcontentoffer struct { 
	// ImageUrl - URL for image displayed to the customer when displaying content offer.
	ImageUrl *string `json:"imageUrl,omitempty"`


	// DisplayMode - The display mode of Genesys Widgets when displaying content offer.
	DisplayMode *string `json:"displayMode,omitempty"`


	// LayoutMode - The layout mode of the text shown to the user when displaying content offer.
	LayoutMode *string `json:"layoutMode,omitempty"`


	// Title - Title used in the header of the content offer.
	Title *string `json:"title,omitempty"`


	// Headline - Headline displayed above the body text of the content offer.
	Headline *string `json:"headline,omitempty"`


	// Body - Body text of the content offer.
	Body *string `json:"body,omitempty"`


	// CallToAction - Properties customizing the call to action button on the content offer.
	CallToAction *Patchcalltoaction `json:"callToAction,omitempty"`


	// Style - Properties customizing the styling of the content offer.
	Style *Patchcontentofferstylingconfiguration `json:"style,omitempty"`

}

func (u *Patchcontentoffer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchcontentoffer

	

	return json.Marshal(&struct { 
		ImageUrl *string `json:"imageUrl,omitempty"`
		
		DisplayMode *string `json:"displayMode,omitempty"`
		
		LayoutMode *string `json:"layoutMode,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Headline *string `json:"headline,omitempty"`
		
		Body *string `json:"body,omitempty"`
		
		CallToAction *Patchcalltoaction `json:"callToAction,omitempty"`
		
		Style *Patchcontentofferstylingconfiguration `json:"style,omitempty"`
		*Alias
	}{ 
		ImageUrl: u.ImageUrl,
		
		DisplayMode: u.DisplayMode,
		
		LayoutMode: u.LayoutMode,
		
		Title: u.Title,
		
		Headline: u.Headline,
		
		Body: u.Body,
		
		CallToAction: u.CallToAction,
		
		Style: u.Style,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Patchcontentoffer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
