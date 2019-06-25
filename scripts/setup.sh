#!/bin/bash

cd $WORKSPACE
echo "****************************************************"
echo "govault :: Fetching dependencies"
echo "****************************************************"
go get -u github.com/boltdb/bolt
go get -u github.com/gorilla/mux