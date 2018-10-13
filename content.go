package contents

import (
	"net/url"
	"time"
)

type Content interface {
	Type() string
	Title() string
	Date() time.Time
	Outline() string
	ImageURL() url.URL
	URL() url.URL
}
