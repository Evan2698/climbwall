package security

import (
	"testing"
)

func Test_KIL(t *testing.T) {

	v := NewEncryptyMethod("raw")
	b := ToBytes(v)

	t.Log(b)

	c, e := FromByte(b)

	t.Log(c.GetName(), c.GetIV())
	t.Log(e)
	t.Log(ToBytes(c))
}
