# !/bin/sh

CGO_ENABLED=0 && cd pinghop && go build  -ldflags="-extldflags=-static" -o pinghop && cd ..
mv ./pinghop/pinghop ./pinghop_docker/pinghop
cd ./pinghop_docker && sudo docker build -t ubuntu:test . && cd ..
sudo docker rm $(sudo docker ps -aq) --force &> /dev/null
sudo docker run -d --name onos -e ONOS_APPS=drivers,gui2,openflow,fwd -p 6653:6653 -p 8181:8181 onosproject/onos:2.7.0
sleep 20
sudo python3 ./topo_1.py