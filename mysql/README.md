<div align="center">
  <a href="https://www.mysql.com/">
    <img alt="mysql" src="../logos/mysql.png"/>
  </a>
  <h1>MySQL</h1>
</div>

# Table of Contents

- [Password Validation Component System Variables](#password-validation-component-system-variables)
- [Individual column's counter for different values](#individual-columns-counter-for-different-values)

## Password Validation Component System Variables

If we want to change the mysql password in our way for a specific user, we face some problems e.g it can't be possible when the `validate_password` component is enabled. To see the configuration of password checking we can run this command in a mysql shell:

```sh
mysql> SHOW VARIABLES LIKE 'validate_password.%';
+--------------------------------------+--------+
| Variable_name                        | Value  |
+--------------------------------------+--------+
| validate_password.check_user_name    | ON     |
| validate_password.dictionary_file    |        |
| validate_password.length             | 8      |
| validate_password.mixed_case_count   | 1      |
| validate_password.number_count       | 1      |
| validate_password.policy             | MEDIUM |
| validate_password.special_char_count | 1      |
+--------------------------------------+--------+
```

So we can Change password validation policy in MySQL like this:

```sh
mysql> SET GLOBAL validate_password.length = 4;
mysql> SET GLOBAL validate_password.policy=LOW;
```

Then, we will be able to change the password by running this command for the `root` user:

```sh
mysql> ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'root';
```

## Individual column's counter for different values

```sql
SELECT age, COUNT(*) FROM campers GROUP BY age;
```
