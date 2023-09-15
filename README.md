# graphqlExample

## Test Server:
Test url: http://localhost:8060/graphql

## Test Queries:

### 1. Mutation:

```
mutation { 
    UpdateDB (ID : 100, Name : "Lana", ComplexSetArgs: {C1param: "iota", C2param: "real"})
    { 
        ID
        Name
        Complex {
            C1param
            C2param
        }
    } 
}
```

### 2. Query:

```
query { 
        getDB {
        ID
        Name
        Complex {
            C1param
            C2param
        }
    }
    setDB (ID : 1, Name : "John", ComplexArg: {C1param: "complex iota", C2param: "complex real"})
    { 
        ID
        Name
        Complex {
            C1param
            C2param
        }
    } 
}

```

## Development:
### Note: Please make sure to name the field same as DB object.