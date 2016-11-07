# Hands on [NATS](http://nats.io/)

Materials for my _Messaging in the Cloud: hands on NATS_ session at UBS Coding Harbor meetup.

## Prelude

- Why yet another messaging broker?
- Choose the right tool for the job!
- Exactly-once delivery... WAAAAT?
- 2 generals problem

## NATS 101

- [NATS protocol](http://nats.io/documentation/internals/nats-protocol/)
- Pure pub-sub
  - At-most-once delivery - fire and forget
  - Always ordered from single publisher
  - Limiting interest
  - Subject wildcards
- Queue groups
- Request-reply
  - Tail latency

## Clustering

- Always on (dial tone)
- Full mesh
- Auto-discovery
- Cluster-aware clients
