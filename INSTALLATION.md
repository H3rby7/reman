# Installation

1. Install GO (follow the default, where it puts go into `your-user-directory/go`)
1. Set the GOPATH variable as suggested to the go root folder
1. Set the GOBIN variable to GOPATH/bin (where the binaries reside)
1. Install buffalo: `go get -u -v github.com/gobuffalo/buffalo/buffalo`
1. Install buffalo-pop: `go get -u -v github.com/gobuffalo/buffalo-pop` (not sure if needed)
1. Git clone this repository into whatever go path you have under `src/github.com/h3rby7/reman`
1. If you use intelliJ make sure to enable VGO with proxyMode = direct
1. Then go over any 'red' import and sync the dependencies. This may take a while.
