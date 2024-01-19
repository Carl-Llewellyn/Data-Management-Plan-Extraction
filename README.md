# Data Management Plan - API for Kafka and Postgres

This project was created to provide an efficient and consistent methodology to collect the requirements of researcher’s project requirements and needs at the Department of Fisheries and Oceans. Many of the questions have been removed on the request of the department. It allows for instant data analytics through Kafka, in addition to separate endpoints for direct Postgres integration and file management.

## Directories

### Routes directory
Contains a "routes" directory for the handlers which is split by Kafka and Postgres. Within the "spi" directory are also all configurations (using a .env file) in addition to structs for json models.

### Logger directory
Contains a logging package to allow for easy logging to a file and console, depending on console flags.

### .devcontainer directory

Docker dcontainer json for remote development environments.

### Tests directory

An example of assertations that can be done on the codebase to confirm functionality.
