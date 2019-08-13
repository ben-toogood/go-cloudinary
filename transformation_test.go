package cloudinary

import (
	"fmt"
	"testing"
)

func TestTransformImage(t *testing.T) {
	srv := Service{cloudName: "demo"}
	publicID := "yellow_tulip.jpg"

	tt := []struct {
		name            string
		url             string
		transformations []Transformation
	}{
		{
			name: "No transformations",
			url:  fmt.Sprintf("%v/%v/image/upload//%v", baseResourceUrl, srv.CloudName(), publicID),
		},
		{
			name:            "Height transformation",
			transformations: []Transformation{SizeTransformation{Height: 100}},
			url:             fmt.Sprintf("%v/%v/image/upload/h_100/%v", baseResourceUrl, srv.CloudName(), publicID),
		},
		{
			name:            "Width transformation",
			transformations: []Transformation{SizeTransformation{Width: 100}},
			url:             fmt.Sprintf("%v/%v/image/upload/w_100/%v", baseResourceUrl, srv.CloudName(), publicID),
		},
		{
			name:            "Width and height transformations",
			transformations: []Transformation{SizeTransformation{Height: 100, Width: 100}},
			url:             fmt.Sprintf("%v/%v/image/upload/w_100,h_100/%v", baseResourceUrl, srv.CloudName(), publicID),
		},
		{
			name:            "Width and invalid height transformations",
			transformations: []Transformation{SizeTransformation{Height: -100, Width: 100}},
			url:             fmt.Sprintf("%v/%v/image/upload/w_100/%v", baseResourceUrl, srv.CloudName(), publicID),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			url := srv.TransformedImageURL(publicID, tc.transformations...)

			if url != tc.url {
				t.Errorf("Expected to get %v but actually got %v", tc.url, url)
			}
		})
	}
}
