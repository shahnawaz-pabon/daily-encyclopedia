<div align="center">
  <a href="https://www.microsoft.com/en-us/sql-server/sql-server-downloads">
    <img alt="mssql" src="../logos/mssql.png"/>
  </a>
  <h1>SQL Server</h1>
</div>

## Update a column from another column's value with concat

```sql
update INSTITUTE set EMAIL=concat(EIIN_NO, '@gmail.com') WHERE EMAIL NOT LIKE '%@%';
```
