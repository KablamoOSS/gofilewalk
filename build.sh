#!/usr/bin/env bash

echo 'filepathwalk'
cd filepathwalk
go install
cd ..

echo 'filepathwalkdir'
cd filepathwalkdir
go install
cd ..

echo 'iafan'
cd iafan
go install
cd ..

echo 'karrick'
cd karrick
go install
cd ..

echo 'michealtjones'
cd michealtjones
go install
cd ..

echo 'readdir'
cd readdir
go install
cd ..

echo 'All Done'