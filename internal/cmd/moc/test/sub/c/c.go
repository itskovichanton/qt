package c

import (
	"github.com/itskovichanton/qt/core"

	_ "github.com/itskovichanton/qt/internal/cmd/moc/test/sub/conf"
)

type StructSubGoC struct{}
type StructSubMocC struct{ core.QObject }
