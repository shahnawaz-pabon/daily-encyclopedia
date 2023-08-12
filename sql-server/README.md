<div align="center">
  <a href="https://www.microsoft.com/en-us/sql-server/sql-server-downloads">
    <img alt="mssql" src="../logos/mssql.png"/>
  </a>
  <h1>SQL Server</h1>
</div>

## Table of Contents

- [Update a column from another column's value with concat](#update-a-column-from-another-columns-value-with-concat)
- [Error IDENTITY_INSERT is set to OFF](#error-identity_insert-is-set-to-off)
- [Get calculated salary with considering other benefits](#get-calculated-salary-with-considering-other-benefits)

### Update a column from another column's value with concat

```sql
update INSTITUTE set EMAIL=concat(EIIN_NO, '@gmail.com') WHERE EMAIL NOT LIKE '%@%';
```

### Error IDENTITY_INSERT is set to OFF

While trying to insert data into the sql server's table, this error occurs: `Cannot insert explicit value for identity column in table 'TABLE_NAME' when IDENTITY_INSERT is set to OFF`. If you genuinely need to insert explicit values into an identity column, you can enable the `IDENTITY_INSERT` option for the table. However, please note that this should be used with caution, as it overrides the default behavior of the identity column.

```sh
SET IDENTITY_INSERT MAPPING_ZONAL_OFFICE_DISTRICT ON;

-- Perform your INSERT statement here

SET IDENTITY_INSERT MAPPING_ZONAL_OFFICE_DISTRICT OFF;
```

Remember, when using the IDENTITY_INSERT option, you need to explicitly specify values for the identity column in your INSERT statement. Ensure that the provided values do not conflict with existing identity values in the column.

### Get calculated salary with considering other benefits

```sql
SELECT
  BCD.[AMOUNT] AS [AMOUNT]
FROM
  IEIMS_DEV_DTE_NON_GOVT.dbo.SALARY_CONFIGURATION SC
  JOIN (
    SELECT
      SC1.SALARY_GRADE_SCALE_MAP_ID,
      SC1.INCREMENT_STEP_NO,
      SUM(
        CASE WHEN BCD1.[CALCULATION_TYPE] = 'FIXED' THEN BCD1.AMOUNT WHEN BCD1.[CALCULATION_TYPE] = 'PERCENTAGE' THEN CEILING(
          (BCD1.AMOUNT / 100) * SC1.BASIC_SALARY
        ) END * CASE WHEN BCD1.[SALARY_BREAKDOWN_NATURE] = 'DEDUCTION' THEN -1 WHEN BCD1.[SALARY_BREAKDOWN_NATURE] = 'ADDITION' THEN 1 ELSE 1 END
      ) + SC1.BASIC_SALARY AS [AMOUNT]
    FROM
      SALARY_GRADE_SCALE_MAPPING SGSM
      JOIN BREAKDOWN_CONFIGURATION BC1 ON SGSM.SALARY_SCALE_ID = BC1.SALARY_PAY_SCALE_ID
      JOIN BREAKDOWN_CONFIGURATION_DETAILS BCD1 ON BC1.BREAKDOWN_CONFIGURATION_ID = BCD1.BREAKDOWN_CONFIGURATION_ID
      JOIN SALARY_BREAKDOWN SB ON SB.SALARY_BREAKDOWN_ID = BCD1.SALARY_BREAKDOWN_ID
      JOIN SALARY_CONFIGURATION SC1 ON SC1.SALARY_GRADE_SCALE_MAP_ID = SGSM.SALARY_GRADE_SCALE_MAP_ID
    WHERE
      BC1.RECORD_STATUS = 'ACTIVE'
      AND BC1.END_DATE IS NULL
      AND ((:arrearType = 'SALARY' AND SB.SALARY_BREAKDOWN_TYPE != 'BONUS') OR (:arrearType = 'BONUS' AND SB.SALARY_BREAKDOWN_TYPE = 'BONUS'))
      AND SGSM.SALARY_GRADE_SCALE_MAP_ID BETWEEN BCD1.FROM_SALARY_GRADE_ID
      AND BCD1.TO_SALARY_GRADE_ID
    GROUP BY
      SC1.SALARY_GRADE_SCALE_MAP_ID,
      SC1.INCREMENT_STEP_NO,
      SC1.BASIC_SALARY
  ) AS BCD ON SC.SALARY_GRADE_SCALE_MAP_ID = BCD.SALARY_GRADE_SCALE_MAP_ID
  AND SC.INCREMENT_STEP_NO = BCD.INCREMENT_STEP_NO
WHERE
  SC.SALARY_GRADE_SCALE_MAP_ID=:gradeMapId
  AND SC.INCREMENT_STEP_NO=:stepNo
```
