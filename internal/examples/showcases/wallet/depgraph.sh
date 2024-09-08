#!/bin/bash

set -ev

$GOPATH/bin/godepgraph -horizontal -s -o github.com/itskovichanton/qt/internal/examples/showcases/wallet github.com/itskovichanton/qt/internal/examples/showcases/wallet | dot -Tpng -o depgraph.png