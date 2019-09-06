module github.com/BCCE/mattermost-plugin-channel-guard

go 1.12

require (
	github.com/go-ldap/ldap v3.0.3+incompatible // indirect
	github.com/go-redis/redis v6.15.5+incompatible // indirect
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/hashicorp/go-hclog v0.9.2 // indirect
	github.com/hashicorp/go-plugin v1.0.1 // indirect
	github.com/lib/pq v1.2.0 // indirect
	github.com/mattermost/mattermost-server v5.12.0+incompatible
	github.com/mattn/go-sqlite3 v1.11.0 // indirect
	github.com/onsi/ginkgo v1.10.1 // indirect
	github.com/onsi/gomega v1.7.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.4.0 // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
	golang.org/x/crypto v0.0.0-20190829043050-9756ffdc2472 // indirect
	google.golang.org/appengine v1.6.2 // indirect
)

// Workaround for https://github.com/golang/go/issues/30831 and fallout.
replace github.com/golang/lint => github.com/golang/lint v0.0.0-20190227174305-8f45f776aaf1

// Override since git.apache.org is down.
replace git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999
