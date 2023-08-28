#!/bin/sh
# eg.1 : sh run.sh
# eg.2, push all release: sh run.sh push
# eg.3, push all release with dev branch: sh run.sh push dev

function push {
    if [ -n "$1" ]; then
        echo "git push origin $1:$1"
        git push origin $1:$1
        
        echo "git push alimy $1:$1"
        git push alimy $1:$1

        echo "git push bitbus $1:$1"
        git push bitbus $1:$1
    else
        push dev
        push r/paopao-ce
        push r/paopao-ce-plus
        push r/paopao-ce-pro
        push r/paopao-ce-xtra
    fi
}

case $1 in
"push")
    push $2
    ;;
"merge")
    echo "merge command"
    ;;
*)
    push $2
    ;;
esac