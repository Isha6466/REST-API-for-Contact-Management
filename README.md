# REST-API-for-Contact-Management
This project involves building a REST API for managing contacts using Go and the Gofr framework. The API offers CRUD operations (Create, Read, Update, Delete) for managing contact information stored in a relational database (presumably MySQL, based on imports).

Functionality in Steps:

1. Initialization:
  Create a Gofr application instance.
  Establish connection to the database (omitted in the provided code).
  Set the Gofr database connection to the established connection.

2. Endpoint Creation:
  Define API endpoints for each CRUD operation:
    1.GET /contacts: Retrieves a list of all contacts.
    2. GET /contacts/:id: Retrieves a specific contact by ID.
    3. POST /contacts: Creates a new contact.
    4. PUT /contacts/:id: Updates an existing contact.
    5. DELETE /contacts/:id: Deletes a contact.

3.Implementation:
  For each endpoint:
  Define a handler function that receives the HTTP request context.
  Extract relevant information from the context, like path parameters for :id.
  Use Gofr methods to interact with the database based on the endpoint:
    1. GET: app.Data.DB.Find for all contacts, app.Data.DB.First for specific contact.
    2. POST: app.Data.DB.Create to save a new contact.
    3. PUT: app.Data.DB.Updates to modify an existing contact.
    4. DELETE: app.Data.DB.Delete to remove a contact.
  Handle potential errors from database operations:
    Distinguish between record not found errors and other database issues.
    Send appropriate HTTP status codes for successful responses and different error types.
  Use Gofr methods to respond to the request with:
    Contacts list or single contact object for GET.
    Confirmation message (e.g., "Contact created") for POST, PUT, and DELETE.
    Error messages with details for failed operations.

4.Error Handling:
  Use different HTTP status codes to indicate successful and error responses:
    1. 200 OK for successful requests.
    2. 400 Bad Request for invalid input data.
    3. 404 Not Found for missing resources (e.g., contact not found).
    4. 500 Internal Server Error for unexpected errors.

5.Unit Tests:
  Include unit tests for critical functionalities like retrieving contacts.
  Use mocks to simulate dependencies like database interactions.
  Verify expected behavior of the handlers and responses.

Additionally
The provided code snippet omits the database connection section and potential ID generation logic.
This is a basic implementation; consider features like authentication, authorization, validation, and pagination for a more robust system.

In summary, this project involves building the backend and API logic for a contact management system in Go. 
We use Gofr to handle routing and responses, Gorm for database interaction, and HTTP status codes for clear communication of success and error conditions.

