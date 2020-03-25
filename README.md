Casbin Enforcers
===

## Introduction

This repository contains an extension to the casbin framework allowing the user to chain an arbitrary number of custom enforcers. This allows functionality to be layered on top of the default casbin enforcer.

Any object which implements the BasicEnforcer or FullEnforcer interface can be used in building a chain. This enables powerful dynamic chaining which can be configured dynamically at runtime. There are currently two enforcers provided by this package: cached and synced.

The synced enforcer is a thread safe wrapper to provide safe access to the enforcer when accessing the API from multiple concurrent goroutines.

The cached enforcer implements basic request caching to speed up common queries. It optionally implements an automatic cache invalidation on policy updates. More granular control over cache invalidation will be added in the future.

