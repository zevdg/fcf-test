#!/bin/bash
cd "$(dirname "$0")"
go mod vendor
gcloud functions deploy FcfTest --runtime go111 --trigger-event providers/cloud.firestore/eventTypes/document.write --trigger-resource "projects/fb-js-sdk-test/databases/(default)/documents/go-func-test/{docId}"
