#!/bin/sh

stress_name="stress"
client_name="client"
server_name="server"

if $1 = "$stress_name"; then
  stress "$2" "$3" "$4" "$5" "$6" "$7"
elif $1 = "$client_name"; then
  client "$2" "$3"
elif $1 = "$server_name"; then
  server "$2"
else
  echo "supporting parameter: $stress_name, $client_name, $server_name"
  exit
fi

stress() {
  cd ./stress || exit
  go build -o stress_main .
  chmod +x stress_main
  pid=$(ps -ef | grep -E "stress_main" | grep -v grep | awk '{print $2}')
  kill -9 "$pid" && nohup ./stress_main "$1" "$2" "$3" "$4" "$5" "$6"  >/dev/null 2>&1 &
}

client() {
  cd ./client || exit
  p=client_main
  go build -o "$p" .
  chmod +x "$p"
  pid=$(ps -ef | grep -E "$p" | grep -v grep | awk '{print $2}')
  kill -9 "$pid" && nohup ./"$p" "$1" "$2" > /dev/null 2>&1 &
}

server() {
  cd ./server || exit
  p=server_main
  go build -o "$p" .
  chmod +x "$p"
  pid=$(ps -ef | grep -E "$p" | grep -v grep | awk '{print $2}')
  kill -9 "$pid" && nohup ./"$p" "$1" > /dev/null 2>&1 &
}
