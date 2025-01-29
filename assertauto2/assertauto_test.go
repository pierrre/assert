package assertauto2

import (
	"testing"

	"github.com/pierrre/pretty"
)

func Test(t *testing.T) {
	Equal(t, "aaaaaa")
	Equal(t, 123456)
	Equal(t, 123.456)
	Equal(t, pretty.DefaultCommonValueWriter)
}
