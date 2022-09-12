package validations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringOrChar(t *testing.T) {
	m := "  hello world  "
	assert := assert.New(t)
	sanitized := Sanitize(m)
	assert.Equal("hello world", sanitized, "The sanitized message should be equal to the expected message")
}
