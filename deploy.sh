#!/bin/bash

cd client
grunt
mv dist ..
cd ..
goapp deploy -oauth
rm -rf dist
