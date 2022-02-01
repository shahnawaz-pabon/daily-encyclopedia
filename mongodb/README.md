<div align="center">
  <a href="https://www.mongodb.com/">
    <img alt="mongodb" src="../logos/mongodb.png"/>
  </a>
  <h1>MongoDB</h1>
</div>

# Table of Contents

- [MongoDB Schema Design Best Practices](https://www.mongodb.com/developer/article/mongodb-schema-design-best-practices/)
- [MongoDB Shell Commands](mongodb-shell-commands)

## MongoDB Shell Commands

```sh
# add key value to all documents
db.collection_name.updateMany({}, {$set:{"like_count": 1}})

# remove key value from all documents
db.collection_name.updateMany({}, {$unset:{"notification_details.emergency_label": ""}})

# remove multiple documents
db.collection_name.remove({"id":{$nin:[1637042527, 1637127944, 1637213711]}})
```
