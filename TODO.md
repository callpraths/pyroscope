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
- Out of the box `make test` failed. Investigation pending.
- Running the [ridshare tutorial](https://grafana.com/docs/pyroscope/latest/get-started/ride-share-tutorial/) fails:
  ```
   > [us-east 2/4] RUN pip3 install flask pyroscope-io==0.8.11 pyroscope-otel==0.4.0:                                                                                             
  2.342 Collecting flask                                                                                                                                                          
  2.391   Downloading flask-3.1.1-py3-none-any.whl (103 kB)                                                                                                                       
  2.414      ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 103.3/103.3 kB 4.7 MB/s eta 0:00:00
  2.486 Collecting pyroscope-io==0.8.11
  2.494   Downloading pyroscope_io-0.8.11-py2.py3-none-manylinux_2_17_x86_64.manylinux2014_x86_64.whl (2.7 MB)
  2.625      ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 2.7/2.7 MB 21.0 MB/s eta 0:00:00
  2.711 Collecting pyroscope-otel==0.4.0
  2.719   Downloading pyroscope_otel-0.4.0-py3-none-any.whl (10 kB)
  2.998 Collecting cffi>=1.6.0
  3.005   Downloading cffi-1.17.1-cp39-cp39-manylinux_2_17_x86_64.manylinux2014_x86_64.whl (445 kB)
  3.029      ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 445.2/445.2 kB 20.9 MB/s eta 0:00:00
  3.085 Collecting opentelemetry-sdk<2.0.0,>=1.27.0
  3.092   Downloading opentelemetry_sdk-1.33.1-py3-none-any.whl (118 kB)
  3.101      ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 119.0/119.0 kB 17.3 MB/s eta 0:00:00
  3.105 INFO: pip is looking at multiple versions of pyroscope-io to determine which version is compatible with other requirements. This could take a while.
  3.106 ERROR: Cannot install pyroscope-io==0.8.11 and pyroscope-otel==0.4.0 because these package versions have conflicting dependencies.
  3.106 
  3.106 The conflict is caused by:
  3.106     The user requested pyroscope-io==0.8.11
  3.106     pyroscope-otel 0.4.0 depends on pyroscope-io==0.8.8
  3.106 
  3.106 To fix this you could try to:
  3.106 1. loosen the range of package versions you've specified
  3.106 2. remove package versions to allow pip attempt to solve the dependency conflict
  3.106 
  3.107 ERROR: ResolutionImpossible: for help visit https://pip.pypa.io/en/latest/topics/dependency-resolution/#dealing-with-dependency-conflicts
  3.233 
  3.233 [notice] A new release of pip is available: 23.0.1 -> 25.1.1
  3.233 [notice] To update, run: pip install --upgrade pip
  ------
  failed to solve: process "/bin/sh -c pip3 install flask pyroscope-io==0.8.11 pyroscope-otel==0.4.0" did not complete successfully: exit code: 1
  ```
  Looks to have broken in [!4179](https://github.com/grafana/pyroscope/pull/4179)
