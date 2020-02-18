sed -e 's/head/bt/g' -e 's/new\ Node(/\&Node{Value:/g' -e 's/);/}/g' -e 's/Node\ //g' -e 's/left/Left/g' -e 's/right/Right/g' data.go > data
mv data data.go
