dump:
	mysqldump -uoge -phogehogeA00 -r studydb.sql --single-transaction studydb

import:
	mysql -uoge -phogehogeA00 studydb < studydb.sql

run:
	go run .

build:
	go build

clean: 
	rm -f server mctrialgo mctrialgo.zip

zip:	clean dump
	zip -9 -r mctrialgo.zip ./*

