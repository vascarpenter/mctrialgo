dump:
	mysqldump -uoge -phogehogeA00 -r studydb.sql --single-transaction studydb

import:
	mysql -uoge -phogehogeA00 studydb < studydb.sql

