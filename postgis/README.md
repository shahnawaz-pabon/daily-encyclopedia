<div align="center">
  <a href="https://postgis.net/">
    <img alt="postgis" src="../logos/postgis.png"/>
  </a>
  <h1>PostGIS</h1>
</div>

# Table of Contents

- [Setup POSTGIS in Docker](#setup-postgis-in-docker)
- [Query for points inside a polygon](#query-for-points-inside-a-polygon)
- [SQL query to have a complete geojson feature from PostGIS](#sql-query-to-have-a-complete-geojson-feature-from-postgis)

## Setup POSTGIS in Docker

```sh
docker run -d -p 5432:5432 --name db kartoza/postgis:13.0
```

**Postgis docker’s Credentials:**

```sh
Username: docker
Password: docker
Database: gis
```

## Query for points inside a polygon

```sql
SELECT e.id FROM institute e WHERE ST_Contains(ST_GeomFromText(polygon_text), st_point(e.y,e.x))
```

Here `polygon_text` is a geometry text like this `POLYGON((-103.30549637500008 20.852735681153252,-103.08103481249998 20.612974162085475,-101.6261045 20.537532106266806,-99.83567868749998 20.395877027062447,-99.80306537500002 22.0572706994358,-99.64994812500004 28.918636198451633,-121.1212769375 8.69559423007209,-103.30549637500008 20.852735681153252))`

## SQL query to have a complete geojson feature from PostGIS

```sql
SELECT jsonb_build_object(
  'type',     'FeatureCollection',
  'features', jsonb_agg(feature)
)
FROM (
  SELECT jsonb_build_object(
    'type',       'Feature',
    'id',         gid,
    'geometry',   ST_AsGeoJSON(geom)::jsonb,
    'properties', to_jsonb(inputs)
  ) AS feature
  FROM (
    SELECT * FROM input_table
  ) inputs
) features;
```
