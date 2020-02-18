#!/usr/bin/env bash
workspace=$(cd $(dirname $0) && pwd -P)

{
    cd $workspace
    find
}