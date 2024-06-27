package entity

import (
	"testing"
	"github.com/stretch/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Product 1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.NotEqual(t, "Product 1", p.Name)
	assert.Equal(t, 10, p.Price)
}
