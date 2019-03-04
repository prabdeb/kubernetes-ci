# Kubernetes CI

WIP: Simple gated CI Server to run on K8's Cluster typically in a hosted environment

## Architecture

### Components

1. Server - Receive notifications on code changes, parse and trigger Pipeline
    * SCM - WebHook (GitHub/BitBucket Server)
    * Pipeline Parser
    * Conditional Driver (Push/Tag/Pull Request/GIT Diff/GIT Branch/File Content)
    * Stage Loop (for/while)
    * Pipeline Execution (K8's Jobs)
    * Data Store & Logger (configMaps/etcd-io bbolt)
2. Server UI - User Interface to View/Admin Server
3. Authenticator - Authenticate and authorize users
    * RBAC
    * Dex
4. Secrets - Store Secrets as K8's secrets
5. API - API for server
6. CLI - CLI for server