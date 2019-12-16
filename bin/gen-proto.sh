#!/bin/bash

case $1 in
  ruby)
    grpc_tools_ruby_protoc -I pb/ --ruby_out=ruby/lib --grpc_out=ruby/lib pb/server_messages.proto
    ;;
esac
