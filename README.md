# Godzilla

[![CircleCI](https://circleci.com/gh/ankur22/godzilla/tree/master.svg?style=svg)](https://circleci.com/gh/ankur22/godzilla/tree/master)
[![codecov](https://codecov.io/gh/ankur22/godzilla/branch/master/graph/badge.svg)](https://codecov.io/gh/ankur22/godzilla)
![Lint everything](https://github.com/ankur22/godzilla/workflows/Lint%20everything/badge.svg)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/a787eb9216b9490bbf6da1e2ae5cbc74)](https://www.codacy.com/manual/akahank/godzilla?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=ankur22/godzilla&amp;utm_campaign=Badge_Grade)
![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)

This is a monorepo which contains Go code for a messenger service. This repo is a playground where I will be playing around with the following:

 - Playing around with monorepos
 - Ensuring the chosen CI tools:
   - run golang-lint
   - to run tests
   - to run integration tests
   - to build the applications
 - Creating a server which accepts HTTP and gRPC requests
 - Creating a CLI client to send and recieve messages
 - Zap for logging
 - Prometheus and Grafana for monitoring
 - Using a modern database (postgres, mongo or both)
 - Github actions

## References

 - The original monorepo template was forked from: https://github.com/labs42io/circleci-monorepo

## Changes to the License

I've kept the license as MIT, and added my name to the copyright.

