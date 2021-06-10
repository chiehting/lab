#!/bin

#copyFrom=goalstrike-clients-service/templates/NOTES.txt
copyFrom=
oldString=goalstrike-clients-service

if [[ ! -n "$copyFrom" ]]; then
  echo 'Missing source file'
  exit
fi

for i in $(ls -d */ | grep -v $oldString)
do
  echo "Copy from $copyFrom to ${copyFrom/$oldString\//$i}"
  cp -f $copyFrom ${copyFrom/$oldString/$i}

  echo "Replace sting to ${copyFrom/$oldString\//$i}"
  sed -i "" "s/$oldString/${i/\//}/g" "${copyFrom/$oldString\//$i}"
done
