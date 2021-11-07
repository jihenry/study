#创建测试表
DROP TABLE IF EXISTS test_flush_log;

CREATE TABLE test_flush_log(id int, name char(50)) ENGINE = innodb;

#创建插入指定行数的记录到测试表中的存储过程
DROP PROCEDURE IF EXISTS proc;

delimiter $$ CREATE PROCEDURE proc(i int) BEGIN
DECLARE s int DEFAULT 1;

DECLARE c CHAR(50) DEFAULT repeat('a', 50);

while s <= i DO START TRANSACTION;

INSERT INTO test_flush_log
VALUES(NULL, c);

COMMIT;

SET s = s + 1;

END WHILE;

END $$ delimiter;