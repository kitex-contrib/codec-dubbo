#!/bin/sh

close_name="close"
close_all_name="all"

stress_name="stress"
stress_dir="./stress"
stress_main="stress_main"

dubbo_client_name="dubbo_client"
dubbo_client_dir="./dubbo/client"
dubbo_client_main="dubbo_client_main"

dubbo_server_name="dubbo_server"
dubbo_server_dir="./dubbo/server"
dubbo_server_main="dubbo_server_main"

kitex_client_name="kitex_client"
kitex_client_dir="./kitex/client"
kitex_client_main="kitex_client_main"

kitex_server_name="kitex_server"
kitex_server_dir="./kitex/server"
kitex_server_main="kitex_server_main"

close_process() {
  pid=$(ps -ef | grep -E "$1" | grep -v grep | awk '{print $2}')
  if [ -n "$pid" ];
  then
    kill -9 "$pid"
  fi
}

close() {
  if test "$1" = "$close_all_name";
  then
    all_process=("$stress_main" "$dubbo_client_main" "$kitex_client_main" "$dubbo_server_main" "$kitex_server_main")
    for p in "${all_process[@]}";
    do
      close_process "$p"
    done
    return
  fi

  if test "$1" = "$stress_name";
  then
    close_process "$stress_main"
  fi
  if test "$1" = "$dubbo_client_name";
  then
    close_process "$dubbo_client_main"
  fi
  if test "$1" = "$kitex_client_name";
  then
    close_process "$kitex_client_main"
  fi
  if test "$1" = "$dubbo_server_name";
  then
    close_process "$dubbo_server_main"
  fi
  if test "$1" = "$kitex_server_name";
  then
    close_process "$kitex_server_main"
  fi
  close_process "$1"
}

stress() {
  cd "$stress_dir" || exit
  p=$stress_main
  go build -o "$p" .
  chmod +x "$p"
  ./"$p" "$1" "$2" "$3" "$4" "$5" "$6" >../stress_log 2>&1
}

client() {
  cd "$1" || exit
  p=$2
  go build -o "$p" .
  chmod +x "$p"
  nohup ./"$p" "$3" "$4" "$5" "$6" >/dev/null 2>&1 &
}

server() {
  cd "$1" || exit
  p=$2
  go build -o "$p" .
  chmod +x "$p"
  nohup ./"$p" "$3" "$4" >/dev/null 2>&1 &
}

if test "$1" = "$stress_name";
then
  stress "$2" "$3" "$4" "$5" "$6" "$7"
elif test "$1" = "$dubbo_client_name";
then
  client "$dubbo_client_dir" "$dubbo_client_main" "$2" "$3" "$4" "$5"
elif test "$1" = "$kitex_client_name";
then
  client "$kitex_client_dir" "$kitex_client_main" "$2" "$3" "$4" "$5"
elif test "$1" = "$dubbo_server_name";
then
  server "$dubbo_server_dir" "$dubbo_server_main" "$2" "$3"
elif test "$1" = "$kitex_server_name";
then
  server "$kitex_server_dir" "$kitex_server_main" "$2" "$3"
elif test "$1" = "$close_name";
then
  close "$2"
else
  echo "supporting parameter: $stress_name, $dubbo_client_name, $kitex_client_name, $dubbo_server_name, $kitex_server_name, $close_name"
  exit
fi
