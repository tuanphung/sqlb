package sqlb

type PlaceholderType int

const (
	Question PlaceholderType = iota
	Dollar
	Unknown
)

var placeholder = Question

func SetPlaceholder(newPlaceholder PlaceholderType) {
	if newPlaceholder >= Unknown {
		// Default is `?``
		placeholder = Question
		return
	}

	placeholder = newPlaceholder
}

func GetPlaceholder() PlaceholderType {
	return placeholder
}
