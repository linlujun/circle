
/*
已经发生修改，不再使用

-- CREATE TABLE user_circle( 
-- 	   uid INT UNSIGNED NOT NULL,
-- 	   cid INT UNSIGNED NOT NULL,
--    	PRIMARY KEY (uid,cid)
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- CREATE TABLE masters( 
-- 	   mid INT UNSIGNED NOT NULL,
-- 	   cid INT UNSIGNED NOT NULL,
-- 	   create_time DATETIME,
--    	PRIMARY KEY (mid)
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- CREATE TABLE comments(
-- 	   id INT UNSIGNED AUTO_INCREMENT,
-- 	   content VARCHAR(160),
-- 	   picurl VARCHAR(100),
-- 	   tid INT UNSIGNED NOT NULL,
-- 	   uid INT UNSIGNED NOT NULL,
-- 	   cid INT UNSIGNED NOT NULL,
-- 	   ccid INT UNSIGNED,
-- 	   create_time DATETIME,
-- 	   like_count INT UNSIGNED,
--    	PRIMARY KEY (id)
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- CREATE TABLE topics(
-- 	   id INT UNSIGNED AUTO_INCREMENT,
-- 	   tid INT UNSIGNED NOT NULL,
-- 	   uid INT UNSIGNED NOT NULL,
-- 	   cid INT UNSIGNED NOT NULL,
-- 	   create_time DATETIME,
-- 	   contenturl VARCHAR(100),
-- 	   comment_count INT UNSIGNED,
--    	PRIMARY KEY (id)
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- CREATE TABLE cusers(
-- 	   id INT UNSIGNED AUTO_INCREMENT,
-- 	   cid INT UNSIGNED NOT NULL,
-- 	   uid INT UNSIGNED NOT NULL,
-- 	   jointime DATETIME,
-- 	   topic_count INT UNSIGNED  DEFAULT 0 ,
-- 	   comment_count INT UNSIGNED  DEFAULT 0 ,
--    	PRIMARY KEY (id)
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- CREATE TABLE circles(
-- 	   id INT UNSIGNED AUTO_INCREMENT,
-- 	   cname VARCHAR(20),
-- 	   description VARCHAR(160),
-- 	   masterid INT UNSIGNED NOT NULL,
-- 	   user_count INT UNSIGNED DEFAULT 0,
-- 	   topic_count INT UNSIGNED  DEFAULT 0 ,
-- 	   comment_count INT UNSIGNED  DEFAULT 0 ,
--    	PRIMARY KEY (id)
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- CREATE TABLE users(
-- 	   id INT UNSIGNED AUTO_INCREMENT,
-- 	   uname VARCHAR(20),
-- 	   pwd VARCHAR(20) NOT NULL,
-- 	   role INT UNSIGNED NOT NULL,
-- 	   nickname VARCHAR(20),
-- 	   userpic VARCHAR(100),
-- 	   description VARCHAR(160),
-- 	   topic_count INT UNSIGNED  DEFAULT 0 ,
-- 	   comment_count INT UNSIGNED  DEFAULT 0 ,
--    	PRIMARY KEY (id)
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8;


*/
-- CREATE TABLE session( 
-- 	   session_id VARCHAR(100),
-- 	   login_name VARCHAR(20),
-- 	   TTL INT
--    	PRIMARY KEY (session_id)
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8;
