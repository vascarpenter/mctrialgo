-- MySQL dump 10.13  Distrib 8.0.19, for osx10.15 (x86_64)
--
-- Host: localhost    Database: studydb
-- ------------------------------------------------------
-- Server version	8.0.19

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `events`
--

DROP TABLE IF EXISTS `events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `events` (
  `ID` int unsigned NOT NULL AUTO_INCREMENT,
  `hospital_id` int unsigned NOT NULL,
  `serialid` int unsigned NOT NULL,
  `eventid` int unsigned NOT NULL,
  `date` date DEFAULT NULL,
  `alive` tinyint(1) DEFAULT NULL,
  `dropout` tinyint(1) DEFAULT NULL,
  `macce` tinyint(1) DEFAULT NULL,
  `bh` int DEFAULT NULL,
  `bw` int DEFAULT NULL,
  `sbp` int DEFAULT NULL,
  `dbp` int DEFAULT NULL,
  `hr` int DEFAULT NULL,
  `event` varchar(80) DEFAULT NULL,
  `diffdays` int DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=128 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events`
--

LOCK TABLES `events` WRITE;
/*!40000 ALTER TABLE `events` DISABLE KEYS */;
INSERT INTO `events` VALUES (1,7,1,1,'2019-03-04',1,0,0,170,50,120,80,60,'test-711',NULL),(2,7,2,1,'2019-03-04',1,0,0,160,80,130,70,70,'test1',NULL),(3,7,2,2,'2019-03-05',1,0,0,150,80,130,70,80,'test2',NULL),(4,7,1,2,'2019-03-05',1,0,0,150,80,130,80,70,'test2',NULL),(5,7,1,3,'2019-03-06',1,0,0,170,65,130,70,60,'３番目',NULL),(7,8,3,1,'2019-03-04',1,0,0,153,67,130,80,60,'event',NULL),(77,7,1,4,'2019-03-07',1,0,0,170,60,140,70,60,'event-7-1-3',NULL),(78,7,1,5,'2019-03-13',1,0,0,170,60,140,70,60,'event-7-1-3',NULL),(79,7,1,6,'2019-03-23',1,0,0,170,60,140,70,60,'event-7-1-3',NULL),(80,7,1,7,'2019-04-02',1,0,0,170,60,140,70,60,'event-7-1-3',NULL),(81,7,1,8,'2019-04-12',1,0,0,170,60,140,70,60,'event-7-1-3',NULL),(82,7,1,9,'2019-04-22',1,0,0,170,60,140,70,60,'event-7-1-3',NULL),(83,7,1,10,'2019-05-02',1,0,0,170,60,140,70,60,'event-7-1-3',NULL),(84,7,1,11,'2019-05-12',1,0,0,170,60,140,70,60,'event-7-1-3',NULL),(85,7,1,12,'2019-05-22',1,0,1,170,60,140,70,60,'event-7-1-3',NULL),(86,7,2,3,'2019-03-07',1,0,0,170,60,140,70,60,'event-7-2-2',NULL),(87,7,2,4,'2019-03-13',1,0,0,170,60,140,70,60,'event-7-2-2',NULL),(88,7,2,5,'2019-03-23',1,0,0,170,60,140,70,60,'event-7-2-2',NULL),(89,7,2,6,'2019-04-02',1,0,0,170,60,140,70,60,'event-7-2-2',NULL),(90,7,2,7,'2019-04-12',1,0,0,170,60,140,70,60,'event-7-2-2',NULL),(91,7,2,8,'2019-04-22',1,0,0,170,60,140,70,60,'event-7-2-2',NULL),(92,7,2,9,'2019-05-02',1,0,0,170,60,140,70,60,'event-7-2-2',NULL),(93,7,2,10,'2019-05-12',1,0,0,170,60,140,70,60,'event-7-2-2',NULL),(94,7,2,11,'2019-05-22',1,0,0,170,60,140,70,60,'event-7-2-2',NULL),(95,7,3,1,'2019-03-03',1,0,0,170,60,140,70,60,'event-7-3-0',NULL),(96,7,3,2,'2019-03-13',1,0,0,170,60,140,70,60,'event-7-3-0',NULL),(97,7,3,3,'2019-03-23',1,0,0,170,60,140,70,60,'event-7-3-0',NULL),(98,7,3,4,'2019-04-02',1,0,0,170,60,140,70,60,'event-7-3-0',NULL),(99,7,3,5,'2019-04-12',1,0,0,170,60,140,70,60,'event-7-3-0',NULL),(100,7,3,6,'2019-04-22',1,0,0,170,60,140,70,60,'event-7-3-0',NULL),(101,7,3,7,'2019-05-02',1,0,0,170,60,140,70,60,'event-7-3-0',NULL),(102,7,3,8,'2019-05-12',1,0,0,170,60,140,70,60,'event-7-3-0',NULL),(103,7,3,9,'2019-05-22',1,0,0,170,60,140,70,60,'event-7-3-0',NULL),(104,8,3,2,'2019-03-03',1,0,0,170,60,140,70,60,'event-8-3-1',NULL),(105,8,3,3,'2019-03-13',1,0,0,170,60,140,70,60,'event-8-3-1',NULL),(106,8,3,4,'2019-03-23',1,0,0,170,60,140,70,60,'event-8-3-1',NULL),(107,8,3,5,'2019-04-02',1,0,0,170,60,140,70,60,'event-8-3-1',NULL),(108,8,3,6,'2019-04-12',1,0,0,170,60,140,70,60,'event-8-3-1',NULL),(109,8,3,7,'2019-04-22',1,0,0,170,60,140,70,60,'event-8-3-1',NULL),(110,8,3,8,'2019-05-02',0,0,0,170,60,140,70,60,'event-8-3-1',NULL),(111,7,4,1,'2019-03-03',1,0,0,170,60,140,70,60,'event-7-4-0',NULL),(112,7,4,2,'2019-03-13',1,0,0,170,60,140,70,60,'event-7-4-0',NULL),(113,7,4,3,'2019-03-23',1,0,0,170,60,140,70,60,'event-7-4-0',NULL),(114,7,4,4,'2019-04-02',1,0,0,170,60,140,70,60,'event-7-4-0',NULL),(115,7,4,5,'2019-04-12',1,0,1,170,60,140,70,60,'event-7-4-0',NULL),(116,7,5,1,'2019-03-03',1,0,0,170,60,140,70,60,'event-7-5-0',NULL),(117,7,5,2,'2019-03-13',1,0,0,170,60,140,70,60,'event-7-5-0',NULL),(118,7,5,3,'2019-03-23',1,0,0,170,60,140,70,60,'event-7-5-0',NULL),(119,7,5,4,'2019-04-02',1,0,1,170,60,140,70,60,'event-7-5-0',NULL),(120,7,6,1,'2019-03-03',1,0,0,170,60,140,70,60,'event-7-6-0',NULL),(121,7,6,2,'2019-03-13',1,0,0,170,60,140,70,60,'event-7-6-0',NULL),(122,7,6,3,'2020-03-03',0,0,0,170,60,140,70,60,'event-7-6-0',NULL),(123,8,2,1,'2019-03-03',1,0,0,170,60,140,70,60,'event-8-2-0',NULL),(124,8,2,2,'2019-03-13',1,0,0,170,60,140,70,60,'event-8-2-0',NULL),(125,8,2,3,'2019-03-23',1,0,0,170,60,140,70,60,'event-8-2-0',NULL),(126,8,2,4,'2019-04-02',1,0,0,170,60,140,70,60,'event-8-2-0',NULL),(127,8,2,5,'2019-04-12',1,0,1,170,60,140,70,60,'event-8-2-0',NULL);
/*!40000 ALTER TABLE `events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `hospitals`
--

DROP TABLE IF EXISTS `hospitals`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `hospitals` (
  `ID` int unsigned NOT NULL AUTO_INCREMENT,
  `hospital_id` int unsigned NOT NULL,
  `name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL,
  `userid` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `userpass` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `mailaddress` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `hospitals`
--

LOCK TABLES `hospitals` WRITE;
/*!40000 ALTER TABLE `hospitals` DISABLE KEYS */;
INSERT INTO `hospitals` VALUES (1,7,'サンプル病院(pass test)','2020-01-30 17:48:57','test','$2a$10$HuKg5hNOrQvlNn1p6Ck00O45fk1MoL4gsikUbbnvCz7bsAGWEUa.y','test@gmail.com'),(2,8,'サンプル病院２(pass test2)','2020-02-02 20:29:28','test2','$2b$10$lKqbAlp39v1LNPR2tpHzMOpCbatpxlOdEqqr2QhsBdid2ATEAk0zu','test2@me.com'),(3,1,'admin','2020-01-30 17:48:57','admin','$2a$10$1BzsC7cdT.pvQ4fgyIWL..dQP3H2X0/ynPKhL0sbl5kEcPzqH4Ffy','admintest@gmail.com');
/*!40000 ALTER TABLE `hospitals` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `patients`
--

DROP TABLE IF EXISTS `patients`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `patients` (
  `ID` int unsigned NOT NULL AUTO_INCREMENT,
  `patient_id` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `hospital_id` int unsigned NOT NULL,
  `serialid` int unsigned NOT NULL,
  `trialgroup` int NOT NULL,
  `initial` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `birthdate` date DEFAULT NULL,
  `female` tinyint(1) DEFAULT NULL,
  `age` int DEFAULT NULL,
  `allowdate` date DEFAULT NULL,
  `startdate` date DEFAULT NULL,
  `dropdate` date DEFAULT NULL,
  `maccedate` date DEFAULT NULL,
  `deaddate` date DEFAULT NULL,
  `finishdate` date DEFAULT NULL,
  `diffdays` int DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `patients`
--

LOCK TABLES `patients` WRITE;
/*!40000 ALTER TABLE `patients` DISABLE KEYS */;
INSERT INTO `patients` VALUES (1,'1',7,1,0,'A.I.','1970-03-04',0,49,'2019-03-03','2019-03-03','2019-05-22',NULL,NULL,NULL,NULL),(2,'2',7,2,0,'B.C','1950-05-06',0,69,'2019-03-03','2019-03-03',NULL,NULL,NULL,'2020-03-08',NULL),(3,'3',7,3,0,'D.E','1960-01-23',0,50,'2019-03-03','2019-03-03',NULL,NULL,NULL,'2020-03-08',NULL),(4,'S2-001',8,3,1,'E.C','1977-04-05',1,42,'2019-03-03','2019-03-03',NULL,NULL,'2019-05-02',NULL,NULL),(5,'7',7,4,0,'T.T','1970-01-01',0,50,'2019-03-03','2019-03-03','2019-04-12',NULL,NULL,NULL,NULL),(6,'8',7,5,0,'F.G','1920-01-01',0,100,'2019-03-03','2019-03-03','2019-04-02',NULL,NULL,NULL,NULL),(7,'6-6',7,6,0,'I.D','1970-01-01',1,50,'2019-03-03','2019-03-03',NULL,NULL,'2020-03-03','2020-03-06',NULL),(8,'S2-002',8,2,1,'I.D','1945-01-01',1,75,'2019-03-03','2019-03-03','2019-04-12',NULL,NULL,NULL,NULL);
/*!40000 ALTER TABLE `patients` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-03-08 17:05:47
