package bucket

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWait(t *testing.T) {
	Convey("新建一个 Bucket 对象", t, func() {
		b := New(1, 1)
		Convey("利用已经取消的 ctx 去 wait", func() {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			err := b.Wait(ctx)
			Convey("应该报错", func() {
				So(err.Error(), ShouldStartWith, "wait error:")
			})
		})
	})
}
