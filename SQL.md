## 設計一張資料表 並撰寫 sql 找出第一次登入後7天內還有登入的使用者 ##
例如：3/10第一次登入，3/12有再登入，滿足第一次登入後7天內還有登入

   - 任何 sql 語言回答皆可 
   - 簡單描述語法邏輯
   - 答案請提供 schema (column, type) 與 sql 

Table Schema:
   user_id:
      dataType: bigint
   login_time:
      dataType: timestamp

   CREATE TABLE user_logins (
      user_id BIGINT NOT NULL,
      login_time TIMESTAMP NOT NULL,
      PRIMARY KEY (user_id, login_time)
   );

思路:
   我把login_time當成PK, 讓他照login_time順序排序, 再找出group by user_id, login_time limit 1的紀錄, 這裡用windows function的RANK去做, 最後再join user_logins並找7天內的第二筆

sql:
   SELECT DISTINCT u.user_id
   FROM user_logins u
   JOIN first_result f ON u.user_id = f.user_id
      (SELECT *
      FROM (
         SELECT *, ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY login_time) AS rn
         FROM user_logins
      ) t
      WHERE rn = 1;) first_result f
   WHERE u.login_time > f.first_login
   AND u.login_time <= f.first_login + INTERVAL '7 days';


