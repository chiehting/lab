#!/bin/sh

cd `dirname $0`/..; pwd
. build/var.sh

set -ex

npm install
npm run build

echo "dev stage prod" | xargs -n1 cp -rf dist
cd dev && find . -type f -exec sed -i "s/$prodAPIURL/$devAPIURL/g" {} \;
cd ../stage && find . -type f -exec sed -i "s/$prodAPIURL/$stageAPIURL/g" {} \;

