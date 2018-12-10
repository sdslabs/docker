package daemon // import "github.com/sdslabs/docker/daemon"

import (
	// Importing packages here only to make sure their init gets called and
	// therefore they register themselves to the logdriver factory.
	_ "github.com/sdslabs/docker/daemon/logger/awslogs"
	_ "github.com/sdslabs/docker/daemon/logger/etwlogs"
	_ "github.com/sdslabs/docker/daemon/logger/fluentd"
	_ "github.com/sdslabs/docker/daemon/logger/gcplogs"
	_ "github.com/sdslabs/docker/daemon/logger/gelf"
	_ "github.com/sdslabs/docker/daemon/logger/jsonfilelog"
	_ "github.com/sdslabs/docker/daemon/logger/logentries"
	_ "github.com/sdslabs/docker/daemon/logger/splunk"
	_ "github.com/sdslabs/docker/daemon/logger/syslog"
)
