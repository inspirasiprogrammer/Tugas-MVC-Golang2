/*
SQLyog Ultimate v12.4.3 (64 bit)
MySQL - 5.7.24 : Database - digibank
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`digibank` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `digibank`;

/*Table structure for table `accounts` */

DROP TABLE IF EXISTS `accounts`;

CREATE TABLE `accounts` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `id_account` longtext,
  `name` longtext,
  `password` longtext,
  `account_number` bigint(20) DEFAULT NULL,
  `saldo` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

/*Data for the table `accounts` */

insert  into `accounts`(`id`,`id_account`,`name`,`password`,`account_number`,`saldo`) values 
(1,'id-414','Irwan Syahputra','$2a$10$Q6uPSqviWn4FtBKCjHcNc.HUWNQUPKSuXsetbkY3Y5NPg2DAAYgqe',832712,2428977),
(2,'id-650','Zaid Mujtaba','$2a$10$q5iKj2GnvlEB3ShSF6Dvu.k80Xy6kfyq8KSvYNNkP3JlkGaiHgFES',577718,1380100);

/*Table structure for table `transactions` */

DROP TABLE IF EXISTS `transactions`;

CREATE TABLE `transactions` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `transaction_type` bigint(20) DEFAULT NULL,
  `transaction_description` longtext,
  `sender` bigint(20) DEFAULT NULL,
  `amount` bigint(20) DEFAULT NULL,
  `recipient` bigint(20) DEFAULT NULL,
  `timestamp` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1;

/*Data for the table `transactions` */

insert  into `transactions`(`id`,`transaction_type`,`transaction_description`,`sender`,`amount`,`recipient`,`timestamp`) values 
(1,0,'',832712,250000,577718,1604470346),
(2,1,'',832712,50000,0,1604470464),
(3,2,'',832712,1500000,0,1604472621),
(4,3,'',832712,0,0,1604473083),
(5,3,'',832712,0,0,1604473167),
(6,3,'',832712,0,0,1604473168),
(7,3,'',832712,0,0,1604473169),
(8,3,'',832712,0,0,1604473228);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
