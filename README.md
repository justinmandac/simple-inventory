Simple Inventory
==================

API Documentation
------------------

### `/categories`

#### GET `/`
**Description**: Retrieves the list of all categories.
##### Sample Response
**Success**:
```
{
  "err" : 0,
  "msg" : [],
  "data" : [
    {
      "name" : "adhesives",
      "id" : 1,
      "parentId" : {
        "Int64" : 0,
        "Valid" : false
      }
    },
    {
      "name" : "tools",
      "id" : 2,
      "parentId" : {
        "Int64" : 0,
        "Valid" : false
      }
    }
    },
    {
      "name" : "power tools",
      "id" : 3,
      "parentId" : {
        "Int64" : 2,
        "Valid" : true
      }
    }
  ]
}
```
**Error**:
```
{
  "err" : 1,
  "msg" : [
    "Sample Error Message",
    ...
  ],
  "data" : null
}
```
#### POST `/`
**Description**: Creates a new Category
##### Request Body:
```
{
  "name" : "foo category",
  "parentId" : null|integer
}
```

#### GET `/{id}`
#### PUT `/{id}`
#### DELETE `/{id}`

### `/items`

#### GET `/`
#### POST `/`

#### GET `/{id}`
#### PUT `/{id}/categories`
#### PUT `/{id}/stock/{stockId}`
