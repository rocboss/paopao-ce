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
        push_all dev r/paopao-ce r/paopao-ce-plus r/paopao-ce-pro r/paopao-ce-xtra
    fi
}

function push_all {
    if [ $# -eq 0 ]; then 
        push
    else 
        while [ $# -gt 0 ]; do
            push $1
            shift
        done
    fi 
}

case $1 in
"push")
    shift 
    push_all $@
    ;;
"merge")
    echo "merge command"
    ;;
*)
    push_all
    ;;
esac