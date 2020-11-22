#! /bin/bash


echo "--- sort /etc/passwd ref \`man sort\` --- "

sort -t ':' -k 3 -n /etc/passwd
echo "--- sort file by file size and order by desc ---"
du -sh * | sort -nr


echo  "---use coproc to run in backgroud ---"

coproc ( sleep 10 ;sleep 2)

