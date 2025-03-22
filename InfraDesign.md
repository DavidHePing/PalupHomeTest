假設公司有個訂單系統, 所有使用者的資料都存在MySQL資料庫底下, 請設計一個報表系統提供分析數據, 
並滿足「在資料量大的情況下系統仍可以正常運作」
試著描述出此系統的大致架構與思維

方法1:
    MySql OLTP DB -> MySql OLAP DB
    用排程或CDC的方式, 將資料從OLTP的主DB搬到專給OLAP的MySql
    針對MySql OLAP DB做partition by create_time
    報表系統去OLAP DB by 時間去撈

方法2:
    MySql OLTP DB -> Column base OLAP DB
    用排程或CDC的方式, 將資料從OLTP的主DB搬到比較適合OLAP的Column base DB, 例如ClickHouse, HBase等
    報表系統從OLAP DB去撈