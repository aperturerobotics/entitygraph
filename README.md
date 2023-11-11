# Entity Graph

> Expose application state as a connected graph of entities for visualization. 

## Introduction

[![Go Reference Widget]][Go Reference] [![Go Report Card Widget]][Go Report Card]

[Go Reference]: https://pkg.go.dev/github.com/aperturerobotics/entitygraph
[Go Reference Widget]:https://pkg.go.dev/badge/github.com/aperturerobotics/entitygraph.svg
[Go Report Card Widget]: https://goreportcard.com/badge/github.com/aperturerobotics/entitygraph
[Go Report Card]: https://goreportcard.com/report/github.com/aperturerobotics/entitygraph

The Entity Graph is a common format for application components to expose their
internal state for monitoring in the form of a graph of entities.

 - Entities can be shared between components, identified by common identifiers
 - Links between entities form edges in the graph.
 - Properties on entities and links describe instantaneous metric values.

At runtime, any visualization frameworks can create a CollectEntityGraph
directive to collect all active entities exposed by running controllers.

## Support

Please open a [GitHub issue] with any questions / issues.

[GitHub issue]: https://github.com/aperturerobotics/entitygraph/issues/new

... or feel free to reach out on [Matrix Chat] or [Discord].

[Discord]: https://discord.gg/KJutMESRsT
[Matrix Chat]: https://matrix.to/#/#aperturerobotics:matrix.org

## License

Apache 2.0
