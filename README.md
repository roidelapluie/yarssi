# Yarrsi

Converts Yaml into IRSSI configuration


## Input Yaml

```
---
servers:
- name: freenode
  address: chat.freenode.org
  ssl: true
  autoconnect: true
  port: 6697
  password: querty
- name: geeknode
  address: irc.geeknode.org
  ssl: true
  port: 6697
  autoconnect: true
  autosendcmd: "/msg C nick identify querty"
- name: bitlbee
  address: localhost
  autoconnect: true
  port: 6667
  password: qwerty
  autosendcmd: "/wait 5000"
nickname: roidelapluie
username: roidelapluie
realname: roidelapluie
theme: custom.theme
hilights:
- roidelapluie
windows:
- name: '#Troll'
  server: geeknode
  show: true
- name: '#technical'
  server: bitlbee
  show: true
- name: '#centos-devel'
  server: freenode
- name: '#centos'
  server: freenode
- name: '#puppet'
  server: freenode
```

This will convert this into irssi config.
