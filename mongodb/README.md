<div align="center">
  <a href="https://www.mongodb.com/">
    <img alt="mongodb" src="../logos/mongodb.png"/>
  </a>
  <h1>MongoDB</h1>
</div>

# Table of Contents

- [Setup mongodb via docker](#setup-mongodb-via-docker)
- [MongoDB Schema Design Best Practices](https://www.mongodb.com/developer/article/mongodb-schema-design-best-practices/)
- [Query Keywords and Functions](#query-keywords-and-functions)
- [MongoDB Shell Commands](#mongodb-shell-commands)
- [Complex query with aggregation and pagination](#complex-query-with-aggregation-and-pagination)
- [Pagination with linked collections](#pagination-with-linked-collections)
- [Query to format data for the line and bar chart](#query-to-format-data-for-the-line-and-bar-chart)
- [Group by query with a specific condition](#group-by-query-with-a-specific-condition)

## Setup mongodb via docker

```sh
# create mongodb_data directory with writable permission by docker process
mkdir -m 777 mongodb_data
export MONGODB_VERSION=6.0-ubi8
docker run --name mongodb -d -p 27017:27017 -v $(pwd)/mongodb_data:/data/db mongodb/mongodb-community-server:$MONGODB_VERSION
```

## Query Keywords and Functions

- **`$unwind`**: This operator breaks down an array field within input documents, producing a separate document for each element.

- **`$match`**: This operator identifies and retrieves existing documents in the database that satisfy specified conditions.

- **`$project`**: This operator enables the extraction of specific fields from the collection.

- **`aggregate`**: Aggregation in MongoDB is like a series of steps (pipeline) for handling lots of documents. Each step can filter, sort, group, reshape, or modify the documents in the collection. It's a way to process and transform data efficiently.

## MongoDB Shell Commands

```sh
# add key value to all documents
db.collection_name.updateMany({}, {$set:{"like_count": 1}});

# remove key value from all documents
db.collection_name.updateMany({}, {$unset:{"notification_details.emergency_label": ""}});

# remove multiple documents
db.collection_name.remove({
  id: { $nin: ['1637042527', '1637127944', '1637213711'] }
});

# getting a user's schedule on Monday whose email is admin@admin.com

db.users.aggregate([
  {
    $match: {
      email: "admin@admin.com"
    }
  },
  {
    $unwind: "$schedule"
  },
  {
    $match: {
      "schedule.dayOfWeek": "Monday"
    }
  },
  {
    $project: {
      _id: 0,
      dayOfWeek: "$schedule.dayOfWeek",
      timeSlots: "$schedule.timeSlots"
    }
  }
]);

# getting list between two dates of a specific id
db.collection_name.find({"createTime": {
  '$gte': "2023-12-01",
  '$lte': "2024-01-06"
},
  "locationId": "your id"
});
```

## Complex query with aggregation and pagination

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

## Pagination with linked collections

```js
const { page = 1, size = 10 } = req.query;

/**
 * Pagination Details
 * User - Collection
 * Portfolio - Collection
 */
const count = await User.countDocuments();
const totalPages = Math.ceil(count / size);
const currentPage = parseInt(page);

// API logic for getting users with pagination
const users = await User.find()
  .skip((currentPage - 1) * size)
  .limit(size)
  .populate("portfolio");
```

## Query to format data for the line and bar chart

This query finds the total reviews along with the individual counts for star rating

```python
def generate_monthly_report(location_id: str, year: int):

    year_str = str(year)

    # Construct start_date and end_date
    start_date = f"{year_str}-01-01"
    end_date = f"{int(year_str) + 1}-01-01"

    pipeline = [
        {
            "$match": {
                "locationId": location_id,
                "createTime": {"$gte": start_date, "$lt": end_date},
            }
        },
        {
            "$addFields": {
                "parsedCreateTime": {
                    "$dateFromParts": {
                        "year": {"$year": {"$dateFromString": {"dateString": "$createTime"}}},
                        "month": {"$month": {"$dateFromString": {"dateString": "$createTime"}}},
                        "day": {"$dayOfMonth": {"$dateFromString": {"dateString": "$createTime"}}},
                        "hour": {"$hour": {"$dateFromString": {"dateString": "$createTime"}}},
                        "minute": {"$minute": {"$dateFromString": {"dateString": "$createTime"}}},
                        "second": {"$second": {"$dateFromString": {"dateString": "$createTime"}}},
                        "millisecond": {"$millisecond": {"$dateFromString": {"dateString": "$createTime"}}},
                    }
                }
            }
        },
        {
            "$group": {
                "_id": {"$month": "$parsedCreateTime"},
                "total": {"$sum": 1},
                "1 star": {"$sum": {"$cond": [{"$eq": ["$starRating", "ONE"]}, 1, 0]}},
                "2 star": {"$sum": {"$cond": [{"$eq": ["$starRating", "TWO"]}, 1, 0]}},
                "3 star": {"$sum": {"$cond": [{"$eq": ["$starRating", "THREE"]}, 1, 0]}},
                "4 star": {"$sum": {"$cond": [{"$eq": ["$starRating", "FOUR"]}, 1, 0]}},
                "5 star": {"$sum": {"$cond": [{"$eq": ["$starRating", "FIVE"]}, 1, 0]}},
            }
        },
        {
            "$project": {
                "name": {"$dateToString": {"format": "%b", "date": {"$dateFromParts": {"year": year, "month": "$_id", "day": 1}}}},
                # "month": "$_id.month",
                "total": 1,
                "1 star": 1,
                "2 star": 1,
                "3 star": 1,
                "4 star": 1,
                "5 star": 1,
                "_id": 0,
            }
        },
        {
            "$sort": {"createTime": 1}
        }
    ]

    result = list(Review.aggregate(pipeline))
    return result

```

## Group by query with a specific condition

```python
def get_location_wise_bad_reviews():
    yesterday = (datetime.now() - timedelta(days=1)
                 ).strftime("%Y-%m-%dT%H:%M:%SZ")

    # Specify the list of ratings considered as bad reviews
    bad_ratings = ["ONE", "TWO"]

    # Group reviews by locationId
    pipeline = [
        {
            "$match": {
                "starRating": {"$in": bad_ratings},
                "createTime": {"$gte": yesterday},
            }
        },
        {
            "$group": {
                "_id": "$locationId",
                "reviews": {"$push": "$$ROOT"}
            }
        }
    ]

    location_reviews = list(Review.aggregate(pipeline))

    return location_reviews

```
