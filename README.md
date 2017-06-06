Simple Inventory
==================

API Documentation
------------------

### `/categories`

#### GET `/`
**Description**: Retrieves the list of all categories.
#####  Response
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
##### Response 
**Success**:
```
{
  "err" : 0,
  "msg" : [],
  "data" : null
}
```
**Error**:
```
{
  "err" : 1,
  "msg" : [
    "Message",
    ...
  ],
  "data" : null
}
```
#### GET `/{id}`
##### Response
**Success**
```
{
  "err" : 0,
  "msg" : [],
  "data" : {
    "name" : "foo category",
    "id" : 24,
    "parentId" : {
      "Int64" : 0,
      "Valid" : false
    }
  }
}
```
**Error**:
```
{
  "err" : 1,
  "msg" : [
    "Message",
    ...
  ],
  "data" : null
}
```
#### PUT `/{id}`
#### DELETE `/{id}`

### `/items`

#### GET `/`
#### POST `/`

#### GET `/{id}`
#### PUT `/{id}/categories`
#### PUT `/{id}/stock/{stockId}`
