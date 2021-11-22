package logging

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
)

func TestLog(t *testing.T) {
	Convey("Testing DB connection", t, func() {
		// log.Info("aaaaa")

		logger := zap.NewExample()
		defer logger.Sync()

		// // zapcore.Field
		// logger.With(
		// 	zapcore.Field("bbbb"),
		// 	"aaaa".(zapcore.Field),
		// ).Info("tracked some metrics")
	})
}
