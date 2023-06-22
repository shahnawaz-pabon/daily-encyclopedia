<div align="center">
  <a href="https://www.microsoft.com/en-us/sql-server/sql-server-downloads">
    <img alt="mssql" src="../logos/mssql.png"/>
  </a>
  <h1>SQL Server</h1>
</div>

# Table of Contents

- [Update a column from another column's value with concat](#update-a-column-from-another-columns-value-with-concat)
- [Error IDENTITY_INSERT is set to OFF](#error-identity_insert-is-set-to-off)

## Update a column from another column's value with concat

```sql
update INSTITUTE set EMAIL=concat(EIIN_NO, '@gmail.com') WHERE EMAIL NOT LIKE '%@%';
```

## Error IDENTITY_INSERT is set to OFF

While trying to insert data into the sql server's table, this error occurs: `Cannot insert explicit value for identity column in table 'TABLE_NAME' when IDENTITY_INSERT is set to OFF`. If you genuinely need to insert explicit values into an identity column, you can enable the `IDENTITY_INSERT` option for the table. However, please note that this should be used with caution, as it overrides the default behavior of the identity column.

```sh
SET IDENTITY_INSERT MAPPING_ZONAL_OFFICE_DISTRICT ON;

-- Perform your INSERT statement here

SET IDENTITY_INSERT MAPPING_ZONAL_OFFICE_DISTRICT OFF;
```

Remember, when using the IDENTITY_INSERT option, you need to explicitly specify values for the identity column in your INSERT statement. Ensure that the provided values do not conflict with existing identity values in the column.
