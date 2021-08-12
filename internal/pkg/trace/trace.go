package trace

import (
	"context"
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

type (
	traceIDCtx   struct{}
)

var incrNum uint64

type Trace struct {

}

func (*Trace) NewId() string {
	return fmt.Sprintf("trace-id-%d-%s-%d",
		os.Getpid(),
		time.Now().Format("2006.01.02.15.04.05.999"),
		atomic.AddUint64(&incrNum, 1))
}

// NewTraceID 创建追踪ID的上下文
func (*Trace) Context(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDCtx{}, traceID)
}

