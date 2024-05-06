#!/bin/bash
clear
start=`date +%s`
make run
./example
exitCode=`echo $?`
end=`date +%s`

runtime=$((end-start))

echo ""
read -p "Program finished in $runtime seconds. Status: $exitCode (Q)uit: " input

case "${input}" in
    (q|Q) exit 0 ;;
    *) ./run.sh ;;
    esac
