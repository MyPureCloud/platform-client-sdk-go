package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Patchcontentoffer) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
