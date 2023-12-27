# Spotify Challenge

## Overview

The Spotify Challenge project is a comprehensive solution that addresses the complexities of interacting with the Spotify API. It is designed to create a robust web application capable of fetching track data from Spotify and seamlessly storing it in a MySQL database.

## Features

- **Create Track**: This endpoint facilitates the creation of a new track. It achieves this by dynamically fetching data from the Spotify API and efficiently storing it in the database.

- **Get Track by ISRC**: This endpoint provides a means to retrieve detailed track information by supplying the International Standard Recording Code (ISRC).

- **Get Tracks by Artist**: With this endpoint, users can conveniently retrieve tracks by specifying the artist's name.

## Prerequisites

Before running the application, ensure that the following prerequisites are met:

- Go (Golang)
- MySQL
- Spotify Developer Account with API credentials

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/kvbendalam/spotifiy-challenge.git
   cd spotifiy-challenge
   ```

2. Run the application:

   ```bash
   go run main.go
   ```

3. Access the API endpoints:

   - **Create Track**: `POST /track`
   - **Get Track by ISRC**: `GET /track/{isrc}`
   - **Get Tracks by Artist**: `GET /tracks/artist/{artist}`

Feel free to explore and utilize these endpoints to interact with the Spotify Challenge project seamlessly.
