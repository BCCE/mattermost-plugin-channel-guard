module github.com/BCCE/mattermost-plugin-channel-guard

go 1.12

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/Masterminds/squirrel v1.1.0 // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/dyatlov/go-opengraph v0.0.0-20180429202543-816b6608b3c8 // indirect
	github.com/go-gorp/gorp v2.0.0+incompatible // indirect
	github.com/go-ldap/ldap v3.0.3+incompatible // indirect
	github.com/go-redis/redis v6.15.5+incompatible // indirect
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/hashicorp/go-hclog v0.9.2 // indirect
	github.com/hashicorp/go-plugin v1.0.1 // indirect
	github.com/lib/pq v1.2.0 // indirect
	github.com/mattermost/gorp v2.0.0+incompatible // indirect
	github.com/mattermost/mattermost-server v5.11.1+incompatible
	github.com/mattn/go-sqlite3 v1.11.0 // indirect
	github.com/onsi/ginkgo v1.10.1 // indirect
	github.com/onsi/gomega v1.7.0 // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/ziutek/mymysql v1.5.4 // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
	golang.org/x/crypto v0.0.0-20190829043050-9756ffdc2472 // indirect
	google.golang.org/appengine v1.6.2 // indirect
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

// Workaround for https://github.com/golang/go/issues/30831 and fallout.
replace github.com/golang/lint => github.com/golang/lint v0.0.0-20190227174305-8f45f776aaf1

// Override since git.apache.org is down.
replace git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999
