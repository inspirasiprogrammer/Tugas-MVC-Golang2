CREATE DATABASE  IF NOT EXISTS `digibank` ;
USE `digibank`;
-- MySQL dump 10.13  Distrib 8.0.21, for macos10.15 (x86_64)
--
-- Host: localhost    Database: digibank
-- ------------------------------------------------------
-- Server version	8.0.21

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `accounts`
--

DROP TABLE IF EXISTS `accounts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `accounts` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `id_account` LONGTEXT,
  `name` LONGTEXT,
  `password` LONGTEXT,
  `account_number` BIGINT DEFAULT NULL,
  `saldo` BIGINT DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=INNODB AUTO_INCREMENT=8 ;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `accounts`
--

LOCK TABLES `accounts` WRITE;
/*!40000 ALTER TABLE `accounts` DISABLE KEYS */;
INSERT INTO `accounts` VALUES (1,'id-414','Alex','$2a$10$i4neEnvSu34lrL94t2YbJeLRCn2pmFrDZ7MaUMfxiq6/rUk47ftMS',832712,106120),(2,'id-650','Budi','$2a$10$osBixkpTm1kjRsLTcXdySO8SDvoq6yIF03ZbkYB6HNRqWAyMXKkxu',577718,106120),(3,'id-957','Ani','$2a$10$E75HKtov1rnNXUzl9ijmlu05wpIXrqwLfGTuh4YSM2MsSchlQm72O',902992,106120),(4,'id-443','Joni','$2a$10$l3OQc8yHitD5LP5iSNRVS.hq82p/O.Ma0nZ6CEeQ7sF.JkBkE5UeS',621232,106120),(5,'id-483','Ahmad','$2a$10$rafGYK39jcUTSDzP3.YBcuBroDccYvGrCa15bSXvjTaoNvW6spRxe',351775,106120);
/*!40000 ALTER TABLE `accounts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transactions` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `transaction_type` BIGINT DEFAULT NULL,
  `transaction_description` LONGTEXT,
  `sender` BIGINT DEFAULT NULL,
  `amount` BIGINT DEFAULT NULL,
  `recipient` BIGINT DEFAULT NULL,
  `timestamp` BIGINT DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=INNODB AUTO_INCREMENT=30 ;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
INSERT INTO `transactions` VALUES (1,0,'beli bakso',902992,15000,832712,1604408082),(2,1,'',902992,10000,0,1604408228),(3,1,'',902992,50000,0,1604408286),(4,2,'',902992,100000,0,1604408368),(27,3,'',902992,2,0,1604452698),(28,3,'',902992,0,0,1604452734),(29,3,'',902992,0,0,1604452823);
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'digibank'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-11-04  8:23:06
