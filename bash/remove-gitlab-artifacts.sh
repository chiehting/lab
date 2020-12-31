#!/bin/bash

token="UxxxxxxxxxxxxxxxxxxD"
server="gitlab.example.com"
#project_ids=$(seq 2 102)
project_ids="12 57 30 28 23 73 21"

#for project_id in $(seq 2 ${project_end_id})
for project_id in ${project_ids}
do
  echo $project_id
  job_ids=$(curl --header PRIVATE-TOKEN:${token} https://${server}/api/v4/projects/${project_id}/jobs|jq '.[].id')

  for job_id in ${job_ids}
  do
    URL="https://$server/api/v4/projects/$project_id/jobs/$job_id/erase"
    curl --request POST --header "PRIVATE-TOKEN:${token}" "$URL"
  done
done

exit

