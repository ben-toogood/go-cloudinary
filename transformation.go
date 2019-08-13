package cloudinary

import (
	"fmt"
	"strings"
)

// Transformation is a type of image transformation which is supported by cloudinary.
// docs: https://cloudinary.com/documentation/image_transformation_reference
type Transformation interface {
	URLEncoded() string
}

// SizeTransformation adjusts the size of an image
type SizeTransformation struct {
	Width  int
	Height int
	// TODO: Add multiple resize options
}

// URLEncoded returns the URL path component required for transformation
func (t SizeTransformation) URLEncoded() string {
	components := []string{"c_fit"}

	if t.Width > 0 {
		components = append(components, fmt.Sprintf("w_%v", t.Width))
	}

	if t.Height > 0 {
		components = append(components, fmt.Sprintf("h_%v", t.Height))
	}

	return strings.Join(components, ",")
}

// TransformedImageURL returns the URL for an image, having applied the transformationsfs
func (s *Service) TransformedImageURL(publicID string, ts ...Transformation) string {
	segments := make([]string, len(ts))
	for i, t := range ts {
		segments[i] = t.URLEncoded()
	}

	joinedSegments := strings.Join(segments, "/")
	return fmt.Sprintf("%s/%s/%s/upload/%s/%s", baseResourceUrl, s.cloudName, imageType, joinedSegments, publicID)
}
