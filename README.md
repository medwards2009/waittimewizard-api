# Wait Time Wizard API

A REST API service that provides theme park wait times and information. This service acts as a proxy to the [Theme Parks API](https://api.themeparks.wiki/), resolving CORS policy issues for client-side applications.

## Features

- **Destinations Endpoint**: Get a list of all available theme park destinations
- **Live Data Endpoint**: Retrieve real-time wait times and operational status for a specific destination

## API Endpoints

```bash
GET /api/destinations      # List all available destinations
GET /api/live/:id         # Get live wait times for a specific destination
```

## Local Development

1. Install dependencies:

```bash
npm install
```

2. Start the development server:

```bash
npm run start:dev
```

3. Access the API documentation:

- Visit `http://localhost:8080/api-docs` for Swagger documentation

## Environment Variables

Create a `.env` file in the root directory:

```bash
PORT=8080
```

## Data Attribution

This project uses data from the [Theme Parks API](https://api.themeparks.wiki/). All wait times and park information are sourced from this API.

### Copyright Notice

This application is for educational purposes and is not monetized. All data is provided by the Theme Parks API under their terms of use. Visit their [documentation](https://api.themeparks.wiki/) for more information about data usage and licensing.

## License

This project is [MIT licensed](LICENSE) with the following exceptions:

- All theme park data is provided by and copyrighted to their respective owners
- Theme park wait times and information are sourced from the Theme Parks API
