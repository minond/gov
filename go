#!/bin/bash
#
# Use this go instead of go-go to get GO111MODULE when you need it.

set -euo pipefail

#=== ENVVARS ==================================================================
#          GOV: Path to G.O.V. root directory.
#           GO: Path to common and the currently linked version of Go.
#==============================================================================
GOV=${GOV:-~/.gov}
GO=${GO:-$GOV/go}

if [ -f go.mod ]; then
  export GO111MODULE=on
fi

$GO/bin/go $*
