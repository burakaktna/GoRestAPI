# GoRestAPI

This is a simple RESTful API written in Go. It allows you to perform CRUD (Create, Read, Update, Delete) operations on articles, tags and article tags.

## Installation

To run this project, you need to have Docker and Docker Compose installed on your system. Then, run the following command:

```
docker-compose up -d --build

```

This will start the API server on port 8096.

## Usage

To use the API, you can send HTTP requests to the following endpoints:

### Articles

- `GET /articles`: Get all articles
- `POST /articles`: Create a new article
- `GET /articles/:id`: Get an article by ID
- `PUT /articles/:id`: Update an article by ID
- `DELETE /articles/:id`: Delete an article by ID

### Tags

- `GET /tags`: Get all tags
- `POST /tags`: Create a new tag
- `GET /tags/:id`: Get a tag by ID
- `PUT /tags/:id`: Update a tag by ID
- `DELETE /tags/:id`: Delete a tag by ID

### Article Tags

- `GET /article_tags`: Get all article tags
- `POST /article_tags`: Create a new article tag
- `GET /article_tags/:id`: Get an article tag by ID
- `PUT /article_tags/:id`: Update an article tag by ID
- `DELETE /article_tags/:id`: Delete an article tag by ID

The API accepts and returns JSON data.

## Contributing

If you would like to contribute to this project, feel free to fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more information.