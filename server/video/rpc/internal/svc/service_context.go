package svc

import (
	"douniu/server/common/consts"
	"douniu/server/common/init_db"
	"douniu/server/user/rpc/userrpc"
	"douniu/server/video/model"
	"douniu/server/video/rpc/internal/config"
	"github.com/bwmarrin/snowflake"
	"github.com/olivere/elastic/v7"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	Snowflake      *snowflake.Node
	SqlConn        sqlx.SqlConn
	RedisClient    *redis.Client
	VideoModel     model.VideoModel
	KqPusherClient *kq.Pusher
	UserRpc        userrpc.UserRpc
	ESClient       *elastic.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MysqlConf.DataSource)
	snowflakeNode, _ := snowflake.NewNode(consts.VideoMachineId)
	return &ServiceContext{
		Config:         c,
		SqlConn:        mysqlConn,
		VideoModel:     model.NewVideoModel(mysqlConn, c.CacheRedis),
		Snowflake:      snowflakeNode,
		RedisClient:    init_db.InitRedis(c.RedisConf.Host, c.RedisConf.Password),
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
		//UserRpc:  userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf)),
		ESClient: init_db.GetESClient(c.ESConf.Host),
	}
}
