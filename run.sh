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
        echo "git push origin dev:dev"
        git push origin dev:dev
        
        echo "git push alimy r/paopao-ce:r/paopao-ce"
        git push alimy r/paopao-ce:r/paopao-ce

        echo "git push bitbus r/paopao-ce-plus:r/paopao-ce-plus"
        git push bitbus r/paopao-ce-plus:r/paopao-ce-plus
          
        echo "git push bitbus r/paopao-ce-pro:r/paopao-ce-pro"
        git push bitbus r/paopao-ce-pro:r/paopao-ce-pro

        echo "git push bitbus r/paopao-ce-xtra:r/paopao-ce-xtra"
        git push bitbus r/paopao-ce-xtra:r/paopao-ce-xtra
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