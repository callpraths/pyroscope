These are small quality of life issues I've encountered setting up the repo for the first time.
Dropping them here to make enough headway to gain confidence that these issues are for real.

- [docs/internal/contributing/README.md](docs/internal/contributing/README.md) Instruction to build should be `make build-dev`, not `make go/bin`.
  Out of the box, `make go bin` fails:
  ```
  pprabhu@oss1:/mnt/data/src/pyroscope$ make go/bin
  grep: warning: ? at start of expression
  GOOS=linux GOARCH=amd64 GOAMD64=v2 CGO_ENABLED=0 go build -tags "netgo embedassets" -ldflags "-extldflags \"-static\" -s -w -X github.com/grafana/pyroscope/pkg/util/build.Branch=main -X github.com/grafana/pyroscope/pkg/util/build.Version=main-0e4976a03 -X github.com/grafana/pyroscope/pkg/util/build.Revision=0e4976a03 -X github.com/grafana/pyroscope/pkg/util/build.BuildDate=2025-05-23T15:22:34+01:00" -gcflags= ./cmd/pyroscope
  public/assets_embedded.go:16:12: pattern build: no matching files found
  ```
  - while there, probably best to scan and fix broken links etc in [the docs](https://grafana.com/docs/pyroscope/latest/) in one go.
- Out of the box `make test` failed. Investigation pending.


## Project

Build a demo website to showcase Pyroscope components running in a distributed setup.

