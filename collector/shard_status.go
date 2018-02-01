package collector

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
	"fmt"
)

var (
	myShardState   = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: Namespace,
		Subsystem: "shardinfo",
		Name:      "my_state",
		Help:      "An integer between 0 and 10 that represents the replica state of the current member",
	}, []string{"id"})
)

type ShardsStatus struct {
	Ok		string		`bson:"ok"`
	Shards		[]ShardInfo	`bson:"shards"`
}

type ShardInfo struct {
	Id		string		`bson:"_id"`
	Host		string		`bson:"host"`
	State		int32		`bson:"state"`
}

// Export exports the replSetGetStatus stati to be consumed by prometheus
func (shardsStatus *ShardsStatus) Export(ch chan<- prometheus.Metric) {
	myShardState.Reset()

	mFailedShardCount := 0
	for _, member := range shardsStatus.Shards {
		ls := prometheus.Labels{
			"id":  member.Id,
		}

		memberState.With(ls).Set(float64(member.State))
		if member.State != 1 {
			mFailedShardCount += 1
		} else {
			//Get Shard detail info
		}
	}
	glog.Info(fmt.Sprintf("Failed Shard Count: %d", mFailedShardCount))
	// collect metrics
	myShardState.Collect(ch)
}

// Describe describes the GetShardStatus metrics for prometheus
func (shardsStatus *ShardsStatus) Describe(ch chan<- *prometheus.Desc) {
	myShardState.Describe(ch)
}

// GetShardStatus returns the replica status info
func GetShardStatus(session *mgo.Session) *ShardsStatus {
	result := &ShardsStatus{}
	err := session.DB("admin").Run(bson.D{{"listShards", 1}}, result)
	if err != nil {
		glog.Error("Failed to get replSet status:" + err.Error())
		return nil
	}
	return result
}
