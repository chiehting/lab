#/bin/sh

cd `dirname $0`

count=3
interval=`expr 9 / ${count}`
domain=https://stage-api.example.com/api/v1/ping

echo $domain
for i in $(seq 1 $count)
do
  printf "##### starting_time:  $(date)\n"

  curl -w @- -s -o /dev/null "$domain" \
   <<'EOF'
      \n
    time_namelookup:  %{time_namelookup}\n
       time_connect:  %{time_connect}\n
    time_appconnect:  %{time_appconnect}\n
      time_redirect:  %{time_redirect}\n
   time_pretransfer:  %{time_pretransfer}\n
 time_starttransfer:  %{time_starttransfer}\n
           ---------------------\n
           time_total:  %{time_total} seconds\n
EOF

  sleep $interval

done

