package collector

import (
	"testing"

	"github.com/qianweicheng/mongodb_exporter/shared"
	"github.com/prometheus/client_golang/prometheus"
	"strings"
	"fmt"
	"github.com/golang/glog"
)

func Test_CollectServerStatus(t *testing.T) {
	shared.ParseEnabledGroups("assers,durability,backgrond_flushing,connections,extra_info,global_lock,index_counters,network,op_counters,memory,locks,metrics,cursors")
	collector := NewMongodbCollector(MongodbCollectorOpts{URI: "localhost"})
	go collector.Collect(nil)
}

func Test_DescribeCollector(t *testing.T) {
	collector := NewMongodbCollector(MongodbCollectorOpts{URI: "localhost"})

	ch := make(chan *prometheus.Desc)
	go collector.Describe(ch)
}

func Test_CollectCollector(t *testing.T) {
	collector := NewMongodbCollector(MongodbCollectorOpts{URI: "localhost"})

	ch := make(chan prometheus.Metric)
	go collector.Collect(ch)
}

func Test_ShardCollectCollector(t *testing.T) {
	k := "rs1/mongodb-shad-a-0.mongodb-shad.default.svc.cluster.local:27018, mongodb-shad-a-1.mongodb-shad.default.svc.cluster.local:27018,mongodb-shad-a-2.mongodb-shad.default.svc.cluster.local:27018"
	hostStr := k[len("rs1")+1:]
	hosts := strings.Split(hostStr, ",")
	glog.Error(fmt.Sprintf("Host: %s", hostStr))
	for _, host := range hosts {
		glog.Error(fmt.Sprintf("Repl Host: %s", host))
	}
}
