<div align="center">
  <a href="https://www.mongodb.com/">
    <img alt="mongodb" src="../logos/mongodb.png"/>
  </a>
  <h1>MongoDB</h1>
</div>

# Table of Contents

- [MongoDB Schema Design Best Practices](https://www.mongodb.com/developer/article/mongodb-schema-design-best-practices/)
- [MongoDB Shell Commands](mongodb-shell-commands)
- [Miscellaneous Queries](#miscellaneous-queries)

## MongoDB Shell Commands

```sh
# add key value to all documents
db.collection_name.updateMany({}, {$set:{"like_count": 1}})

# remove key value from all documents
db.collection_name.updateMany({}, {$unset:{"notification_details.emergency_label": ""}})

# remove multiple documents
db.collection_name.remove({"id":{$nin:[1637042527, 1637127944, 1637213711]}})
```

## Miscellaneous Queries

Complex query with aggregation and pagination

```python
# page_number will be given
# pagination starts
skip = 15
limit = 15

if page_number > 0:
    page_number = page_number - 1
else:
    page_number = 0

skip = page_number * skip

# End of pagination

pipeline = [
    {
        "$match": {"is_deleted": False},
    },
    {
        "$lookup": {
            "from": "tasks",
            "let": {"cohort_id": "$cohort_id"},
            "pipeline": [
                {"$match": {"$expr": {"$eq": ["$$cohort_id", "$cohort_id"]}}},
            ],
            "as": "tasks",
        }
    },
    {
        "$project": {
            "_id": 0,
            "cohort_id": 1,
            "title": 1,
            "description": 1,
            "total_send": {"$sum": "$tasks.total_send"},
            "total_received": {"$sum": "$tasks.total_success"},
            "created_at": 1,
            "type": 1,
            "status": {
                "$cond": {
                    "if": {"$eq": ["$type", "Repeat"]},
                    "then": "Repeat",
                    "else": {
                        "$cond": {
                            "if": {
                                "$eq": [
                                    {"$arrayElemAt": ["$tasks.status", 0]},
                                    "RUNNING",
                                ]
                            },
                            "then": "Active",
                            "else": {
                                "$cond": {
                                    "if": {
                                        "$eq": [
                                            {"$arrayElemAt": ["$tasks.status", 0]},
                                            "PENDING",
                                        ]
                                    },
                                    "then": "Upcoming",
                                    "else": {
                                        "$cond": {
                                            "if": {
                                                "$eq": [
                                                    {
                                                        "$arrayElemAt": [
                                                            "$tasks.status",
                                                            0,
                                                        ]
                                                    },
                                                    "DONE",
                                                ]
                                            },
                                            "then": "Done",
                                            "else": {
                                                "$cond": {
                                                    "if": {
                                                        "$eq": [
                                                            {
                                                                "$arrayElemAt": [
                                                                    "$tasks.status",
                                                                    0,
                                                                ]
                                                            },
                                                            "KILLED",
                                                        ]
                                                    },
                                                    "then": "Inactive",
                                                    "else": "created",
                                                }
                                            },
                                        }
                                    },
                                }
                            },
                        }
                    },
                }
            },
        },
    },
    {
        "$addFields": {
            "success_rate": {
                "$round": [
                    {
                        "$multiply": [
                            {
                                "$cond": [
                                    {"$eq": ["$total_send", 0]},
                                    0,
                                    {"$divide": ["$total_received", "$total_send"]},
                                ]
                            },
                            100,
                        ]
                    },
                    0,
                ]
            }
        }
    },
    {"$sort": {"cohort_id": -1}},
    {"$skip": skip},
    {"$limit": limit},
]

response = collection_name.aggregate(pipeline).to_list(length=None)

# return response
```
