# ************************************************************
# Sequel Pro SQL dump
# Version 5446
#
# https://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: localhost (MySQL 8.0.15)
# Database: studydb
# Generation Time: 2020-02-23 13:52:57 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table events
# ------------------------------------------------------------

DROP TABLE IF EXISTS `events`;

CREATE TABLE `events` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `hospitalid` int(11) DEFAULT NULL,
  `serialid` int(11) DEFAULT NULL,
  `eventid` int(11) DEFAULT NULL,
  `date` date DEFAULT NULL,
  `alive` tinyint(1) DEFAULT NULL,
  `bh` int(11) DEFAULT NULL,
  `bw` int(11) DEFAULT NULL,
  `sbp` int(11) DEFAULT NULL,
  `dbp` int(11) DEFAULT NULL,
  `hr` int(11) DEFAULT NULL,
  `event` varchar(80) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `events` WRITE;
/*!40000 ALTER TABLE `events` DISABLE KEYS */;

INSERT INTO `events` (`id`, `hospitalid`, `serialid`, `eventid`, `date`, `alive`, `bh`, `bw`, `sbp`, `dbp`, `hr`, `event`)
VALUES
	(1,7,1,1,'2020-01-01',1,170,50,120,80,60,'test'),
	(2,7,2,1,'2020-01-01',0,160,80,130,70,70,'test1'),
	(3,7,2,2,'2020-01-02',1,150,80,130,70,80,'test2');

/*!40000 ALTER TABLE `events` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table hospitals
# ------------------------------------------------------------

DROP TABLE IF EXISTS `hospitals`;

CREATE TABLE `hospitals` (
  `hospital_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL,
  `userid` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `userpass` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `mailaddress` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  PRIMARY KEY (`hospital_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `hospitals` WRITE;
/*!40000 ALTER TABLE `hospitals` DISABLE KEYS */;

INSERT INTO `hospitals` (`hospital_id`, `name`, `created_at`, `userid`, `userpass`, `mailaddress`)
VALUES
	(7,'サンプル病院(pass test)','2020-01-30 17:48:57','test','$2b$10$k6ICyWTb73z6td4yO2fmaOyIgnLyVv/xmjhvWiC796WEQiOblRVDy','test@gmail.com'),
	(8,'サンプル病院２','2020-02-02 20:29:28','test2','$2b$10$lKqbAlp39v1LNPR2tpHzMOpCbatpxlOdEqqr2QhsBdid2ATEAk0zu','test2@me.com');

/*!40000 ALTER TABLE `hospitals` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table patients
# ------------------------------------------------------------

DROP TABLE IF EXISTS `patients`;

CREATE TABLE `patients` (
  `patient_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `hospital_id` int(11) DEFAULT NULL,
  `serialid` int(11) DEFAULT NULL,
  `initial` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `birthdate` date DEFAULT NULL,
  `female` tinyint(1) DEFAULT NULL,
  `age` int(11) DEFAULT NULL,
  `allowdate` date DEFAULT NULL,
  `startdate` date DEFAULT NULL,
  `dropdate` date DEFAULT NULL,
  `dropout` tinyint(1) DEFAULT NULL,
  `finishdate` date DEFAULT NULL,
  PRIMARY KEY (`patient_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `patients` WRITE;
/*!40000 ALTER TABLE `patients` DISABLE KEYS */;

INSERT INTO `patients` (`patient_id`, `hospital_id`, `serialid`, `initial`, `birthdate`, `female`, `age`, `allowdate`, `startdate`, `dropdate`, `dropout`, `finishdate`)
VALUES
	(1,7,1,'A.I.','1970-03-04',0,49,NULL,NULL,NULL,NULL,NULL),
	(2,7,2,'B.C','1950-05-06',0,69,NULL,NULL,NULL,NULL,NULL),
	(3,7,3,'D.E','1960-01-23',0,50,NULL,NULL,NULL,NULL,NULL),
	(4,8,1,'E.C','1980-04-05',1,10,NULL,NULL,NULL,NULL,NULL),
	(7,7,4,'T.T','1970-01-01',0,50,NULL,NULL,NULL,NULL,NULL),
	(8,7,5,'F.G','1920-01-01',0,100,NULL,NULL,NULL,NULL,NULL);

/*!40000 ALTER TABLE `patients` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
