module github.com/onsi/ginkgo

go 1.16

require (
	github.com/go-task/slim-sprig v0.0.0-20210107165309-348f09dbbbc0
	github.com/nxadm/tail v1.4.8
	github.com/onsi/gomega v1.18.1
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e
	golang.org/x/tools v0.1.10
)

retract v1.16.3 // git tag accidentally associated with incorrect git commit
