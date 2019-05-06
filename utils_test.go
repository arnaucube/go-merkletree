package merkletree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSetBitmap(t *testing.T) {

	var v Hash
	setbitmap(v[:], 7)
	setbitmap(v[:], 8)
	setbitmap(v[:], 255)
	expected := "0x8000000000000000000000000000000000000000000000000000000000000180"
	assert.Equal(t, expected, v.Hex())

	assert.Equal(t, false, testbitmap(v[:], 6))
	assert.Equal(t, true, testbitmap(v[:], 7))
	assert.Equal(t, true, testbitmap(v[:], 8))
	assert.Equal(t, false, testbitmap(v[:], 9))
	assert.Equal(t, true, testbitmap(v[:], 255))

}

func TestHashBytes(t *testing.T) {
	h := HashBytes([]byte("test")).Hex()
	assert.Equal(t, "0x9c22ff5f21f0b81b113e63f7db6da94fedef11b2119b4088b89664fb9a3cb658", h)

	h = HashBytes([]byte("authorizeksign")).Hex()
	assert.Equal(t, "0x353f867ef725411de05e3d4b0a01c37cf7ad24bcc213141a05ed7726d7932a1f", h)
}

func TestUint32ToBytes(t *testing.T) {
	b := Uint32ToBytes(999)
	assert.Equal(t, []byte{0xe7, 0x3, 0x0, 0x0}, b)
}
func TestBytesToUint32(t *testing.T) {
	u := BytesToUint32([]byte{0xe7, 0x3, 0x0, 0x0})
	assert.Equal(t, uint32(999), u)
}
