# BackendShedulerProGolang

## There is an endpoint in the project - /api

### Query types

- `GET` - /api/mailing/
- `GET` - /api/mailing/:id

- `POST` - /api/mailing/
  - type: int
  - filter: str
  - groupFlag: bool
  - singleFileFlag: bool
  - address: str
  - chatId: str
- `PUT` - /api/mailing/:id
  - type: int
  - filter: str
  - groupFlag: bool
  - singleFileFlag: bool
  - address: str
  - chatId: str
- `DELETE` - /api/mailing/:id

### To run the project on your device, you need to clone the repository and create a database. After that, install all the necessary packages.

> [!IMPORTANT]
> The `.env` file has been attached as an example.
