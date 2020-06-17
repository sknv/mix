#!/usr/bin/env bash

PROTO_DIR=app/proto
THIRD_PARTY_DIR=third_party/protoc/include

# Account
protoc -I ${PROTO_DIR} -I ${THIRD_PARTY_DIR} \
    --gofast_out=plugins=grpc:${PROTO_DIR} \
    ${PROTO_DIR}/*.proto
