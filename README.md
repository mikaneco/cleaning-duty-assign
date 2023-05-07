# Weekly Cleaning Duty Assignment Program
This program retrieves a list of members using the Google Sheets API and assigns cleaning duties on a weekly basis. Assignments are done in order, with each member being assigned on a week-by-week basis.

## Required Environment Variables
Define the following environment variables in a .env file:

- `SHEET_ID`: The sheet ID of the Google Sheets.
- `API_KEY`: The API key for accessing the Google Sheets API.
- `SHEET_NAME`: The name of the sheet containing the member list.

In the Google Spreadsheet, please define the id in column A and the name in column B. The program will fetch this information to assign cleaning duties on a weekly basis. Make sure to follow this format for the program to work correctly.
|id|name|
|----|----|
|1|John|
|2|Alex|
|3|Chris|

## Usage
1. Install the Go programming language.
2. Clone this repository.
3. In the directory of the cloned repository, run go mod init to initialize the dependencies.
4. Create a .env file and define the environment variables mentioned above.
5. Execute go run main.go to run the program. The cleaning duty assignments for each week will be output to the console.

## Note
This program is subject to rate limiting by the Google Sheets API. Please avoid sending a large number of requests in a short period. Additionally, keep your API key secret and set appropriate restrictions to prevent unauthorized access.
