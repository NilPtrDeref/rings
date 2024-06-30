package rings

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

const LogEntryKey = "log-entry"

func (r *Ring) SetLog(request *http.Request, entry *logrus.Entry) {
	ctx := request.Context()
	ctx = context.WithValue(ctx, LogEntryKey, entry)
	*request = *request.WithContext(ctx)
}

func (r *Ring) GetLog(request *http.Request) *logrus.Entry {
	ctx := request.Context()
	value := ctx.Value(LogEntryKey)

	entry, ok := value.(*logrus.Entry)
	if ok {
		return entry
	}

	entry = logrus.NewEntry(r.Logger)
	r.SetLog(request, entry)
	return entry
}

func (r *Ring) UpdateLog(request *http.Request, key string, value any) {
	entry := r.GetLog(request)
	entry = entry.WithField(key, value)
	r.SetLog(request, entry)
}
