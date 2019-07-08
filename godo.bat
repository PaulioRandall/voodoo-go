@echo off

cls
del /F "./cmd/voodoo.exe"

set build=0
if /I [%1] EQU [build] set build=1
if /I [%1] EQU [run] set build=1
if /I %build% EQU 1 (
	echo Building...
	cd cmd
	go build voodoo.go
	cd ..
)

set run=0
if [%1] EQU [run] set run=1
if [%run%] EQU [1] (
	echo Running...
	cd cmd
	voodoo.exe exe scroll_1.voo
	cd ..
)