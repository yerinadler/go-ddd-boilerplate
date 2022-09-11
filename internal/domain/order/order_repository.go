package order

import (
	"github.com/yerinadler/go-ddd/pkg/core"
)

type OrderRepository interface {
	core.Repository[Order]
}
