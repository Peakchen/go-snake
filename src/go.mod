module go-snake

go 1.15

require (
	github.com/Peakchen/xgameCommon v0.0.6-0.20210310091059-6c4643710177
	github.com/Shopify/sarama v1.19.0 // indirect
	github.com/Shopify/toxiproxy v2.1.4+incompatible // indirect
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/axgle/mahonia v0.0.0-20180208002826-3358181d7394 // indirect
	github.com/bwmarrin/snowflake v0.3.0 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/davyxu/ulog v1.0.1
	github.com/eapache/go-resiliency v1.1.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.1 // indirect
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db // indirect
	github.com/gomodule/redigo v1.8.3
	github.com/gonutz/ide v0.0.0-20200517034207-df64a3832118 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.14.6 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/nats-io/nats-server/v2 v2.2.0 // indirect
	github.com/nats-io/nats.go v1.10.1-0.20210228004050-ed743748acac
	github.com/panjf2000/ants/v2 v2.4.3
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a // indirect
	github.com/rs/xid v1.2.1 // indirect
	github.com/samuel/go-zookeeper v0.0.0-20200724154423-2164a8ac840e
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/segmentio/ksuid v1.0.3 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/sony/sonyflake v1.0.0 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	go.etcd.io/etcd v0.5.0-alpha.5.0.20210123184945-d51c6c689ba3 //v1.2.3
	google.golang.org/appengine v1.6.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.26.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/ini.v1 v1.62.0
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.8
)

replace github.com/coreos/bbolt v1.1.0 => go.etcd.io/bbolt v1.1.0 //v1.3.4

replace go.etcd.io/bbolt v1.1.0 => github.com/coreos/bbolt v1.1.0

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

replace google.golang.org/grpc v1.27.0 => google.golang.org/grpc v1.27.0
