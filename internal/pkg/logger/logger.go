package logger

import (
	"context"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

type log = logrus.Logger

var (
	AddHook = logrus.AddHook
	SetReportCaller = logrus.SetReportCaller
	WithFields = logrus.WithFields
)

const (
	TraceIdKey = "trace_id"
	UserIdKey  = "user_id"
	TagKey     = "tag"
	VersionKey = "version"
	StackKey   = "stack"

	RequestBodyKey = "request_body_key"
	ResponseBodyKey = "response_body_key"

	PathKey = "path"
	IPKey = "ip"
	TimeConsumeKey = "time_consume"
)

type (
	traceIDKey struct{}
	userIDKey  struct{}
	tagKey     struct{}
	stackKey   struct{}
)

var (
	version string
)

// SetVersion 设定版本
func SetVersion(v string) {
	version = v
}


func init() {
	logrus.SetOutput(ioutil.Discard)
}

// NewTraceIdContext 创建跟踪ID上下文
func NewTraceIdContext(ctx context.Context, traceId string) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceId)
}

// FromTraceIdContext 从上下文中获取跟踪ID
func FromTraceIdContext(ctx context.Context) string {
	v := ctx.Value(traceIDKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// NewTagContext 创建Tag上下文
func NewTagContext(ctx context.Context, tag string) context.Context {
	return context.WithValue(ctx, tagKey{}, tag)
}

// FromTagContext 从上下文中获取Tag
func FromTagContext(ctx context.Context) string {
	v := ctx.Value(tagKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func WithContext(ctx context.Context) *logrus.Entry {
	if ctx == nil {
		ctx = context.Background()
	}

	fields := map[string]interface{}{
		VersionKey: version,
	}

	if v := FromTraceIdContext(ctx); v != "" {
		fields[TraceIdKey] = v
	}

	//if v := FromUserIDContext(ctx); v != "" {
	//	fields[UserIDKey] = v
	//}
	//
	if v := FromTagContext(ctx); v != "" {
		fields[TagKey] = v
	}
	//
	//if v := FromStackContext(ctx); v != nil {
	//	fields[StackKey] = fmt.Sprintf("%+v", v)
	//}




	return logrus.WithContext(ctx).WithFields(fields)
}