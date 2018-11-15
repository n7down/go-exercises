package jwtapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIifQ.UIZchxQD36xuhacrJF9HQ5SIUxH5HBiv9noESAacsxU"
	claims, err := Parse(tokenString)
	if err != nil {
		assert.Fail(t, "%s", err)
	}
	assert.Equal(t, claims["foo"], "bar")
}

func TestParseClaims(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIifQ.UIZchxQD36xuhacrJF9HQ5SIUxH5HBiv9noESAacsxU"
	claims, err := ParseClaims(tokenString)
	if err != nil {
		assert.Fail(t, "%s", err)
	}
	assert.Equal(t, claims.Foo, "bar")
}
