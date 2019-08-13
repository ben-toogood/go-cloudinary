package cloudinary

type Transformation interface {
	UrlEncoded() string
}

func (s *Service) TransformImage(publicId string, ts ...Transformation) string {
	return ""
}
