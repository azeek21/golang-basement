## Notes
My notes for golang intensive at School21 team project 1.

### Task
- Make a replicated/distributed database and a client to work with it.


### Execution
- A clint that works as a frontend for user. Also sends heart beats every N second to leader. If leader doesn't respond it should move on to the next node in the list (we get that list in heartbeats)
- A Cluster that consists of a Leader and N nodes, replication_factor, min_aggree_count (default is more than half of nodes)
- A leader that accepts tasks and replicates to other in the cluster through commit and apply Raft algorithm
- Node that accepts and logs (commiting) and applies when told by leader. Also if it doesn't hear from others in N duration should start a leader election.
- Election timeout should vary between nodes by a bit so not everyone starts leader election at the same time
- If a Leader node accepts an election request, it should disagree
- Every action needs N > NodesInCluster/2 agreements to be done. e.g: leader election, applying changes

> [!NOTE]
> This is a work in progress and subject to changes

