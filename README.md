# Tax Calculator

This project will let you to compare your taxes for new tax Regime, based on 2025 budget.

The project uses ReactJs as frontend and Go as the backend service to do the calculation.

Enjoy!!

## Prerequisites

In order to run the project you need [Git](https://git-scm.com/downloads) and [Docker](https://www.docker.com/get-started/) installed on your system.
Follow the links to find installation instructions.

## 1. Clone the repository

```sh
git clone https://github.com/drshott/tax-service.git
```

## 2. Build and run the image using Dockerfile

```sh
cd tax-service
docker build -t tax-service .
docker run -itd -p 80:8080 tax-ctr tax-service
```

## 3. Access the application

You can now access the application on your browser using http://localhost/