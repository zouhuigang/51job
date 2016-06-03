set time_zone='+8:00'; 
    -- --------------------------------------------------
    --  Table Structure for `51job/model.User`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `51job_user` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `id51` varchar(255) NOT NULL DEFAULT ''  UNIQUE,
        `jobyear` varchar(255),
        `age` varchar(255),
        `sex` varchar(255),
        `address` varchar(255),
        `major` varchar(255),
        `study` varchar(255),
        `date51` datetime NOT NULL,
        `created` datetime NOT NULL,
        `updated` datetime NOT NULL
    ) ENGINE=INNODB  DEFAULT CHARSET=utf8 COMMENT="简历表";

    -- --------------------------------------------------
    --  Table Structure for `51job/model.UserKeyword`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `51job_user_keyword` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `id51` varchar(255) NOT NULL DEFAULT '' ,
        `file_address` varchar(255) NOT NULL DEFAULT '' ,
        `date51` datetime NOT NULL,
        `created` datetime NOT NULL,
        `user_id` integer NOT NULL,
        `keyword_id` integer NOT NULL
    ) ENGINE=INNODB  DEFAULT CHARSET=utf8 COMMENT="简历关键字表";

    -- --------------------------------------------------
    --  Table Structure for `51job/model.Userinfo`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `51job_userinfo` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `id51` varchar(255) NOT NULL DEFAULT ''  UNIQUE,
        `date51` datetime NOT NULL,
        `content` longtext NOT NULL,
        `created` datetime NOT NULL,
        `updated` datetime NOT NULL,
        `user_id` integer UNIQUE
    ) ENGINE=INNODB  DEFAULT CHARSET=utf8 COMMENT="简历详情表";

    -- --------------------------------------------------
    --  Table Structure for `51job/model.Keyword`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `51job_keyword` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `keyword` varchar(255) NOT NULL DEFAULT '' ,
        `address` varchar(255) NOT NULL DEFAULT '' ,
        `kind` varchar(255) NOT NULL DEFAULT '' ,
        `created` datetime NOT NULL,
        `updated` datetime NOT NULL,
        `time51` integer NOT NULL DEFAULT 1 
    ) ENGINE=INNODB  DEFAULT CHARSET=utf8 COMMENT="关键字表";
