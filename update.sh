./lite --test https://raw.githubusercontent.com/chengaopan/AutoMergePublicNodes/master/list.txt \
--config ./config.json

jq -r '.nodes.[]|select(.ping!="0")| .link' output.json >> ./tmp
sort -u ./tmp > nodes
