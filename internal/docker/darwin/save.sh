#!/bin/bash

docker save itskovichanton/qt:darwin | gzip -n > darwin.tar.gz
