
rm ./tmp

while IFS= read -r line; do
./lite --test $line --config ./config.json

jq -r '.nodes.[]|select(.ping!="0")| .link' output.json >> ./tmp

done < ./list.txt


sort -u ./tmp > nodes
