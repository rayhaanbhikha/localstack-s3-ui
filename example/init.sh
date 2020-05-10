#!/bin/bash

which aw

state=$?
if [ $state -ne 0 ]
then
  echo "Download aws cli first"
  exit $state
fi

export AWS_ACCESS_KEY_ID=mock-key
export AWS_SECRET_ACCESS_KEY=mock-key
export AWS_PAGER=''

aws --endpoint-url=http://localhost:4566 s3api create-bucket --bucket my-bucket
aws --endpoint-url=http://localhost:4566 s3 cp ./web-app s3://my-bucket/web-app --recursive