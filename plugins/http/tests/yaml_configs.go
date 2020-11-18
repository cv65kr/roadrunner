package tests

var t1 string = `
server:
  command: "php psr-worker.php"
  user: ""
  group: ""
  env:
    "RR_HTTP": "true"
  relay: "pipes"
  relayTimeout: "20s"

http:
  debug: true
  address: 0.0.0.0:8080
  maxRequestSize: 200
  middleware: [ "" ]
  uploads:
    forbid: [ ".php", ".exe", ".bat" ]
  trustedSubnets: [ "10.0.0.0/8", "127.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16", "::1/128", "fc00::/7", "fe80::/10" ]
  pool:
    numWorkers: 4
    maxJobs: 0
    allocateTimeout: 60s
    destroyTimeout: 60s

  ssl:
    port: 8888
    redirect: true
    cert: fixtures/server.crt
    key: fixtures/server.key
  #    rootCa: root.crt
  fcgi:
    address: tcp://0.0.0.0:6920
  http2:
    enabled: false
    h2c: false
    maxConcurrentStreams: 128
`
