#!/usr/bin/python
"""
This is the most simple example to showcase Containernet.
"""
from mininet.net import Containernet
from mininet.node import Controller, RemoteController
from mininet.cli import CLI
from mininet.link import TCLink
from mininet.log import info, setLogLevel
setLogLevel('info')

controller = RemoteController(name='c0', port=6653, ip="192.168.0.106")
net = Containernet(controller=RemoteController, link=TCLink)
# info('*** Adding controller\n')

net.addController(controller)
info('*** Adding docker containers\n')
h1 = net.addDocker('h1', ip='192.168.0.120/24', dimage="ubuntu:test")
h2 = net.addDocker('h2', ip='192.168.0.121/24', dimage="ubuntu:test")
h3 = net.addDocker('h3', ip='192.168.0.123/24', dimage="ubuntu:test")
h4 = net.addDocker('h4', ip='192.168.0.124/24', dimage="ubuntu:test")
h5 = net.addDocker('h5', ip='192.168.0.125/24', dimage="ubuntu:test")
info('*** Adding switches\n')
s1 = net.addSwitch('s1', protocols='OpenFlow13')
s2 = net.addSwitch('s2', protocols='OpenFlow13')
s3 = net.addSwitch('s3', protocols='OpenFlow13')
s4 = net.addSwitch('s4', protocols='OpenFlow13')
s5 = net.addSwitch('s5 ', protocols='OpenFlow13')
info('*** Creating links\n')
net.addLink(h1, s1)
net.addLink(h2, s2)
net.addLink(h3, s3)
net.addLink(h4, s4)
net.addLink(h5, s5)
net.addLink(s1, s2, cls=TCLink)
net.addLink(s2, s3, cls=TCLink)
net.addLink(s3, s4, cls=TCLink)
net.addLink(s4, s5, cls=TCLink)
net.addLink(s5, s1, cls=TCLink)
net.addLink(s3, s5, cls=TCLink)
info('*** Starting network\n')
net.start()
info('*** Running CLI\n')
CLI(net)
info('*** Stopping network')
net.stop()

