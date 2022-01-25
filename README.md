# Mongo data removal agent

Service for sheduling removes of entities in Mongo.

Repository also contains:
- Tool for seeding database
- Simple API for testing request to database

## Agent usage

### CLI agruments
| Argument | Description | Required | Default |
|----------|------------|----------|---------|
| `--uri` | URI to MongoDB | true  | none |
| `--db` | Database name | true  | none |
| `--col` | Collection name | true  | none |
| `--workers` | URI to MongoDB | false  | 10 |
| `--port` | Amount of workers that handle deletions| false  | 80 |

### API
| Method | URI | Description |
|--------|-----|-------------|
| POST | `delete/{clientId}` | Delete data for whole Client |
| POST | `delete/{clientId}/{userId}` | Delete data for specific User |
| GET | `exists/{clientId}` | Check if Client has any data |
| GET | `exists/{clientId}/{userId}` | Check if User has any data |