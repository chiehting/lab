#/bin/sh

#curl -sS -X GET -H "Authorization:uennfk231" -H "Content-Type: application/x-www-form-urlencoded" "https://adminapi.bb-bcgames.com/user/balance?account=99999999999" 2>&1

cd `dirname $0`

count=3
interval=`expr 60 / ${count}`

for i in $(seq 1 $count)
do

  printf "##### starting_time:  $(date)\n"
  curl -w @- -s  -o null "$@" <<'EOF'
      \n
      time_namelookup:  %{time_namelookup}\n
         time_connect:  %{time_connect}\n
      time_appconnect:  %{time_appconnect}\n
     time_pretransfer:  %{time_pretransfer}\n
        time_redirect:  %{time_redirect}\n
   time_starttransfer:  %{time_starttransfer}\n
           ---------------------\n
           time_total:  %{time_total} seconds\n
EOF

  sleep $interval

done

find . -iname '*.log' -type f -mtime +10 -exec rm '{}' \;
