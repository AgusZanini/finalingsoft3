# Proyecto final Ingeniería de Software III

**Objetivo**: El objetivo proyecto es automatizar la construccion y el despliegue de una aplicacion, asi como las pruebas unitarias y de integracion.

**Desarrollo**: Para el desarrollo del proyecto se uso una aplicacion simple de desarrollo propio que tiene dos funcionalidades:
- Multiplicar dos numeros, mostrar el resultado y almacenar la operacion
- Buscar por id una operacion almacenada y mostrarla.

La aplicacion esta constituida por tres partes fundamentales:
- Backend: Golang, tiene 2 endpoints:
  - `/operations` Con un POST recibe un json donde van a estar los dos numeros a multiplicar, va almacenar la operacion y devolver el resultado para que el frontend lo muestre.
  - `/operations/:id` Con un GET recibe un id y va a traer la operacion que tenga ese id si es que existe.
- Frontend: React.js, se encarga de hacer las peticiones necesarias al backend, de mostrar el resultado de nuevas operaciones y de mostrar las operaciones buscadas por su id.
- Base de datos: Mysql, tiene una tabla `operaciones` que tendra los dos numeros implicados en cada operacion, el resultado y su id correspondiente.

Para el desarrollo del proceso de CI/CD, se opto por usar github actions y se hizo un workflow `ci-cd` donde se tienen los siguientes jobs:
- build-test-publish: Se encarga de la construccion de la aplicacion, correr los test unitarios y publicar las imagenes de los servicios a dockerhub
- deploy: Se encarga de desplegar los servicios de backend y frontend en Google Cloud Run y la base de datos en una instancia de Cloud SQL
- integration-tests: Se encarga de correr los test de integracion usando codeceptjs
