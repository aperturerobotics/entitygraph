# Entity Graph

> Expose application state as a connected graph of entities for visualization. 

## Introduction

The Entity Graph is a common format for application components to expose their
internal state for monitoring in the form of a graph of entities.

 - Entities can be shared between components, identified by common identifiers
 - Links between entities form edges in the graph.
 - Properties on entities and links describe instantaneous metric values.

At runtime, any visualization frameworks can create a CollectEntityGraph
directive to collect all active entities exposed by running controllers.
