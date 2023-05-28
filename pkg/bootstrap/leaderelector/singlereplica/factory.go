package singlereplica

import (
	"context"

	"github.com/solo-io/gloo/pkg/bootstrap/leaderelector"
	"github.com/solo-io/go-utils/contextutils"
)

var _ leaderelector.ElectionFactory = new(singleReplicaElectionFactory)

// singleReplicaElectionFactory runs leader election for components that do not support true leader election
// The election is a no-op and returns an Identity that is always considered the "leader" since there is only one
type singleReplicaElectionFactory struct {
}

func NewElectionFactory() *singleReplicaElectionFactory {
	return &singleReplicaElectionFactory{}
}

func (f *singleReplicaElectionFactory) StartElection(ctx context.Context, electionConfig *leaderelector.ElectionConfig) (leaderelector.Identity, error) {
	contextutils.LoggerFrom(ctx).Debugf("Starting Single Replica Leader Election")

	if electionConfig == nil {
		contextutils.LoggerFrom(ctx).Debugf("(4)-> nil electionConfig")
	} else {
		contextutils.LoggerFrom(ctx).Debugf("(4)-> electionConfig not nil")
	}

	return Identity(), nil
}

// Identity returns the Identity used in single replica elections
// Since there is only 1 replica, the identity is always considered the "leader"
func Identity() leaderelector.Identity {
	elected := make(chan struct{})
	close(elected) // immediately signal the identity as the leader
	return leaderelector.NewIdentity(elected)
}
