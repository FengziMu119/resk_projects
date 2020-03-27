package resk

import (
	"resk_projects/apis/gorpc"
	_ "resk_projects/apis/web"
	_ "resk_projects/core/accounts"
	_ "resk_projects/core/envelopes"
	"resk_projects/infra"
	"resk_projects/infra/base"
	"resk_projects/jobs"
	_ "resk_projects/public/ui"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.GoRPCStarter{})
	infra.Register(&gorpc.GoRpcApiStarter{})
	infra.Register(&jobs.RefundExpiredJobStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiStarter{})
	infra.Register(&base.HookStarter{})
}
