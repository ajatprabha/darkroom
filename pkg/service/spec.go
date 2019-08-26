package service

type Spec struct {
	// Scope defines a scope for the image manipulation job, it can be used for logging/mertrics collection purposes
	Scope string
	// ImageData holds the actual image contents to processed
	ImageData []byte
	// Params hold the key-value pairs for the processing job and tells the manipulator what to do with the image
	Params map[string]string
	// Formats have the information of accepted formats, whether darkroom can return the image using webp or not
	Formats []string
}

type SpecBuilder interface {
	WithScope(scope string) SpecBuilder
	Build() Spec
}

type specBuilder struct {
	scope     string
	imageData []byte
	params    map[string]string
	formats   []string
}

func (sb *specBuilder) WithScope(scope string) SpecBuilder {
	sb.scope = scope
	return sb
}

func (sb *specBuilder) Build() Spec {
	return Spec{
		Scope: sb.scope,
	}
}

func NewSpecBuilder() SpecBuilder {
	return &specBuilder{}
}
