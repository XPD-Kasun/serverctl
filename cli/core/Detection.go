package core

type Detector[T DetectionContext, R DetectionResult] interface {
	Detect(context *T) *R
}

type DetectionType string

const (
	Installation  DetectionType = "Installation"
	File          DetectionType = "File"
	Remote        DetectionType = "Remote"
	Configuration DetectionType = "Configuration"
)

type DetectionContext interface {
	DetectionType() DetectionType
}

type DetectionResult interface {
	DetectionType() DetectionType
	GetResultData() map[string][]string
	foundTarget() bool
}
