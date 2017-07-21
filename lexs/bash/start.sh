#!/bin/bash

cd ./_output/local

rm -rf ./nohup.out

mkdir -p ./data

nohup ./bin/conf.x &

nohup ./bin/account.x &

nohup ./bin/image.x &

nohup ./bin/post.x &

nohup ./bin/feeds.x &

nohup ./bin/recsys.x &

nohup ./bin/search.x &

