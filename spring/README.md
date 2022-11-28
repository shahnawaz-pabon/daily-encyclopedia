<div align="center">
  <a href="https://spring.io/">
    <img alt="spring" src="../logos/spring.png" height="96"/>
  </a>
  <h1>Spring</h1>
</div>

# Table of Contents

- [Parse datetime with specific locale](#parse-datetime-with-specific-locale)

## Parse datetime with specific locale

```java
// Import
import java.util.Date;
import java.text.DateFormat;
import java.util.Locale;

Locale localeEn = new Locale("en", "US");
DateFormat dateFormatEn = DateFormat.getDateTimeInstance(DateFormat.DEFAULT, DateFormat.SHORT, localeEn);
String dateEn = dateFormatEn.format(new Date());

Locale localeBn = new Locale("bn", "BD");
DateFormat dateFormatBn = DateFormat.getDateTimeInstance(DateFormat.DEFAULT, DateFormat.SHORT, localeBn);
String dateBn = dateFormatBn.format(new Date());
```
