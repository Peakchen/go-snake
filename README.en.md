####Software architecture


1. Data security protection of gate server, load balancing and distribution of client messages

2. Game server message processing, data storage

3. Login server login service

4. Robot server client simulator

5. db mysql

6. ...


####Instructions for use

Serverapp parameter list:

"args": [

"-app",

"gate",

"-ver",

"1",

"-webhost",

"127.0.0.1:19921",

"-tcphost",

"127.0.0.1:19922",

"-pprof",

"127.0.0.1",

"-log",

"1",

"-clients",

"1"

]



All server modules have init, run and other initial start modes.