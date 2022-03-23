module github.com/aliyun/aliyun-cli

go 1.16

replace github.com/alibabacloud-go/tea => ./local/tea

replace github.com/aliyun/credentials-go => ./local/credentials-go

replace github.com/syndtr/goleveldb => ./local/goleveldb

replace github.com/onsi/ginkgo => ./local/ginkgo

replace github.com/onsi/ginkgo/v2 => ./local/ginkgo_v2

replace github.com/smartystreets/goconvey => ./local/goconvey

replace golang.org/x/mod => ./local/mod

require (
	github.com/alibabacloud-go/tea v1.1.17
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1529
	github.com/aliyun/aliyun-oss-go-sdk v2.2.1+incompatible
	github.com/aliyun/credentials-go v1.2.1
	github.com/alyu/configparser v0.0.0-20191103060215-744e9a66e7bc
	github.com/droundy/goopt v0.0.0-20170604162106-0b8effe182da
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.18.1
	github.com/stretchr/testify v1.7.0
	github.com/syndtr/goleveldb v1.0.0
	golang.org/x/crypto v0.0.0-20220321153916-2c7772ba3064
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c
	gopkg.in/ini.v1 v1.66.2
)

require (
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b // indirect
	golang.org/x/time v0.0.0-20220224211638-0e9765cccd65 // indirect
)
